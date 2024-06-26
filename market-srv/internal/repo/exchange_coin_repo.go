package repo

import (
	"context"
	"market-srv/internal/model"
)

type ExchangeCoinRepo interface {
	FindVisible(ctx context.Context) (list []*model.ExchangeCoin, err error)
	FindBySymbol(ctx context.Context, symbol string) (*model.ExchangeCoin, error)
}
