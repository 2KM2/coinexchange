package handler

import "market-api/internal/svc"

func RoutesHandler(r *Routers, serverCtx *svc.ServiceContext) {
	rate := NewExchangeRateHandler(serverCtx)
	rateGroups := r.Group()
	rateGroups.Post("/market/exchange-rate/usd/:unit", rate.UsdRate)
}
