package defer_statement

import "fmt"

type product struct{
	productName string
	unitPrice int
}

func (p product) save(){
	fmt.Println("Ürün Kaydedildi:",p.productName)
	defer Log()
}

func Demo3(){
	p:=product{productName:"Laptop",unitPrice:5000}
	//p.save()
	//defer p.save()//buradayken laptop çıktısı verir
	p=product{productName:"Mouse",unitPrice:45}
	defer p.save()//buradaki mouse , son değeri alır
	fmt.Println("İşlem Başarılı")
}

func Log(){
	fmt.Println("Loglandı")
}