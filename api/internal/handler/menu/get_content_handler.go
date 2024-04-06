package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/qmcloud/admin-core/api/internal/logic/menu"
	"github.com/qmcloud/admin-core/api/internal/svc"
)

// swagger:route get /menu/content-data menu GetContent
//

//

//
// Responses:
//  200: BaseMsgResp

func GetContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetContentLogic(r.Context(), svcCtx)
		resp, err := l.GetContent()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
