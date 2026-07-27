//Вопросы:
// 1) Что будет если один из воркеров зависнет? Поможет ли тут закрытие канала?
// 2) Нужен ли тут контекст, если да, то какой?
// 3) Как поправить код, чтобы не бояться зависания воркеров ?
// 4) Где должен проверяться `<-ctx.Done()` ?

// 1) горутина зависнет, значит будет утечка памяти. закрытие не поможет потому что воркер зависнет уже в блоке for range
// 2) нужен. WithCancel для ручной отмены WithTimeout/WithDeadline для отмены по времени
// 3) добавить селект с проверкой завершения контекста
// 4) должен проверяться в момент отправки джобы, внутри воркера в момент взятия джобы и в момент обработки джобы (на случай если долгая операция)
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, jobs <-chan int) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case <-time.After(2 * time.Second):
				fmt.Printf("job %v one", job)
			}
		case <-ctx.Done():
			return
		}
	}
	//for job := range jobs {
	//	fmt.Println("processing", job)
	//	time.Sleep(2 * time.Second)
	//}
}

func main() {
	jobs := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < 5; i++ {
		select {
		case jobs <- i:
			go worker(ctx, jobs)
		case <-ctx.Done():
			break
		}

		//jobs <- i
	}

}
