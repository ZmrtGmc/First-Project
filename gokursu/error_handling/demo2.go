package error_handling

import (
	"fmt"
	"os"//paket 
)
//hatanın türü/tipi
//type assertion

func Demo2(){
	f,err := os.Open("demo1.txt")
	if err!=nil{
		if pErr,ok:=err.(*os.PathError); ok{
		//eğer işlem true ise burası çalışır, ok : true
			fmt.Println("Dosya bulunamadı: ",pErr.Path)
			return 
		//dosya yolunda , pointeer olarak tanımladık yolunu
		}else {
			fmt.Println("Dosya Bulunamadı")
			return
		}
	} else{
		fmt.Println(f.Name())
	}
}

