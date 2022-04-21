package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type FuzzyGetGoodsReq struct {
	g.Meta `path:"/fuzzyGetGoods" tags:"Goods" method:"get" summary:"FuzzyGetGoods"`
	Name   string `v:"required|length:1,20#请输入Name|Name长度为1到20位"`
}

type FuzzyGetGoodsRes struct {
	g.Meta `mime:"application/json" example:"{'code':0}"`
	Goods  []*Goods `json:"goods,omitempty"`
}

type Goods struct {
	Id     uint64  `json:"_id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Unit   string  `json:"unit,omitempty"`
	Price  float64 `json:"price,omitempty"`
	Tariff float64 `json:"tariff,omitempty"`
}
