package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
	// worker "github.com/tuancamtbtx/learning-go/app/worker/drop_table"
)

func Task() {
	fmt.Println("ahihi do cho ne")
}

func main() {

	s := gocron.NewScheduler()
	s.Every(1).Day().At("10:25").Do(Task)
	<-s.Start()
}
