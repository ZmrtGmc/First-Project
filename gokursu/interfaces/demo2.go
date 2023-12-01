package interfaces

import "fmt"
//kredi hesaplaması/ yüzdeler değşibilir
//interface örneği
type Mortgage struct{//ev kredisi
	creditPaymentTotal float64
	address string
	rate float64
}

type Car struct{
	creditPaymentTotal float64
	carInfo string
	rate float64
}
//ikisinde de aynı işlemi yapacağımız için ikisini aynı
//interface bağlıyoruz yani hesaplama için 

type CreditCalculator interface{
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate/12 
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate/12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64{
	//müşterinin birden fazla kredisi olabilir o yüzden array olarka kullanılmıştır ve interface gönderiliyor
	paymentTotal:=0.0
	for i:=0; i< len(credits);i++{
		paymentTotal=paymentTotal+credits[i].Calculate()
	}

	return paymentTotal
}

func Demo2(){
	credit1:=Mortgage{rate:10,creditPaymentTotal:100000,address:"İstanbul"}
	credit2:=Mortgage{rate:12,creditPaymentTotal:50000,address:"Ankara"}
	credit3:=Car{rate:15,creditPaymentTotal:6000,carInfo:"Polo"}

	credits:=[]CreditCalculator{credit1,credit2,credit3}
	total:=CalculateMonthlyPayment(credits)

	fmt.Println("Toplam Ödeme: ",total)
}

/*main kısmı
interfaces.Demo2()

ekran çıktısı :
toplam ödemesi 140833.333333
*/
