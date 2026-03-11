package nablesimulator

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/threatmate/restfulwrapper"
)

type ContextKey string

const (
	ContextKeyCurrentUser ContextKey = "currentUser"
)

func init() {
	restfulwrapper.Register("custom.currentUser", func(apiTagValue string, field reflect.StructField, info *restfulwrapper.RestfulFunctionInfo) (restfulwrapper.InputFieldFunction, error) {
		switch field.Type.String() {
		case "nablesimulator.APIUser":
		case "*nablesimulator.APIUser":
		default:
			return nil, fmt.Errorf("bad type for field %s: %s", field.Name, field.Type.String())
		}

		return func(v reflect.Value, req *restful.Request, metadataValue reflect.Value) error {
			ctx := req.Request.Context()

			contextValue := ctx.Value(ContextKeyCurrentUser)
			if contextValue == nil {
				return restfulwrapper.NewAPIResponseError(http.StatusUnauthorized, "Unauthorized")
			}

			currentUser := contextValue.(*APIUser)
			switch v.Interface().(type) {
			case APIUser:
				v.Set(reflect.ValueOf(*currentUser))
			case *APIUser:
				v.Set(reflect.ValueOf(currentUser))
			default:
				return restfulwrapper.NewAPIResponseError(http.StatusInternalServerError, fmt.Sprintf("Bad type for field %s", field.Name))
			}
			return nil
		}, nil
	})
}

func handleAuthentication(universe *Universe) func(builder *restful.RouteBuilder) {
	return func(builder *restful.RouteBuilder) {
		builder.Filter(func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
			authorization := req.Request.Header.Get("Authorization")
			slog.InfoContext(req.Request.Context(), "Found authorization header", "Authorization", authorization)
			if strings.HasPrefix(authorization, "Bearer ") {
				apiKey := strings.TrimPrefix(authorization, "Bearer ")
				slog.InfoContext(req.Request.Context(), "Found API key", "API Key", apiKey)
				for _, apiUser := range universe.APIUsers {
					for _, accessToken := range apiUser.accessTokens {
						if accessToken.Token == apiKey && accessToken.ExpiresAt.After(time.Now()) {
							slog.InfoContext(req.Request.Context(), "Found API user", "API User", apiUser.Username)
							ctx := req.Request.Context()
							ctx = context.WithValue(ctx, ContextKeyCurrentUser, apiUser)
							req.Request = req.Request.WithContext(ctx)
							break
						}
					}
				}
			}

			chain.ProcessFilter(req, resp)
		})
	}
}
