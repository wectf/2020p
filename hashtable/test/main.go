package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strconv"
)

const TableSize = 10000

var TableSizeBI = big.NewInt(int64(TableSize))

const MaxCollision = 10

type LinkedList struct {
	Content       [MaxCollision]int
	InsertedCount int // count of element in linked list
}

type HashTable struct {
	Content      [TableSize]*LinkedList // array for mapping hash to the linked list
	HashParam1   *big.Int               // p1 for hashing
	HashParam2   *big.Int               // p2 for hashing
	ElementCount int                    // count of all elements in hash table
}

func (t *HashTable) hash(value int) uint {
	// hash = p1 * value % p2 % table size
	v := big.NewInt(int64(value))
	var h big.Int
	h.Exp(v, t.HashParam1, t.HashParam2)
	h.Mod(&h, TableSizeBI)
	return uint(h.Uint64())
}

func (t *HashTable) insert(value int) bool {
	var elementHash = t.hash(value)                // get hash
	var linkedListForHash = t.Content[elementHash] // get linked list for the slot
	linkedListForHash.InsertedCount++              // increase count for the linked list
	if linkedListForHash.InsertedCount > 10 {
		fmt.Println(linkedListForHash.Content)
		return true
	}
	t.ElementCount++                                                     // increase count for the hash table
	linkedListForHash.Content[linkedListForHash.InsertedCount-1] = value // insert
	return false                                                         // <10 collisions
}

func (t *HashTable) generateRandForHashing() {
	//currentTime :=time.Now().UnixNano()
	x, _ := strconv.Atoi(os.Args[1])
	rand.Seed(int64(x))
	//fmt.Println(currentTime)
	t.HashParam1 = big.NewInt(int64(rand.Intn(1 << 32))) // allow 16 bits, so 16+20 = 32 => no overflow
	t.HashParam2 = big.NewInt(int64(rand.Intn(1 << 32)))
}

func (t *HashTable) recreate() {
	t.generateRandForHashing()
	//for !t.HashParam1.ProbablyPrime(0) || !t.HashParam2.ProbablyPrime(0) {
	//	t.generateRandForHashing()
	//}
	for i := 0; i < TableSize; i++ {
		t.Content[i] = &LinkedList{[MaxCollision]int{}, 0}
	}
}

func main() {
	var t HashTable
	t.recreate()
	for i := 1 << 13; i < 1<<16; i++ {
		if t.insert(i) {
			break
		}
	}
}
