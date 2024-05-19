package repo

import (
	"context"
	"ucenter-srv/internal/model"
)

// MemberRepo domain操作的是接口
// 具体查询操作在DAO层实现
type MemberRepo interface {
	FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error)
	Save(ctx context.Context, mem *model.Member) error
	UpdateLoginCount(ctx context.Context, id int64, step int) error
	FindMemberById(ctx context.Context, memberId int64) (mem *model.Member, err error)
}
