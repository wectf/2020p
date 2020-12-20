package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haisum/recaptcha"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// P(success by bruteforce) = 1 - e^(blablabla) << 0.001 by birthday paradox
// hope it's correct math...
// anyways, it is impossible to bruteforce :)
const TableSize = 10000

var TableSizeBI = big.NewInt(int64(TableSize))

const MaxCollision = 10

// fake linked list...
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

type SystemTimeInfo struct {
	HashTableInitTime int64 // hash table recreate/init time
	StartTime         int64 // system start time
}

var hashTable HashTable
var systemTimeInfo SystemTimeInfo
var RecaptchaVerifier recaptcha.R
var FLAG string

// initializations for the hashtable
func (t *HashTable) generateRandForHashing() {
	currentTime := time.Now().UnixNano()
	rand.Seed(currentTime)
	systemTimeInfo.HashTableInitTime = currentTime
	t.HashParam1 = big.NewInt(int64(rand.Intn(1 << 32))) // allow 16 bits, so 16+20 = 32 => no overflow
	t.HashParam2 = big.NewInt(int64(rand.Intn(1 << 32)))
}

func (t *HashTable) recreate() {
	t.generateRandForHashing()
	for !t.HashParam1.ProbablyPrime(0) || !t.HashParam2.ProbablyPrime(0) {
		t.generateRandForHashing()
	}
	for i := 0; i < TableSize; i++ {
		t.Content[i] = &LinkedList{[MaxCollision]int{}, 0}
	}
	t.ElementCount = 0
}

// input: value => a value to be hashed
// output: hash of the value
func (t *HashTable) hash(value int) uint {
	// hash = p1 ** value % p2 % table size
	v := big.NewInt(int64(value))
	var h big.Int
	h.Exp(v, t.HashParam1, t.HashParam2)
	h.Mod(&h, TableSizeBI)
	return uint(h.Uint64())
}

// input: value => a value to be inserted
// output: whether max collision happens (i.e. whether there are
// 		   10 elements in a single slot)
// This is an O(1) operation
func (t *HashTable) insert(value int) bool {
	if t.find(value) {
		return false // already in the table
	}
	var elementHash = t.hash(value)                // get hash
	var linkedListForHash = t.Content[elementHash] // get linked list for the slot
	linkedListForHash.InsertedCount++              // increase count for the linked list
	if linkedListForHash.InsertedCount >= 10 {
		return true // 10 collisions
	}
	t.ElementCount++                                                     // increase count for the hash table
	linkedListForHash.Content[linkedListForHash.InsertedCount-1] = value // insert
	return false                                                         // <10 collisions
}

// input: value => a value to be found
// output: whether this value is inside hash table
// This is an O(1) operation
func (t *HashTable) find(value int) bool {
	var elementHash = t.hash(value)                // get hash
	var linkedListForHash = t.Content[elementHash] // get linked list for the slot
	for i := 0; i < linkedListForHash.InsertedCount; i++ {
		if linkedListForHash.Content[i] == value {
			return true // found
		}
	}
	return false // not found
}

// helper function for dealing with http server...
// output: the value in the GET request
// 		   if the value is not int or is not valid (1048576 > value > 4096), return -1
func processValue(context *gin.Context) int {
	value, isInt := strconv.Atoi(context.DefaultQuery("value", "0"))
	// check whether the value is integer
	if isInt != nil {
		context.String(http.StatusInternalServerError, "Not an int")
		return -1
	}
	// check (1048576 > value > 4096)
	if value < (1<<12) || value > (1<<20) {
		context.String(http.StatusInternalServerError, "Value has to be < 1048576 && > 4096 :(")
		return -1
	}
	return value
}

func main() {
	// initialize global variables
	RecaptchaVerifier.Secret = os.Getenv("RECAPTCHA_SECRET") // recaptcha
	FLAG = os.Getenv("FLAG")                                 // flag
	systemTimeInfo.StartTime = time.Now().UnixNano()         // system start time
	hashTable.recreate()                                     // initialize hash table
	router := gin.Default()                                  // initialize web server
	router.LoadHTMLGlob("templates/*")                       // load template for web server
	// / => index
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"sys_time":   systemTimeInfo.StartTime,
			"table_time": systemTimeInfo.HashTableInitTime,
			"total_elem": hashTable.ElementCount,
		})
	})
	// /insert?value=[value]
	router.GET("/insert", func(context *gin.Context) {
		// check recaptcha; for testing pls remove following 5 lines
		isRecaptchaCorrect := RecaptchaVerifier.Verify(*context.Request)
		if !isRecaptchaCorrect {
			context.String(http.StatusForbidden, "Incorrect Recaptcha")
			return
		}
		value := processValue(context)
		if value == -1 {
			return
		}
		if hashTable.insert(value) {
			// if there are 10 collisions, give the flag
			hashTable.recreate()
			context.String(http.StatusOK, FLAG)
			return
		}
		context.String(http.StatusOK, "Ok")
	})
	// /find?value=[value]
	router.GET("/find", func(context *gin.Context) {
		value := processValue(context)
		if value == -1 {
			return
		}
		if hashTable.find(value) {
			context.String(http.StatusOK, "In table")
		} else {
			context.String(http.StatusNotFound, "Not in table")
		}
	})
	_ = router.Run(":8080")
}
