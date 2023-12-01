package channels

/*Kanallar(Channels), Go dilinde asenkron programlama
 yaparken değer aktarımı yapabileceğimiz hatlardır*/

//import "fmt"

func CiftSayilar(ciftSayiCn chan int) {
	toplam:=0
	for i:=0; i<=10; i += 2 {
		toplam=toplam+i
	}
	ciftSayiCn <- toplam //atama yapıyoruz işlemin bittiğini
}

func TekSayilar(tekSayiCn chan int) {
	toplam:=0
	for i:=1; i<10; i += 2 {
		toplam=toplam+i
	}
	 tekSayiCn <- toplam
}
//farklı iki tane fonksiyonu toplamının çarpımını bulalım

/*	
main kısmı
		ciftSayiCn  := make(chan int)
		tekSayiCn := make(chan int)
		go channels.CiftSayilar(ciftSayiCn)
		go channels.TekSayilar(tekSayiCn)

		ciftSayiToplam,tekSayiToplam:= <-ciftSayiCn, <-tekSayiCn

		carpim := ciftSayiToplam*tekSayiToplam
		fmt.Println("Çarpım: ",carpim)
*/