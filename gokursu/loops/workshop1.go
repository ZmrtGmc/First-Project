package loops

import "fmt"
//ödevi var--- girilen sayının 1-100 arasında kontrolü
//1 den 100 a kadar bir sayı söylenecek
//sayı tahmin oyunu 

//1-3 arasındaki tahmin sayısı : süper
// 4-6 arasındaki tahmin sayısı güzel
// 6 ve daha üstünde ise: fena değil

func Demo3(){
	aklimdakiSayi:= 80
	tahminEdilenSayi:= 0
	skor :=0

	fmt.Println("Bir sayı tahmin ediniz:")
	fmt.Scanln(&tahminEdilenSayi)//kullanıcıdan okunan sayı
	skor=skor+1

	for aklimdakiSayi!=tahminEdilenSayi{
		if tahminEdilenSayi < aklimdakiSayi{
		fmt.Println("Daha Büyük Bir Sayı Giriniz")
		fmt.Scanln(&tahminEdilenSayi)//kullanıcıdan okunan sayı
		skor=skor+1
		} else {
		fmt.Println("Daha Küçük Bir Sayı Giriniz")
		fmt.Scanln(&tahminEdilenSayi)
		skor=skor+1
	}
}
	durum:=""
	if skor<=0 || skor<=3 {
		durum="Süper"

	} else if skor<=6{
		durum="Güzel"

	} else if skor>6 {
		durum="Fena Değil"
	}

    if tahminEdilenSayi == aklimdakiSayi{
	//fmt.Println("Tebrikler Bildiniz")
	fmt.Printf("Tebrikler Bildiniz. %v tahmin : %v",skor,durum)
    }

}