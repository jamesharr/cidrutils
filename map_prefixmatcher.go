package cidrutil

import (
	"fmt"
	"net"
	"sort"
)

// MapPrefixTableStruct - asdf
type MapPrefixTableStruct struct {
	// maskUsed indicates whether a prefix length is used
	maskUsed [128]bool

	// A reverse-sorted list of prefix lengths
	maskList []int

	// An array of prefixTables maps based on prefixLength
	// prefixTable[myPrefixLength][netAddr] -> Value
	prefixTable [128]map[ip6addr]Value
}

// MapPrefixTable creates a PrefixMatcher based on a map data structure
// TODO: When is this good to use?
func MapPrefixTable() *MapPrefixTableStruct {
	return &MapPrefixTableStruct{}
}

// Set a prefix
func (mpm *MapPrefixTableStruct) Set(prefix net.IPNet, v Value) error {
	// Convert to 16 byte values
	v6mask := castip6mask(prefix.Mask)
	v6net := castip6addr(prefix.IP)

	maskLen, size := v6mask.Size()
	if size != 128 {
		return fmt.Errorf(ErrNonCanonicalMask)
	}

	// Initialize prefixTable if it needs it
	if !mpm.maskUsed[maskLen] {
		// Mark as in-use
		mpm.maskUsed[maskLen] = true

		// Add to our search mask list
		mpm.maskList = append(mpm.maskList, maskLen)
		sort.Sort(sort.Reverse(sort.IntSlice(mpm.maskList)))

		// Initialize map
		mpm.prefixTable[maskLen] = make(map[ip6addr]Value)
	}

	// Set in table
	mpm.prefixTable[maskLen][v6net] = v

	// Return nil
	return nil
}

// Delete removes an entry
func (mpm *MapPrefixTableStruct) Delete(prefix net.IPNet) error {
	// Convert to 16 byte values
	v6mask := castip6mask(prefix.Mask)
	v6net := castip6addr(prefix.IP)

	maskLen, size := v6mask.Size()
	if size != 128 {
		return fmt.Errorf(ErrNonCanonicalMask)
	}

	delete(mpm.prefixTable[maskLen], v6net)

	return nil
}

// MatchAll returns all CIDRs matching a specific value
func (mpm *MapPrefixTableStruct) MatchAll(ip net.IP) []Value {
	// v6net, maskLen, err

	var k ip6addr
	copy(k[:], ip.To16())
	return nil
}

// MatchLPM performs a Longest Prefix Match, aka most specific prefix
func (mpm *MapPrefixTableStruct) MatchLPM(ip net.IP) Value {
	ip = ip.To16()
	for _, maskLen := range mpm.maskList {
		prefix := ip.Mask(len2mask[maskLen])
		prefixArr := castip6addr(prefix)
		v, found := mpm.prefixTable[maskLen][prefixArr]
		if found {
			return v
		}
	}
	return nil
}

// MatchSPM performs a Shortest Prefix Match, aka least specific prefix
func (mpm *MapPrefixTableStruct) MatchSPM(ip net.IP) Value {
	ip = ip.To16()
	for _, maskLen := range mpm.maskList {
		prefix := ip.Mask(len2mask[maskLen])
		prefixArr := castip6addr(prefix)
		v, found := mpm.prefixTable[maskLen][prefixArr]
		if found {
			return v
		}
	}
	return nil
}

// MatchExact returns the exact match or no match
func (mpm *MapPrefixTableStruct) MatchExact(prefix net.IPNet) (Value, error) {
	// Convert to 16 byte values
	v6mask := castip6mask(prefix.Mask)
	v6net := castip6addr(prefix.IP)

	maskLen, size := v6mask.Size()
	if size != 128 {
		return nil, fmt.Errorf(ErrNonCanonicalMask)
	}

	rv, _ := mpm.prefixTable[maskLen][v6net]
	return rv, nil
}

// Convert an IPMask to a 128 bit address
func castip6mask(mask net.IPMask) net.IPMask {
	ones, bits := mask.Size()
	if bits == 32 {
		return net.CIDRMask(ones+96, 128)
	} else if bits == 128 {
		return mask
	}
	panic("Mask size is neither 32 nor 128")
}

// Convert an IPMask to a 128 bit address
func castip6addr(ip net.IP) ip6addr {
	var rv ip6addr
	copy(rv[:], ip.To16())
	return rv
}

// Create len2mask
func makeLen2Mask() (rv [128]net.IPMask) {
	for i := range rv {
		rv[i] = net.CIDRMask(i, 128)
	}
	return
}

// Since the conversion between length and mask is going to be used often
// this is a pre-initialized table of them.
var len2mask [128]net.IPMask = makeLen2Mask()
