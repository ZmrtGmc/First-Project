package for_range

import "fmt"

func Demo3(){
	sozluk:=map[string]string{"book":"kitap","table":"masa"}
//k her zaman key
//v value
	for k,v:=range sozluk{
		fmt.Print(k)
		fmt.Println(v)
	}
}