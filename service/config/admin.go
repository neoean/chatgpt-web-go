package config

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func AdminConfigList(ctx context.Context) (configList []*model.Config, err error) {
	dc := dao.Q.Config
	configList, err = dc.WithContext(ctx).Find()
	if err != nil {
		logs.Error("config list error: %v", err)
		return nil, err
	}

	return
}

func ConfigUpdate(ctx context.Context, configMap map[string]string) error {
	dc := dao.Q.Config
	for cName, cValue := range configMap {
		find, err := dc.WithContext(ctx).Where(dc.Name).Find()
		if err != nil {
			logs.Error("config find name:%v error: %v", cName, err)
			return err
		}

		if len(find) == 0 {
			logs.Error("config find none, name: %v", cName)
			continue
		}

		resultInfo, err := dc.WithContext(ctx).Where(dc.Name.Eq(cName)).Update(dc.Value, cValue)
		if err != nil {
			logs.Error("config update error: %v", err)
			return err
		}
		if resultInfo.RowsAffected < 1 {
			logs.Error("config update fail: RowsAffected < 1")
			continue
		}
	}

	return nil
}
