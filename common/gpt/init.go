package gpt

import (
	"chatgpt-web-new-go/common/config"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"context"
	"github.com/robfig/cron/v3"
	gogpt "github.com/sashabaranov/go-openai"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Init() {

	// init first
	doInitClient()

	// cron job
	c := cron.New()
	_, err := c.AddFunc("30 * * * *", doInitClient)
	if err != nil {
		logs.Error("cron add func error: %v", err)
		panic(err)
	}
	c.Start()
}

func doInitClient() {
	logs.Debug("doInitClient start len: %v", len(config.Gpt))

	dk := dao.Q.Aikey
	aiKeys, err := dk.Where(dk.IsDelete.Eq(0), dk.Status.Eq(1)).Find()
	if err != nil {
		logs.Error("ai keys get error: %v", err)
		panic(err)
	}

	cnf := config.Config.Gpt

	var clientList []*gogpt.Client
	for _, ak := range aiKeys {
		gptConfig := gogpt.DefaultConfig(ak.Key)

		// proxy
		addProxy(cnf.Proxy, gptConfig)

		// 自定义gptConfig.BaseURL
		if ak.Host != "" {
			gptConfig.BaseURL = ak.Host
		}

		clientList = append(clientList, gogpt.NewClientWithConfig(gptConfig))
	}

	config.Gpt = clientList

	logs.Debug("doInitClient start len: %v", len(config.Gpt))
}

func addProxy(proxy string, gptConfig gogpt.ClientConfig) {
	if proxy != "" {
		transport := &http.Transport{}

		if strings.HasPrefix(proxy, "socks5h://") {
			// 创建一个 DialContext 对象，并设置代理服务器
			dialContext, err := newDialContext(proxy[10:])
			if err != nil {
				panic(err)
			}
			transport.DialContext = dialContext
		} else {
			// 创建一个 HTTP Transport 对象，并设置代理服务器
			proxyUrl, err := url.Parse(proxy)
			if err != nil {
				panic(err)
			}
			transport.Proxy = http.ProxyURL(proxyUrl)
		}
		// 创建一个 HTTP 客户端，并将 Transport 对象设置为其 Transport 字段
		gptConfig.HTTPClient = &http.Client{
			Transport: transport,
		}
	}
}

type dialContextFunc func(ctx context.Context, network, address string) (net.Conn, error)

func newDialContext(socks5 string) (dialContextFunc, error) {
	baseDialer := &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 60 * time.Second,
	}

	if socks5 != "" {
		// split socks5 proxy string [username:password@]host:port
		var auth *proxy.Auth = nil

		if strings.Contains(socks5, "@") {
			proxyInfo := strings.SplitN(socks5, "@", 2)
			proxyUser := strings.Split(proxyInfo[0], ":")
			if len(proxyUser) == 2 {
				auth = &proxy.Auth{
					User:     proxyUser[0],
					Password: proxyUser[1],
				}
			}
			socks5 = proxyInfo[1]
		}

		dialSocksProxy, err := proxy.SOCKS5("tcp", socks5, auth, baseDialer)
		if err != nil {
			return nil, err
		}

		contextDialer, ok := dialSocksProxy.(proxy.ContextDialer)
		if !ok {
			return nil, err
		}

		return contextDialer.DialContext, nil
	} else {
		return baseDialer.DialContext, nil
	}
}
