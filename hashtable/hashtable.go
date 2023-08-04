package hashtable

import (
	"log"
)

//https://www.youtube.com/watch?v=S5NY1fqisSY
//https://samwho.dev/hashing/

// Hashtable with no deletions allowed. Duplicates overwrite the existing value. Values are of
// type V and keys are strings -- one extension is to adapt this class to use other types as keys.
//
// The underlying data is stored in the array `arr', and the actual values stored are pairs of
// (key, value). This is so that we can detect collisions in the hash function and look for the next
// location when necessary.
type Hashtable[K string, V any] struct {
	arr       []pair[K, V] //an array of pair objects, where each pair contains the key and value stored in the hashtable.
	max       int          //the size of arr. This should be a prime number
	itemCount int          //the number of items stored in arr
	maxLoad   float64      //the maximum load factor
	probeType PROBE_TYPE   //the type of probe to use when dealing with collisions
	seive     map[int]bool
}

const hugeNumber = 10e9

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
func (table Hashtable[K, V]) Get(key K) (val V, ok bool, collisionCount int) {
	hash := table.hash(key)
	if table.arr[hash].Key == "" {
		return val, false, collisionCount
	}
	val, ok, collisionCount = table.find(hash, key, 0)
	if !ok {
		return val, false, collisionCount
	}
	return val, true, collisionCount
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
	if table.maxLoad < float64(table.itemCount)/float64(table.max) {
		table.resize()
	}
	hash := table.hash(key)
	if table.arr[hash].Key == key {
		table.arr[hash].Value = value
		return
	}
	pear := pair[K, V]{key, value}
	idx := table.findEmpty(table.hash(pear.Key), pear.Key, 0)
	table.arr[idx] = pear
}

func (table Hashtable[K, V]) HasKey(key K) bool {
	hash := table.hash(key)
	if table.arr[hash].Key == "" {
		return false
	}
	return true
}

func (table Hashtable[K, V]) Capacity() int {
	return table.max
}

func (table Hashtable[K, V]) LoadFactor() float64 {
	return table.maxLoad
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
func (table Hashtable[K, V]) find(startPos int, key K, stepNum int) (val V, found bool, collisionCount int) {
	if table.arr[startPos].Key == "" {
		return val, false, stepNum
	}
	switch table.arr[startPos].Key {
	case key:
		val, found = table.arr[startPos].Value, true
		return val, found, stepNum
	default:
		stepNum++
		nextLocation := table.getNextLocation(startPos, stepNum, key)
		return table.find(nextLocation, key, stepNum)
	}
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
	if table.arr[startPos].Key == "" {
		return startPos
	}
	stepNum++
	nextLocation := table.getNextLocation(startPos, stepNum, key)
	return table.findEmpty(nextLocation, key, stepNum)

}

func (table Hashtable[K, V]) getNextLocation(startPos int, stepNum int, key K) int {
	step := startPos
	switch table.probeType {
	case LINEAR_PROBE:
		step++
	case DOUBLE_HASH:
		step = table.doubleHash(key)
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
func (table Hashtable[K, V]) hash(key K) int {
	// Approach 1: 4940 collisions
	//return len(key) % table.Capacity()

	// Approach 2: 998 collisions

	//var hashVal uint8 = 0
	//var pow27 uint8 = 1
	//for i := len(key) - 1; i >= 0; i-- {
	//	c := key[i] - 96
	//	hashVal += pow27 * c
	//	pow27 *= 27
	//}

	// Approach 3 (Horner's method with maths and stuff for speed): 998 collisions

	//var hashVal = key[0] - 96
	//for i := 1; i < len(key); i++ {
	//	c := key[i] - 96
	//	hashVal = hashVal*27 + c
	//}

	//Approach 4 (Horner's method moving hashVal inside loop to stop it getting huge): 834 collisions

	//var hashVal = key[0] - 96
	//for i := 1; i < len(key); i++ {
	//	c := key[i] - 96
	//	hashVal = (hashVal*27 + c) % uint8(table.Capacity())
	//}

	return len(key) * hugeNumber % table.Capacity()

}
func (table Hashtable[K, V]) doubleHash(key K) int {
	return 27 - (table.hash(key) % 27)
}

func (table Hashtable[K, V]) isPrime(n int) bool {
	// noob implementation:
	// for i := 2; i < n; i++ {
	//	if n%i == 0 {
	//		return false
	//	}
	//	return true
	//}

	// less noobie still bad. filtering evens
	// for i := 2; i < n/2; i++ {
	//	if n%i == 0 {
	//		return false
	//	}
	//	return true
	//}

	//  filtering sq roots
	// for i := 2; i*i < n; i++ {
	//	if n%i == 0 {
	//		return false
	//	}
	//	return true
	//}

	// better to have sieve that holds primes in memory
	if n > len(table.seive) {
		panic("array too short")
	}
	return table.seive[n]
}

func (table Hashtable[K, V]) nextPrime(n int) bool {
	if n > len(table.seive) {
		panic("array too short")
	}
	return table.seive[n+1]
}

func (table Hashtable[K, V]) SetProbe(probeType PROBE_TYPE) {
	table.probeType = probeType
}

// resize the hashtable, to be used when the load factor exceeds maxLoad. The new size of
// the underlying array should be the smallest prime number which is at least twice the size
// of the old array.
func (table Hashtable[K, V]) resize() {
	table.max = table.max * 2
	table.arr = make([]pair[K, V], table.max)
}

func (table Hashtable[K, V]) GetBiggestPrime(limit int) int {
	for i := len(table.seive) - 1; i > 0; i-- {
		if table.seive[i] {
			return i
		}
	}
	return 0
}

// MakeSeive creates a seive of Eratosthenes
func (table *Hashtable[K, V]) MakeSeive(n int) {
	primes := make(map[int]bool)
	for i := 0; i < n; i++ {
		primes[i] = true
	}
	primes[0], primes[1] = false, false // 2 is smallest prime
	for i := 2; i < len(primes); i++ {
		// if i is prime its multiples are not
		if primes[i] {
			for j := 2; i*j < len(primes); j++ {
				primes[i*j] = false // set all multiples of all primes to false
			}
		}
	}
	table.seive = primes
}

func (table *Hashtable[K, V]) SetCapacity(n int) {
	table.max = table.GetBiggestPrime(n)
}
