	/*ilk uygulama
	fmt.Print("Merhaba") //mesajı yazdığında alt satıra geçmez
	fmt.Println("bu benim ilk uygulamam")//mesajın yazdıktan sonra alt satıra geçer
	fmt.Print("Ben go öğreniyorum")
	fmt.Print(" ZG")

	/*ikinci uygulama
	var metin string="Zümrüt Gemici"
	fmt.Println(metin)
	var kdv int=10
	fmt.Println(100+100*kdv/100)

	kdv3 :=25 //kısa tanımlama 
	kdv4 :="metin" //bu sadece go programa özeldir
	kdv5 :=0.5  //değişken tanımlama ve veri tipi
	fmt.Println(kdv3)
	fmt.Println(kdv4)
	fmt.Println(kdv5)
	fmt.Printf("veri tipi: %T ",kdv3)//kontrol yaparak ekrana yazdırdığımız için printf şeklinde kullanılıyoruz
	fmt.Printf("veri tipi: %T ",kdv4)
	fmt.Printf("veri tipi: %T ",kdv5)
	// 3.uygulama
	zmrt :=false
	fmt.Printf("veri tipi: %T\n",zmrt)//bool tipi

	var durum bool=false // var durum bool şeklinde de tanımlanır
	var metin1 string="Zümrüt"
	var metin2 string="Gemici"
	var metin3 string="Gemici"

	durum =metin1==metin2 //eşitlik oporötür yapılıyor
	fmt.Println(durum)
	durum= metin2!=metin3// eşit değildir dimi? sorusu
	fmt.Println(!durum)//mantıksal değerin tam tersini yazmasını istediğimizde başına ünlem koyuluyor
	*/
	package variables
	import "fmt"

	func Demo1(){
		kdv3 :=25 //kısa tanımlama 
		kdv4 :="metin" //bu sadece go programa özeldir
		kdv5 :=0.5  //değişken tanımlama ve veri tipi
		fmt.Println(kdv3)
		fmt.Println(kdv4)
		fmt.Println(kdv5)
		fmt.Printf("veri tipi: %T ",kdv3)//kontrol yaparak ekrana yazdırdığımız için printf şeklinde kullanılıyoruz
		fmt.Printf("veri tipi: %T ",kdv4)
		fmt.Printf("veri tipi: %T ",kdv5)

	}

