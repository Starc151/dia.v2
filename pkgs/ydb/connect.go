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
}

func NewConnect(m func ()) error {
	connectDB := connected{}
	tt := 
	
	err :=  connectDB.insert(dbName, dataInsert)
	return err
}

// func NewConnect(dbName string, dataInsert map[string]float64) error {
// 	connectDB := connected{}
	
// 	err :=  connectDB.insert(dbName, dataInsert)
// 	return err
// }

func (c *connected) connect() error {
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




