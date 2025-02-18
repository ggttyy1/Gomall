package utils

import "context"

func GetUserIdfromCtx(ctx context.Context) int32 {
	userid := ctx.Value(SessionUserId)
	if userid == nil {
		return 0
	}
	return userid.(int32)
}
