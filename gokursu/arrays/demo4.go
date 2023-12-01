package arrays

import "fmt"
//çok boyutlu diziler
//matrix yapısı dama tahtası gibidir 3 satır 5 sütun a1 a2 a3 şeklinde



func Demo4() {
	var sayilar [2][3]int //ilki satır ikincisi sütun
	sayilar[0][0]=1
	sayilar[0][1]=3
	sayilar[0][2]=5
	sayilar[1][0]=2
	sayilar[1][1]=4
	sayilar[1][2]=6
	//fmt.Println(sayilar)//1 3 5  2 4 6
	//fmt.Println(sayilar[1][1])//4

	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			fmt.Print(sayilar[i][j])
			fmt.Print(" ")//boşluk bırakmak için
		}
		fmt.Println()//alt satıra geç
	}
}