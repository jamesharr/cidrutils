package cidrutil

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func cidr(input string) net.IPNet {
	_, cidrPtr, _ := net.ParseCIDR(input)
	var cidr net.IPNet = *cidrPtr
	return cidr
}
func ip(input string) net.IP {
	return net.ParseIP(input)
}

// func TestLPM_trie(t *testing.T) {
// 	testLPM(t, TriePrefixTable())
// }

func TestLPM_mpt(t *testing.T) {
	testLPM(t, MapPrefixTable())
}

func testLPM(t *testing.T, mpm PrefixTable) {
	// For Values stored in the PrefixTable we're using sentinel integers to ensure we get the correct field back

	// Check to find non-specific matches
	mpm.Set(cidr("1.2.3.0/24"), 1)
	mpm.Set(cidr("2.3.0.0/16"), 2)
	assert.Equal(t, 1, mpm.MatchLPM(ip("1.2.3.7")))
	assert.Equal(t, 2, mpm.MatchLPM(ip("2.3.255.255")))
	assert.Nil(t, mpm.MatchLPM(ip("3.3.3.3")))

	mpm.Set(cidr("0.0.0.0/0"), 3)
	assert.Equal(t, 3, mpm.MatchLPM(ip("4.4.4.4")))
	assert.Equal(t, 1, mpm.MatchLPM(ip("1.2.3.8")))

	// Check a v6
	assert.Nil(t, mpm.MatchLPM(ip("::1")), "A v4 default route should not match a v6 route")

	// Delete a route
	mpm.Delete(cidr("0.0.0.0/0"))
	assert.Nil(t, mpm.MatchLPM(ip("3.3.3.3")))

	// Add a subsequent v6 default and see if a v4 matches (it should)
	mpm.Set(cidr("::/0"), 4)
	assert.Equal(t, 4, mpm.MatchLPM(ip("::1")))

	// Delete a CIDR that doesn't exist (this should not error out)
	mpm.Delete(cidr("6.6.6.6/17"))

	// Add a specific route
	mpm.Set(cidr("1.2.3.4/32"), 5)
	assert.Equal(t, 5, mpm.MatchLPM(ip("1.2.3.4")))
}
