package client

import (
	"errors"
	"net/http"
	"time"
)

var httpClient *http.Client

func noRedirect(req *http.Request, via []*http.Request) error {
	return errors.New("Don't redirect!")
}

func init() {
	// http client
	httpClient = &http.Client{
		Timeout:       time.Second,
		CheckRedirect: noRedirect,
	}
}
