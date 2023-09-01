package persona

import (
	"chatgpt-web-new-go/common/bizError"
	"chatgpt-web-new-go/common/logs"
	"chatgpt-web-new-go/dao"
	"chatgpt-web-new-go/model"
	"context"
)

func List(ctx context.Context, page, size int) (list []*model.Persona, count int64, err error) {
	du := dao.Q.Persona

	list, count, err = du.WithContext(ctx).Where(du.IsDelete.Eq(0)).FindByPage((page-1)*size, size)
	if err != nil {
		logs.Error("Persona list error: %v", err)
		return nil, 0, err
	}
	return
}

func Delete(ctx context.Context, id int64) error {
	du := dao.Q.Persona

	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(id)).Update(du.IsDelete, 1)
	if err != nil {
		logs.Error("Persona delete error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("Persona delete fail: RowsAffected < 1")
		return bizError.CommonDeleteError
	}

	return nil
}

func Add(ctx context.Context, n *model.Persona) error {
	dn := dao.Q.Persona

	err := dn.WithContext(ctx).Create(n)
	if err != nil {
		logs.Error("Persona create error: %v", err)
		return err
	}

	return nil
}
func Update(ctx context.Context, u *model.Persona) error {
	du := dao.Q.Persona
	resultInfo, err := du.WithContext(ctx).Where(du.ID.Eq(u.ID)).Updates(u)
	if err != nil {
		logs.Error("Persona update error: %v", err)
		return err
	}
	if resultInfo.RowsAffected < 1 {
		logs.Error("Persona update fail: RowsAffected < 1")
		return bizError.CommonUpdateError
	}

	return nil
}
