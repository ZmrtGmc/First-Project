package conditionals

import (
	"fmt"
)
//if blokları 
func Demo1() {
	var hesap float64=1000
	var cekilmekIstenen float64=900

	if cekilmekIstenen > hesap {
		fmt.Print("Hesaptaki para yetersiz")
	}
    if cekilmekIstenen == hesap{
		fmt.Println("para çekilebilir")
	}
	if cekilmekIstenen <= hesap{
		fmt.Println("para çekilebilir")
	}

	fmt.Println("Bitti, Hesaptaki Para: "+ fmt.Sprintf("%v", hesap))
	fmt.Printf("bitti,hesaptaki para: %v", hesap)

}