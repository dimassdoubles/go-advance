package tables

import (
	"git.solusiteknologi.co.id/goleaf/goleafcore/glentity"
	"github.com/shopspring/decimal"
)

type LearnProduct struct {
	ProductId   int64           `json:"productId"`
	ProductName string          `json:"productName"`
	Price       decimal.Decimal `json:"price"`

	glentity.MasterEntity

	// Extends glentity.MasterEntity untuk langsung menambahkan kolom :
	//      active, active_datetime, non_active_datetime, create_datetime,
	//      update_datetime, create_user_id, update_user_id, version

	// Extends glentity.BaseEntity jika hanya kolom :
	//      create_datetime, update_datetime, create_user_id, update_user_id, version
}
