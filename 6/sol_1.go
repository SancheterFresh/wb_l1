package main

//остановка полученным сигналом

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			//case для остановки функции
			case <-stop:
				return
			}
		}
	}()

	//отправка остановки
	stop <- true
	close(stop)
}
