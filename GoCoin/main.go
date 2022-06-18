package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/donggni0712/GoCoin/blockchain"
)

const port string = ":4000"

type HomeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//이를 자동화 :
	//tmpl, err := template.ParseFiles("templates/home.html") => template.Must
	Data := HomeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(w, Data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listrning on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
	//http.ListenAndServe : 서버 실행, 에러가 있으면 에러 리턴
	//log.Fatal : os.Exit(1)이후에 erro를 출력
	//os.Exit(1) : error code 1으로 종료
}
