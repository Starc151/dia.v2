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

func (c *connected) SelectAll() ([][]string, error) {
	defer c.cancel()
	defer c.db.Close(c.ctx)
	result := [][]string{}
	c.err = c.db.Table().Do(c.ctx,
		func(ctx context.Context, s table.Session) (err error) {
			_, res, err := s.Execute(ctx, table.DefaultTxControl(), "SELECT * FROM "+c.dbName+";", nil)
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
				resList := []string{
					entity.Date_Time.Format("02.01.2006"),
					entity.Date_Time.Format("15:04"),
					f32toStr(entity.Current_Glucose, "Glucose"),
					f32toStr(entity.Bread_Unit, "Bread unit"),
					f32toStr(entity.Bolus, "Bolus"),
				}
				result = append(result, resList)
			}
			return res.Err()
		},
	)
	if c.err != nil {
		return nil, c.err
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
	return result, nil
}

func f32toStr(num float32, desc string) string {
	return fmt.Sprintf("%s: %.1f", desc, num)
}