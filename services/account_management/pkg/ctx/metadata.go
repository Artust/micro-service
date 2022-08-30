package ctxutil

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/metadata"
)

func ExtractHeaders(ctx context.Context) metadata.MD {
	headers, _ := metadata.FromIncomingContext(ctx)
	return headers
}

func ExtractUid(ctx context.Context) (int64, error) {
	md := ExtractHeaders(ctx)
	uidS := md.Get("uid")[0]
	if uidS == "" {
		return 0, fmt.Errorf("user not authorized")
	}
	uid, err := strconv.ParseInt(uidS, 10, 64)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func ExtractResetToken(ctx context.Context) (string, error) {
	md := ExtractHeaders(ctx)
	token := md.Get("reset_password_token")[0]
	if token == "" {
		return "", fmt.Errorf("user not authorized")
	}
	return token, nil
}
