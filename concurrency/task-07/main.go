// Вопросы:
// 1) До запуска кода можно сразу сказать где ошибка? 
// 2) Почему так нельзя делать? Что может случиться?

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			fmt.Println("working")
		}()
	}

	wg.Wait()
}

// 1) Горутина добавляется в WaitGroup внутри себя же, а должна снаружи
// 2) Программа может пробежать wg.Wait до того как горутины добявятся в wg
