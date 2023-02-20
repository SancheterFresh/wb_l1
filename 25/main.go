package main

import (
	"time"
)

func sleep(s time.Duration) {
	<-time.After(s)
}

func main() {
	seconds := 10 * time.Second
	sleep(seconds)
}
