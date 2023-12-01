package interfaces
/*Interface, farklı tasarlanmış ama sonuç olarak
aynı amaç için kullanılan
structlar için bir standart tanımlamasıdır. */
//geometrik şeklin alanını hesalanacak
import (
	"fmt"
	"math" //kütüphane eklememiz
	)
//shape : şekil , istediğimiz isim yazabiliriz belli kod değildir

type shape interface{//interface tanımlama
	area() float64 //alanı
}

type rectangle struct{//dikdörtgen
	width,height float64//yüksekilk ve genişlik
}

type circle struct{//daire
	radius float64 //yarıçap
}

//structlara metot yazabiliriz

func (r rectangle) area() float64{
	return r.width * r.height
}

func (c circle) area() float64{//ismin aynı olmasına dikkat ediyoruz interfacedekiyle
	return c.radius * c.radius * math.Pi
}

func school(s shape){
	fmt.Println("Şeklin Alanı: ",s.area())
}

func Demo1(){
	r:=rectangle{width:10,height:6}
	school(r)

	c:=circle{radius:10}
	school(c)
}


/*main kısmı
	interfaces.Demo1()

	çalışması için yeterlidir ekran çıktısı
	Şeklin Alanı:  60
Şeklin Alanı:  314.1592653589793*/
