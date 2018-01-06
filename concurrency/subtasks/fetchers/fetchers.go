package fetchers

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Fetcher interface {
	Fetch(url string) (string, error)
	GetName() string
}

type GoogleFetcher struct {
	Name string
}

func (g *GoogleFetcher) Fetch(url string) (string, error) {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%s is fetching %s", g.Name, url), nil
}

func (g *GoogleFetcher) GetName() string {
	return g.Name
}

func NewGoogleFetcher(name string) *GoogleFetcher {
	return &GoogleFetcher{Name: name}
}

type BingFetcher struct {
	Name string
}

func (b *BingFetcher) Fetch(url string) (string, error) {
	time.Sleep(time.Second * 3)
	return fmt.Sprintf("%s is fetching %s", b.Name, url), nil
}

func (b *BingFetcher) GetName() string {
	return b.Name
}

func NewBingFetcher(name string) *BingFetcher {
	return &BingFetcher{Name: name}
}

type DuckDuckGoFetcher struct {
	Name string
}

func (d *DuckDuckGoFetcher) Fetch(url string) (string, error) {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("%s is fetching %s", d.Name, url), nil
}

func (d *DuckDuckGoFetcher) GetName() string {
	return d.Name
}

func NewDuckDuckGoFetcherFetcher(name string) *DuckDuckGoFetcher {
	return &DuckDuckGoFetcher{Name: name}
}
func FetchResults(url string, fetchers []Fetcher, timeout time.Duration) ([]string, []error) {
	chStr := make(chan string, len(fetchers))
	chErr := make(chan error, len(fetchers))
	for _, f := range fetchers {
		go func(f Fetcher) {
			s, err := f.Fetch(url)
			if err != nil {
				chErr <- err
			} else {
				chStr <- s
			}
		}(f)
	}
	stringResults := make([]string, 0, len(fetchers))
	errorResults := make([]error, 0, len(fetchers))
	for range fetchers {
		select {
		case s := <-chStr:
			stringResults = append(stringResults, s)
		case e := <-chErr:
			errorResults = append(errorResults, e)
		}
	}
	return stringResults, errorResults
}

func RunFetchers() {
	fetchers := []Fetcher{NewGoogleFetcher("Google"), NewGoogleFetcher("Bing"), NewGoogleFetcher("Duck Duck Go")}
	r, e := FetchResults("http://www.abc.com", fetchers, time.Millisecond*100)
	spew.Dump(r, e)
}
