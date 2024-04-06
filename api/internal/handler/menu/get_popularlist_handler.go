package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/qmcloud/admin-core/api/internal/logic/menu"
	"github.com/qmcloud/admin-core/api/internal/svc"
)

// swagger:route get /menu/popularlist menu GetPopularlist
//

//

//
// Responses:
//  200: BaseMsgResp

func GetPopularlistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetPopularlistLogic(r.Context(), svcCtx)
		resp, err := l.GetPopularlist()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
