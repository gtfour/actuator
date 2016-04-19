package main

import "fmt"
import "sync"
import "runtime"

type T int

func main() {

    var slice []T
    var wg sync.WaitGroup

    queue := make(chan T, 1)
    wg.Add(100)
    for i:=0; i<100 ; i++ {




    }



}
/*
package main

import "fmt"
import "sync"
import "runtime"

type T int

func main() {
	var slice []T
	var wg sync.WaitGroup

	queue := make(chan T, 1)

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()

			// Do stuff
			runtime.Gosched()

			queue <- T(i)
		}(i)
	}

	go func() {
		defer wg.Done()
		for t := range queue {
			slice = append(slice, t)
		}
	}()

	wg.Wait()
	fmt.Println(slice)
}


*/
