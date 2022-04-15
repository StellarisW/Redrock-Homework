package api

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"main/utils/reptile"
)

type IndexApi struct {
	BaseApi
}

var insIndex = IndexApi{}

func (a *IndexApi) Get(r *ghttp.Request) {

}

func (a *IndexApi) Update(r *ghttp.Request) {
	reptile.Get("csgo")
}
