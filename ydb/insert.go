package ydb

import (
	"context"
	"fmt"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func Insert(dataInsert map[string]float64) error {
	cRes, err := connect()
    if err != nil {
        return err
    }
    defer cRes.cancel()
	defer cRes.db.Close(cRes.ctx)
	dateTime := uint32(time.Now().Unix())+10800 //поправка на часовой пояс Мск
	dbName := "result_bolus"
	err = cRes.db.Table().DoTx(cRes.ctx,
		func(ctx context.Context, tx table.TransactionActor) (err error) {
			res, err := tx.Execute(ctx, `
			DECLARE $Date_Time AS Datetime;
			DECLARE $Bolus AS Float;
			DECLARE $Current_Glucose AS Float;
			DECLARE $Bread_Unit AS Float;
			INSERT INTO ` + dbName + ` ( Date_Time, Bolus, Current_Glucose, Bread_Unit )
			VALUES ( $Date_Time, $Bolus, $Current_Glucose, $Bread_Unit );
		`,
				table.NewQueryParameters(
					table.ValueParam("$Date_Time", types.DatetimeValue(dateTime)),
					table.ValueParam("$Bolus", types.FloatValue(float32(dataInsert["Bolus"]))),
					table.ValueParam("$Current_Glucose", types.FloatValue(float32(dataInsert["Current_Glucose"]))),
					table.ValueParam("$Bread_Unit", types.FloatValue(float32(dataInsert["Bread_Unit"]))),
				),
			)
			if err != nil {
				return err
			}
			if err = res.Err(); err != nil {
				return err
			}
			return res.Close()
		}, table.WithIdempotent(),
	)
	if err != nil {
		return fmt.Errorf("RECORDING FAIlED")
	}
	return nil
}
