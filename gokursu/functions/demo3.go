package functions

//import "fmt" programda burada aktif kullanılmayacaksa yazmak hata veriyor 

//variadic fonksiyonlar
//kaç tane parametre ile işlem yapacağımızı bilmediğimiz zamanlarda kullanılır
//parametreyi dizi olarak tanımlar tanımlanması sayilar ...int şeklinde yapılır
func ToplaVariadic(sayilar ...int) int {
	toplam:=0
	for i:=0;i<len(sayilar);i++{
		toplam=toplam+sayilar[i]
	}
	return toplam
}

/*main kısmı 
	var sonuc int =  functions.ToplaVariadic(5,7,8)
	var sonuc1 int = functions.ToplaVariadic(5,7,8,9,6,45)
	var sonuc2 int = functions.ToplaVariadic()
	fmt.Println(sonuc)
	fmt.Println(sonuc1)
	fmt.Println(sonuc2)
	number:=[]int{4,6,7,0,11}
	var sonuc3 int = functions.ToplaVariadic(number...)
	fmt.Println(sonuc3)//parademik olduğunu belirtmek zorundayız

	*/