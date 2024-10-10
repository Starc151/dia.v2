package ydb

import (
	"context"
	"fmt"
	"os"

	
	"github.com/ydb-platform/ydb-go-sdk/v3"
	yc "github.com/ydb-platform/ydb-go-yc"
)

type connected struct {
	db *ydb.Driver
	ctx	context.Context
	cancel context.CancelFunc
	dbName string
	err error
}

func (c *connected) connect() {
	c.dbName = "result_bolus"
	path := "pkgs/ydb/token/"
	f, _ := os.ReadFile(path + "dsn.txt")
	
	dsn := string(f)
	
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.db, c.err = ydb.Open(c.ctx, dsn,
		yc.WithServiceAccountKeyFileCredentials(path + "authorized_key.txt"),
	)
	if c.err != nil {
		c.err = fmt.Errorf("нет соединения")
	}
}
