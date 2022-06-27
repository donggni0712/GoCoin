package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func main() {
	//explorer.Start(4000)
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
