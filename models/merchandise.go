package models

type Merchandise struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

/*
**Мерч** — это продукт, который можно купить за монетки.
Всего в магазине доступно 10 видов мерча.
Каждый товар имеет уникальное название и цену.
Ниже приведён список наименований и их цены.

| Название     | Цена |
|--------------|------|
| t-shirt      | 80   |
| cup          | 20   |
| book         | 50   |
| pen          | 10   |
| powerbank    | 200  |
| hoody        | 300  |
| umbrella     | 200  |
| socks        | 10   |
| wallet       | 50   |
| pink-hoody   | 500  |

Предполагается, что в магазине бесконечный запас каждого вида мерча.
*/

var MerchandiseList = map[string]Merchandise{
	"t-shirt":    {"t-shirt", 80},
	"cup":        {"cup", 20},
	"book":       {"book", 50},
	"pen":        {"pen", 10},
	"powerbank":  {"powerbank", 200},
	"hoody":      {"hoody", 300},
	"umbrella":   {"umbrella", 200},
	"socks":      {"socks", 10},
	"wallet":     {"wallet", 50},
	"pink-hoody": {"pink-hoody", 500},
}
