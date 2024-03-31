package ydb

import (
	"context"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
)
type Table struct{
    Date    string
    Time    string
    Bolus   float32
    Glucose float32
    Xe      float32
} 

func Select() []Table {
    loc, _ := time.LoadLocation("Europe/GMT")
    time.Local = loc
    dateTime := time.Now()
    resYDB := []Table{}
    
	db, ctx, cancel := connect()
	defer cancel()
	defer db.Close(ctx)

    db.Table().Do(ctx, func(ctx context.Context, s table.Session) (err error) {
        _, res, _ := s.Execute(ctx, table.DefaultTxControl(), "SELECT * FROM res;", nil)

        defer res.Close()

        if err = res.NextResultSetErr(ctx); err != nil {
            return err
        }
        var Table Table
        for res.NextRow() {
            res.ScanNamed(
                named.OptionalWithDefault("date_time", &dateTime),
                named.OptionalWithDefault("bolus", &Table.Bolus	),
                named.OptionalWithDefault("glucose", &Table.Glucose),
                named.OptionalWithDefault("xe", &Table.Xe),
            )
            Table.Date = dateTime.Format("02 Jan 06")
            Table.Time = dateTime.Format("15:04")
            resYDB = append(resYDB, Table)
        }
        return res.Err()
    })
    for i, j := 0, len(resYDB)-1; i < j; i, j = i+1, j-1 {
        resYDB[i], resYDB[j] = resYDB[j], resYDB[i]
    }
    return resYDB
}