package logic

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Password  string
	Host      string
	Proxy     string
}

type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type Kline struct {
}

func NewKline() {

}
