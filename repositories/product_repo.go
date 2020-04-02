package repositories

import (
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"github.com/yaboyou/facai-admin/datamodel"
	"github.com/yaboyou/facai-admin/datasource"
)

type ProductRepo interface {
	GetProductList(m map[string]interface{}) (total int, productList []datamodel.Product)
	SaveProduct(product datamodel.Product) (err error)
	DeleteProduct(id uint) (err error)
	GetProduct(id uint) (product datamodel.Product, err error)
}

func NewProductRepo() ProductRepo {

	return &productRepository{}
}

type productRepository struct{}

func (p productRepository) GetProduct(id uint) (product datamodel.Product, err error) {
	err = db.First(&product, id).Error
	return
}

func (p productRepository) GetProductList(m map[string]interface{}) (total int, productList []datamodel.Product) {
	var db = datasource.GetDB()
	db.Table("product").Count(&total)
	err := db.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"]) - 1) * cast.ToInt(m["size"])).Find(&productList).Error
	if err != nil {
		panic("select Error")
	}
	return
}

func (p productRepository) SaveProduct(product datamodel.Product) (err error) {
	if product.PID != 0 {
		err := db.Save(&product).Error
		return err
	} else {
		err := db.Create(&product).Error
		var buffer bytes.Buffer
		sql := "INSERT INTO `productSpecs` (`s_name`,`s_state`,`s_price`,`s_brief`,`s_stock`,`s_productid`) VALUES "
		if _, err := buffer.WriteString(sql); err != nil {
			return err
		}
		for i, e := range product.ProductSpecs {
			if i == len(product.ProductSpecs)-1 {
				buffer.WriteString(fmt.Sprintf("('%s',%d,%f,'%s',%d,%d);", e.SName, e.SState, e.SPrice, e.SBrief, e.SStock, product.PID))
			} else {
				buffer.WriteString(fmt.Sprintf("('%s',%d,%f,'%s',%d,%d),", e.SName, e.SState, e.SPrice, e.SBrief, e.SStock, product.PID))
			}

		}
		err = db.Exec(buffer.String()).Error
		return err
	}
}

func (p productRepository) DeleteProduct(id uint) (err error) {
	return db.Where("p_id = ?", id).Delete(&datamodel.Product{}).Error
}
