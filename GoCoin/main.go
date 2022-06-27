package main

import (
	explorer "github.com/donggni0712/GoCoin/explorer/templates"
	rest "github.com/donggni0712/GoCoin/restapi"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
	// rest에서도 '/'를 다루고 explorer에서도 '/'를 다뤄서 오류가 남
	//port가 달라도 url함수를 매칭해주는 건 똑같음
	//Mux = multiflexer url을 지켜보고 원하는 함수를 실행
	//
}
