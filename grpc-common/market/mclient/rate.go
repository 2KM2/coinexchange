package mclient

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"grpc-common/market/types/rate"
)

type (
	RateReq = rate.RateReq
	RateRes = rate.RateRes

	ExchangeRate interface {
		UsdRate(ctx context.Context, in *RateReq, opts ...grpc.CallOption) (*RateRes, error)
	}

	defaultExchangeRate struct {
		cli zrpc.Client
	}
)

func NewExchangeRate(cli zrpc.Client) ExchangeRate {
	return &defaultExchangeRate{
		cli: cli,
	}
}

func (m *defaultExchangeRate) UsdRate(ctx context.Context, in *RateReq, opts ...grpc.CallOption) (*RateRes, error) {
	client := rate.NewExchangeRateClient(m.cli.Conn())
	return client.UsdRate(ctx, in, opts...)
}
