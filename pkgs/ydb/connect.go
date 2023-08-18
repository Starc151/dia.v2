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
func (c *Connected) connect() error {
	path := "pkgs/ydb/token/"
	f, err := os.ReadFile(path + "dsn.txt")
	if err != nil {
		return fmt.Errorf("ошибка файла ключей")
	}
	dsn := string(f)
	
	c.ctx, c.cancel = context.WithCancel(context.Background())
	c.db, err = ydb.Open(c.ctx, dsn,
		yc.WithServiceAccountKeyFileCredentials(path + "authorized_key.txt"),
	)
	if err != nil {
		return fmt.Errorf("нет соединения")
	}
	return nil
}