package conditionals


/* üç adet int değişken tanımla,
	ekrana en büyük sayıyı yazıdırınız*/
import "fmt"

func Demo3() {
	var sayi int=100
	var sayi1 int=200
	var sayi2 int=150
//tek satırda tanımlama yapılabilir var sayi,sayi1,sayi2 int =10,5,2

//en küçük sayıyı bulan 
	if sayi < sayi1 && sayi < sayi2{
		fmt.Println("En Küçük Sayı: ", sayi)
	} else if sayi1 < sayi && sayi1 < sayi2{
		fmt.Println("En Küçük Sayı: ", sayi1)
	} else if sayi2 < sayi && sayi2 < sayi1{
		fmt.Println("En Küçük Sayı: ", sayi2)
	}
//en büyük sayı
	if sayi > sayi1 && sayi > sayi2{
		fmt.Println("En Büyük Sayı: ", sayi)
	} else if sayi1 > sayi && sayi1 > sayi2{
		fmt.Println("En Büyük Sayı: ", sayi1)
	} else if sayi2 > sayi && sayi2 > sayi1{
		fmt.Println("En Büyük Sayı: ", sayi2)
	}

}