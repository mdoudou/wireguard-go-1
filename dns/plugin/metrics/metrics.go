// Package metrics implement a handler and plugin that provides Prometheus metrics.
package metrics

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/inverse-inc/wireguard-go/dns/plugin"
	"github.com/inverse-inc/wireguard-go/dns/plugin/pkg/reuseport"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics holds the prometheus configuration. The metrics' path is fixed to be /metrics .
type Metrics struct {
	Next plugin.Handler
	Addr string
	Reg  *prometheus.Registry

	ln      net.Listener
	lnSetup bool

	mux *http.ServeMux
	srv *http.Server

	zoneNames []string
	zoneMap   map[string]struct{}
	zoneMu    sync.RWMutex
}

// New returns a new instance of Metrics with the given address.
func New(addr string) *Metrics {
	met := &Metrics{
		Addr:    addr,
		Reg:     prometheus.DefaultRegisterer.(*prometheus.Registry),
		zoneMap: make(map[string]struct{}),
	}

	return met
}

// MustRegister wraps m.Reg.MustRegister.
func (m *Metrics) MustRegister(c prometheus.Collector) {
	err := m.Reg.Register(c)
	if err != nil {
		// ignore any duplicate error, but fatal on any other kind of error
		if _, ok := err.(prometheus.AlreadyRegisteredError); !ok {
			log.Fatalf("Cannot register metrics collector: %s", err)
		}
	}
}

// AddZone adds zone z to m.
func (m *Metrics) AddZone(z string) {
	m.zoneMu.Lock()
	m.zoneMap[z] = struct{}{}
	m.zoneNames = keys(m.zoneMap)
	m.zoneMu.Unlock()
}

// RemoveZone remove zone z from m.
func (m *Metrics) RemoveZone(z string) {
	m.zoneMu.Lock()
	delete(m.zoneMap, z)
	m.zoneNames = keys(m.zoneMap)
	m.zoneMu.Unlock()
}

// ZoneNames returns the zones of m.
func (m *Metrics) ZoneNames() []string {
	m.zoneMu.RLock()
	s := m.zoneNames
	m.zoneMu.RUnlock()
	return s
}

// OnStartup sets up the metrics on startup.
func (m *Metrics) OnStartup() error {
	ln, err := reuseport.Listen("tcp", m.Addr)
	if err != nil {
		log.Errorf("Failed to start metrics handler: %s", err)
		return err
	}

	m.ln = ln
	m.lnSetup = true

	m.mux = http.NewServeMux()
	m.mux.Handle("/metrics", promhttp.HandlerFor(m.Reg, promhttp.HandlerOpts{}))

	// creating some helper variables to avoid data races on m.srv and m.ln
	server := &http.Server{Handler: m.mux}
	m.srv = server

	go func() {
		server.Serve(ln)
	}()

	ListenAddr = ln.Addr().String() // For tests.
	return nil
}

// OnRestart stops the listener on reload.
func (m *Metrics) OnRestart() error {
	if !m.lnSetup {
		return nil
	}
	u.Unset(m.Addr)
	return m.stopServer()
}

func (m *Metrics) stopServer() error {
	if !m.lnSetup {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := m.srv.Shutdown(ctx); err != nil {
		log.Infof("Failed to stop prometheus http server: %s", err)
		return err
	}
	m.lnSetup = false
	m.ln.Close()
	return nil
}

// OnFinalShutdown tears down the metrics listener on shutdown and restart.
func (m *Metrics) OnFinalShutdown() error { return m.stopServer() }

func keys(m map[string]struct{}) []string {
	sx := []string{}
	for k := range m {
		sx = append(sx, k)
	}
	return sx
}

// ListenAddr is assigned the address of the prometheus listener. Its use is mainly in tests where
// we listen on "localhost:0" and need to retrieve the actual address.
var ListenAddr string

// shutdownTimeout is the maximum amount of time the metrics plugin will wait
// before erroring when it tries to close the metrics server
const shutdownTimeout time.Duration = time.Second * 5

var buildInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: plugin.Namespace,
	Name:      "build_info",
	Help:      "A metric with a constant '1' value labeled by version, revision, and goversion from which CoreDNS was built.",
}, []string{"version", "revision", "goversion"})
