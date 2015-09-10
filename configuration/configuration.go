package configuration

import (
	"net/url"
	"time"
)

type Configuration struct {
	Headers            map[string]string
	Methods            []string
	OutputFile         string
	ReportDirectory    string
	UrlResponseTimeout time.Duration
	WorkerWaitPeriod   time.Duration
	FuzzSetFile        string
	BaseURL            url.URL
}
