//for döngüsü
//loop : döngü

package loops

import "fmt"

func Demo1() {
	var metin string ="Merhaba Zümrüt"
	//fmt.Println(metin)
	//fmt.Println(metin)
	i:=1
//for un bir çok kullanım şekli vardır 
//yazma şekli olarak en basit hali
	for i<=5{//infinite loop(çünkü i değeri artılırılmamış henüz)
		fmt.Println(metin)
		i=i+1
	}


}