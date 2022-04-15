package model

type Product struct {
	Id                int64   `json:"id" orm:"id"`
	Name              string  `json:"name" orm:"name"`
	Appid             int64   `json:"appid" orm:"appid"`
	BuyMaxPrice       float32 `json:"buy_max_price" orm:"buy_max_price"`
	BuyNum            int     `json:"buy_num" orm:"buy_num"`
	QuickPrice        float32 `json:"quick_price" orm:"quick_price"`
	SellMinPrice      float32 `json:"sell_min_price" orm:"sell_min_price"`
	SellNum           int     `json:"sell_num" orm:"sell_num"`
	SteamPrice        float32 `json:"steam_price" orm:"steam_price"`
	SteamAveragePrice float32 `json:"steam_average_price" orm:"steam_average_price"`
	Discount          float32 `json:"discount" orm:"discount"`
}
