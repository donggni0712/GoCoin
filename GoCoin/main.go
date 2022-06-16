package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from home!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listrning on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
	//http.ListenAndServe : 서버 실행, 에러가 있으면 에러 리턴
	//log.Fatal : os.Exit(1)이후에 erro를 출력
	//os.Exit(1) : error code 1으로 종료
}
