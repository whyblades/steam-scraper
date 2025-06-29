package skin

import (
	"fmt"
	"strings"
)

var Charms = map[string]float64{
	//`"Брелок | Карманная AWP"`:      0.46,
	//`"Брелок | Крошка Ава"`:         0.49,
	//`"Брелок | Крошка Крэсс"`:       0.55,
	//`"Брелок | Здоровяк Кев"`:       0.57,
	//`"Брелок | Резной резачок"`:     0.55,
	//`"Брелок | Мини-калаш"`:         0.69,
	//`"Брелок | Острый соус"`:        0.66,
	//`"Брелок | Щепотка соли"`:       0.70,
	`"Брелок | Крошка SAS"`:         0.89,
	`"Брелок | Дисколет-пулемёт"`:   1.35,
	`"Брелок | Крошка Спорт-сквоч"`: 1.40,
	`"Брелок | Стреляй, пока горячо"`: 1.62,
	`"Брелок | Поп-арт"`: 1.76,
	`"Брелок | Крошка Пляжник"`: 1.86,
	`"Брелок | Гламурный стрелок"`: 2.21,
	`"Брелок | Бешеный банан"`: 2.32,
	`"Брелок | Крошка Усатик"`: 2.49,
	`"Брелок | Крошка Цыпа"`: 2.85,
}

type Skin struct {
	Name string
	Price float64
	Charm string
	PriceWithoutCharm float64
}

func (s *Skin) CheckForProfit() (bool, string) {
	profit := ((s.PriceWithoutCharm + Charms[s.Charm]) / 1.15) - s.Price
	if profit > 0.5 && !strings.Contains(s.Name, "Charm"){
		str := fmt.Sprintf("Found profitable skin!\n\tName: %s\n\tCharm: %s\n\tProfit: %v\n", s.Name, s.Charm, profit)
		return true, str
	}
	return false, ""
}