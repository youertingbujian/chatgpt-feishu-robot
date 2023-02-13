package config

type Config struct {
	AppID        string `json:"app_id"`
	AppSecret    string `json:"app_secret"`
	EncryptKey   string `json:"encrypt_key"`
	OpenAiApiKey string `json:"openai_api_key"`
}

var Conf *Config

// replace conf when in use
const (
	appID        = "" //demo
	secret       = "" // demo
	encryptKey   = "" // demo
	openAiApiKey = ""
)

func init() {
	Conf = &Config{
		AppID:        appID,
		AppSecret:    secret,
		EncryptKey:   encryptKey,
		OpenAiApiKey: openAiApiKey,
	}
}
