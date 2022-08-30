package util

import (
	"context"
	"fmt"
)

type key int

const (
	uid key = iota
)

func ExtractUid(ctx context.Context) (int64, error) {
	ret, ok := ctx.Value(uid).(int64)
	if !ok {
		return 0, fmt.Errorf("user not authorized")
	}
	return ret, nil
}
