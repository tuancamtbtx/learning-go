package main

import (
	"fmt"
	"time"
	// worker "github.com/tuancamtbtx/learning-go/app/worker/drop_table"
)

func Task() {
	fmt.Println("ahihi do cho ne")
}

func main() {

	// s := gocron.NewScheduler()
	// s.Every(1).Day().At("10:25").Do(Task)
	// <-s.Start()
	current_time := time.Now()
	fmt.Println("Running Time: ", current_time.Format("20060102"))
}
