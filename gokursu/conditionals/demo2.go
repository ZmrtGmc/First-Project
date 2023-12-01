package conditionals

import (
	"fmt"
)
//if else blokları 
//if else, if else if yapısı
func Demo2(){
	var hesap float64=1000
	var cekilmekIstenen float64=90

	if cekilmekIstenen > hesap {
		fmt.Println("hesaptaki para yetersiz")
	} else if cekilmekIstenen == hesap{
		fmt.Println("paranız hazırlanıyor")
	} else {
		fmt.Println("paranız hazırlanıyor")
	}

}

/*note
iç içe ifyapısı
if sart{
	if sart{
			koddd
	}
}*/