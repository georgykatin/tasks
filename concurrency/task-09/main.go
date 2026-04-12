// Вопросы:
// 1) Что будет если один из воркеров зависнет? Поможет ли тут закрытие канала?
// 2) Нужен ли тут контекст, если да, то какой?
// 3) Как поправить код, чтобы не бояться зависания воркеров ?
// 4) Где должен проверяться `<-ctx.Done()` ?
package main

import (
	"context"
	"fmt"
	"time"
)

//	func worker(jobs <-chan int) {
//		for job := range jobs {
//			fmt.Println("processing", job)
//			time.Sleep(2 * time.Second)
//		}
//	}

func worker(ctx context.Context, jobs <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case <-time.After(2 * time.Second):
				fmt.Println("processed", job)
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //
	defer cancel()                                                          //

	jobs := make(chan int)

	go worker(ctx, jobs)

	for i := 0; i < 5; i++ {
		jobs <- i
	}
}

// 1) Если воркер зависнет во время обработки, продюсер может зависнуть на jobs <- i,
//    потому что канал небуферизованный и следующий jobs <- i требует живого получателя.
//    close(jobs) поможет, потому что зависание в воркере
// 2) Да, context.WithTimeout
// 3) Сделал выше
// 4) <-ctx.Done() надо проверять везде где код может зависнуть, тоесть
//    ожидание job из канала и ожидание окончания обработки
