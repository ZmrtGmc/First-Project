package error_handling

import (
	"fmt"
	"os"
)
//hata yakalama 
//bir metin dosyasının okuma üzerine hata yakalam
//bir çok konuyu bir araya getirir

func Demo1(){
	f,err := os.Open("demo1.txt")//os go nun sağladığı bir koddur
	//open bize iki değer dönebilir, f dosyayı yoksa hatayı döner
	
	//dosya bulunursa err : değeri (nil)
	if err!=nil{//yani err nilden farklıysa, hata yoksa nil olur error
		fmt.Println("Dosya Bulunamadı")
	} else{
		fmt.Println(f.Name())
	}
}

//dosyayı bulabilmesi için
/* demo1.txt ismindeki dosya main dosyasının yanında olması gerekiyor
o şekilde olduğunda dosyanın adını yaz*/
