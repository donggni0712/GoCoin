package main

import (
	"fmt"
	"os"

	explorer "github.com/donggni0712/GoCoin/explorer/templates"
	rest "github.com/donggni0712/GoCoin/restapi"
)

func usage() {
	fmt.Printf("Wellcom\n\n")
	fmt.Printf("Please use the following commands\n\n")
	fmt.Printf("explorer: 	Start the Explorer\n")
	fmt.Printf("rest:		Start the REST API (recommned)\n")
	os.Exit(0)
}

func main() {
	// go explorer.Start(3000)
	// rest.Start(4000)
	// rest에서도 '/'를 다루고 explorer에서도 '/'를 다뤄서 오류가 남
	//port가 달라도 url함수를 매칭해주는 건 똑같음
	//Mux = multiflexer url을 지켜보고 원하는 함수를 실행
	//

	//CLI 만들 때 cobra framework 사용하면 편함
	//fmt.Println(os.Args)
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
		explorer.Start(3000)
	case "rest":
		fmt.Println("Start REST API")
		rest.Start(4000)
	default:
		usage()
	}
}
