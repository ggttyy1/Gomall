package utils

import (
	"context"
)

func GetUserIdfromCtx(ctx context.Context) int32 {
	// userid := ctx.Value(SessionUserId)
	id := ctx.Value(JWTUserId)

	if id.(int) <= 0 {
		return 0
	}
	return int32(id.(int))
}

func GetUserRolefromCtx(ctx context.Context) string {
	// userid := ctx.Value(SessionUserId)
	Role := ctx.Value(JWTRole)

	return Role.(string)
}
