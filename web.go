package goscrapper

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// initialize new scrapper instance
func NewScrapper(url string) (*Web, error) {
	web := Web{}
	web.URL = url
	web.ctx = context.Background()

	if err := web.Fetch(); err != nil {
		return nil, err
	}

	return &web, nil
}

// initialize new scrapper instance with context
func NewContextScrapper(ctx context.Context, url string) (*Web, error) {
	web := Web{
		ctx: ctx,
		URL: url,
	}

	if err := web.Fetch(); err != nil {
		return nil, err
	}

	return &web, nil
}

type Web struct {
	URL string
	ctx context.Context
	Doc *goquery.Document
}

func (w *Web) Fetch() error {
	req, err := http.NewRequestWithContext(w.ctx, http.MethodGet, w.URL, nil)
	if err != nil {
		return fmt.Errorf("cound not create new request with ctx: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not get %s: %v", w.URL, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusTooManyRequests {
			return fmt.Errorf("you are being rate limited")
		}

		return fmt.Errorf("bad response from server: %s", res.Status)
	}

	w.Doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return fmt.Errorf("could not parse: %v", err)
	}

	select {
	case <-w.ctx.Done():
		return w.ctx.Err()
	default:
		return nil
	}
}
