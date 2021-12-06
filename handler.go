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
