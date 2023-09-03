package euler

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/imroc/req/v3"

	"github.com/ifooth/projecteuler-go/euler/problems"
)

//go:generate go run ./gen/main.go -noconsts -novars -notests -notypes problems

// Euler ..
type Euler struct{}

// NewEuler ..
func NewEuler() *Euler {
	return &Euler{}
}

// Calculate ..
func (e *Euler) Calculate(problemId int) (int64, error) {
	problemKey := fmt.Sprintf("Problem%d", problemId)
	pFunc, ok := problems.Functions[problemKey]
	if !ok {
		return 0, errors.New("no solved")
	}
	result := pFunc.Call([]reflect.Value{})
	if len(result) != 1 {
		return 0, fmt.Errorf("not valid result, %s", result)
	}
	answer, ok := result[0].Interface().(int64)
	if !ok {
		return 0, fmt.Errorf("not valid answer, %s", result)
	}
	return answer, nil
}

// GetProblemContent ..
func GetProblemContent(problemId int) (int64, error) {
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

	if !resp.IsSuccessState() {
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

	fmt.Printf("Euler:    https://projecteuler.net/problem=%d\n", problemId)
	fmt.Printf("Euler-CN: http://pe-cn.github.io/%d/\n", problemId)

	var answer int
	err = errors.New("not login")
	doc.Find("#problem_answer .data_entry").Each(func(i int, s *goquery.Selection) {
		fmt.Println("====")
		answer, err = strconv.Atoi(s.Find(".strong").Text())

		notice := s.Find(".small_notice").Text()
		fmt.Println("Answer:", answer)
		fmt.Println(notice)
	})

	fmt.Println("==== End ====")
	return int64(answer), err
}
