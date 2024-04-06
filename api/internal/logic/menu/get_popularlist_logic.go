package menu

import (
	"context"

	"github.com/qmcloud/admin-core/api/internal/svc"
	"github.com/qmcloud/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPopularlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPopularlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPopularlistLogic {
	return &GetPopularlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPopularlistLogic) GetPopularlist() (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
