package ydb

import (
	"context"
	"fmt"
	"os"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	yc "github.com/ydb-platform/ydb-go-yc"
)

type Connected struct {
	db *ydb.Driver
	ctx	context.Context
	cancel context.CancelFunc
}
func connect() (*Connected, error) {
	cRes := &Connected{}


	f, err := os.ReadFile("ydb/token/dsn.txt")
	if err != nil {
		return cRes, fmt.Errorf("INVALID KEY")
	}
	dsn := string(f)
	
	cRes.ctx, cRes.cancel = context.WithCancel(context.Background())
	cRes.db, err = ydb.Open(cRes.ctx, dsn,
		yc.WithServiceAccountKeyFileCredentials("ydb/token/authorized_key.txt"),
	)
	if err != nil {
		return cRes, fmt.Errorf("NO CONNECT")
	}
	return cRes, nil
}
