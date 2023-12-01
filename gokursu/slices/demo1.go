package slices

import "fmt"
//slice : arraylerin yerine kullanır
//arrayler belli bi aralık vardır ama slice de bu söz konusu değildir 
//eklenebilir veya çıkarabilir 

func Demo1() {
	isimler:=make([]string,3) //kaç elemanla başlatılacağını söylüyoruz
	//slice yapısı
	//make yazmamız slice olduğunu belirtiyoruz
	isimler[0]="Zümrüt"
	isimler[1]="Kaan"
	isimler[2]="Recep"
	//isimler[3]="Kasavet" bunu tanımamız gerekiyor
	isimler =append(isimler,"Kasavet")
	//append fonksiyonu yeni bir slice oluşturuyor

	fmt.Println(isimler)
	fmt.Println(len(isimler))//eleman sayısını verir
	
}