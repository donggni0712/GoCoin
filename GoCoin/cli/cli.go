package cli

import (
	"flag"
	"fmt"
	"os"

	explorer "github.com/donggni0712/GoCoin/explorer/templates"
	rest "github.com/donggni0712/GoCoin/restapi"
)

func usage() {
	fmt.Printf("Wellcom\n\n")
	fmt.Printf("Please use the following flags\n\n")
	fmt.Printf("-port=4000: 	Set the PORT of ther server\n")
	fmt.Printf("-mode=rest:		Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	// go explorer.Start(3000)
	// rest.Start(4000)
	// rest에서도 '/'를 다루고 explorer에서도 '/'를 다뤄서 오류가 남
	//port가 달라도 url함수를 매칭해주는 건 똑같음
	//Mux = multiflexer url을 지켜보고 원하는 함수를 실행
	//

	/*
		//CLI 만들 때 cobra framework 사용하면 편함
		//fmt.Println(os.Args)
		if len(os.Args) < 2 {
			usage()
		}

			//FlagSet : Go에게 어떤 command가 어떤 flag를 가질 것인지 알려줌
			// go run main.go rest -port=4000 -mode=https -v -t .. 처럼 argument가 많을 때 유용
			rest := flag.NewFlagSet("rest", flag.ExitOnError)

			portFlag := rest.Int("port", 4000, "Set the port of ther server")

			switch os.Args[1] {
			case "explorer":
				fmt.Println("Start Explorer")
			case "rest":
				rest.Parse(os.Args[2:]) // Args에서 "port"를 찾고 값이 int인지 체크하고 넣어줌
			default:
				usage()
			}
			if rest.Parsed() {
				fmt.Println(portFlag)
				fmt.Printf("Start server %d", *portFlag)
			}
	*/

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}

	fmt.Println(*port, *mode)

}
