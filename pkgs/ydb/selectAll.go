package ydb

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
)

type entity struct {
	Date_Time       time.Time
	Bolus           float32
	Current_Glucose float32
	Bread_Unit      float32
}

func (c *connected) SelectAll() ([]entity, error) {
	resYDB := []entity{}
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
				resYDB = append(resYDB, entity)
			}
			return res.Err()
		},
	)
	if c.err != nil {
		return resYDB, c.err
	}
	for i, j := 0, len(resYDB)-1; i < j; i, j = i+1, j-1 {
        resYDB[i], resYDB[j] = resYDB[j], resYDB[i]
    }
	rr := []string{}
	vv := reflect.ValueOf(resYDB)
	for k, v := range vv.va{
		rr = append(rr, "bnm")
	}

	date := resYDB[0].Date_Time.Format("02.01.2006")
	time := resYDB[0].Date_Time.Format("15:04:05")
	firstRes := []string{date, time}

	for k := range resYDB {
		if date != resYDB[k].Date_Time.Format("02.01.2006") {
			result = append(result, )
			fmt.Println(resYDB[k].Date_Time.Format("02.01.2006"))
			date = resYDB[k].Date_Time.Format("02.01.2006")
		}
	}
	// fmt.Println(date)
	return resYDB, nil
}