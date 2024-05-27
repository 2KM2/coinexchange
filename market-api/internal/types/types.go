package types

type RateRequest struct {
	Unit string `path:"unit" json:"unit"`
	Ip   string `json:"ip,optional"`
}

type RateResponse struct {
	Rate float64 `json:"rate"`
}