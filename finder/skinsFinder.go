package finder

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"os"
	"steam-scraper/skin"
	"time"
)

var cookieSteam string

func ScrapeCharmInfo(skinsCh chan<- skin.Skin) {
	cookieSteam = os.Getenv("COOKIE_STEAM")
	for charmName, _ := range skin.Charms {
		go func(charmName string) {
			fmt.Printf("goroutine with %s started\n", charmName)
			c := &http.Client{}
			for {
				for pageNum := range 6 {
					for {
						results, err := getListings(c, charmName, int64(pageNum*10))
						if err == nil {
							for _, res := range results {
								skinsCh <- res
							}
							break
						} else {
							time.Sleep(2 * time.Minute)
							fmt.Printf("error in goroutine with %s: %s, page: %d\n", charmName, err.Error(), pageNum)
						}
					}
				}
			}
		}(charmName)
	}
}

func getListings(client *http.Client, charmName string, start int64) ([]skin.Skin, error) {
	apiURL := fmt.Sprintf("https://steamcommunity.com/market/search/render" +
		"?norender=1&sort_column=price&sort_dir=asc&search_descriptions=True&appid=730&" +
		"start=%d&query=%s", start, url.QueryEscape(charmName))

	request, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, errors.New("error while trying to form Steam api request")
	}
	request.Header.Set("Cookie", cookieSteam)

	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New("failed to request Steam api")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return nil, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read Steam api response body ")
	}

	results := gjson.Get(string(body), "results")
	var skins []skin.Skin
	for _, result := range results.Array() {
		name := gjson.Get(result.String(), "hash_name")
		price := gjson.Get(result.String(), "sell_price")
		skins = append(skins, skin.Skin{Name: name.String(), Price: price.Float() / 100, Charm: charmName})
	}

	return skins, nil
}
