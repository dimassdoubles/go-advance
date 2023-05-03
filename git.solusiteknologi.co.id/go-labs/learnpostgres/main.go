package main

import (
	"errors"
	"fmt"

	"git.solusiteknologi.co.id/go-labs/gopostgres/dao/productdao"
	"git.solusiteknologi.co.id/go-labs/gopostgres/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glconstant"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

func countProduct(trx pgx.Tx) (int64, error) {
	var count int64
	err := gldb.SelectOneQTx(trx, *gldb.NewQBuilder().
		Add("SELECT COUNT(1) FROM learn_product").
		Log("QUERY count product : "), &count)
	if err != nil {
		return count, errors.New(fmt.Sprint("error query count product ", err))
	}

	return count, nil
}

func getProducts(trx pgx.Tx, keyword, active string, limit, offset int64) ([]*tables.LearnProduct, error) {
	results := make([]*tables.LearnProduct, 0)

	// dokumentasi query builder ( QBuilder ) sudah ada di tiap method, bisa dicoba2 sendiri
	// setiap query dalam QBuilder bisa memasukkan :paramName
	// untuk nantinya bisa diset nilainya menggunakan SetParam("paramName", value)
	err := gldb.SelectQTx(trx, *gldb.NewQBuilder().
		Add(" SELECT A.product_id, A.product_name, A.price, ").
		Add("   A.active, A.active_datetime, A.non_active_datetime, A.create_datetime, ").
		Add("   A.update_datetime, A.create_user_id, A.update_user_id, A.version  ").
		Add(" FROM ", tables.LEARN_PRODUCT, " A ").
		Add(" WHERE true ").
		AddILike(" AND ", "keyword", keyword, "A.product_name").
		AddIfNotEmpty(active, " AND A.active = :active ").
		Add(" ORDER BY A.product_name ASC ").
		Add(" LIMIT :limit OFFSET :offset ").
		SetParam("active", active).
		SetParam("limit", limit).
		SetParam("offset", offset).
		Log("QUERY getProduct: "), &results)
	// .Log(...) digunakan untuk menampilkan hasil akhir query pada Log console
	if err != nil {
		return results, errors.New(fmt.Sprint("error query get product ", err))
	}

	return results, nil
}

func main() {
	// init log
	glinit.InitLog(glinit.LogConfig{
		LogFile:  "log/gopostgres.log",
		LogLevel: glinit.DEFAULT_LOG_LEVEL,
	})

	glinit.InitDb(glinit.DbConfig{
		User:              "sts",
		Password:          "Awesome123!",
		Port:              14555,
		Host:              "localhost",
		Name:              "gopostgresdb",
		ApplicationName:   "LearnGoPostegres",
		PoolMaxConnection: 1,
		PoolMinConnection: 1,
	})

	err := gldb.BeginTrx(func(trx pgx.Tx) error {
		// parameter "trx" adalah
		// semua koneksi yang ingin dijalankan dalam 1 transaksi yang sama harus menggunakan nilai trx yan sama

		// sample penggunaan selectOne
		count, err := countProduct(trx)
		if err != nil {
			return err
		}
		logrus.Debug("Result count product : ", count)

		// productList, err := getProducts(trx, keyword, active, limit, offset)
		// if err != nil {
		//	return err
		// }

		productList, err := getProducts(trx, "kecap", glconstant.YES, 100, 0)
		if err != nil {
			return err
		}

		// hanya sekedar debug bisa gunakan koding berikut
		logrus.Debug("Result products : ", goleafcore.Dto{
			"productList": productList,
		}.PrettyString())

		inputAdd := productdao.InputAdd{
			Tx:            trx,
			AuditUserId:   52,
			AuditDatetime: glutil.DateTimeNow(),
			ProductName:   "Kopi Kapal Api",
			Price:         decimal.NewFromFloat(2500.01),
		}
		learnProduct, err := productdao.Add(inputAdd)
		if err != nil {
			return err
		}

		logrus.Debug("Result products : ", goleafcore.Dto{
			"learnProduct": learnProduct,
		}.PrettyString())

		return nil
	})

	if err != nil {
		logrus.Debug("Error connection occured : ", err)
	}

}
