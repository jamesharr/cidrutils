package cidrutil

import (
	"net"
)

// Value kept in the PrefixTable
type Value interface{}

// PrefixTable TODO docs. This API is unstable at the moment.
type PrefixTable interface {
	Set(net net.IPNet, v Value) error
	Delete(net net.IPNet) error
	MatchAll(ip net.IP) []Value
	MatchLPM(ip net.IP) Value
	MatchSPM(ip net.IP) Value
	MatchExact(net net.IPNet) (Value, error)
}

// Used internally as a fixed-size array (instead of a slice)
// that can be used as a map key. In many structures, addresses
// are normalized to IPv6 addresses for simplicity.
type ip6addr [16]byte

// ErrNonCanonicalMask is given back from errors when a mask is non-canonical
const ErrNonCanonicalMask = "Mask is non-canonical"
