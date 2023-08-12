package ydb

import (
	"context"
	"fmt"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
)

type entity struct {
	Date_Time       time.Time
	Bolus           float32
	Current_Glucose float32
	Bread_Unit      float32
}

func (c *Connected) SelectAll(dbName string) ([]entity, error) {
	resYDB := []entity{}
	err := c.connect()
	if err != nil {
		return nil, err
	}
	defer c.cancel()
	defer c.db.Close(c.ctx)

	err = c.db.Table().Do(c.ctx, func(ctx context.Context, s table.Session) (err error) {
		_, res, err := s.Execute(ctx, table.DefaultTxControl(), "SELECT * FROM "+dbName+";", nil)
		if err != nil {
			return fmt.Errorf("ошибка запроса")
		}
		defer res.Close()

		if err = res.NextResultSetErr(ctx); err != nil {
			return err
		}
		var entity entity
		for res.NextRow() {
			res.ScanNamed(
				named.OptionalWithDefault("Date_Time", &entity.Date_Time),
				named.OptionalWithDefault("Bolus", &entity.Bolus),
				named.OptionalWithDefault("Current_Glucose", &entity.Current_Glucose),
				named.OptionalWithDefault("Bread_Unit", &entity.Bread_Unit),
			)
			resYDB = append(resYDB, entity)
		}
		return res.Err()
	})
	if err != nil {
		return resYDB, err
	}
	return resYDB, nil
}
