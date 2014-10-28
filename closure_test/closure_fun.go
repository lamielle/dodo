// Compare go and js closures

package main

import "fmt"
import "sync"

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		local := i
		go func() {
			fmt.Println("i=", i)
			fmt.Println("local=", local)
			wg.Done()
		}()
	}
	wg.Wait()
}
