package functions

//import "fmt"
//çoklu return döndüren fonksiyonlar
func DortIslem(x int,y int) (int,int,int,float32){

	toplam:=x+y
	cikarma:=x-y
	carpim:=x*y
	bolum:=float32(x/y)//veritipini değiştirmek

	return toplam,cikarma,carpim,bolum
}

/*	sonuc1,sonuc2,sonuc3,sonuc4:=functions.DortIslem(10,6)
	//hepsini değilde bir kaç tanesini kullanmak istersek
	// 	sonuc1,sonuc2,sonuc3, _ :=functions.DortIslem(10,6)
	//_ : alt çizgi koyduğumuzda bunu işleme alma anlamı veriyoruz
	fmt.Println("toplam: ",sonuc1)
	fmt.Println("çıkarma: ",sonuc2)
	fmt.Println("çarpma: ",sonuc3)
	fmt.Println("bölüm: ",sonuc4)

	main kısmı */