package productdao

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/gopostgres/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glconstant"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

type InputAdd struct {
    Tx            pgx.Tx
    AuditUserId   int64
    AuditDatetime string
 
    ProductName string
    Price       decimal.Decimal
}

func Add(input InputAdd) (*tables.LearnProduct, error) {
    result := tables.LearnProduct{}
 
    err := gldb.SelectOneQTx(input.Tx, *gldb.NewQBuilder().
        Add(" INSERT INTO ", tables.LEARN_PRODUCT, " (  ").
        Add("   product_name, price, ").
        Add("   active, active_datetime, non_active_datetime, create_datetime, ").
        Add("   update_datetime, create_user_id, update_user_id, version ").
        Add(" ) VALUES ( ").
        Add("   :productName, :price, ").
        Add("   :active, :datetime , :datetime , :datetime , ").
        Add("   :datetime, :userId, :userId, :version ").
        Add(" ) RETURNING ").
        Add("   product_id, product_name, price, ").
        Add("   active, active_datetime, non_active_datetime, create_datetime, ").
        Add("   update_datetime, create_user_id, update_user_id, version ").
        SetParam("productName", input.ProductName).
        SetParam("price", input.Price).
        SetParam("active", glconstant.YES).
        SetParam("datetime", input.AuditDatetime).
        SetParam("userId", input.AuditUserId).
        SetParam("version", 0),
 
        &result)
 
    if err != nil {
        return nil, errors.New(fmt.Sprint("error add product ", err))
    }
 
    return &result, nil
}
