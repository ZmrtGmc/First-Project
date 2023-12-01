package string_functions

import (
	"fmt"
	"strings"
)

func Demo1(){
	isim:="Zümrüt"
	fmt.Println(strings.Count(isim,"ü"))
	fmt.Println(strings.Contains(isim,"ü"))
	fmt.Println(strings.Index(isim,"m"))//ekran çıktısı 3 oluyor
	fmt.Println(strings.Index(isim,"a"))//eğer yoksa -1 gönderiri
	
	sonuc:=strings.Index(isim,"a")
	if sonuc !=-1{
		fmt.Println("a harfi vardır")
	} else {
		fmt.Println("a harfi yok")
	}
	fmt.Println(strings.ToLower(isim))
	fmt.Println(strings.ToUpper(isim))

	//tolower: küçük harfe çeviri
	//toUpper: büyük harfe çevir
	//ekran çıktısı ü: count(string , aradaığımız harf)
	//iki tane ü olduğu için 2 olur ekran çıktısı
}
//case sensitive: büyük küçük harf duyarlıdır
//go dili büyük küçük harf duyarlıdır
//ascii : e E farklıdır değerleri
//alias