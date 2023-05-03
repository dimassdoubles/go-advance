package productdao

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/gopostgres/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

type InputEdit struct {
	Tx            pgx.Tx
	AuditUserId   int64
	AuditDatetime string

	ProductId   int64
	ProductName string
	Active      string
	Price       decimal.Decimal
	Version     int64
}

func Edit(input InputEdit) (*tables.LearnProduct, error) {
	result := tables.LearnProduct{}

	err := gldb.SelectOneQTx(input.Tx, *gldb.NewQBuilder().
		Add(" UPDATE ", tables.LEARN_PRODUCT, " SET  ").
		AddSetNext("product_name", input.ProductName).
		AddSetNext("price", input.Price).
		AddPrepareUpdateAudit(input.AuditDatetime, input.AuditUserId).
		Add(" WHERE product_id = :productId ").
		AddFilterVersion(input.Version). // otomatis tambahkan filter version harus sesuai
		Add(" RETURNING ").
		Add("   product_id, product_name, price, ").
		Add("   active, active_datetime, non_active_datetime, create_datetime, ").
		Add("   update_datetime, create_user_id, update_user_id, version ").
		SetParam("productId", input.ProductId).
		SetParam("productName", input.ProductName).
		SetParam("price", input.Price).
		SetParam("active", input.Active).
		SetParam("datetime", input.AuditDatetime).
		SetParam("userId", input.AuditUserId).
		SetParam("version", 0).
		Log("QUERY edit product : "),

		&result)

	if err != nil {
		return nil, errors.New(fmt.Sprint("error add product ", err))
	}

	return &result, nil
}
