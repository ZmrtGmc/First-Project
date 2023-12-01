package defer_statement

import "fmt"

func Demo2(sayi int) string{
	defer DeferFunc()
	if sayi%2==0{
		//çift sayıysa
		return "Çift Sayıdır"
	}
	if sayi%2!=0{
		return "Tek Sayıdır"
	}
	//DeferFunc()//bu şekilde yazıldığında bir sayı döndürülüdüğü için bu hiç bir zaman çalışmayackakı
	//ancak önüne defer yazdığımızda her türlü bu yazılan fonskiyon çalışacaktır
	return "Belli Değil"
}

func Test(){
	fmt.Println(Demo2(9))
	fmt.Println(Demo2(10))
}

func DeferFunc(){
	fmt.Println("Bitti")
}
/*Bitti
Tek Sayıdır
Bitti
Çift Sayıdır 
ekran çıktısında önce bitti yazmasının sebebi
fonksiyon aslında bitiyor ama ekrana bir şey yazılmadığı için sanki önce bitti yazılıp sonra 
fonksiyon yazılıyor sanılıyor ama öyle değil
demo2 return olduktan sonra bitti yazdı ardından demo2 çıktısını ekrana yazdırdık
*/


