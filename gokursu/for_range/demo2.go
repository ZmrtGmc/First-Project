package for_range

import "fmt"

//1-10 arasındaki sayılardan tek olanları toplayan program
//ama for_range ile yapılacaktır

func Demo2() {
	//sayilar :=[10]int {1,2,3,4,5,6,7,8,9,10}normal array
	sayilar:=[]int{1,2,3,4,5,6,7,8,9,10}//slice yapısı
	var toplam int=0
//birinci yol
/*
	for i:=0;i<len(sayilar);i++{
		if sayilar[i]%2 != 0{
			//çift sayı olmasın
			toplam=toplam+sayilar[i]	
		}
	}
	fmt.Printf("Tek Sayıların Toplamı: %v",toplam)
*/
	//for_range ie yapımı
	//i,sehir:=range sehiler
	for _,sayi:=range sayilar{
		if sayi%2 != 0{
			//çift sayı olmasın
			toplam=toplam+sayilar[sayi]	
		}
	}
	fmt.Printf("Tek Sayıların Toplamı: %v",toplam)

}

