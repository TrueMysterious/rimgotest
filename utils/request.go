package utils

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

func GetJSON(url string) (gjson.Result, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return gjson.Result{}, err
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Origin", "https://imgur.com")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:91.0) Gecko/20100101 Firefox/91.0")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return gjson.Result{}, err
	}
	rateLimitRemaining := res.Header.Get("X-RateLimit-UserRemaining")
	if rateLimitRemaining != "" {
		ratelimit, _ := strconv.Atoi(rateLimitRemaining)
		if ratelimit <= 0 {
			return gjson.Result{}, fmt.Errorf("ratelimited by imgur")
		}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return gjson.Result{}, err
	}

	switch (res.StatusCode) {
	case 200:
		return gjson.Parse(string(body)), nil
	case 429:
		return gjson.Result{}, fmt.Errorf("ratelimited by imgur")
	default:
		return gjson.Result{}, fmt.Errorf("received status %s, expected 200 OK.\n%s", res.Status, string(body))
	}
}