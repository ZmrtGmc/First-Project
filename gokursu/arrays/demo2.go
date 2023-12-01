package arrays

import "fmt"

func Demo2() {
	var sehirler [5]string //dizi tanımlama
	sehirler[0]="Ankara"
	sehirler[1]="İstanbul"
	sehirler[2]="İzmir"
	sehirler[3]="Adana"
	sehirler[4]="Diyerbakır"
	//fmt.Println(sehirler)
	
	for i:=0;i<5;i++{
		fmt.Println(sehirler[i])//dizi elemanlarını alt alta yazdırma
	}
}