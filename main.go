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

func ParseArgs() int {
	problemId := flag.Int("n", 0, "euler problem num")
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

	return *problemId
}

func main() {
	problemId := ParseArgs()
	eulerProject := euler.NewEuler()
	st := time.Now()
	result, err := eulerProject.Calculate(problemId)
	if err != nil {
		log.Println(result, "duration", time.Since(st))
		os.Exit(0)
	}
	log.Println(result, "duration", time.Since(st))
}
