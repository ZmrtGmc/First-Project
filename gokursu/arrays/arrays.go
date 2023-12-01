package arrays

import "fmt"

func Demo1() {
	var sayilar [5]int //sayilar dizinin adı, dizi veri tipi int, eleman sayısı 5 
	sayilar[2]=50 //dizilerde sayma sayısı sıfırdan başlar
	//yani aslında 5 elemanlı dizide 0-1-2-3-4 şeklindedir
	fmt.Println(sayilar)
	fmt.Println(sayilar[2])//2.elemanı ekrana yaz
}