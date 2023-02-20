package main

//остановка полученным сигналом

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			//case для работы программы
			case <-stop:
				return
			}
		}
	}()

	//отправка остановки
	stop <- true
	close(stop)
}
