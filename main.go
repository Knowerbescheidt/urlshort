package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	db := ConnectToDb()
	pathsToUrls := RetrieveData(db)
	//yaml_path := flag.String("yaml", "", "please specify location for the yaml file you want to use")
	// json_path := flag.String("json", "", "please specify location for the json file you want to use")
	// // do not forget to parse flags
	// flag.Parse()
	// jsonFile, err := ioutil.ReadFile(*json_path)
	// if err != nil {
	// 	os.Exit(2)
	// }
	// pathsToUrls := make(map[string]string)
	// err2 := json.Unmarshal(jsonFile, &pathsToUrls)
	// if err2 != nil {
	// 	log.Fatal("Error during parsing yaml data ", err2)
	// }

	mux := defaultMux()
	// pathsToUrls := map[string]string{
	// 	"/git":        "https://github.com/",
	// 	"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	// }
	mapHandler := MapHandler(pathsToUrls, mux)
	// when writing yaml make sure there is no indentation tab just spaces
	// 	yaml_str := `
	// - path: /urlshort
	//   url: https://github.com/Knowerbescheidt/urlshort
	// - path: /gin-swagger
	//   url: https://github.com/swaggo/gin-swagger
	// `

	// mapHandler, err := urlshort.MapHandler(mapHandler, hello)
	// if err != nil {
	// 	panic(err)
	// }

	http.ListenAndServe(":8080", mapHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
