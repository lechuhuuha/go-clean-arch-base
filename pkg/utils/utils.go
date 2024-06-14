package utils

import (
	"io"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
)

const (
	Per600 = 0600
	Per640 = 0640
	Per644 = 0644
	Per700 = 0700
	Per711 = 0711
	Per750 = 0750
	Per755 = 0755
	Per777 = 0777
)

func SetTimezone() {
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		panic(err)
	}
	time.Local = loc // -> this is setting the global timezone
}

func IsValidURL(domain string) bool {
	_, err := url.Parse(domain)
	return err == nil
}

func FormatDomainURL(domain string) (string, error) {
	// Check if the domain is a valid host (FQDN)
	parsedURL, err := url.Parse(domain)
	if err != nil {
		return domain, err
	}
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "http"
	}
	return parsedURL.String(), nil
}

func GetDocumentTitle(docBody io.ReadCloser) string {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(docBody)
	if err != nil {
		log.Error().Err(err).Msg("Error when get new document")
		return ""
	}
	title := doc.Find("title").Text()
	return title
}

func GetURLField(d, field string) string {
	domain, err := url.Parse(d)
	result := ""
	if err != nil {
		return result
	}
	switch field {
	case "scheme":
		result = domain.Scheme
	case "userinfo":
		result = domain.User.String()
	case "path":
		result = domain.Path
	case "query":
		result = domain.Query().Encode()
	case "fragment":
		result = domain.Fragment
	default:
		result = domain.Host
	}
	if result == "" {
		result = domain.String()
	}
	return result
}
