package main

/* func main() {

	//设置API认证信息并创建实例
	s := kucoin.NewApiService(
		kucoin.ApiKeyOption("secret"),
		kucoin.ApiSecretOption("secret"),
		kucoin.ApiPassPhraseOption("passphrase"),
		kucoin.ApiKeyVersionOption("key_id"),
	)

	//获取WebSocket连接令牌并建立连接
	rsp, err := s.WebSocketPublicToken()
	tk := &kucoin.WebSocketTokenModel{}

	if err != nil {
		//
	}

	if err := rsp.ReadData(tk); err != nil {
		//
		fmt.Println(err.Error())
	}

	// c := s.NewWebSocketClient(tk)
	// mc, ec, err := c.Connect()
	//订阅特定交易对市场行情
} */
