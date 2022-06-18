package explorer

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/donggni0712/GoCoin/blockchain"
)

const templateDir string = "explorer/templates/"

var templates *template.Template

type HomeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	//tmpl := template.Must(template.ParseFiles("templates/pages/home.gohtml"))

	// if err != nil {
	// 	log.Fatal(err)
	// }
	//이를 자동화 :
	//tmpl, err := template.ParseFiles("templates/home.html") => template.Must
	data := HomeData{"Home", blockchain.GetBlockchain().AllBlocks()}

	//tmpl.Execute(w, data)
	templates.ExecuteTemplate(w, "home", data)
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("BlockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(portNum int) {
	port := fmt.Sprintf(":%d", portNum)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)
	fmt.Printf("Listrning on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
	//http.ListenAndServe : 서버 실행, 에러가 있으면 에러 리턴
	//log.Fatal : os.Exit(1)이후에 erro를 출력
	//os.Exit(1) : error code 1으로 종료
}
