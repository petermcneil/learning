package hashtable

import "log"

//https://www.youtube.com/watch?v=S5NY1fqisSY
//https://samwho.dev/hashing/

// Hashtable with no deletions allowed. Duplicates overwrite the existing value. Values are of
// type V and keys are strings -- one extension is to adapt this class to use other types as keys.
//
// The underlying data is stored in the array `arr', and the actual values stored are pairs of
// (key, value). This is so that we can detect collisions in the hash function and look for the next
// location when necessary.
type Hashtable[K string, V any] struct {
	arr       []pair[K, V] //an array of pair objects, where each pair contains the key and value stored in the hashtable
	max       int          //the size of arr. This should be a prime number
	itemCount int          //the number of items stored in arr
	maxLoad   float64      //the maximum load factor
	probeType PROBE_TYPE   //the type of probe to use when dealing with collisions
}

type pair[K string, V any] struct {
	Key   K
	Value V
}

type PROBE_TYPE int

const (
	LINEAR_PROBE PROBE_TYPE = iota
	QUADRATIC_PROBE
	DOUBLE_HASH
)

func New[V any](capacity int) Hashtable[string, V] {
	return Hashtable[string, V]{
		arr:     make([]pair[string, V], capacity),
		max:     capacity,
		maxLoad: 0.6,
	}
}

// Get the value associated with key, or return null if key does not exists. Use the find method to search the
// array, starting at the hashed value of the key, stepNum of zero and the original key.
// @param key
// @return
func (table Hashtable[K, V]) Get(key K) (V, bool) {
	panic("Method not implemented")
}

// Put - Store the value against the given key. If the loadFactor exceeds maxLoad, call the resize
// method to resize the array. the If key already exists then its value should be overwritten.
// Create a new pair item containing the key and value, then use the findEmpty method to find an unoccupied
// position in the array to store the pair. Call findEmpty with the hashed value of the key as the starting
// position for the search, stepNum of zero and the original key.
// containing
// @param key
// @param value
func (table Hashtable[K, V]) Put(key K, value V) {
	panic("Method not implemented")
}

func (table Hashtable[K, V]) HasKey(key K) bool {
	panic("Method not implemented")
}

func (table Hashtable[K, V]) Capacity() int {
	panic("Method not implemented")
}

func (table Hashtable[K, V]) LoadFactor() float64 {
	panic("Method not implemented")
}

// find the value stored for this key, starting the search at position startPos in the array. If
// the item at position startPos is null, the Hashtable does not contain the value, so return null.
// If the key stored in the pair at position startPos matches the key we're looking for, return the associated
// value. If the key stored in the pair at position startPos does not match the key we're looking for, this
// is a hash collision so use the getNextLocation method with an incremented value of stepNum to find
// the next location to search (the way that this is calculated will differ depending on the probe type
// being used). Then use the value of the next location in a recursive call to find.
// @param startPos
// @param key
// @param stepNum
// @return
func (table Hashtable[K, V]) find(startPos int, key K, stepNum int) V {
	panic("Method not implemented")
}

// findEmpty - Find the first unoccupied location where a value associated with key can be stored, starting the
// search at position startPos. If startPos is unoccupied, return startPos. Otherwise use the getNextLocation
// method with an incremented value of stepNum to find the appropriate next position to check
// (which will differ depending on the probe type being used) and use this in a recursive call to findEmpty.
// @param startPos
// @param stepNum
// @param key
// @return
func (table Hashtable[K, V]) findEmpty(startPos int, key K, stepNum int) int {
	panic("Method not implemented")
}

func (table Hashtable[K, V]) getNextLocation(startPos int, stepNum int, key K) int {
	step := startPos
	switch table.probeType {
	case LINEAR_PROBE:
		step++
	default:
		log.Fatalf("Ask Pete - not implemented getNextLocation for %d", table.probeType)
	}
	return step % table.max
}

// Return an int value calculated by hashing the key. See the lecture slides for information
// on creating hash functions. The return value should be less than max, the maximum capacity
// of the array
// @param key
// @return
func (table Hashtable[K, V]) hash(key K) bool {
	panic("Method not implemented")
}

func (table Hashtable[K, V]) isPrime() {
	panic("Method not implemented")
}

func (table Hashtable[K, V]) nextPrime(n int) int {
	panic("Method not implemented")
}

// resize the hashtable, to be used when the load factor exceeds maxLoad. The new size of
// the underlying array should be the smallest prime number which is at least twice the size
// of the old array.
func (table Hashtable[K, V]) resize() {
	panic("Method not implemented")
}
