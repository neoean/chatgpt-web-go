package config

import (
	"github.com/go-redis/redis"
	"github.com/robfig/cron/v3"
	gogpt "github.com/sashabaranov/go-openai"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

var (
	Config      *Configuration
	DB          *gorm.DB // DB instance
	Redis       *redis.Client
	Gpt         *gogpt.Client
	Gcron       *cron.Cron // cron
	EmailDialer *gomail.Dialer
)
