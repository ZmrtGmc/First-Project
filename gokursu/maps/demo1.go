package maps

import "fmt"
//map yapısı: key-value mimarisi üzerine kurulmuş yapılardır

func Demo1(){
	sozluk := make(map[string]string)//anahatar ve değerini oluşturuyoruz
	sozluk["table"]="Masa"
	sozluk["book"]="Kitap"
	sozluk["pencil"]="Kalem"//mape değer atama

	fmt.Println(sozluk["book"])//bize book karşılığını verir yani ekrana kitap yazar
	fmt.Println("eleman sayısı: ",len(sozluk))
	fmt.Println("Sözlük :",sozluk)//tüm elemanları ekrana yazdırır
	delete(sozluk,"book")//sözlükten book silecek
	fmt.Println("eleman sayısı: ",len(sozluk))
	fmt.Println("Sözlük :",sozluk)


//sözlüğün içerisinde var mı yok mu işlemi
	//deger := sozluk["table"]//iki değer döndürür, deger: table varmi true yada false
	deger, varMi := sozluk["table"]
	fmt.Println(deger)
	fmt.Println("listede olma durumu: ",varMi)

	deger1,varMii := sozluk["zmrt"]//iki değer döndürür, deger: table varmi true yada false
	fmt.Println("listede olma durumu: ",varMii)
	fmt.Println(deger1)

	//make kullanmadan da maps yapılabilir
	sozluk2:=map[string]string{"glass":"bardak","microphone":"mikrofon"}
	fmt.Println(sozluk2)
}