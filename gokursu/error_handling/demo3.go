package error_handling

import (
	"fmt"
	"errors"
)
/*bir arayüzün temel türüne erişmeye ve bir 
arayüz değişkenindeki belirsizliği ortadan kaldırmaya yardımcı olur.
*/

func TahminEt(tahmin int) (string,error) {//geri dönüş değeri farklı istediğimizde bu şekilde tanımlayabiliyoruz
	//error go dilinde tanımlı oldpu için tanımlama yapmamız gerekmiyor
	if tahmin < 1 || tahmin > 100{
		return "", errors.New("1-100 arasında bir sayı giriniz.")
	}
	return "Bildiniz",nil //nil : hata yok demekti

}

func Demo3() {
	mesaj, _:=TahminEt(50)
	fmt.Println(mesaj)
	fmt.Println(TahminEt(101))
	fmt.Println(TahminEt(-10))
}