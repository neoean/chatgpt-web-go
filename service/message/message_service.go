package message

import (
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/common/types"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"chatgpt-web-new-go/service/gpt"
	"context"
)

func MessageList(ctx context.Context, uid int64) (result []*Message, err error) {
	dm := dao.Q.Message

	messages, err := dm.WithContext(ctx).Where(dm.UserID.Eq(uid)).Find()
	if err != nil {
		logs.Error("message find error: %v", err)
		return
	}

	pMsgMap := make(map[string][]*model.Message)
	for _, m := range messages {
		m := m
		pID := m.ParentMessageID

		pList, found := pMsgMap[pID]
		if !found {
			pList = []*model.Message{}
		}
		pList = append(pList, m)
		pMsgMap[pID] = pList
	}

	for pID, pList := range pMsgMap {
		head := pList[0]

		msgItem := &Message{
			ID:        pID,
			Path:      pID,
			Name:      head.Content,
			PersonaID: head.PersonaID,
		}

		var msgDatas []*MessageData
		for _, msgItem := range pList {
			msgDataItem := &MessageData{
				UserID:     msgItem.UserID,
				DateTime:   msgItem.CreateTime.Format(types.TimeFormatDate),
				Role:       msgItem.Role,
				Status:     "pass",
				Text:       msgItem.Content,
				PersonaID:  msgItem.PersonaID,
				PluginID:   msgItem.PluginID,
				PluginInfo: nil,
				RequestOptions: &RequestOptions{
					ParentMessageID: msgItem.ParentMessageID,
					Prompt:          msgItem.Content,
					Options: &Options{
						FrequencyPenalty: msgItem.FrequencyPenalty,
						MaxTokens:        msgItem.MaxTokens,
						Model:            msgItem.Model,
						PresencePenalty:  msgItem.PresencePenalty,
						Temperature:      msgItem.Temperature,
					},
				},
			}

			msgDatas = append(msgDatas, msgDataItem)
		}

		msgItem.Data = msgDatas

		result = append(result, msgItem)
	}

	return
}

func AdminMessageList(ctx context.Context) (result []*model.Message, err error) {
	dm := dao.Q.Message

	result, err = dm.WithContext(ctx).Where(dm.IsDelete.Eq(0)).Find()
	if err != nil {
		logs.Error("admin message find error: %v", err)
		return
	}

	return
}

func MessageAdd(ctx context.Context, uid int64, r *gpt.Request, msgList []*ChatProcessResponse) {
	if len(msgList) == 0 {
		return
	}
	msgHead := msgList[0]

	content := ""
	for _, m := range msgList {
		content += m.Content
	}

	// insert message
	msg := &model.Message{
		UserID:           uid,
		Content:          content,
		PersonaID:        types.InterfaceToInt64(r.PersonaId),
		Role:             msgHead.Role,
		FrequencyPenalty: int32(r.Options.FrequencyPenalty),
		MaxTokens:        int32(r.Options.MaxTokens),
		Model:            r.Options.Model,
		PresencePenalty:  int32(r.Options.PresencePenalty),
		Temperature:      int32(r.Options.Temperature),
		ParentMessageID:  r.ParentMessageId,
	}

	dm := dao.Q.Message
	err := dm.WithContext(ctx).Create(msg)
	if err != nil {
		logs.Error("message create error: %v", err)
	}
}
