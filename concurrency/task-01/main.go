// Вопросы:
// 1) Почему результат почти всегда не 1000?
// 2) Что покажет запуск с -race флагом?
// 3) Как можно исправить код, чтобы счетчик был 1000? Здесь одно решение или несколько? Какое оптимальное?

package main

import (
	"fmt"
)

func main() {
	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}

	fmt.Println("counter:", counter)
}

// 1) Из-за race condition. counter не потокобезопасен
// 2) WARNING: DATA RACE
// 3) Либо использовать mutex, либо atomic. Для счетчика лучше использовать atomic.
