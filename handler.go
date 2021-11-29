package urlshort

import "net/http"

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
