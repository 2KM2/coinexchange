package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"grpc-common/market/types/market"
	"market-api/internal/svc"
	"market-api/internal/types"
)

type MarketLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMarketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MarketLogic {
	return &MarketLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *MarketLogic) SymbolThumbTrend(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	//var thumbs []*market.CoinThumb
	//thumb := l.svcCtx.Processor.GetThumb()
	symbolThumbRes, err := l.svcCtx.MarketRpc.FindSymbolThumbTrend(context.Background(), &market.MarketReq{Ip: req.Ip})
	if err != nil {
		return nil, err
	}
	if err := copier.Copy(&list, symbolThumbRes.List); err != nil {
		return nil, err
	}
	return
}
func (l *MarketLogic) SymbolThumb(req *types.MarketReq) (list []*types.CoinThumbResp, err error) {
	return nil, err
}
