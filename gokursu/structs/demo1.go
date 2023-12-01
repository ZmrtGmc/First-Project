package structs

import "fmt"
//structs: class konusuna benzer
//
func Demo1(){//structları sırasıyla yazmak gerekiyor veri türlerine uygun şekilde
	fmt.Println(product{"Bilgisayar",500,"XYZ",200})
	fmt.Println(product{"Bilgisayar",500,"XYZ",150})
	fmt.Println(product{name:"Bilgisayar",unitPrice:500})//default olarak verilmesini istersek bu şekilde tanımlıyoruz
	
}


type product struct{//product diye bir struct mevcuttur
//içerğine özellikleri koyacağız
//product kısmına istediğimizi yazabiliriz struct ismidir
	name string//değişim adı ve türü
	unitPrice float64
	brand string
	discountRate int
}