package euler

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
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
		return 0, errors.New("no solved")
	}
	result := pFunc()
	return result, nil
}

func init() {
	solvedProblems = map[int]problems.ProblemFunc{
		1: problems.Problem1,
		2: problems.Problem2,
		3: problems.Problem3,
		4: problems.Problem4,
		5: problems.Problem5,
		6: problems.Problem6,
	}
}

func GetProblemContent(problemId int) {
	// Request the HTML page.
	client := req.C().SetTimeout(time.Second * 60)
	if os.Getenv("EULER_DEBUG") != "" {
		client = client.DevMode()
	}
	url := fmt.Sprintf("https://projecteuler.net/problem=%d", problemId)

	req := client.R()

	sessionId := os.Getenv("EULER_SESSION_ID")
	keepAliveId := os.Getenv("EULER_KEEPALIVE_ID")
	if sessionId != "" && keepAliveId != "" {
		sessionCookie := &http.Cookie{
			Name:   "PHPSESSID",
			Value:  sessionId,
			Domain: "projecteuler.net",
		}
		keepAliveCookie := &http.Cookie{
			Name:   "keep_alive",
			Value:  keepAliveId,
			Domain: "projecteuler.net",
		}
		req = req.SetCookies(sessionCookie, keepAliveCookie)
	}

	resp, err := req.Get(url)
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
	doc.Find("#problem_answer .data_entry").Each(func(i int, s *goquery.Selection) {
		fmt.Println("====")
		answer := s.Find(".strong").Text()
		notice := s.Find(".small_notice").Text()
		fmt.Println("Answer:", answer)
		fmt.Println(notice)
	})

	fmt.Println("==== End ====")
}
