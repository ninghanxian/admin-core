package menu

import (
	"context"
	"github.com/qmcloud/admin-common/i18n"

	"github.com/qmcloud/admin-core/api/internal/svc"
	"github.com/qmcloud/admin-core/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentLogic {
	return &GetContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetContentLogic) GetContent() (resp *types.ContentListResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.ContentListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data = []types.ResContentResp{{X: "2024-04-04", Y: 5800}, {X: "2024-04-05", Y: 2780}, {X: "2024-04-06", Y: 3950}, {X: "2024-04-07", Y: 6800}, {X: "2024-04-08", Y: 2800}}
	return resp, nil
}
