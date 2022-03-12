package euler

import (
	"errors"

	"github.com/ifooth/projecteuler-go/euler/problems"
)

var solvedProblems map[int]problems.ProblemFunc

type Euler struct {
}

func NewEuler() *Euler {
	return &Euler{}
}

func (e *Euler) Calculate(problemId int) (int64, error) {
	pFunc, ok := solvedProblems[problemId]
	if !ok {
		return 0, errors.New("not finish")
	}
	result := pFunc()
	return result, nil
}

func init() {
	solvedProblems = map[int]problems.ProblemFunc{
		1: problems.Problem1,
		2: problems.Problem2,
	}
}
