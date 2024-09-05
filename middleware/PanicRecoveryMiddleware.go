package middleware

import (
	"github.com/yanshicheng/opsx-common/handler/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func PanicRecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// 记录日志
				logx.Errorf("panic error: %+v", err)
				// panic 错误统一返回  errorx.ServerErr 类型错误
				httpx.ErrorCtx(r.Context(), w, errorx.ServerErr)
			}
		}()
		next(w, r)
	}
}
