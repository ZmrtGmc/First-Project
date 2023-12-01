package restful

import (
	"fmt"
	"io/ioutil"//diğer kütüphane
	"net/http"//paket kodu
	"encoding/json"
	"bytes"
)
/*Bugünkü yazımızda Golang ile JSON parse etmeye
bakacağız. Hepimizin bildiği gibi günümüzde bir
API (application programming interface) a
veri göndermede ya da veri çekmede en sık kullanılan veri formatı
JSON (javascript object notation) dur.
Golang ile de kendi oluşturduğumuz verimizi
(Golang struct) JSON’a dönüştürüp bir API’a request olarak
gönderebilir ya da bir API’dan gelen
JSON verisini Go programımızda kullanabiliriz. */

type Todo struct{
	UserId int `json:"userId"`//apideki neyi netlediğimi belirliyoruz
	Id int `json:"userId"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}
//test amaçlı kullanılmaktadır
//gönderdiğimizin aynısını gönderirse sorun yoktur

func Demo1(){
	response,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	//istekte bulunuyoruz
	if err != nil{
		fmt.Println(err)//hata varsa
	}
	defer response.Body.Close()//işlemin bittiğinde kanalı kapat

	bodyBytes,_ := ioutil.ReadAll(response.Body)
	//byte formatında geliyor
	bodyString := string(bodyBytes)//değişken türünü değiştirme
	fmt.Println(bodyString)
	
	var todo Todo
	json.Unmarshal(bodyBytes,&todo)//aslında bu bodystging kısmının aynısını yapıyoruz
	fmt.Println(todo)
	//jsondan todo ya çevirme
}

func Demo2(){
	todo:=Todo{1,2,"codewithzmrt",false}
	//todo dan jsona çevirme
	jsonTodo,err:=json.Marshal(todo)
	if err != nil{
		fmt.Println(err)//hata varsa
	}
	response,_:=http.Post("http://jsonplaceholder.typicode.com/todos",
	"application/json;charset=utf-8",bytes.NewBuffer(jsonTodo))
//content type application/json;charset=utf-8
	defer response.Body.Close()

	
	bodyBytes,_ := ioutil.ReadAll(response.Body)
	//byte formatında geliyor
	bodyString := string(bodyBytes)//değişken türünü değiştirme
	fmt.Println(bodyString)
	
	var todoResponse Todo
	json.Unmarshal(bodyBytes,&todoResponse)//aslında bu bodystging kısmının aynısını yapıyoruz
	fmt.Println(todoResponse)

}


/*GET isteği, verileri URL üzerinden gönderen ve sunucudan veri almak için kullanılan bir HTTP yöntemidir. GET, web tarayıcıları tarafından en sık kullanılan yöntemdir ve aşağıdaki özelliklere sahiptir:

Veri Gönderimi: GET isteği, verileri URL’nin sorgu parametreleri aracılığıyla gönderir. Örneğin, bir web sayfasına “http://example.com/page?param1=value1&param2=value2” URL’sini kullanarak GET isteği yapabilirsiniz.
Veri Sınırlaması: GET isteği, sınırlı miktarda veriyi URL üzerinden gönderir. Bu nedenle, genellikle URL sorgu parametreleriyle gönderilen veriler küçük ve sınırlıdır.
Yer İşareti (Bookmark) Kullanımı: GET ile yapılan istekler sıkça yer işareti (bookmark) olarak kaydedilir ve tarayıcı geçmişinde görüntülenir.
Önbellekleme: GET istekleri, sonuçları önbelleğe alınabilir. Bu, aynı isteğin tekrarlanması durumunda sunucuya yeniden talep göndermek yerine önbellekten verilerin alınmasını sağlar.
-------------------------------------------------------------------
POST isteği, verileri HTTP isteğinin gövdesinde (body) gönderen ve sunucuya veri göndermek için kullanılan bir HTTP yöntemidir. POST’un özellikleri şunlardır:

Veri Gönderimi: POST isteği, verileri HTTP isteğinin gövdesinde taşır. Bu, daha büyük veri miktarlarının ve dosyaların gönderilmesine olanak tanır.
Veri Sınırlaması: POST isteği, veri boyutu sınırlamalarına tabi değildir ve daha büyük veri gönderimlerine izin verir.
Güvenlik: POST, verilerin URL üzerinden görünür olmasını engellediği için, duyarlı veya gizli bilgilerin iletilmesi için daha güvenli bir seçenektir.
Önbellekleme: POST istekleri genellikle önbelleğe alınmaz, çünkü her talep farklı veriler içerebilir ve sunucuya önbellekten veri alınması riskli olabilir.

GET ve POST yöntemleri, farklı kullanım senaryolarına sahiptir. Genel olarak:

GET, sunucudan veri almak veya verileri paylaşmak için kullanılır. Örneğin, bir web sayfasını görüntülemek için veya bir arama sorgusu göndermek için GET kullanabilirsiniz.
POST, sunucuya veri göndermek için kullanılır ve gönderilen verilerin URL’de görünmesini engeller. Bu, kullanıcı kimlik bilgileri, ödeme bilgileri ve büyük veri gönderimleri gibi duyarlı verilerin iletilmesi için kullanışlıdır.
Sonuç olarak, GET ve POST, farklı amaçlar için kullanılan HTTP yöntemleridir ve her biri belirli kullanım senaryolarına sahiptir. Hangi yöntemi kullanmanız gerektiği, uygulamanızın ihtiyaçlarına ve güvenlik gereksinimlerine bağlı olarak değişebilir.


*/


