package auth

import (
	"bytes"
	"context"
	"encoding/json"
	auth_service "github.com/fidesy-pay/facade/pkg/auth-service"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"
)

const (
	usernameCtxName = "username"
)

var noNeedAuth = []string{"Login", "SignUp"}

func Auth(authClient auth_service.AuthServiceClient) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body := GraphQLRequestBody{}
			data, _ := io.ReadAll(r.Body)
			r.Body.Close()
			r.Body = io.NopCloser(bytes.NewBuffer(data))

			if string(data) == "" {
				next.ServeHTTP(w, r)
				return
			}

			err := json.Unmarshal(data, &body)
			if err != nil {
				log.Printf("json.Unmarhsal: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if strings.Contains(*body.Query, "IntrospectionQuery") {
				next.ServeHTTP(w, r)
				return
			}

			if body.OperationName == nil {
				log.Println("black swan, operation name is nil")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if slices.Contains(noNeedAuth, *body.OperationName) {
				rw := &CapturingResponseWriter{
					ResponseWriter: w,
				}
				next.ServeHTTP(rw, r)
				return

			}

			authToken := extractAuthorizationToken(r.Header)

			authResp, err := authClient.ValidateToken(r.Context(), &auth_service.ValidateTokenRequest{
				Token: authToken,
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				body, _ := json.Marshal(map[string]string{"error": err.Error()})
				w.Write(body)
				return
			}

			ctx := context.WithValue(r.Context(), usernameCtxName, authResp.GetUsername())

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func extractAuthorizationToken(headers http.Header) string {
	values := strings.Split(headers.Get("Cookie"), ";")
	if len(values) == 0 {
		return ""
	}

	for _, val := range values {
		splitVal := strings.Split(val, "=")
		if len(splitVal) < 2 {
			continue
		}

		if strings.TrimSpace(strings.ToLower(splitVal[0])) == "authorization" {
			return splitVal[1]
		}
	}

	return ""
}
