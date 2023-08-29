package config

type Configuration struct {
	Port        int                `json:"port"`
	Db          *dbConfig          `json:"db"`
	Redis       *redisConfig       `json:"redis"`
	Gpt         *gptConfig         `json:"gpt"`
	EmailServer *emailServerConfig `json:"emailServer"`
}

type dbConfig struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Host     string `json:"host"` // and port
	HostR1   string `json:"host_r_1"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type redisConfig struct {
	Addr     string `json:"addr"` // host and port
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type gptConfig struct {
	ApiKey           string  `json:"api_key"`
	Proxy            string  `json:"proxy"`
	ApiURL           string  `json:"api_url"`
	BotDesc          string  `json:"bot_desc"`
	Model            string  `json:"model"`
	MaxTokens        int     `json:"max_tokens"`
	TopP             float32 `json:"top_p"`
	FrequencyPenalty float32 `json:"frequency_penalty"`
	PresencePenalty  float32 `json:"presence_penalty"`
}

type emailServerConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	SenderName string `json:"sender_name"`
	User       string `json:"user"`
	Password   string `json:"password"`
}
