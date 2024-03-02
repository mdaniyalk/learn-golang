package main 
import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
var results = []string{}
var results2 = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}
	fmt.Println("The time taken is: ", time.Since(t0))

	fmt.Println("with wait group")

	t1 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall2(i)
	}
	wg.Wait()
	fmt.Println("The time taken is: ", time.Since(t1))

	fmt.Println("returning output")

	fmt.Println("without mutex")

	t2 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall3(i)
	}
	wg.Wait()
	fmt.Println("The time taken is: ", time.Since(t2))
	fmt.Println("the result are ", results)

	fmt.Println("with mutex")

	t3 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall4(i)
	}
	wg.Wait()
	fmt.Println("The time taken is: ", time.Since(t3))
	fmt.Println("the result are ", results2)



}


func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the db is: ", dbData[i])
}

func dbCall2(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the db is: ", dbData[i])
	wg.Done()
}

func dbCall3(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the db is: ", dbData[i])
	results = append(results, dbData[i])
	wg.Done()
}

func dbCall4(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the db is: ", dbData[i])
	m.Lock()
	results2 = append(results2, dbData[i])
	m.Unlock()
	wg.Done()
}