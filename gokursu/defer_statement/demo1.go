package defer_statement

import "fmt"
//defer: deferden önce kullanılan kod blogunda hata olsa bile
//defer sayesinde o kod blogu bittiğinde defer olarak tanımlan diğer kod blokları çalışmaya devam eder

func A(){
	fmt.Println("a fonksiyonu çalıştı")
}

func B(){
	defer A()
	fmt.Println("b fonksiyonu çalıştı")
	defer C()//bu kodun üzeerinde ne olursa olsun, b fonksiyonun bitiminden sonra çlışmasını istediğimizi belirtiriz
	defer D()
}
//en son kutuya giren ilk kutudan çıkan olur
//deferlerin sırayla çalışması yukarıda çalışma şekli
//ekrana b den sonra , c yazılır sonra a yazılır
/* ekran çıktısı :
b fonksiyonu çalıştı
d fonksiyonu çalıştı
c fonksiyonu çalıştı
a fonksiyonu çalıştı
*/

func C(){
	fmt.Println("c fonksiyonu çalıştı")
}

func D(){
	fmt.Println("d fonksiyonu çalıştı")
}