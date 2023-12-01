package project

import (
	"fmt"
	"io/ioutil"//diğer kütüphane
	"net/http"//paket kodu
	"encoding/json"
	"bytes"
)
type Product struct{
	Id int `json:"id"`
	ProductName string `json:"productName"`
	CategoryId int `json:"categoryId"`
	UnitPrice float64 `json:"unitPrice"`
}

type Category struct{
	Id int `json:"id"`
	CategoryName string `json:"categoryName"`
}

//tüm ürünleri listele
func GetAllProducts(){
	response, err:=http.Get("http://localhost:3000/products")
	if err!=nil{
		fmt.Println(err)
	}
	defer response.Body.Close()//defer her şartta kapansın

	bodyBytes,_:=ioutil.ReadAll(response.Body)

	//tüm gelen verileri bir struct yapısına okumamız gerekiyo
	var products []Product
	json.Unmarshal(bodyBytes,&products)
	fmt.Println(products)
}

//yeni ürün ekleme
func AddProduct(){
	product:=Product{Id:4,ProductName:"Telephone",CategoryId:1,UnitPrice:6000.0}
	//şu anda bir struct ama bunu json yapısına çevirmemiz gerekir
	jsonProduct,err:=json.Marshal(product)
	//şimdi json olan yapıyı api göndermemiz gerekiyor
	response,err:=http.Post(" http://localhost:3000/products",
	"application/json;charset=utf-8",bytes.NewBuffer(jsonProduct))

	if err!=nil{
		fmt.Println(err)
	}
	defer response.Body.Close()

	//veritabnına yani api kaynağına ekleme yapar
	fmt.Println("Kaydedildi")
}