package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

func main() {
	const songsDir = "songs"
	const port = 8080

	// Register custom MIME types for HLS streaming
	mime.AddExtensionType(".m3u8", "application/vnd.apple.mpegurl")
	mime.AddExtensionType(".ts", "video/MP2T")

	// Serve files with proper headers
	http.Handle("/", addHeaders(http.FileServer(http.Dir(songsDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", songsDir, port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// Middleware to add headers for CORS and set MIME types dynamically
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		ext := filepath.Ext(r.URL.Path)
		if mimeType := mime.TypeByExtension(ext); mimeType != "" {
			w.Header().Set("Content-Type", mimeType)
		}

		h.ServeHTTP(w, r)
	}
}
