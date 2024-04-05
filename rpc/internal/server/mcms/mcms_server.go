// Code generated by goctl. DO NOT EDIT.
// Source: core.proto

package server

import (
	"context"

	"github.com/qmcloud/admin-core/rpc/internal/logic/mcms/email"
	"github.com/qmcloud/admin-core/rpc/internal/logic/mcms/emaillog"
	"github.com/qmcloud/admin-core/rpc/internal/logic/mcms/emailprovider"
	"github.com/qmcloud/admin-core/rpc/internal/logic/mcms/sms"
	"github.com/qmcloud/admin-core/rpc/internal/logic/mcms/smslog"
	"github.com/qmcloud/admin-core/rpc/internal/logic/mcms/smsprovider"
	"github.com/qmcloud/admin-core/rpc/internal/svc"
	"github.com/qmcloud/admin-core/rpc/types/core"
)

type McmsServer struct {
	svcCtx *svc.ServiceContext
	core.UnimplementedMcmsServer
}

func NewMcmsServer(svcCtx *svc.ServiceContext) *McmsServer {
	return &McmsServer{
		svcCtx: svcCtx,
	}
}

func (s *McmsServer) SendEmail(ctx context.Context, in *core.EmailInfo) (*core.BaseUUIDResp, error) {
	l := email.NewSendEmailLogic(ctx, s.svcCtx)
	return l.SendEmail(in)
}

// EmailLog management
func (s *McmsServer) CreateEmailLog(ctx context.Context, in *core.EmailLogInfo) (*core.BaseUUIDResp, error) {
	l := emaillog.NewCreateEmailLogLogic(ctx, s.svcCtx)
	return l.CreateEmailLog(in)
}

func (s *McmsServer) UpdateEmailLog(ctx context.Context, in *core.EmailLogInfo) (*core.BaseResp, error) {
	l := emaillog.NewUpdateEmailLogLogic(ctx, s.svcCtx)
	return l.UpdateEmailLog(in)
}

func (s *McmsServer) GetEmailLogList(ctx context.Context, in *core.EmailLogListReq) (*core.EmailLogListResp, error) {
	l := emaillog.NewGetEmailLogListLogic(ctx, s.svcCtx)
	return l.GetEmailLogList(in)
}

func (s *McmsServer) GetEmailLogById(ctx context.Context, in *core.UUIDReq) (*core.EmailLogInfo, error) {
	l := emaillog.NewGetEmailLogByIdLogic(ctx, s.svcCtx)
	return l.GetEmailLogById(in)
}

func (s *McmsServer) DeleteEmailLog(ctx context.Context, in *core.UUIDsReq) (*core.BaseResp, error) {
	l := emaillog.NewDeleteEmailLogLogic(ctx, s.svcCtx)
	return l.DeleteEmailLog(in)
}

// EmailProvider management
func (s *McmsServer) CreateEmailProvider(ctx context.Context, in *core.EmailProviderInfo) (*core.BaseIDResp, error) {
	l := emailprovider.NewCreateEmailProviderLogic(ctx, s.svcCtx)
	return l.CreateEmailProvider(in)
}

func (s *McmsServer) UpdateEmailProvider(ctx context.Context, in *core.EmailProviderInfo) (*core.BaseResp, error) {
	l := emailprovider.NewUpdateEmailProviderLogic(ctx, s.svcCtx)
	return l.UpdateEmailProvider(in)
}

func (s *McmsServer) GetEmailProviderList(ctx context.Context, in *core.EmailProviderListReq) (*core.EmailProviderListResp, error) {
	l := emailprovider.NewGetEmailProviderListLogic(ctx, s.svcCtx)
	return l.GetEmailProviderList(in)
}

func (s *McmsServer) GetEmailProviderById(ctx context.Context, in *core.IDReq) (*core.EmailProviderInfo, error) {
	l := emailprovider.NewGetEmailProviderByIdLogic(ctx, s.svcCtx)
	return l.GetEmailProviderById(in)
}

func (s *McmsServer) DeleteEmailProvider(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := emailprovider.NewDeleteEmailProviderLogic(ctx, s.svcCtx)
	return l.DeleteEmailProvider(in)
}

func (s *McmsServer) SendSms(ctx context.Context, in *core.SmsInfo) (*core.BaseUUIDResp, error) {
	l := sms.NewSendSmsLogic(ctx, s.svcCtx)
	return l.SendSms(in)
}

// SmsLog management
func (s *McmsServer) CreateSmsLog(ctx context.Context, in *core.SmsLogInfo) (*core.BaseUUIDResp, error) {
	l := smslog.NewCreateSmsLogLogic(ctx, s.svcCtx)
	return l.CreateSmsLog(in)
}

func (s *McmsServer) UpdateSmsLog(ctx context.Context, in *core.SmsLogInfo) (*core.BaseResp, error) {
	l := smslog.NewUpdateSmsLogLogic(ctx, s.svcCtx)
	return l.UpdateSmsLog(in)
}

func (s *McmsServer) GetSmsLogList(ctx context.Context, in *core.SmsLogListReq) (*core.SmsLogListResp, error) {
	l := smslog.NewGetSmsLogListLogic(ctx, s.svcCtx)
	return l.GetSmsLogList(in)
}

func (s *McmsServer) GetSmsLogById(ctx context.Context, in *core.UUIDReq) (*core.SmsLogInfo, error) {
	l := smslog.NewGetSmsLogByIdLogic(ctx, s.svcCtx)
	return l.GetSmsLogById(in)
}

func (s *McmsServer) DeleteSmsLog(ctx context.Context, in *core.UUIDsReq) (*core.BaseResp, error) {
	l := smslog.NewDeleteSmsLogLogic(ctx, s.svcCtx)
	return l.DeleteSmsLog(in)
}

// SmsProvider management
func (s *McmsServer) CreateSmsProvider(ctx context.Context, in *core.SmsProviderInfo) (*core.BaseIDResp, error) {
	l := smsprovider.NewCreateSmsProviderLogic(ctx, s.svcCtx)
	return l.CreateSmsProvider(in)
}

func (s *McmsServer) UpdateSmsProvider(ctx context.Context, in *core.SmsProviderInfo) (*core.BaseResp, error) {
	l := smsprovider.NewUpdateSmsProviderLogic(ctx, s.svcCtx)
	return l.UpdateSmsProvider(in)
}

func (s *McmsServer) GetSmsProviderList(ctx context.Context, in *core.SmsProviderListReq) (*core.SmsProviderListResp, error) {
	l := smsprovider.NewGetSmsProviderListLogic(ctx, s.svcCtx)
	return l.GetSmsProviderList(in)
}

func (s *McmsServer) GetSmsProviderById(ctx context.Context, in *core.IDReq) (*core.SmsProviderInfo, error) {
	l := smsprovider.NewGetSmsProviderByIdLogic(ctx, s.svcCtx)
	return l.GetSmsProviderById(in)
}

func (s *McmsServer) DeleteSmsProvider(ctx context.Context, in *core.IDsReq) (*core.BaseResp, error) {
	l := smsprovider.NewDeleteSmsProviderLogic(ctx, s.svcCtx)
	return l.DeleteSmsProvider(in)
}
