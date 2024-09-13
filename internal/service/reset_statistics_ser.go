package service

func (s service) ResetStatistics() int8 {
	ch := make(chan int8)

	go func(ch chan int8) {
		ch <- s.mongodb.ResetStatistics()
	}(ch)

	if <-ch == 1 {
		return 1
	} else {
		return 0
	}
}
