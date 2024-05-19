package dao

import (
	"common/msdb"
	"common/msdb/gorms"
	"context"
	"errors"
	"gorm.io/gorm"
	"ucenter-srv/internal/model"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func NewMemberDao(db *msdb.MsDB) *MemberDao {
	return &MemberDao{
		conn: gorms.New(db.Conn),
	}
}

func (m MemberDao) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	//TODO implement me
	session := m.conn.Session(ctx)
	mem = model.NewMember() // 初始化 mem 为 model.Member 的指针
	err = session.Model(mem).Where("mobile_phone=?", phone).Limit(1).Take(mem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return mem, err
}

func (m MemberDao) Save(ctx context.Context, mem *model.Member) error {
	//TODO implement me
	session := m.conn.Session(ctx)
	err := session.Save(mem).Error
	return err
}

func (m MemberDao) UpdateLoginCount(ctx context.Context, id int64, step int) error {
	//TODO implement me
	panic("implement me")
}

func (m MemberDao) FindMemberById(ctx context.Context, memberId int64) (mem *model.Member, err error) {
	//TODO implement me
	panic("implement me")
}
