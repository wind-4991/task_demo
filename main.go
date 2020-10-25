package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Set struct {
	timer time.Time
	f func()
}

var wg sync.WaitGroup

func main() {
	s1 := &Set{time.Now().Add(time.Second * 10), func() {
		fmt.Println("i ma s1, current time :", time.Now().Add(time.Second * 10))
	}}

	s2 := &Set{time.Now().Add(time.Second * 20), func() {
		fmt.Println("i ma s2, current time :", time.Now().Add(time.Second * 20))
	}}

	s3 := &Set{time.Now().Add(time.Second * 30), func() {
		fmt.Println("i ma s3, current time :", time.Now().Add(time.Second * 30))
	}}

	ruleList := []*Set{s1, s2, s3}
	wg.Add(1)
	go task(ruleList)
	wg.Wait()
}


// task
func task(taskList []*Set) {
	completeNum := len(taskList)
	ticker := time.NewTicker(time.Second)

	for range ticker.C{
		if (completeNum == 0) {
			fmt.Println("任务执行完毕")
			ticker.Stop()
			os.Exit(0)
		}

		for _, v := range taskList {
			if v.timer.Unix() == time.Now().Unix() {
				v.f()
				completeNum --
			}
		}
	}

	wg.Done()
}