package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/seanr9191/AdventOfCode/pkg/challenge/y2022"
	"github.com/seanr9191/AdventOfCode/pkg/concurrency/worker"
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

		d3 := y2022.Day3{
			Year:      2022,
			Day:       3,
			InputFile: "./assets/2022/day3/input.txt",
			Logger:    sugar,
		}
		job = worker.NewJob(3, d3.Solve)
		pool.SubmitJob(job)

		d4 := y2022.Day4{
			Year:      2022,
			Day:       4,
			InputFile: "./assets/2022/day4/input.txt",
			Logger:    sugar,
		}
		job = worker.NewJob(4, d4.Solve)
		pool.SubmitJob(job)

		d5 := y2022.Day5{
			Year:      2022,
			Day:       5,
			InputFile: "./assets/2022/day5/input.txt",
			Logger:    sugar,
		}
		job = worker.NewJob(5, d5.Solve)
		pool.SubmitJob(job)

		d6 := y2022.Day6{
			Year:      2022,
			Day:       6,
			InputFile: "./assets/2022/day6/input.txt",
			Logger:    sugar,
		}
		job = worker.NewJob(6, d6.Solve)
		pool.SubmitJob(job)

		d7 := y2022.Day7{
			Year:      2022,
			Day:       7,
			InputFile: "./assets/2022/day7/input.txt",
			Logger:    sugar,
		}
		job = worker.NewJob(7, d7.Solve)
		pool.SubmitJob(job)

		d8 := y2022.Day8{
			Year:      2022,
			Day:       8,
			InputFile: "./assets/2022/day8/input.txt",
			Logger:    sugar,
		}
		job = worker.NewJob(8, d8.Solve)
		pool.SubmitJob(job)

		d9 := y2022.Day9{
			Year:      2022,
			Day:       9,
			InputFile: "./assets/2022/day9/input.txt",
			Logger:    sugar,
		}
		job = worker.NewJob(9, d9.Solve)
		pool.SubmitJob(job)

		d10 := y2022.Day10{
			Year:      2022,
			Day:       10,
			InputFile: "./assets/2022/day10/input.txt",
			Logger:    sugar,
		}
		job = worker.NewJob(10, d10.Solve)
		pool.SubmitJob(job)

		pool.Stop()
	}()

	pool.Start()
}
