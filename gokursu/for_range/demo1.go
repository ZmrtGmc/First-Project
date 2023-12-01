package for_range

import "fmt"
//for_range dönügüsü: 
//map yapısında for_range kullanılması bir avantajdır

func Demo1(){
	sehiler:=[]string{"Ankara","İstanbul","İzmir"}//slice yapısı
	//for yapısıyla yapılması
	for i:=0;i<len(sehiler);i++{
		fmt.Println(sehiler[i])
	}

	//for range olarak yapılması//alt çizgi i olarak algılanır
	for i,sehir:=range sehiler{//sehirleir benim için gez demek
		//o an gezdiği istanbulsa sehir: istanbul oluyor
		fmt.Println(sehir)
		fmt.Print(i+1)
	}
	//for _,sehir şeklinde de aynı şekilde çalışır üstteki kod parçası
}