// +build gofuzz

package cache

import (
	"github.com/inverse-inc/wireguard-go/dns/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	return fuzz.Do(New(), data)
}
