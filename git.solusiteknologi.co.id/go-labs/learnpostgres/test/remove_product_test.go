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

func TestRemoveProduct(t *testing.T) {
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

		removedProduct, err := productdao.Remove(productdao.InputRemove{
			Tx:            tx,
			AuditUserId:   -1,
			AuditDatetime: glutil.DateTimeNow(),
			ProductId:     product.ProductId,
		})
		if err != nil {
			return err
		}

		assert.AssertEquals(product.ProductId, removedProduct.ProductId)
		assert.AssertEquals(product.ProductName, removedProduct.ProductName)
		assert.AssertEquals(product.Price, removedProduct.Price)
		assert.AssertEquals(product.Active, removedProduct.Active)
		assert.AssertEquals(product.Version, removedProduct.Version)

		logrus.Debug("REMOVED product : ", goleafcore.NewOrEmpty(removedProduct).PrettyString())

		return nil
	}, func(assert *gltest.Assert, tx pgx.Tx) interface{} {
		glinit.InitLog(glinit.LogConfig{
			LogFile:  "log/gopostgres.log",
			LogLevel: glinit.DEFAULT_LOG_LEVEL,
		})

		return nil
	})
}
