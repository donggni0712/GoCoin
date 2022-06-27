package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/donggni0712/GoCoin/blockchain"
	"github.com/donggni0712/GoCoin/utils"
)

const port string = ":4000"

type URL string

//MarshalText Interface
func (u URL) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("http://localhost%s%s", port, u)), nil
}

//String Interface => URL이 string으로 바뀔때 패키지가 인터페이스를 호출해줌
// func (u URL) String() string{
// 	return "hi"
// }

type URLDescription struct {
	URL         URL    `json:"url"` //field struct tag
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty : 빈 값은 출력x #띄어쓰면 안됨
} // `json:"-"` => field를 무시

type AddBlockBody struct {
	Message string
}

func documentation(w http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "See Documentation",
			Payload:     "data:string",
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

func main() {
	//explorer.Start(4000)
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
