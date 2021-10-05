package database

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	IDProduct   string    `gorm:"column:id_product;type:varchar(25)"`
	ProductName string    `gorm:"column:product_name;type:varchar(25)"`
	Price       int       `gorm:"column:price;type:int"`
	Description string    `gorm:"column:description;type:text"`
	Quantity    int       `gorm:"column:quantity;type:int"`
	CreatedDate time.Time `gorm:"column:created_datetime"`
	UpdatedDate time.Time `gorm:"column:updated_datetime"`
	Active      bool      `gorm:"column:active;type:bool"`
}

func (p *Product) Create(db *gorm.DB, IDProduct, productName, description string, price, quantity int, createdDate time.Time, active bool) error {
	p.IDProduct = IDProduct
	p.ProductName = productName
	p.Price = price
	p.Description = description
	p.Quantity = quantity
	p.CreatedDate = createdDate
	p.Active = active

	return db.Table("product").Create(&p).Error
}

func (p Product) ProductListTime(db *gorm.DB) ([]Product, error) {
	products := []Product{}
	err := db.Table("product").Order("created_datetime desc").Find(&products).Error
	return products, err
}

func (p Product) ProductPriceHigh(db *gorm.DB) ([]Product, error) {
	products := []Product{}
	err := db.Table("product").Order("price desc").Find(&products).Error
	return products, err
}

func (p Product) ProductPriceLow(db *gorm.DB) ([]Product, error) {
	products := []Product{}
	err := db.Table("product").Order("price asc").Find(&products).Error
	return products, err
}

func (p Product) ProductNameAZ(db *gorm.DB) ([]Product, error) {
	products := []Product{}
	err := db.Table("product").Order("product_name asc").Find(&products).Error
	return products, err
}

func (p Product) ProductNameZA(db *gorm.DB) ([]Product, error) {
	products := []Product{}
	err := db.Table("product").Order("created_datetime desc").Find(&products).Error
	return products, err
}

func (p *Product) GetByID(db *gorm.DB, id_product string) error {
	return db.Table("product").Where("id_product=?", id_product).Last(&p).Error
}

func (p *Product) Updateproduct(db *gorm.DB, idproduct, namaBarang, deskripsi string, jumlahBarang int, hargaSatuan, gambarBarang string, updated_datetime time.Time) error {
	sql := "update product set nama_barang=?, deskripsi_barang=?, jumlah_barang=?, harga_barangsatuan=?, gambar_barang=?, updated_datetime=? where id_product=?"
	if err := db.Table("product").Exec(sql, namaBarang, deskripsi, jumlahBarang, hargaSatuan, gambarBarang, updated_datetime, idproduct).Error; err != nil {
		return err
	}

	return nil
}

func (p *Product) Deleteproduct(db *gorm.DB, id_product string) error {
	return db.Table("product").Where("id_product=?", id_product).Delete(&p).Error
}
