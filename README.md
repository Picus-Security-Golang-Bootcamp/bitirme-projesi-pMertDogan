# Patika - Picus Go Web Development Final Project

On this final project we develop  REST API Backend built with Go to handle all Basket Apps possible CRUD operations. Tech Stack is : Go , Gin , GORM, ZAP and PostreSQL

Detailed endpoint definations can be found on Postman or Swagger(Auto Generated from Postman) document.

[Postman Document](https://documenter.getpostman.com/view/11892665/UVyvvEFT)

Also Swagger Document available.

[Swagger](https://app.swaggerhub.com/apis/763/BasketApp/1.0.0)


# Database ERD 
![Database ERD Image](dbERD.png?raw=true "Title")


# Startup Configs

This project is using a config file to get startup or const parameters from config file. To preview default one check this link > [Config File](https://github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-pMertDogan/blob/main/pkg/config/config-local.yaml)

You can specify Database config. Its should be updated with your own credentials..

DBConfig:
  Host: localhost
  Port: 5432
  Dbname: basketStore
  Username: postgres
  Password: 123456
  
  ## Disable dev flag 
  
  To avoid debug logs on console disable dev flag
  
  -Logger:
  Development: true
  Encoding: json
  Level: info
  
  to 
  
  -Logger:
  Development: false
  Encoding: json
  Level: info

## JWT Config
JWT session time periods can be change from config. Dont forget change SecretKey to stronger one. Bcrypto used to hash & salt

JWTConfig:
  SecretKey: dummySecretKey
  RefreshTokenLifeMinute : 1440
  AccesTokenLifeMinute : 100


# Summary of the supported Endpoints

## 1. Sign-up
- Client send email,password and name to register user. If its valid user registered.
## 2. Login
- With email and password combination user can login and get acces & refresh token for logged user
## 3. Create Bulk Category
- Admins can generate bulk category from csv file. Example category file can be found on project root.
[ExampleCategoryCSV](https://github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-pMertDogan/blob/main/exampleCategory.csv)
## 4. List Category
- List avialable categorys with pagination
## 5. AddToBasket
- Users who logged can add products to basket. Authorization requires valid JWT acces token.
## 6. List Basket Items
- User can be list his basket items
## 7. Update/Delete Basket Items
- Update or Delete basket items
## 8. Complete Order
- User can buy requested baskets. BasketID is required to know which items will be buyed 
## 9. List Orders
- User can display orders, pagination applied
## 10. Cancel Order
- User can cancel order is order created time is not older than 14 days.If its older than 14 days client get denied from api
## 11. Create Product
- Admins can create a product. Bulk Product creation is available using CSV file. Check API documents for more info!
## 12. List Product
- Everyone can list avaiable products.Pagination applied
## 13. Search Product
- Search product by sku,desc,name..
## 14. Delete Product
- Admins can delete product. Deletes are soft delete.
## 15. Update Product
- Admins can update product with JSON body

# Tech Stack
- GO
- Gin - to Handle HTTP Requests
- Postgres - our Database run on Postgres
- GORM  - ORM liblary to help us on basic querys
- JWT   - to add support Authorization 
- ZAP   - Level based logging avaiable with the help of the ZAP!


# Current Avaiable Endpoints are (Auto generated by Gin)

|HTTP Method|Endpoint               |Detail                                                                                                                    |
|-----------|-----------------------|--------------------------------------------------------------------------------------------------------------------------|
|POST       | /category/upload      |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/category.rysFromCSV (4 handlers)        |
|GET        | /category/            |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/category.riesWithPagination (3 handlers)|
|GET        | /check/status         |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/check.statusCheck (3                    |
|GET        | /check/ready          |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/check.ready (3 handlers)                |
|POST       | /login                |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/auth.(*authHandler).login-fm               |
|POST       | /register             |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/auth.(*authHandler).(3 handlers)           |
|POST       | /token/fresh          |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/auth.TokenControllerDef.func1              |
|GET        | /product/             |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.tWithPagination (3 handlers)    |
|POST       | /product/             |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.CreateProduct (4                |
|POST       | /product/bulk         |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.CreateBulkProduct               |
|POST       | /product/search       |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.Search (3 handlers)             |
|DELETE     | /product/:id          |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.Delete (4 handlers)             |
|PATCH      | /product/:id          |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.Update (4 handlers)             |
|POST       | /basket/:id           |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket.AddToBasket (4                   |
|GET        | /basket/:id           |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket.GetBasket (4                     |
|PATCH      | /basket/:id           |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket.UpdateBasket (4                  |
|DELETE     | /basket/:id/:basketID |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket.DeleteBasket (4                  |
|POST       | /order/:id/           |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/order.CompleteOrder (4                  |
|GET        | /order/:id/           |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/order.GetOrders (4                      |
|POST       | /order/:id/:orderID   |--> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/order.CancelOrder (4 handlers)          |
