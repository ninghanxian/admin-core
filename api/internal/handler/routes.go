// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	api "github.com/qmcloud/admin-core/api/internal/handler/api"
	authority "github.com/qmcloud/admin-core/api/internal/handler/authority"
	base "github.com/qmcloud/admin-core/api/internal/handler/base"
	captcha "github.com/qmcloud/admin-core/api/internal/handler/captcha"
	department "github.com/qmcloud/admin-core/api/internal/handler/department"
	dictionary "github.com/qmcloud/admin-core/api/internal/handler/dictionary"
	dictionarydetail "github.com/qmcloud/admin-core/api/internal/handler/dictionarydetail"
	emaillog "github.com/qmcloud/admin-core/api/internal/handler/emaillog"
	emailprovider "github.com/qmcloud/admin-core/api/internal/handler/emailprovider"
	menu "github.com/qmcloud/admin-core/api/internal/handler/menu"
	messagesender "github.com/qmcloud/admin-core/api/internal/handler/messagesender"
	oauthprovider "github.com/qmcloud/admin-core/api/internal/handler/oauthprovider"
	position "github.com/qmcloud/admin-core/api/internal/handler/position"
	publicuser "github.com/qmcloud/admin-core/api/internal/handler/publicuser"
	role "github.com/qmcloud/admin-core/api/internal/handler/role"
	smslog "github.com/qmcloud/admin-core/api/internal/handler/smslog"
	smsprovider "github.com/qmcloud/admin-core/api/internal/handler/smsprovider"
	task "github.com/qmcloud/admin-core/api/internal/handler/task"
	tasklog "github.com/qmcloud/admin-core/api/internal/handler/tasklog"
	token "github.com/qmcloud/admin-core/api/internal/handler/token"
	user "github.com/qmcloud/admin-core/api/internal/handler/user"
	"github.com/qmcloud/admin-core/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/core/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/core/init/job_database",
				Handler: base.InitJobDatabaseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/core/init/mcms_database",
				Handler: base.InitMcmsDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/role/create",
					Handler: role.CreateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/update",
					Handler: role.UpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/delete",
					Handler: role.DeleteRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/list",
					Handler: role.GetRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role",
					Handler: role.GetRoleByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: publicuser.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login_by_email",
				Handler: publicuser.LoginByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login_by_sms",
				Handler: publicuser.LoginBySmsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: publicuser.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register_by_email",
				Handler: publicuser.RegisterByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register_by_sms",
				Handler: publicuser.RegisterBySmsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/reset_password_by_email",
				Handler: publicuser.ResetPasswordByEmailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/reset_password_by_sms",
				Handler: publicuser.ResetPasswordBySmsHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/create",
					Handler: user.CreateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/update",
					Handler: user.UpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/delete",
					Handler: user.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user",
					Handler: user.GetUserByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/change_password",
					Handler: user.ChangePasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/perm",
					Handler: user.GetUserPermCodeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/profile",
					Handler: user.GetUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/profile",
					Handler: user.UpdateUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/menu/create",
					Handler: menu.CreateMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/update",
					Handler: menu.UpdateMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/delete",
					Handler: menu.DeleteMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/list",
					Handler: menu.GetMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/role/list",
					Handler: menu.GetMenuListByRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/content-data",
					Handler: menu.GetContentHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/popularlist",
					Handler: menu.GetPopularlistHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/captcha",
				Handler: captcha.GetCaptchaHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/captcha/email",
				Handler: captcha.GetEmailCaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/captcha/sms",
				Handler: captcha.GetSmsCaptchaHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/create",
					Handler: api.CreateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/update",
					Handler: api.UpdateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/delete",
					Handler: api.DeleteApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/list",
					Handler: api.GetApiListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api",
					Handler: api.GetApiByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/authority/api/create_or_update",
					Handler: authority.CreateOrUpdateApiAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/api/role",
					Handler: authority.GetApiAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/menu/create_or_update",
					Handler: authority.CreateOrUpdateMenuAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/menu/role",
					Handler: authority.GetMenuAuthorityHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/create",
					Handler: dictionary.CreateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/update",
					Handler: dictionary.UpdateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/delete",
					Handler: dictionary.DeleteDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary/list",
					Handler: dictionary.GetDictionaryListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary",
					Handler: dictionary.GetDictionaryByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/oauth/login",
				Handler: oauthprovider.OauthLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/oauth/login/callback",
				Handler: oauthprovider.OauthCallbackHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/create",
					Handler: oauthprovider.CreateOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/update",
					Handler: oauthprovider.UpdateOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/delete",
					Handler: oauthprovider.DeleteOauthProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider/list",
					Handler: oauthprovider.GetOauthProviderListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth_provider",
					Handler: oauthprovider.GetOauthProviderByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/token/create",
					Handler: token.CreateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/update",
					Handler: token.UpdateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/delete",
					Handler: token.DeleteTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/list",
					Handler: token.GetTokenListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token",
					Handler: token.GetTokenByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/logout",
					Handler: token.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/department/create",
					Handler: department.CreateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/update",
					Handler: department.UpdateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/delete",
					Handler: department.DeleteDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/list",
					Handler: department.GetDepartmentListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department",
					Handler: department.GetDepartmentByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/position/create",
					Handler: position.CreatePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/update",
					Handler: position.UpdatePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/delete",
					Handler: position.DeletePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/list",
					Handler: position.GetPositionListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position",
					Handler: position.GetPositionByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/dictionary_detail/create",
					Handler: dictionarydetail.CreateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary_detail/update",
					Handler: dictionarydetail.UpdateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary_detail/delete",
					Handler: dictionarydetail.DeleteDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary_detail/list",
					Handler: dictionarydetail.GetDictionaryDetailListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dictionary_detail",
					Handler: dictionarydetail.GetDictionaryDetailByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/dict/:name",
					Handler: dictionarydetail.GetDictionaryDetailByDictionaryNameHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/task/create",
					Handler: task.CreateTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/task/update",
					Handler: task.UpdateTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/task/delete",
					Handler: task.DeleteTaskHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/task/list",
					Handler: task.GetTaskListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/task",
					Handler: task.GetTaskByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/task_log/create",
					Handler: tasklog.CreateTaskLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/task_log/update",
					Handler: tasklog.UpdateTaskLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/task_log/delete",
					Handler: tasklog.DeleteTaskLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/task_log/list",
					Handler: tasklog.GetTaskLogListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/task_log",
					Handler: tasklog.GetTaskLogByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/email_log/create",
					Handler: emaillog.CreateEmailLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/email_log/update",
					Handler: emaillog.UpdateEmailLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/email_log/delete",
					Handler: emaillog.DeleteEmailLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/email_log/list",
					Handler: emaillog.GetEmailLogListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/email_log",
					Handler: emaillog.GetEmailLogByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/sms_log/create",
					Handler: smslog.CreateSmsLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sms_log/update",
					Handler: smslog.UpdateSmsLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sms_log/delete",
					Handler: smslog.DeleteSmsLogHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sms_log/list",
					Handler: smslog.GetSmsLogListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sms_log",
					Handler: smslog.GetSmsLogByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/sms_provider/create",
					Handler: smsprovider.CreateSmsProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sms_provider/update",
					Handler: smsprovider.UpdateSmsProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sms_provider/delete",
					Handler: smsprovider.DeleteSmsProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sms_provider/list",
					Handler: smsprovider.GetSmsProviderListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sms_provider",
					Handler: smsprovider.GetSmsProviderByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/email_provider/create",
					Handler: emailprovider.CreateEmailProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/email_provider/update",
					Handler: emailprovider.UpdateEmailProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/email_provider/delete",
					Handler: emailprovider.DeleteEmailProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/email_provider/list",
					Handler: emailprovider.GetEmailProviderListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/email_provider",
					Handler: emailprovider.GetEmailProviderByIdHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/sms/send",
					Handler: messagesender.SendSmsHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/email/send",
					Handler: messagesender.SendEmailHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
