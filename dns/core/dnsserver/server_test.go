package dnsserver

import (
	"context"
	"testing"

	"github.com/inverse-inc/wireguard-go/dns/plugin"
	"github.com/inverse-inc/wireguard-go/dns/plugin/pkg/log"
	"github.com/inverse-inc/wireguard-go/dns/plugin/test"

	"github.com/miekg/dns"
)

type testPlugin struct{}

func (tp testPlugin) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	return 0, nil
}

func (tp testPlugin) Name() string { return "testplugin" }

func testConfig(transport string, p plugin.Handler) *Config {
	c := &Config{
		Zone:        "example.com.",
		Transport:   transport,
		ListenHosts: []string{"127.0.0.1"},
		Port:        "53",
		Debug:       false,
	}

	c.AddPlugin(func(next plugin.Handler) plugin.Handler { return p })
	return c
}

func TestNewServer(t *testing.T) {
	_, err := NewServer("127.0.0.1:53", []*Config{testConfig("dns", testPlugin{})})
	if err != nil {
		t.Errorf("Expected no error for NewServer, got %s", err)
	}

	_, err = NewServergRPC("127.0.0.1:53", []*Config{testConfig("grpc", testPlugin{})})
	if err != nil {
		t.Errorf("Expected no error for NewServergRPC, got %s", err)
	}

	_, err = NewServerTLS("127.0.0.1:53", []*Config{testConfig("tls", testPlugin{})})
	if err != nil {
		t.Errorf("Expected no error for NewServerTLS, got %s", err)
	}
}

func TestDebug(t *testing.T) {
	configNoDebug, configDebug := testConfig("dns", testPlugin{}), testConfig("dns", testPlugin{})
	configDebug.Debug = true

	s1, err := NewServer("127.0.0.1:53", []*Config{configDebug, configNoDebug})
	if err != nil {
		t.Errorf("Expected no error for NewServer, got %s", err)
	}
	if !s1.debug {
		t.Errorf("Expected debug mode enabled for server s1")
	}
	if !log.D.Value() {
		t.Errorf("Expected debug logging enabled")
	}

	s2, err := NewServer("127.0.0.1:53", []*Config{configNoDebug})
	if err != nil {
		t.Errorf("Expected no error for NewServer, got %s", err)
	}
	if s2.debug {
		t.Errorf("Expected debug mode disabled for server s2")
	}
	if log.D.Value() {
		t.Errorf("Expected debug logging disabled")
	}
}

func BenchmarkCoreServeDNS(b *testing.B) {
	s, err := NewServer("127.0.0.1:53", []*Config{testConfig("dns", testPlugin{})})
	if err != nil {
		b.Errorf("Expected no error for NewServer, got %s", err)
	}

	ctx := context.TODO()
	w := &test.ResponseWriter{}
	m := new(dns.Msg)
	m.SetQuestion("aaa.example.com.", dns.TypeTXT)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.ServeDNS(ctx, w, m)
	}
}
