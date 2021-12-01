package urlshort

import (
	"net/http"

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

func YAMLHandler(yaml_in []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. parse the yaml
	var pus []pathUrl
	err := yaml.Unmarshal(yaml_in, &pus)
	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string)
	for _, pu := range pus {
		pathsToUrls[pu.Path] = pu.URL
	}
	// 2. Convert yaml array into map
	//  3. return a map handler using the map

	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url`
}
