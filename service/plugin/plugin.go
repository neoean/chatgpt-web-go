package plugin

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func PluginList(ctx context.Context) (pluginList []*model.Plugin, err error) {
	dp := dao.Q.Plugin

	pluginList, err = dp.WithContext(ctx).Find()
	if err != nil {
		logs.Error("plugin list error: %v", err)
		return nil, err
	}
	return
}
