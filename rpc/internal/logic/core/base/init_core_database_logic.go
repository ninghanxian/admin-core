package base

import (
	"context"

	"github.com/qmcloud/admin-core/rpc/internal/svc"
	"github.com/qmcloud/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitCoreDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitCoreDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitCoreDatabaseLogic {
	return &InitCoreDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitCoreDatabaseLogic) InitCoreDatabase(in *core.Empty) (*core.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &core.BaseResp{}, nil
}
