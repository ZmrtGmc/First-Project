//struct metodlarıyla çalışma
package structs

import "fmt"

func Demo2(){
	c:=customer{firtName:"Zmrt",lastName:"Gmc",age:23}
	c.save()
	fmt.Println("adı :",c.firtName)//ekran çıktısı zmrt
	c.update()
}

type customer struct{//customer istediğimiz ismi verebiliriz
	firtName string
	lastName string
	age int

	//buraya fonkisyon yazamıyoruz
	//func diyerek işlem yapılamıyor
	//veritabanı olarak düşünelim
	//eklee çıkarma, silme işlemleri yapılabilir
}

func (c customer) save(){//save isimdir sadece
	fmt.Println("Eklendi")//c customer : ilişkilendirdiğimiz
}

func (c customer) update(){//save isimdir sadece
	fmt.Println("Güncellendi")//c customer : ilişkilendirdiğimiz
}