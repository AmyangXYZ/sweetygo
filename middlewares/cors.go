package middlewares

import (
	"net/http"
	"strings"

	"github.com/AmyangXYZ/sweetygo"
)

// CORSOpt is options of CORS middleware.
type CORSOpt struct {
	Skipper      Skipper
	AllowOrigins []string
	AllowMethods []string
}

// CORS returns a Cross-Origin Resource Sharing (CORS) middleware.
func CORS(opt CORSOpt) sweetygo.HandlerFunc {
	return func(ctx *sweetygo.Context) error {
		if opt.Skipper == nil {
			opt.Skipper = DefaultSkipper
		}
		if opt.AllowOrigins == nil {
			opt.AllowOrigins = []string{"*"}
		}
		if opt.AllowMethods == nil {
			opt.AllowMethods = []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete}
		}

		if opt.Skipper(ctx) == true {
			ctx.Next()
			return nil
		}
		allowOrigins := strings.Join(opt.AllowOrigins, ",")
		allowMethods := strings.Join(opt.AllowMethods, ",")

		ctx.Resp.Header().Set("Access-Control-Allow-Origin", allowOrigins)
		ctx.Resp.Header().Set("Access-Control-Allow-Methods", allowMethods)
		ctx.Next()
		return nil
	}
}
