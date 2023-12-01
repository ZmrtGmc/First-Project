package loops

import "fmt"

// arkadaş sayı bulma
//arkadaş sayı: iki sayı birbirinin kendiis hariç bölenleri toplamına eşitse bu sayılara arkadaş sayı denir
// 220: 1+2+4+5+10+11+20+22+44+55+110 =284 (yani 220 ve 284 bunlar arkadaş sayıdır)
// 284 bölenleri ; 1+2+4+71+142=220

//ekran çıktısı: 220 ve 284 arkadaş sayılardır
// 10 ve 65 arkadaş sayı değildir

func Demo5(){
	sayi1:=220
	sayi2:=285
	toplam1:=0
	toplam2:=0
	for i:=1;i<sayi1;i++{
		if sayi1%i ==0{
			toplam1=toplam1+i
		}
	}

	for i:=1;i<sayi2;i++{
		if sayi2%i ==0{
			toplam2=toplam2+i
		}
	}
	//toplam1 : sayi2 olmalı, toplam2 : sayi1 ise bunlar arkadaş
	if toplam1==sayi2 && toplam2==sayi1{
		fmt.Printf("%v ile %v arkadaş sayılardır",sayi1,sayi2)
	} else {
		fmt.Printf("%v ile %v arkadaş sayı değildir",sayi1,sayi2)
	}
}