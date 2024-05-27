package server

import (
	"context"
	"grpc-common/market/types/rate"
	"market-srv/internal/logic"
	"market-srv/internal/svc"
)

type ExchangeRateServer struct {
	svcCtx *svc.ServiceContext
	rate.UnimplementedExchangeRateServer
}

func NewExchangeRateServer(ctx *svc.ServiceContext) *ExchangeRateServer {
	return &ExchangeRateServer{
		svcCtx: ctx,
	}
}

// UsdRate 重载grpc的实现
func (e *ExchangeRateServer) UsdRate(ctx context.Context, req *rate.RateReq) (*rate.RateRes, error) {
	l := logic.NewExchangeRateLogic(ctx, e.svcCtx)
	return l.UsdRate(req)
}
