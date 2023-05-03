package test_test

import (
	"testing"

	"git.solusiteknologi.co.id/go-labs/gopostgres/dao/productdao"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glconstant"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func TestEditProduct(t *testing.T) {
	gltest.TestDb(t, func(tx pgx.Tx) error {
		product, err := productdao.Add(productdao.InputAdd{
			Tx:            tx,
			AuditUserId:   -1,
			AuditDatetime: glutil.DateTimeNow(),
			ProductName:   "Sirup Marjan",
			Price:         decimal.NewFromInt(45000),
		})
		if err != nil {
			return err
		}

		// otomatisasi cek apakah yg diinsert benar
		assert := gltest.NewAssert(t)

		// productId tidak dapat dipastikan karena selalu beda menggunakan sequence
		assert.AssertEquals("Sirup Marjan", product.ProductName)
		assert.AssertEquals(decimal.NewFromInt(45000), product.Price)
		assert.AssertEquals(glconstant.YES, product.Active)
		assert.AssertEquals(0, product.Version)

		logrus.Debug("ADDED product : ", goleafcore.NewOrEmpty(product).PrettyString())

		editedProduct, err := productdao.Edit(productdao.InputEdit{
			Tx:            tx,
			AuditUserId:   -1,
			AuditDatetime: glutil.DateTimeNow(),
			ProductName:   "Sirup Marjan Edited",
			Price:         decimal.NewFromInt(50000),
			ProductId:     product.ProductId,
			Active:        product.Active,
			Version:       product.Version,
		})
		if err != nil {
			return err
		}

		assert.AssertEquals(product.ProductId, editedProduct.ProductId)
		assert.AssertEquals("Sirup Marjan Edited", editedProduct.ProductName)
		assert.AssertEquals(decimal.NewFromInt(50000), editedProduct.Price)
		assert.AssertEquals(product.Version + 1, editedProduct.Version)

		logrus.Debug("EDITED product : ", goleafcore.NewOrEmpty(editedProduct).PrettyString())

		return nil
	}, func(assert *gltest.Assert, tx pgx.Tx) interface{} {
		glinit.InitLog(glinit.LogConfig{
			LogFile:  "log/gopostgres.log",
			LogLevel: glinit.DEFAULT_LOG_LEVEL,
		})

		return nil
	})
}
