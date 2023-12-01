package interfaces

import "fmt"

//tip doğrulama 
/* programamın akışına göre değişkenin istediğimiz türde mi 
onu kontrol etmemize sağlıyor */

func dogrula(i interface{}){//sayi tipini doğrulamaya çalışıyoruz
	sayi,ok :=i.(int)
	fmt.Println(sayi,ok)
}

func Demo3(){
	var sayi1 interface{} = 5
	dogrula(sayi1)

	var sayi2 interface{} = "zmrt"
	dogrula(sayi2)
}