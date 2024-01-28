package auth

import "context"

const (
	SessionCtxName = "session"
)

type Session struct {
	Username string
	ClientID string
}

func GetSession(ctx context.Context) Session {
	return ctx.Value(SessionCtxName).(Session)
}
