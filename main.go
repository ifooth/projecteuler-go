package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/ifooth/projecteuler-go/euler"
)

var (
	version string
)

func parseArgs() (int, bool) {
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
	problemId, showContent := parseArgs()
	eulerProject := euler.NewEuler()
	st := time.Now()
	var answer int64
	var err error
	if showContent {
		answer, err = euler.GetProblemContent(problemId)
		if err != nil {
			slog.Info("not login", "duration", time.Since(st))
		} else {
			slog.Info("show content done", "duration", time.Since(st))
		}
	}

	st = time.Now()
	result, err := eulerProject.Calculate(problemId)
	if err != nil {
		slog.Info("calculate failed", "duration", time.Since(st), "err", err)
		os.Exit(0)
	}
	if err == nil {
		slog.Info(fmt.Sprintf("result %d(solve) = %d(euler) is %t", result, answer, result == answer), "duration", time.Since(st))
		os.Exit(0)
	}
	slog.Info("calculate done", "result", result, "duration", time.Since(st))
}

func init() {
	textHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.LevelInfo,
		ReplaceAttr: SourceAttr,
	})
	logger := slog.New(textHandler)
	slog.SetDefault(logger)
}

// SourceAttr source 格式化为 dir/file:line 格式
func SourceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key != slog.SourceKey {
		return a
	}

	source, ok := a.Value.Any().(*slog.Source)
	if !ok {
		return a
	}

	dir, file := filepath.Split(source.File)
	source.File = filepath.Join(filepath.Base(dir), file)
	return a
}
