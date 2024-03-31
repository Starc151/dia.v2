package ydb

import (
	"context"
	"fmt"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func (c *Connected) Insert(dataInsert map[string]float64) error {
	dbName := "result_bolus"
	_, zone := time.Now().Zone()
	tnu := int(time.Now().Unix())
	timeNow := uint32(tnu + zone)

	err := c.connect()
    if err != nil {
        return err 
    }
	defer c.cancel()
	defer c.db.Close(c.ctx)

	err = c.db.Table().DoTx(c.ctx,
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
					table.ValueParam("$Date_Time", types.DatetimeValue(timeNow)),
					table.ValueParam("$Bolus", types.FloatValue(float32(dataInsert["bolus"]))),
					table.ValueParam("$Current_Glucose", types.FloatValue(float32(dataInsert["glucose"]))),
					table.ValueParam("$Bread_Unit", types.FloatValue(float32(dataInsert["bUnit"]))),
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
		return fmt.Errorf("ошибка записи")
	}
	return nil
}
