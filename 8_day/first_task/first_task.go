package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		time.Sleep(30 * time.Millisecond) // Имитируем работу
		results <- job * 2
	}
}

func main() {
	const numJobs = 20
	const numWorkers = 3

	// 1. Создай буферизированный канал jobs емкостью numJobs
	// 2. Создай буферизированный канал results емкостью numJobs
	// ТВОЙ КОД ЗДЕСЬ

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// 3. Запусти в цикле ровно 3 воркера (numWorkers) через ключевое слово go.
	// Не забывай перед запуском каждого делать wg.Add(1) и передавать &wg по указателю!
	// ТВОЙ КОД ЗДЕСЬ
	for id := range numWorkers {
		wg.Add(1)
		go worker(id+1, jobs, results, &wg)
	}
	// 4. В отдельном цикле забрось в канал jobs числа от 1 до numJobs.
	// Сразу после этого ОБЯЗАТЕЛЬНО закрой канал jobs через close(jobs),
	// чтобы воркеры поняли, что задачи закончились, и вышли из своего цикла range!
	// ТВОЙ КОД ЗДЕСЬ
	for job := 1; job <= numJobs; job++ {
		jobs <- job
	}
	close(jobs)

	// 5. Запусти отдельную горутину, которая будет ждать завершения всех воркеров
	// через wg.Wait() и ПОСЛЕ этого закроет канал results через close(results).
	// (Загадка на подумать: почему wg.Wait() и close(results) нельзя написать просто в основном потоке main?)
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(results)
	}()
	// wait нельзя в main чтобы не блокировать весь код, а close(results) потмоу что канал может закрыться раньше чем все воркеры отработают

	// 6. Читай результаты из канала results через цикл for range и выводи их на экран.
	// Как только канал results закроется, этот цикл завершится, и вся программа финиширует.
	for res := range results {
		fmt.Println("Получен результат:", res)
	}
}
