package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/rate"
	"market-srv/internal/domain"
	"market-srv/internal/svc"
)

type ExchangeRateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	exchangeRateDomain *domain.ExchangeRateDomain
}

func NewExchangeRateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExchangeRateLogic {
	return &ExchangeRateLogic{
		ctx:                ctx,
		svcCtx:             svcCtx,
		Logger:             logx.WithContext(ctx),
		exchangeRateDomain: domain.NewExchangeRateDomain(),
	}
}

func (l *ExchangeRateLogic) UsdRate(req *rate.RateReq) (*rate.RateRes, error) {
	usdRate := l.exchangeRateDomain.UsdRate(req.Unit)
	return &rate.RateRes{
		Rate: usdRate,
	}, nil
}
