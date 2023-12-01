package loops

import "fmt"

//1.kullanıcıdan bir sayı girmesini isteyiniz
//asal sayı bulan program
//for example: 23 - asalsayıdır
//asal sayi : kendisine ve sadece 1 bölünen sayı denir


func Demo4(){
	sayi:=0
	fmt.Println("Bir sayı giriniz:")
	fmt.Scanln(&sayi)

	asalMi:=true
	for i:=2; i<sayi; i++{
		if sayi%i == 0{//asal sayı olmuyor 
			asalMi=false//farklı sayılara bölündüğü için
	}
}

	if asalMi == true{
		fmt.Println("Asaldır")
	} else {
		fmt.Println("Asal Değildir")
	}




}