package utils

import (
	resty "github.com/go-resty/resty/v2"
	"time"
)

func GetRestyClient() *resty.Client {
	// Create a Resty Client
	client := resty.New()

	// Retries are configured per client
	client.
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(20 * time.Second)

	return client
}
