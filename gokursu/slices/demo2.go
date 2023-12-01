package slices

import "fmt"

func Demo2() {
	//slice kopyalanabilir
	//slice make yapısını kullanarak oluşturabiliyorduk ancak dizi oluşturabilir gibi
	//tanımlanabailir örneğin; sehirler :=[]string{"Ankara"}
	sehirler:=[]string{"Ankara","İstanbul","İzmir"}
	fmt.Println(sehirler)

	sehirlerKopya:=make([]string,len(sehirler))
	fmt.Println(sehirlerKopya)
	//copy(dest, src)- kopyalancak dosya, yapıştırılan dosya
	copy(sehirlerKopya,sehirler)
	fmt.Println(sehirlerKopya)
	sehirler=append(sehirler,"Adana")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[1:3])//1.indisten başla 3 e kadar al
	//üstekinin ekran çıkısı: istanbul izmirdir
	//en sondaki yazdığımız mesela 3 dahil değildir
	fmt.Println(sehirler[1:])//1den başla sonuna kadar al
	fmt.Println(sehirler[:2])//baştan başlar 2 ye kadar dahil eder.
	


}