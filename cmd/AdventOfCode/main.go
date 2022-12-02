package main

import (
	"AdventOfCode/pkg/challenge/y2022"
	"AdventOfCode/pkg/concurrency/worker"
	"go.uber.org/zap"
	"log"
)

func main() {
	config := zap.NewProductionConfig()
	config.Sampling = nil
	logProvider, err := config.Build()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	sugar := logProvider.Sugar()
	defer sugar.Sync()

	sugar.Info("Welcome to Advent of Code!")

	pool := worker.NewPool(1, sugar)

	go func() {

		d1 := y2022.Day1{
			Year:      2022,
			Day:       1,
			InputFile: "./assets/2022/day1/input.txt",
			Logger:    sugar,
		}
		job := worker.NewJob(1, d1.Solve)
		pool.SubmitJob(job)

		d2 := y2022.Day2{
			Year:      2022,
			Day:       2,
			InputFile: "./assets/2022/day2/input.txt",
			Logger:    sugar,
		}
		job = worker.NewJob(2, d2.Solve)
		pool.SubmitJob(job)

		pool.Stop()
	}()

	pool.Start()
}
