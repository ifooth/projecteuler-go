package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ifooth/projecteuler-go/euler"
)

var (
	version string
)

func ParseArgs() (int, bool) {
	problemId := flag.Int("n", 0, "euler problem num")
	showContent := flag.Bool("c", false, "show euler problem content")
	flagVersion := flag.Bool("v", false, "print version")

	flag.Parse()

	if *flagVersion {
		fmt.Println("version:", version)
		os.Exit(0)
	}

	if *problemId == 0 {
		flag.Usage()
		os.Exit(0)
	}

	return *problemId, *showContent
}

func main() {
	problemId, showContent := ParseArgs()
	eulerProject := euler.NewEuler()
	st := time.Now()
	if showContent {
		euler.GetProblemContent(problemId)
		log.Println("show content done", "duration", time.Since(st))
	}

	st = time.Now()
	result, err := eulerProject.Calculate(problemId)
	if err != nil {
		log.Println(err, "duration", time.Since(st))
		os.Exit(0)
	}
	log.Println(result, "duration", time.Since(st))
}
