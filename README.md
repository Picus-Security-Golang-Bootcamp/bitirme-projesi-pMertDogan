# bitirme-projesi


DB 
https://www.figma.com/file/y4zTgzP0721nSmujTu0lD9/Untitled?node-id=0%3A1


1. Sign-up
- Kullanicidan gerekli bilgileri alip veri tabaninda bir kullanici olusturmalisiniz ve
cevap olarak JWT token donmelisiniz.
2. Login
- Veri tabaninda kayitli kullanicilar email ve sifre ile sisteme giris yapmali eger iki
bilgi de dogruyse JWT token olusturup kullaniciya donmelisiniz.
3. Create Bulk Category
- Bu endpoint ile sadece admin rolundeki kullanicilar CSV dosyasi yukleyerek yeni
bir kategori yaratmali
4. List Category
- Veri tabaninda aktif ve silinmemis olan tum kategoriler listelenmeli
5. AddToBasket
- Sisteme giris yapmis ve Token suresi gecmemis kullanicilar urunlerini sepete
ekleyebilir
6. List Basket Items
- Kullanicilar sepetine ekledigi urunleri listeleyebilir
7. Update/Delete Basket Items
- Kullanicilar sepetine ekledigi urunlerin adetini guncelleyebilir ya da urunu
sepetten cikarabilir
8. Complete Order
- Kullanicilar sepetine ekledigi urunler ile bir siparis olusturabilir
9. List Orders
- Musteriler kendisine ait gecmis siparisleri goruntuleyebilir
10. Cancel Order
- Eger musterinin siparis tarihini henuz 14 gunu gecmediyse musteri siparisini iptal
edebilir. Siparis olusma tarihinden sonra 14 gun gectiyse iptal istegi gecersiz olmalidir.
11. Create Product
- Admin rolündeki kullanıcılar sisteme tekil olarak ürün oluşturabilmeli.
12. List Product
- Kullanıcılar tüm ürünleri listeleyebilmeli. Burada rol kontrolü yok.
13. Search Product�
- Kullanıcılar ürünler içerisinde arama yapabilmeli. Arama kısmında ürün adı, sku
vb. gibi alanlarda arama yapılabilir.
14. Delete Product
- Admin rolündeki kullanıcılar ürün silebilir.
15. Update Product
- Admin rolündeki kullanıcılar ürünü güncelleyebilir.
Proje boyunca sizden kullanmanizi bekledigimiz teknolojiler;
- Gin
- Postgres
- GORM
- JWT
Kategori ve urun listeleme kisminda pagination yani sayfalama yapmaniz gerekiyor. Normal
sartlarda bu endpointlerden cok fazla veri gelebileceginden dolayi tum veriyi bi anda ekrana
bastiramazsiniz. Bu yuzden burada sayfalama yapmaniz gerekiyor.


#Current Avaiable Endpoints are (Auto generated by Gin)

POST   /category/upload          --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/category.rysFromCSV (4 handlers)
GET    /category/                --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/category.riesWithPagination (3 handlers)
GET    /check/status             --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/check.statusCheck (3 
GET    /check/ready              --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/check.ready (3 handlers)
POST   /login                    --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/auth.(*authHandler).login-fm 
POST   /register                 --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/auth.(*authHandler).(3 handlers)
POST   /token/fresh              --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/pkg/auth.TokenControllerDef.func1 
GET    /product/                 --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.tWithPagination (3 handlers)
POST   /product/                 --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.CreateProduct (4 
POST   /product/bulk             --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.CreateBulkProduct 
POST   /product/search           --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.Search (3 handlers)
DELETE /product/:id              --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.Delete (4 handlers)
PATCH  /product/:id              --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/product.Update (4 handlers)
POST   /basket/:id               --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket.AddToBasket (4 
GET    /basket/:id               --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket.GetBasket (4 
PATCH  /basket/:id               --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket.UpdateBasket (4 
DELETE /basket/:id/:basketID     --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket.DeleteBasket (4 
POST   /order/:id/               --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/order.CompleteOrder (4 
GET    /order/:id/               --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/order.GetOrders (4 
POST   /order/:id/:orderID       --> github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/order.CancelOrder (4 handlers)

