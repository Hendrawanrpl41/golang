package main
import (
	"fmt"
)
func main(){
	//mencegah async
	//buat chanel buffer
	// hello <- true (sender/pengirim) #harus ada penerima
	// <- hello (recevier/penerima) untuk menerima dari pengirim
	//setiap chanel harus di close bila ingin pakai "for"
	hello := make(chan string, 2) //2 jumlah buffer/variabel

	hello <- "hendrawan"
	hello <- "koding"
	close(hello)

	// fmt.Println(<- hello)
	// fmt.Println(<- hello)

	for v := range hello{
		fmt.Println(v)
	}
	

}