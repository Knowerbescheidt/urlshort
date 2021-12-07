package urlshort

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-yaml/yaml"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// the ok is true or false depending on if it finds a value in the map
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		// if we can mathc a path
		// redirect to it
		// else
		fallback.ServeHTTP(w, r)
	}
}

func YamlHandler(yaml_in []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrlSlice, err := parseYaml(yaml_in)
	if err != nil {
		return nil, err
	}
	pathsToUrls := convertArrayToMap(pathUrlSlice)
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pus []pathUrl
	err := yaml.Unmarshal(data, &pus)
	if err != nil {
		return nil, err
	}
	return pus, err
}

func convertArrayToMap(pus []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pus {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url`
}

func DbHandler() {
	connectToDb()
}

func connectToDb() {
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "gophercise"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}
	sql_get := "SELECT * FROM pathmap"
	data, err := db.Query(sql_get)
	if err != nil {
		fmt.Println("Erroor with executing query")
		os.Exit(2)
	}
	fmt.Println(data)
}
