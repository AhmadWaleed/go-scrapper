package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reg = regexp.MustCompile(`([a-zA-Z0-9._-]+@([a-zA-Z0-9_-]+\.)+[a-zA-Z0-9_-]+)`)

func (w *Web) emails() ([]string, error) {
	body, err := w.Doc.Html()
	if err != nil {
		return nil, fmt.Errorf("could not get body (html) %v", err)
	}

	var emails []string
	buf := bytes.NewBufferString(body).Bytes()
	parseEmails(buf, &emails)

	return emails, nil
}

func parseEmails(body []byte, scrapedEmails *[]string) {
	res := reg.FindAll(body, -1)
	for _, r := range res {
		email := string(r)
		if !validEmail(email) {
			continue
		}

		*scrapedEmails = append(*scrapedEmails, strings.TrimSpace(email))
	}
}

func validEmail(email string) bool {
	split := strings.Split(email, ".")
	if len(split) < 2 {
		return false
	}

	ending := split[len(split)-1]

	if len(ending) < 2 {
		return false
	}

	if _, err := strconv.Atoi(ending); err == nil {
		return false
	}

	return true
}
