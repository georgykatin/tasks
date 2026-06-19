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

func worker(ctx context.Context, jobs <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if ok {
				fmt.Println("processing", job)
				time.Sleep(2 * time.Second)
			} else {
				return
			}
		}
	}
}

func main() {
	jobs := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker(ctx, jobs)

	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

}

// 1) Закрытие не поможет.
// 2) Тут подойдет контекст с отменой. Чтобы в случае чего его можно было отменить.
// 3) Как поправить код, чтобы не бояться зависания воркеров ?
// 4) Где должен проверяться `<-ctx.Done()` ?
