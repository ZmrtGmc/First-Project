package pointers

import "fmt"

func Demo2(sayilar []int){//parametre olarak array verilmesini istiyoruz
	sayilar[0]=100
	fmt.Println("demodaki sayı: ",sayilar[0])
}

/*main kısmı
	sayilar :=[]int{1,2,3}
	pointers.Demo2(sayilar)
	fmt.Println("maindeki sayı: ",sayilar[0])

	ekran çıktısı: 100

HER ARRAY BİR POİNTERDIR AMA HER POİNTER BİR ARRAY DEĞİLDİR

	*/
