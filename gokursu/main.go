package main

//import "fmt"  //kullanmak istediğimizde tnaımlıyoruz, genelde kullanılması zorunlu oluyor
//modül yapısı
import (
	//"fmt"
	//"golesson/variables" 
	//"golesson/conditionals"
	//"golesson/loops"
	//"golesson/arrays"
	//"golesson/slices"
	  //"golesson/functions"
	 // "golesson/maps"
	 //"golesson/for_range"
	// "golesson/pointers"
	//"golesson/structs"
	//"golesson/goroutines"
	//"time"
	//"golesson/channels"
	//"golesson/interfaces"
	//"golesson/defer_statement"
	//"golesson/error_handling"
	//"golesson/string_functions"
	//"golesson/restful"
	"golesson/project"


)//bu modülün altındaki variables komutunu kullanıyorum

func main() {

	//variables.Demo1()//fonksiyonu çalıştırıyorum
	//fmt.Print()
	//conditionals.Demo3()
	//arrays.Demo4()
	//slices.Demo2()
	//functions.SelamVer()
	//functions.Topla(2,3)//parametre veriyoruz
	//sonuc1,sonuc2,sonuc3,sonuc4:=functions.DortIslem(10,6)
	//hepsini değilde bir kaç tanesini kullanmak istersek
	// 	sonuc1,sonuc2,sonuc3, _ :=functions.DortIslem(10,6)
	//_ : alt çizgi koyduğumuzda bunu işleme alma anlamı veriyoruz
	/*fmt.Println("toplam: ",sonuc1)
	fmt.Println("çıkarma: ",sonuc2)
	fmt.Println("çarpma: ",sonuc3)
	fmt.Println("bölüm: ",sonuc4)*/

		//hepsi çalışır
	/*var sonuc int =  functions.ToplaVariadic(5,7,8)
	var sonuc1 int = functions.ToplaVariadic(5,7,8,9,6,45)
	var sonuc2 int = functions.ToplaVariadic()
	fmt.Println(sonuc)
	fmt.Println(sonuc1)
	fmt.Println(sonuc2)
	number:=[]int{4,6,7,0,11}
	var sonuc3 int = functions.ToplaVariadic(number...)
	fmt.Println(sonuc3)//parademik olduğunu belirtmek zorundayız

*/
//maps.Demo1()
	//for_range.Demo2()
	//for_range.Demo3()
	/*sayi:=20
	pointers.Demo1(&sayi)
	fmt.Println("maindeki sayı: ",sayi)
	sayilar :=[]int{1,2,3}
	pointers.Demo2(sayilar)
	fmt.Println("maindeki sayı: ",sayilar[0])*/
		//structs.Demo2()
		/*go goroutines.CiftSayilar()
		go goroutines.TekSayilar()
		time.Sleep(5*time.Second)//5 saniye olarak bekletiliyor
		fmt.Println("Main bitti: ")
		ciftSayiCn  := make(chan int)
		tekSayiCn := make(chan int)
		go channels.CiftSayilar(ciftSayiCn)
		go channels.TekSayilar(tekSayiCn)

		ciftSayiToplam,tekSayiToplam:= <-ciftSayiCn, <-tekSayiCn

		carpim := ciftSayiToplam*tekSayiToplam
		fmt.Println("Çarpım: ",carpim)*/
	//interfaces.Demo2()
	//defer_statement.Demo3()
	//error_handling.Demo2()
	//fmt.Println(error_handling.TahminEt2(102))
	//string_functions.Demo2()
	//restful.Demo2()
	project.AddProduct()
	project.GetAllProducts()

}
//camel case   camelCases
// pascal case  PascalCase

/*go da moodül tanımlama terminalde
go mod init golesson şeklinde yapılır
golesso isimdir , kullanıcı istediği ismi yazabilir
init, başlatmak anlamında gelir*/