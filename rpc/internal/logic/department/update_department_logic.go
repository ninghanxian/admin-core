package department

import (
	"context"

	"github.com/qmcloud/admin-common/utils/pointy"

	"github.com/qmcloud/admin-core/rpc/internal/svc"
	"github.com/qmcloud/admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/qmcloud/admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/qmcloud/admin-common/i18n"
)

type UpdateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDepartmentLogic {
	return &UpdateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDepartmentLogic) UpdateDepartment(in *core.DepartmentInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.Department.UpdateOneID(*in.Id).
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilSort(in.Sort).
		SetNotNilName(in.Name).
		SetNotNilAncestors(in.Ancestors).
		SetNotNilLeader(in.Leader).
		SetNotNilPhone(in.Phone).
		SetNotNilEmail(in.Email).
		SetNotNilRemark(in.Remark).
		SetNotNilParentID(in.ParentId).
		Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
