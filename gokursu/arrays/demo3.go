package arrays

import "fmt"

func Demo3() {
	sayilar :=[5]int {20,30,50,10,2} //dizinin farklı tanımlama şekli
	fmt.Println(sayilar)

	//en büyük sayı hangisi bulan uygulama
	//lenght: uzunluk
	//len(sayilar): dizinin uzunluğunu hesaplar
	enBuyuk:=sayilar[0]
	for i:=0;i<len(sayilar);i++{
		if sayilar[i]>enBuyuk{
			enBuyuk=sayilar[i]
		}
	}
	fmt.Println("En büyük: ",enBuyuk)


}