package aoc

import (
	"fmt"
	"github.com/blambo10/advent-of-code-22/pkg/config"
	"github.com/blambo10/advent-of-code-22/pkg/utils"
	"net/http"
)

func getDayData(day string) (*string, error) {

	cfg := config.GetAOCConfig()
	client := utils.GetRestyClient().R()
	url := fmt.Sprintf("https://%s/2022/day/%s/input", cfg.URL, day)

	resp, err := client.
		SetCookie(&http.Cookie{
			Name:  "session",
			Value: cfg.Cookie,
		}).
		Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("an error occured fetching inputs")
	}

	data := string(resp.Body())

	return &data, nil
}
