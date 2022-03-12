package euler

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/ifooth/projecteuler-go/euler/problems"
	"github.com/imroc/req/v3"
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

func GetProblemContent(problemId int) {
	// Request the HTML page.
	client := req.C().SetTimeout(time.Second * 60)
	url := fmt.Sprintf("https://projecteuler.net/problem=%d", problemId)

	resp, err := client.R().
		Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccess() {
		log.Fatal(err)
	}

	body, err := resp.ToBytes()
	if err != nil {
		log.Fatal(err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	title := doc.Find("#content h2").Text()
	content := doc.Find(".problem_content").Text()
	fmt.Println("==== Project Euler Problem ====")
	fmt.Println(title)
	fmt.Println(content)
	fmt.Println("==== End ====")
}
