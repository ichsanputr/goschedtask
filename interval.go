package goschedtask

import "time"

func Second(i int) time.Duration {
	return time.Second * time.Duration(i)
}

func Minute(i int) time.Duration {
	return time.Minute * time.Duration(i)
}

func Hour(i int) time.Duration {
	return time.Hour * time.Duration(i)
}
