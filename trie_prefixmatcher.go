package cidrutil

// TODO - Try implementing a trie.
// I started working on this, but quickly determined that it'd be better to
// finish the map_prefixmatcher.go implementation

// type trieStruct struct {

// 	// Matches at this node, indexed by how many bits were used to match this node
// 	matchMaster [8]Value

// 	// Matches at this node, non-matches removed
// 	matchCompress []Value

// 	// Children of this node
// 	children [256]*trieStruct
// }

// //TriePrefixTable creates a 256-ary trie
// func TriePrefixTable() PrefixTable {
// 	rv := &trieStruct{}
// 	return rv
// }

// // Set a prefix
// func (pt *trieStruct) Set(prefix net.IPNet, v Value) error {
// 	ip := prefix.IP.To16()
// 	maskLen := calcMaskLen(prefix.Mask)
// 	pt.set(ip, maskLen, v)
// 	return nil
// }

// // Recursive function for Set()
// func (pt *trieStruct) set(ip []byte, maskLen int, v Value) {
// 	if maskLen < 8 {
// 		pt.matchMaster[maskLen] = v
// 		// TODO - rebuild matchCompress
// 	} else {

// 	}
// 	currBit := ip[0]
// }

// // Delete a prefix
// func (pt *trieStruct) Delete(net net.IPNet) error {
// 	panic("not implemented") // TODO: Implement
// 	return nil
// }

// // MatchLPM
// func (pt *trieStruct) MatchLPM(ip net.IP) Value {
// 	panic("not implemented") // TODO: Implement
// 	return nil
// }

// // Matchall finds all matches, ordered from least-specific to most-specific
// func (pt *trieStruct) MatchAll(ip net.IP) []Value {
// 	panic("not implemented") // TODO: Implement
// 	return nil
// }

// // Find the shortest prefix match possible
// func (pt *trieStruct) MatchSPM(ip net.IP) Value {
// 	panic("not implemented") // TODO: Implement
// 	return nil
// }

// // MatchExact blah
// func (pt *trieStruct) MatchExact(net net.IPNet) (Value, error) {
// 	panic("not implemented") // TODO: Implement
// 	return nil, nil
// }

// // How many IPs will be in a match if we have N bits of the mask left
// var trieSliceSize = [8]int{255, 127, 63, 31, 15, 7, 3, 1}
