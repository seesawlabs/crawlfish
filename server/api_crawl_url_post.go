package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"strings"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

func (a *ApiV1CrawlHandler) apiV1CrawlUrlPost(c *echo.Context) error {
	crawlRequest := CrawlRequest{}

	fmt.Println("We have got request !!")

	if err := c.Bind(&crawlRequest); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	crawlRequest.SplitWords()
	crawlRequest.TrimUrl()

	go func() {
		response := crawlPayload(&crawlRequest)
		a.Firebase.Push(response)
	}()

	return c.JSON(http.StatusOK, nil)
}

type CrawlRequest struct {
	Url      string   `json:"url"`
	Words    string   `json:"words"`
	WordList []string `json:"-"`
}

func (c *CrawlRequest) SplitWords() {
	for _, s := range strings.Split(c.Words, "\n") {
		c.WordList = append(c.WordList, strings.TrimSpace(s))
	}
}

func (c *CrawlRequest) TrimUrl() {
	c.Url = strings.TrimSpace(c.Url)
}

func (c *CrawlRequest) WordsTotal() int {
	return len(c.WordList)
}

type CrawlResponse struct {
	Website           string     `json:"website"`
	WordsFound        WordsFound `json:"words_found"`
	PagesSearched     Links      `json:"pages_searched"`
	PagesFound        Links      `json:"pages_found"`
	TotalTimeTaken    float64    `json:"total_time_taken"`
	PercentageOfFound float64    `json:"percentage_of_found"`
}

func crawlPayload(crawlRequest *CrawlRequest) *CrawlResponse {
	logrus.Info("Start crawling....")

	ext := &Ext{
		crawlRequest,
		map[string]Links{},
		Links{},
		Links{},
		&gocrawl.DefaultExtender{},
	}

	//Set custom options
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogError
	opts.SameHostOnly = true

	// TODO: Remove after development
	opts.MaxVisits = 2

	craw := gocrawl.NewCrawlerWithOptions(opts)

	startCrawling := time.Now()
	err := craw.Run(crawlRequest.Url)
	if err != nil {
		// TODO: Log error
		fmt.Println(err.Error())
	}

	totalTimeElapsed := time.Since(startCrawling)

	return &CrawlResponse{
		Website:           crawlRequest.Url,
		WordsFound:        ext.WordsFound,
		PagesSearched:     ext.PagesSearched,
		PagesFound:        ext.PagesFound,
		TotalTimeTaken:    totalTimeElapsed.Seconds(),
		PercentageOfFound: ext.WordsFound.PercentOfWordsFound(crawlRequest.WordsTotal()),
	}
}

type Link struct {
	Url         string  `json:"ulr"`
	TimeElapsed float64 `json:"time_elapsed"`
	WordCount   *int    `json:"word_count,omitempty"`
}

type Links []Link
type WordsFound map[string]Links

func (w WordsFound) PercentOfWordsFound(total int) float64 {
	found := len(w)

	return float64((found / total) * 100)
}

type Ext struct {
	CrawlPayload  *CrawlRequest
	WordsFound    WordsFound
	PagesSearched Links
	PagesFound    Links
	*gocrawl.DefaultExtender
}

func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
	start := time.Now()

	url := ctx.URL().String()

	pageHasAnyWord := false

	body, _ := ioutil.ReadAll(res.Body)
	for _, word := range e.CrawlPayload.WordList {
		if bytes.Contains(body, []byte(word)) {
			pageHasAnyWord = true

			startWordSearch := time.Now()
			count := bytes.Count(body, []byte(word))
			endWordSearch := time.Since(startWordSearch)

			e.WordsFound[word] = append(e.WordsFound[word], Link{url, endWordSearch.Seconds(), &count})
		}
	}

	elapsedTime := time.Since(start)

	// Add to PagesSearched
	e.PagesSearched = append(e.PagesSearched, Link{url, elapsedTime.Seconds(), nil})

	// Add to PagesFound
	if pageHasAnyWord {
		e.PagesFound = append(e.PagesFound, Link{url, elapsedTime.Seconds(), nil})
	}

	return nil, true
}

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	if isVisited {
		return false
	}

	return true
}
