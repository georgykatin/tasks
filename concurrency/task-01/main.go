// Вопросы:
// 1) Почему результат почти всегда не 1000?
// 2) Что покажет запуск с -race флагом?
// 3) Как можно исправить код, чтобы счетчик был 1000? Здесь одно решение или несколько? Какое оптимальное?

// Ответы:
// 1) из-за даты рейса, простой инкремент не атомарная операция
// 2) выдаст ошибку дата рейс
// 3) использовать мутекс или атомики. в данном случае атомики быстрее

package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	// var counter int64
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//atomic.AddInt64(&counter, 1)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter)
}
