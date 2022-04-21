package service

import (
	"antd-vegetable-go/internal/model/entity"
	"antd-vegetable-go/internal/service/internal/dao"
	"context"
)

var insGoods = sGoods{}

type sGoods struct {
}

func Goods() *sGoods {
	return &insGoods
}

func (s *sGoods) FuzzyGetGoods(ctx context.Context, name string) ([]*entity.Goods, error) {
	var (
		goods []*entity.Goods
		m     = dao.Goods.Ctx(ctx)
	)
	err := m.WhereLike("name", "%"+name+"%").Scan(&goods)
	if err != nil {
		return nil, err
	}
	return goods, nil
}
