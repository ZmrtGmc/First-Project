package functions

import "fmt"

//fonksiyonlar


func Topla(sayi1 int, sayi2 int) {
	var toplam =sayi1+sayi2
	fmt.Println("Sonuç: ",toplam)
}

func SelamVer(){
	fmt.Println("Selam babyy")
}

//geri değer döndüren fonksiyon 
func Topla1(sayi1 int, sayi2 int) int{
	var toplam=sayi1+sayi2
	return toplam//geri değer döndürür
	/*var sonuc=functions.Topla1(22,6)
	fmt.Println(sonuc*10)*/
}
func SelamVer1(kullaniciAdi string){
	//functions.SelamVer1("Zümrüt")
	fmt.Println("Merhaba ",kullaniciAdi)
}

