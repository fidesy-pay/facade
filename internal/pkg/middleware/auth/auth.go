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
	UsernameCtxName = "username"
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

			ctx := context.WithValue(r.Context(), UsernameCtxName, authResp.GetUsername())

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func extractAuthorizationToken(headers http.Header) string {
	authToken := strings.Split(
		headers.Get("Authorization"), " ",
	)

	if len(authToken) == 2 {
		return authToken[1]
	}

	cookies := headers.Get("Cookie")
	for _, cookie := range strings.Split(cookies, ";") {
		parts := strings.Split(cookie, "=")
		if len(parts) < 2 {
			continue
		}

		if strings.ToLower(strings.TrimSpace(parts[0])) == "authorization" {
			return parts[1]
		}
	}

	return ""
}
