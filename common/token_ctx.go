package common

import "context"

func WithUserToken(ctx context.Context, token string) context.Context {
	if token == "" {
		return ctx
	}

	return context.WithValue(ctx, TokenHeaderName, token)
}

func UserTokenFromContext(ctx context.Context) string {
	token := ctx.Value(TokenHeaderName)

	return token.(string)
}
