package repositories

import (
	"bytes"
	"errors"
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

	GetProductTypeList(m map[string]interface{}) (total int, typeList []datamodel.ProductType)
	GetProductTypeListByAll(m map[string]interface{}) (typeList []datamodel.ProductType)
	SaveProductType(pType datamodel.ProductType) (err error)
	DeleteProductType(id uint) (err error)
	GetProductType(id uint) (pType datamodel.ProductType, err error)
}

func NewProductRepo() ProductRepo {

	return &productRepository{}
}

type productRepository struct{}

var dbProduct = datasource.GetDB()

func (p productRepository) GetProductTypeListByAll(m map[string]interface{}) (typeList []datamodel.ProductType) {
	err := dbProduct.Find(&typeList).Error
	if err != nil {
		panic("select Error")
	}
	return
}

func (p productRepository) GetProductTypeList(m map[string]interface{}) (total int, typeList []datamodel.ProductType) {
	dbProduct.Table("productType").Count(&total)
	err := dbProduct.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"]) - 1) * cast.ToInt(m["size"])).Find(&typeList).Error
	if err != nil {
		panic("select Error")
	}
	return
}

func (p productRepository) SaveProductType(pType datamodel.ProductType) (err error) {
	if pType.PtID != 0 {
		err := dbProduct.Save(&pType).Error
		return err
	} else {
		err := dbProduct.Create(&pType).Error
		return err
	}
}

func (p productRepository) DeleteProductType(id uint) (err error) {

	var productList []datamodel.Product
	dbProduct.Where("p_typeid = ?", id).Find(&productList)
	if len(productList) != 0 {
		return errors.New("product by typeid not null")
	}
	productType, err := p.GetProductType(id)
	if err != nil {
		return err
	}
	return dbProduct.Save(&productType).Error
}

func (p productRepository) GetProductType(id uint) (pType datamodel.ProductType, err error) {
	err = dbProduct.First(&pType, id).Error
	return
}

func (p productRepository) GetProduct(id uint) (product datamodel.Product, err error) {
	err = dbProduct.First(&product, id).Error
	return
}

func (p productRepository) GetProductList(m map[string]interface{}) (total int, productList []datamodel.Product) {

	dbProduct.Table("product").Count(&total)
	err := dbProduct.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"]) - 1) * cast.ToInt(m["size"])).Find(&productList).Error
	if err != nil {
		panic("select Error")
	}
	return
}

func (p productRepository) SaveProduct(product datamodel.Product) (err error) {
	if product.PID != 0 {
		err := dbProduct.Save(&product).Error
		if err != nil {
			return err
		}
		err = dbProduct.Where("s_productid = ?", product.PID).Delete(&datamodel.ProductSpecs{}).Error
		if err != nil {
			return err
		}
		return createProductSpecs(product.PID, product.ProductSpecs)

	} else {
		err := dbProduct.Create(&product).Error
		if err != nil {
			return err
		}
		return createProductSpecs(product.PID, product.ProductSpecs)

	}
}

func createProductSpecs(pid int, specs []datamodel.ProductSpecs) (err error) {
	var buffer bytes.Buffer
	sql := "INSERT INTO `productSpecs` (`s_name`,`s_state`,`s_price`,`s_brief`,`s_stock`,`s_productid`) VALUES "
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, e := range specs {
		if i == len(specs)-1 {
			buffer.WriteString(fmt.Sprintf("('%s',%d,%f,'%s',%d,%d);", e.SName, e.SState, e.SPrice, e.SBrief, e.SStock, pid))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s',%d,%f,'%s',%d,%d),", e.SName, e.SState, e.SPrice, e.SBrief, e.SStock, pid))
		}
	}
	err = dbProduct.Exec(buffer.String()).Error
	return err
}

func (p productRepository) DeleteProduct(id uint) (err error) {
	product, err := p.GetProduct(id)
	if err != nil {
		return err
	}
	product.PState = 2
	err = dbProduct.Save(&product).Error
	if err != nil {
		return err
	}

	return nil
}
