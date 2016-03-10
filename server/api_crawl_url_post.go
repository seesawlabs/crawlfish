package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"strings"

	"github.com/PuerkitoBio/gocrawl"
	"github.com/PuerkitoBio/goquery"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

func (a *ApiV1CrawlHandler) apiV1CrawlUrlPost(c *echo.Context) error {
	crawlRequest := CrawlRequest{}

	if err := c.Bind(&crawlRequest); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	crawlRequest.SplitWords()
	crawlRequest.TrimUrl()

	startDate := time.Now()

	response := CrawlResponse{
		CrawlingStartDate: &startDate,
		Words:             crawlRequest.WordList,
		Website:           crawlRequest.Url,
		InProgress:        true,
	}

	fullUrl, err := a.Firebase.Push(response)
	if err != nil {
		logrus.Error(err)
	}

	response.FirebasePath = fullUrl

	u, err := url.Parse(*fullUrl)
	path := u.Path

	go func() {
		response := crawlPayload(&crawlRequest)

		err := a.Firebase.Update(path, response)
		if err != nil {
			logrus.Error(err)
		}

	}()

	return c.JSON(http.StatusOK, response)
}

type CrawlRequest struct {
	Url      string   `json:"url"`
	Words    string   `json:"words"`
	WordList []string `json:"-"`
}

func (c *CrawlRequest) SplitWords() {
	for _, s := range strings.Split(c.Words, "\n") {
		word := strings.TrimSpace(s)
		if word != "" {
			c.WordList = append(c.WordList, word)
		}
	}
}

func (c *CrawlRequest) TrimUrl() {
	c.Url = strings.TrimSpace(c.Url)
}

func (c *CrawlRequest) WordsTotal() int {
	return len(c.WordList)
}

type CrawlResponse struct {
	FirebasePath      string     `json:"firebase_path,omitempty"`
	CrawlingStartDate *time.Time `json:"crawling_start_date,omitempty"`
	CrawlingEndDate   *time.Time `json:"crawling_end_date,omitempty"`
	Words             []string   `json:"words,omitempty"`
	Website           string     `json:"website,omitempty"`
	WordsFound        WordsFound `json:"words_found,omitempty"`
	PagesSearched     Links      `json:"pages_searched,omitempty"`
	TotalTimeTaken    float64    `json:"total_time_taken,omitempty"`
	PercentageOfFound float64    `json:"percentage_of_found,omitempty"`
	InProgress        bool       `json:"in_progress"`
}

func crawlPayload(crawlRequest *CrawlRequest) *CrawlResponse {
	logrus.Info("Start crawling....")

	ext := &Ext{
		crawlRequest,
		make(map[string]Links, 1),
		Links{},
		&gocrawl.DefaultExtender{},
	}

	//Set custom options
	opts := gocrawl.NewOptions(ext)
	opts.CrawlDelay = 1 * time.Second
	opts.LogFlags = gocrawl.LogError
	opts.SameHostOnly = true

	opts.MaxVisits = 50

	craw := gocrawl.NewCrawlerWithOptions(opts)

	startCrawling := time.Now()
	err := craw.Run(crawlRequest.Url)
	if err != nil {
		logrus.Error(err)
		fmt.Println(err.Error())
	}

	totalTimeElapsed := time.Since(startCrawling)

	endDate := time.Now()
	return &CrawlResponse{
		CrawlingEndDate:   &endDate,
		InProgress:        false,
		WordsFound:        ext.WordsFound,
		PagesSearched:     ext.PagesSearched,
		TotalTimeTaken:    totalTimeElapsed.Seconds(),
		PercentageOfFound: ext.WordsFound.PercentOfWordsFound(crawlRequest.WordsTotal()),
	}
}

type Link struct {
	Url         string  `json:"ulr"`
	TimeElapsed float64 `json:"time_elapsed"`
	WordCount   *int    `json:"word_count,omitempty"`
	Found       *bool   `json:"found,omitempty"`
}

type Links []Link
type WordsFound map[string]Links

func (w WordsFound) PercentOfWordsFound(total int) float64 {
	found := len(w)

	return (float64(found) / float64(total)) * 100
}

type Ext struct {
	CrawlPayload  *CrawlRequest
	WordsFound    WordsFound
	PagesSearched Links
	*gocrawl.DefaultExtender
}

func (e *Ext) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {

	start := time.Now()

	url := ctx.URL().String()

	fmt.Printf("Crawling URL: %s || Status: %s \n \n ", url, res.Status)

	pageHasAnyWord := false

	body, _ := ioutil.ReadAll(res.Body)
	for _, word := range e.CrawlPayload.WordList {
		if bytes.Contains(body, []byte(word)) {
			pageHasAnyWord = true

			startWordSearch := time.Now()
			count := bytes.Count(body, []byte(word))
			endWordSearch := time.Since(startWordSearch)

			e.WordsFound[word] = append(e.WordsFound[word], Link{url, endWordSearch.Seconds(), &count, nil})
		}
	}

	elapsedTime := time.Since(start)

	// Add to PagesSearched
	e.PagesSearched = append(e.PagesSearched, Link{
		url,
		elapsedTime.Seconds(),
		nil,
		&pageHasAnyWord,
	})

	return nil, true
}

func (e *Ext) Filter(ctx *gocrawl.URLContext, isVisited bool) bool {
	if isVisited {
		return false
	}

	return true
}
