package checker

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"os"
	"steam-scraper/skin"
)

var cookieMT string

func CheckPrice(skinsCh <-chan skin.Skin, resCh chan<- skin.Skin) {
	cookieMT = os.Getenv("COOKIE_MT")
	go func() {
		fmt.Println("MarketTracker API goroutine started")
		c := &http.Client{}
		for s := range skinsCh {
			price, err := requestApi(c, s.Name)
			if err != nil {
				fmt.Println(err.Error())
			}
			resCh <- skin.Skin{Name: s.Name, Price: s.Price, Charm: s.Charm, PriceWithoutCharm: price}
		}
	}()
}

func requestApi(c *http.Client, query string) (float64, error) {
	apiURL := "https://steammt.ru/api/stats/items?search=" + url.QueryEscape(query)
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return 0, errors.New("error while trying to form MarketTracker api")
	}
	req.Header.Set("Cookie", cookieMT)

	resp, err := c.Do(req)
	if err != nil {
		return 0, errors.New("failed to request MarketTracker api")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.New("failed to read MarketTracker response body")
	}

	results := gjson.Get(string(body), "items.#.priceHistory.0.price")
	var lowestPrice float64
	for i, v := range results.Array() {
		if i == 0 || v.Float() < lowestPrice {
			lowestPrice = v.Float()
		}
	}
	return lowestPrice, nil
}