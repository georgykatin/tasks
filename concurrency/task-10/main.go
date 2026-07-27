// Вопросы:
// 1) Почему тут дедлок?
// 2) Как поправить чтобы дедлок не было?
// 3) Нужен ли `defer` в горутине?

// 1) заполняется буфер канала, нет читателя
// 2) добавить чтение из канала
// 3) нужен, на случай если паника или ошибка освободить место в канале
package main

import (
	"fmt"
	"sync"
)

func main() {
	sem := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		sem <- struct{}{}
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()

			fmt.Println("job", i)
		}(i)
	}
	wg.Wait()
}
