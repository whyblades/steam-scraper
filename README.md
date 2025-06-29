
# Steam Scraper

Tool for finding profitable cs2 skins with charms on Steam Market using SteamMarketTracker(steammt.ru) API and Internal Steam API (credits to [InternalSteamWebAPI](https://github.com/Revadike/InternalSteamWebAPI) repo)

## Steps for getting started:
1. Clone repo: ```git clone https://github.com/whyblades/steam-scraper.git```

2. Create ```.env``` file with ```COOKIE_MT``` and ```COOKIE_STEAM```

3. Authorize on [steammt.ru](https://steammt.ru/), copy ```Cookie``` header's value using browser devdtools and assign this value to ```COOKIE_MT```

4. Authorize on Steam. Account currency must be USD and language must be set to Russian

5. Make any request to [steamcommunity.com/market](https://steamcommunity.com/market/), get ```Cookie``` header's value and assign it to ```COOKIE_STEAM```
6. (optional) If you want to get notifications about found skins to Telegram Bot, follow simple steps in [Telegram Bot](###Telegram-Bot) section

7. Now you can just ```go run .```


### Telegram Bot
1. Create bot using ```@BotFather```  in Telegram

2. When bot is created, copy Bot Api token and set it as value of ```TELEGRAMBOT_APITOKEN``` variable in ```.env``` file
3. Get chat id of chat with your bot ([getting chat id](https://gist.github.com/nafiesl/4ad622f344cd1dc3bb1ecbe468ff9f8a))
4. Create ```CHAT_ID``` in ```.env``` and assign it the value of chat id from previous step