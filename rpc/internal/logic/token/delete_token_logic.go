package token

import (
	"context"

	"github.com/qmcloud/admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/qmcloud/admin-common/i18n"

	"github.com/qmcloud/admin-core/rpc/ent/token"

	"github.com/qmcloud/admin-core/rpc/internal/svc"
	"github.com/qmcloud/admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/qmcloud/admin-core/rpc/types/core"
)

type DeleteTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTokenLogic {
	return &DeleteTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTokenLogic) DeleteToken(in *core.UUIDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Token.Delete().Where(token.IDIn(uuidx.ParseUUIDSlice(in.Ids)...)).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
