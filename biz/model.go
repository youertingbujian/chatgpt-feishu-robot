package biz

// ContentText 消息内容
type ContentText struct {
	Text string
}

// ReceiveEventEncrypt 加密订阅事件结构体
type ReceiveEventEncrypt struct {
	Encrypt string `json:"encrypt"` // 加密字符串
}

// DecryptToken 订阅事件结构体
type DecryptToken struct {
	Challenge string `json:"challenge"` // 应用需要在响应中原样返回的值
	Token     string `json:"token"`     // 即 Verification Token
	Type      string `json:"type"`      // 表示这是一个验证请求
}
