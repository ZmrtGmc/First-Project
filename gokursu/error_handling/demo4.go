package error_handling

import (
	"fmt"
	//"errors"
)
//kendi hata yapılarımzıı oluşturmak

type borderException struct{
	parameter int
	message string
}

func (b *borderException) Error() string{
	return fmt.Sprintf("%d --- %s",b.parameter,b.message)
}

func TahminEt2(tahmin int) (string,error){
	if tahmin < 1 || tahmin > 100{
		return "",&borderException{tahmin,"Sınırların dışında kaldı"}
	}
	return "Bildiniz",nil
}

/*main kısmı
	fmt.Println(error_handling.TahminEt2(102))
ekran çıktısı : 102 --- Sınırların dışında kaldı
*/

