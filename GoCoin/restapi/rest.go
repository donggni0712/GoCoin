package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/donggni0712/GoCoin/blockchain"
	"github.com/donggni0712/GoCoin/utils"
	"github.com/gorilla/mux"
)

type url string

var port string

//MarshalText Interface
func (u url) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("http://localhost%s%s", port, u)), nil
}

//String Interface => URL이 string으로 바뀔때 패키지가 인터페이스를 호출해줌
// func (u URL) String() string{
// 	return "hi"
// }

type urlDescription struct {
	URL         url    `json:"url"` //field struct tag
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty : 빈 값은 출력x #띄어쓰면 안됨
} // `json:"-"` => field를 무시

type AddBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMesseage string `json:"errorMesseage"`
}

func documentation(w http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "See All Blocks",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "message:string",
		},
	}

	w.Header().Add("Content-Type", "application/json")
	// firfox등에서 이쁘게 꾸며줌. json이라고 인지

	// json.Marshal : struct => JSON
	// b는 []byte

	/*
		b, err := json.Marshal(data)
		utils.HandleErr(err)
		fmt.Fprintf(w, "%s", b)
		위 3줄 = json.NewEncoder(w).Encode(data)
	*/
	json.NewEncoder(w).Encode(data)

}

func blocks(w http.ResponseWriter, r *http.Request) {
	//request의 method를 구분하여 처리
	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		w.WriteHeader(http.StatusOK)
	}
}

func block(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	block, err := blockchain.GetBlockchain().GetBlock(id)
	encoder := json.NewEncoder(w)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

//아래 함수는 http.Handler를 리턴해야하는데 http.HandlerFunc는 Type임
// http.Handler는 interface
// adapter에게 적절한 argument를 보내주면 dapter는 필요한 것들을 구현해 줌.
// adapter pattern 검색해보기
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "aplication/json")
		next.ServeHTTP(w, r)
	})
}

func Start(aPort int) {
	// port = fmt.Sprintf(":%d", aPort)
	// http.HandleFunc("/", documentation)
	// http.HandleFunc("/blocks", blocks)
	// fmt.Printf("Listening on http://localhost%s\n", port)
	// log.Fatal(http.ListenAndServe(port, nil))

	// 새 mux를 생성해주고 http => handler로 변경해준다.
	//nil => handler
	// handler := http.NewServeMux()
	// port = fmt.Sprintf(":%d", aPort)
	// handler.HandleFunc("/", documentation)
	// handler.HandleFunc("/blocks", blocks)
	// fmt.Printf("Listening on http://localhost%s\n", port)
	// log.Fatal(http.ListenAndServe(port, handler))

	//gorilla mux 사용
	//다른 method로부터 보호해줄 수 있고 parameter 넘겨줄 수 있음
	router := mux.NewRouter()

	//middlerware 사용
	router.Use(jsonContentTypeMiddleware)

	port = fmt.Sprintf(":%d", aPort)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/block/{id:[0-9]+}", block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}
