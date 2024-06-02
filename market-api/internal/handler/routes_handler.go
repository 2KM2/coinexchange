package handler

import "market-api/internal/svc"

func RoutesHandler(r *Routers, serverCtx *svc.ServiceContext) {
	rate := NewExchangeRateHandler(serverCtx)
	rateGroups := r.Group()
	rateGroups.Post("/market/exchange-rate/usd/:unit", rate.UsdRate)

	market := NewMarketHandler(serverCtx)
	marketGroups := r.Group()
	marketGroups.Post("/market/symbol-thumb-trend", market.SymbolThumbTrend)
}
