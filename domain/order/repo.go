package order

import (
	// "errors"

	"encoding/json"

	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket"
	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderRepository struct {
	db *gorm.DB
}

//create a sigleton of the repo instance
var singleton *OrderRepository = nil

//initilaze the repo with gorm db
func OrderRepoInit(db *gorm.DB) *OrderRepository {
	if singleton == nil {
		singleton = &OrderRepository{db}
	}
	return singleton
}

//Before using this you need initialize the repo
func Repo() *OrderRepository {
	return singleton
}

//Migrate curent values if exist on current DB
func (c *OrderRepository) Migrations() {
	c.db.AutoMigrate(&Order{})
	//https://gorm.io/docs/migration.html#content-inner
	//https://gorm.io/docs/migration.html#Auto-Migration
}

/*
	1.Open Transaction and lock product rows
	2.Check if current quantity of each product is enough for order
	3.If not enough quantity then rollback transaction and return error
	4.If enough quantity then update product quantity
	5.Update add order to orders table
	6.Delete completed basket items
	7.Commit transaction
*/
func (c *OrderRepository) CompleteOrder(baskets basket.Baskets, comment, shipingAddress, billingAddress string) error {

	//create producIDArray From basket
	productIdQuantityMap := baskets.GenerateProductIDTotalQuantityMap()

	//create productQuantityMap keys as array

	productIDs := make([]int, len(productIdQuantityMap))

	//this one create a slice of productIDs from keys of the map
	i := 0
	for k := range productIdQuantityMap {
		productIDs[i] = k
		i++
	}
	zap.L().Info("productIDs", zap.Any("productIDs", productIDs))

	//https://www.postgresql.org/docs/9.4/explicit-locking.html
	//lock product rows using productIDs array
	result := c.db.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id in ?", productIDs).Find(&product.Products{})

	if result.Error != nil {
		return errors.Wrap(result.Error, "Error while locking product rows")
	}

	//Open Transaction
	tx := c.db.Begin()
	var buyedProducts product.Products
	//check if current quantity of each product is enough for order
	for productID, quantity := range productIdQuantityMap {
		var product product.Product
		result = tx.Where("products.id = ?", productID).Joins("Category").Joins("Store").First(&product)

		if result.Error != nil {
			tx.Rollback()
			return errors.Wrap(result.Error, "Error while finding product")
		}
		if product.StockCount < quantity {
			//we can not complete order
			tx.Rollback()
			return errors.New("Not enough quantity")
		} else {
			//update product quantity and save it.If error then rollback transaction will help us on this update
			product.StockCount = product.StockCount - quantity
			//we add them to our buyedProducts array
			buyedProducts = append(buyedProducts, product)
			tx.Save(&product)
		}
	}





	//we dont need this to make add extra query. Transaction will do it for us if there is an error
	// //If enough quantity then update product quantity
	// for productID, quantity := range productQuantityMap {
	// 	var product product.Product
	// 	result = tx.Where("id = ?", productID).First(&product)

	// 	if result.Error != nil {
	// 		tx.Rollback()
	// 		return result.Error
	// 	}

	// 	product.StockCount = product.StockCount - quantity
	// 	tx.Save(&product)
	// }

	var orders Orders
	var order Order
	
	order.UserID = baskets[0].UserID
	order.Comment = comment
	order.ShippingAdress = shipingAddress
	order.BillingAddress = billingAddress
	// for each buyedProducts we add order to orders table
	for _, product := range buyedProducts {
		//generate product snapshot from product
		v, err := json.Marshal(product)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "Error while marshalling product")
		}
		order.ProductSnapshot = v
		order.ProductID = int(product.ID)
		order.Quantity = productIdQuantityMap[int(product.ID)]
		orders = append(orders, order)
		//Add quantity information to order
	}
	
	//add orders to orders table
	tx.Save(&orders)

	//delete completed basket items
	tx.Delete(&baskets)
	
	//end transaction
	return tx.Commit().Error
}

//get orders of customer with pagination
func (c *OrderRepository) GetOrders(userID int,page , pageSize int) (Orders, error) {
	var orders Orders
	result := c.db.Where("user_id = ?", userID).
	Scopes(domain.Paginate(page, pageSize)).
	Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}