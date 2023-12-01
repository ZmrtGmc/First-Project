package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2(){
	isim:="Zümrüt"
	fmt.Println(s.HasPrefix(isim,"Zm"))//zm ile başlıyor mu?
	fmt.Println(s.HasSuffix(isim,"üt"))//üt ile bitiyor mu?	
	fmt.Println(s.Index(isim,"Züm"))//Züm kaçıncı index	
	fmt.Println(s.Index(isim,"züm"))//züm kaçıncı index	
	harfler:=[]string{"z","ü","m","r","ü","t"}
	fmt.Println(s.Join(harfler,"*"))//elemanları nasıl ayıracağımı belirtiyorum
	fmt.Println(s.Join(harfler,""))//elemanları nasıl ayıracağımı belirtiyorum

	//belli bir ifade başka bir şeyle değiştirme durumudur
	//örneğin c yi k yap gbi
	sonuc:=s.Join(harfler,"-")
	fmt.Println(sonuc)
	fmt.Println(s.Replace(sonuc,"-","+",-1))//- yerine + olacak
	//-1 ne görürsen gör hepsini değiştir demektir

	//join tam tersi
	//2881111111,02022021,50000 tc-tarih-tutar
	fmt.Println(s.Split(sonuc,"-"))//çizgilere göre ayır demektir
	fmt.Println(s.Repeat(sonuc,5))//count kısmı 5 gibi, tekrar eder sayı kadar
	


}

