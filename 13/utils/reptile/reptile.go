package reptile

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"main/utils"
	"math/rand"
	"sort"
	"time"
)

type Product struct {
	Id                int64   `json:"id" orm:"id"`
	Name              string  `json:"name" orm:"name"`
	Appid             int64   `json:"appid" orm:"appid"`
	BuyMaxPrice       float64 `json:"buy_max_price" orm:"buy_max_price"`
	BuyNum            int     `json:"buy_num" orm:"buy_num"`
	QuickPrice        float64 `json:"quick_price" orm:"quick_price"`
	SellMinPrice      float64 `json:"sell_min_price" orm:"sell_min_price"`
	SellNum           int     `json:"sell_num" orm:"sell_num"`
	SteamPrice        float64 `json:"steam_price" orm:"steam_price"`
	SteamAveragePrice float64 `json:"steam_average_price" orm:"steam_average_price"`
	Discount          float64 `json:"discount" orm:"discount"`
}

var userAgentList = []string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
	"Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36",
}

func Get(game string) {
	ctx := gctx.New()
	categoryUrls := getCategoryUrl(game)
	c := g.Client()
	c.SetHeader("cookie", getCookie())
	c.SetAgent(getRandomUserAgent())
	for _, categoryUrl := range categoryUrls {
		categoryContent := c.GetContent(ctx, categoryUrl)
		fmt.Println(categoryContent)
		totalPageString, _ := gregex.MatchString("\"page_num\":\\s(?s:(.*?)),", categoryContent)
		totalPage := gconv.Int(totalPageString[0][1])
		for i := 1; i <= totalPage; i++ {
			categoryUrl := gstr.Replace(categoryUrl, "page_num=", "page_num="+gconv.String(i))
			fmt.Println(categoryUrl)
			categoryContent = c.GetContent(ctx, categoryUrl)
			appIdString, _ := gregex.MatchAllString("\"appid\":\\s(?s:(.*?)),", categoryContent)
			buyMaxPriceString, _ := gregex.MatchAllString("\"buy_max_price\":\\s\"(?s:(.*?))\",", categoryContent)
			buyNumString, _ := gregex.MatchAllString("\"buy_num\":\\s(?s:(.*?)),", categoryContent)
			idString, _ := gregex.MatchAllString("\"id\":\\s(?s:(.*?)),", categoryContent)
			nameString, _ := gregex.MatchAllString("\"name\":\\s\"(?s:(.*?))\",", categoryContent)
			quickPriceString, _ := gregex.MatchAllString("\"quick_price\":\\s\"(?s:(.*?))\",", categoryContent)
			sellMinPriceString, _ := gregex.MatchAllString("\"sell_min_price\":\\s\"(?s:(.*?))\",", categoryContent)
			sellNumString, _ := gregex.MatchAllString("\"sell_num\":\\s(?s:(.*?)),", categoryContent)
			for i := range appIdString {
				var product = Product{
					Appid:        gconv.Int64(appIdString[i][1]),
					BuyMaxPrice:  gconv.Float64(buyMaxPriceString[i][1]),
					BuyNum:       gconv.Int(buyNumString[i][1]),
					Id:           gconv.Int64(idString[i][1]),
					Name:         utils.UnicodeMarshal(nameString[i][1]),
					QuickPrice:   gconv.Float64(quickPriceString[i][1]),
					SellMinPrice: gconv.Float64(sellMinPriceString[i][1]),
					SellNum:      gconv.Int(sellNumString[i][1]),
				}
				productUrl := getProductUrl(product.Id)
				productContent := c.GetContent(ctx, productUrl)
				var steamPrices []float64
				steamPriceString, _ := gregex.MatchAllString("\\[\\n\\s*(?s:(.*?)),.*\\n\\s*(?s:(.*?))\\n.*", productContent)
				for j := range steamPriceString {
					steamPrices = append(steamPrices, gconv.Float64(steamPriceString[j][2]))
				}
				sort.Float64s(steamPrices)
				var steamAveragePrice float64
				for j := 1; j < len(steamPrices); j++ {
					steamAveragePrice += steamPrices[j]
				}
				product.SteamAveragePrice = steamAveragePrice / float64(len(steamPrices)-2)
				product.Discount = product.SellMinPrice / (0.87 * product.SteamAveragePrice)
				fmt.Println(product)
				if count, err := g.Model("product").Where("id=?", product.Id).Count(); count == 0 || err != nil {
					g.Model("product").Data(product).Insert()
				} else {
					g.Model("product").Data(product).Where("id=?", product.Id).Update()
				}
			}
			time.Sleep(time.Second * 1)
		}
	}
}

func getCategoryUrl(game string) (urls []string) {
	url := "https://buff.163.com/api/market/goods?game="
	url += game
	url += "&page_num=&category=&sort_by=price.asc"
	path := gfile.Pwd()
	content := gfile.GetContents(path + "/app/resource/public/resource/upload/" + game + "/category.txt")
	categories := gstr.SplitAndTrim(content, "\"")
	for _, v := range categories {
		urls = append(urls, gstr.Replace(url, "category=", "category="+v))
	}
	return urls
}

func getProductUrl(id int64) (url string) {
	url = "https://buff.163.com/api/market/goods/price_history?game=csgo&goods_id=&currency=CNY"
	url = gstr.Replace(url, "goods_id=", "goods_id="+gconv.String(id))
	return url
}

func getCookie() string {
	ctx := gctx.New()
	cookie, _ := g.Cfg().Get(ctx, "reptile.cookie")
	return cookie.String()
}

func getRandomUserAgent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgentList[r.Intn(len(userAgentList))]
}
