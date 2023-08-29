package config

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"context"
	"encoding/json"
)

var (
	indexConfigs = []string{
		"shop_introduce",
		"user_introduce",
		"notifications",
		"website_title",
		"website_description",
		"website_keywords",
		"website_logo",
		"website_footer",
		"invite_introduce",
		"random_personas",
		"chat_models",
		"draw_models",
	}
)

func ConfigList(ctx context.Context) (result map[string]interface{}, err error) {
	dc := dao.Q.Config

	configs, err := dc.WithContext(ctx).Where(dc.Name.In(indexConfigs...)).Find()
	if err != nil {
		logs.Error("config get error: %v", err)
		return nil, err
	}

	// trans
	result = make(map[string]interface{})
	for _, c := range configs {
		result[c.Name] = c.Value
	}

	// notification
	dn := dao.Q.Notification
	notifications, err := dn.WithContext(ctx).Where(dn.Status.Eq(1)).Find()
	if err != nil {
		logs.Error("Notification get error: %v", err)
	} else {
		result["notifications"] = notifications
	}

	// personas
	dp := dao.Q.Persona
	personas, err := dp.WithContext(ctx).Where(dp.Status.Eq(1)).Find()
	if err != nil {
		logs.Error("Persona get error: %v", err)
	} else {
		result["random_personas"] = personas
	}

	// chat_models
	if _, found := result["chat_models"]; found {
		chatModels := result["chat_models"].(string)
		var chatModelsMap []map[string]string
		err := json.Unmarshal([]byte(chatModels), &chatModelsMap)
		if err != nil {
			logs.Error("chatModels json error: %v", err)
		} else {
			result["chat_models"] = chatModelsMap
		}
	}

	// draw models
	if _, found := result["draw_models"]; found {
		drawModels := result["draw_models"].(string)
		var drawModelsMap []map[string]string
		err := json.Unmarshal([]byte(drawModels), &drawModelsMap)
		if err != nil {
			logs.Error("drawModels json error: %v", err)
		} else {
			result["draw_models"] = drawModelsMap
		}
	}

	return
}
