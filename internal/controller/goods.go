package controller

import (
	v1 "antd-vegetable-go/api/v1"
	"context"

	"antd-vegetable-go/internal/model/entity"
	"antd-vegetable-go/internal/service"
)

var (
	Goods = cGoods{}
)

type cGoods struct{}

func (c *cGoods) FuzzyGetGoods(ctx context.Context, req *v1.FuzzyGetGoodsReq) (res *v1.FuzzyGetGoodsRes, err error) {
	results, err := service.Goods().FuzzyGetGoods(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return &v1.FuzzyGetGoodsRes{
			Goods: []*v1.Goods{},
		}, nil
	}
	var Goods []*v1.Goods
	for _, v := range results {
		Goods = append(Goods, entity2Api(v))
	}
	return &v1.FuzzyGetGoodsRes{
		Goods: Goods,
	}, nil
}

func entity2Api(g *entity.Goods) *v1.Goods {
	return &v1.Goods{
		Id:     g.Id,
		Name:   g.Name,
		Type:   g.Type,
		Unit:   g.Unit,
		Price:  g.Price,
		Tariff: g.Tariff,
	}
}
