package goscrapper

import (
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
	"testing"
)

func TestEmails(t *testing.T) {
	tt := []struct {
		name   string
		html   string
		emails []string
	}{
		{
			name: "single email",
			html: `
<html>
	<head>
		<title>Test document</title>
	</head>
	<body>
		<p>john@example.com</p>
	</body>
</html>`,
			emails: []string{"john@example.com"},
		},
		{
			name: "multiple emails",
			html: `
<html>
	<head>
		<title>Test document</title>
	</head>
	<body>
		<p>john@example.com</p>
		<p>jane@gmail.com</p>
	</body>
</html>`,
			emails: []string{"john@example.com", "jane@gmail.com"},
		},
		{
			name: "invalid email",
			html: `
<html>
	<head>
		<title>Test document</title>
	</head>
	<body>
		<p>john@example.com</p>
		<p>gmail.com</p>
		<p>gmail.</p>
		<p>343243@23343.23432</p>
	</body>
</html>`,
			emails: []string{"john@example.com"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.html))
			if err != nil {
				t.Fatalf("could not create doc reader %v", err)
			}

			web := Web{Doc: doc}
			emails, err := web.emails()
			if err != nil {
				t.Fatalf("could not fetch emails %v", err)
			}
			if eq := reflect.DeepEqual(emails, tc.emails); !eq {
				t.Errorf("heading should be: %v, got: %v", tc.emails, emails)
			}
		})
	}
}
