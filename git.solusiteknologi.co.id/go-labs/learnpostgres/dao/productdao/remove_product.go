package productdao

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/gopostgres/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/jackc/pgx/v4"
)

type InputRemove struct {
	Tx            pgx.Tx
	AuditUserId   int64
	AuditDatetime string

	ProductId int64
}

func Remove(input InputRemove) (*tables.LearnProduct, error) {
	result := tables.LearnProduct{}

	err := gldb.SelectOneQTx(input.Tx, *gldb.NewQBuilder().
		Add(" DELETE FROM ", tables.LEARN_PRODUCT).
		Add(" WHERE product_id = :productId ").
		Add(" RETURNING ").
		Add("   product_id, product_name, price, ").
		Add("   active, active_datetime, non_active_datetime, create_datetime, ").
		Add("   update_datetime, create_user_id, update_user_id, version ").
		SetParam("productId", input.ProductId).
		Log("QUERY remove product: "),

		&result)

	if err != nil {
		return nil, errors.New(fmt.Sprint("error add product ", err))
	}

	return &result, nil
}
