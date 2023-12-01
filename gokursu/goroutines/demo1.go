//go routines
package goroutines

import "fmt"
//Asenkron aynı anda çalışması fonkisyonun //sistemin kitlenmemesi gerekiyor
/*Asenkron, "eş zamanlı" veya "aynı anda" 
anlamına gelen "synchronous" teriminin aksine, 
işlemlerin kendi hızlarında gerçekleştirildiği bir yöntemdir*/
func CiftSayilar() {
	for i:=0; i<=10; i += 2 {//i =i+2 aynı şekildedir
		//çift sayıları yazacak
		fmt.Println("Çift Sayı: ",i)
		//time.Sleep(1*time.Second)//işlemciden işlemin süresi
	}
}

func TekSayilar() {
	for i:=1; i<10; i += 2 {
		fmt.Println("Tek Sayı: ",i)
	}
}

//ikisinin aynanda çalışmasını istersek
/*iki fonksiyonunun aynı anda çalışmasını istersek
ve birbirlerinin ççalışmasını etkilemesini istersek 
sadece main bölümünde normalde package name ve fonnskiyonun adını
yazdığımızın yazının başına go yazmamız yeterlidir
örnek: goroutines.CiftSayilar()
bunun yerine go goroutines.CiftSayilar()

bittiğini ilişkilendirmemiz gerekiyor : time paketinen yararlanır

*/

/*	
main kısmı
		go goroutines.CiftSayilar()
		go goroutines.TekSayilar()
		time.Sleep(5*time.Second)//5 saniye olarak bekletiliyor
		fmt.Println("Main bitti: ") 

		ekran çıktısı önce tek sayıları yazdı, sonra çift sayıları yazdı
		hangisinde işlemleri hızlıysa onu yazıyor önce ekrana 
		*/