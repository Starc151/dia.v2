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
    Current_Glucose float32
    Bread_Unit      float32
} 

func Select() ([]Table, error ){
    loc, _ := time.LoadLocation("Europe/GMT")
    time.Local = loc
    dateTime := time.Now()
    resYDB := []Table{}
    dbName := "result_bolus"
	cRes, err := connect()
    if err != nil {
        return resYDB, err
    }
    defer cRes.cancel()
	defer cRes.db.Close(cRes.ctx)

    cRes.db.Table().Do(cRes.ctx, func(ctx context.Context, s table.Session) (err error) {
        _, res, err := s.Execute(ctx, table.DefaultTxControl(), "SELECT * FROM " + dbName + ";", nil)
        if err != nil {
            return err
        }
        defer res.Close()

        if err = res.NextResultSetErr(ctx); err != nil {
            return err
        }
        var Table Table
        for res.NextRow() {
            res.ScanNamed(
                named.OptionalWithDefault("Date_Time", &dateTime),
                named.OptionalWithDefault("Bolus", &Table.Bolus	),
                named.OptionalWithDefault("Current_Glucose", &Table.Current_Glucose),
                named.OptionalWithDefault("Bread_Unit", &Table.Bread_Unit),
            )
            Table.Date = dateTime.Format("02 Jan 06")
            Table.Time = dateTime.Format("15:04")
            resYDB = append(resYDB, Table)
        }
        return res.Err()
    })
    return resYDB, nil
}