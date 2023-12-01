package pointers

import "fmt"
//pointers: maindeki değerinde etkilenmesi için kullanılır
//pointer bellek üzerinden çalışacağımı belirtiyoruz
//pointer başına * koyularak işlem yapılır
//* (yıldız ) adres belirtir
func Demo1(sayi *int){
	*sayi=*sayi+1
	fmt.Println("demodaki sayı: ",*sayi)//21 ekran çıktısı
//pointer değer göndermek için & bu işaret kullanılır
}

/* main kısmı
	sayi:=20
	pointers.Demo1(&sayi)//gönderilen sayının bellekteki yerini *
	fmt.Println("maindeki sayı: ",sayi)*/