package main

import (
	"fmt"
	"time"
	"sync"
)

/* ------------------ Poiter in go ------------------*/
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var results = []string{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	/*
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("Pointer p: %v\n", *p)
	fmt.Printf("No pointer i: %v\n", i)

	p = &i
	*p = 1

	fmt.Printf("Pointer p: %v\n", *p)
	fmt.Printf("No pointer i: %v\n", i)

	var k int32= 2
	i = k
	fmt.Printf("No pointer i: %v\n", i)
*/
/*
	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	// ? [1 2 4] ?
	*/

	/*
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memoryy location of the thing1 array is: %p", &thing1)
	// var result [5]float64 = square(thing1)
	var result [5]float64 = square(&thing1)

	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
*/
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nTotal execution time: %v", results)

}

/* ------------------ Goroutines ------------------*/
func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()

}
func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memoryy location of the thing2 array is: %p", &thing2)

	for i := range thing2{
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

