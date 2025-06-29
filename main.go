package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"steam-scraper/checker"
	"steam-scraper/finder"
	"steam-scraper/skin"
	"steam-scraper/tgbot"
)

func main (){
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	notifyBot := tgbot.InitBot()
	finderCh := make(chan skin.Skin)
	resultCh := make(chan skin.Skin)

	finder.ScrapeCharmInfo(finderCh)
	checker.CheckPrice(finderCh, resultCh)

	for v := range resultCh {
		isProfitable, str := v.CheckForProfit()
		if isProfitable {
			fmt.Print(str)
			tgbot.NotifyNewSkin(notifyBot, str)
		}
	}

	fmt.Println("--------------ENDED SCRAPING--------------")
}
