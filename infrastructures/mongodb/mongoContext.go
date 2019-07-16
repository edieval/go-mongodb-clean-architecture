package mongodb

import (
	"context"
	"time"
)

func GetConnectionContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func GetQueryContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	return ctx
}
