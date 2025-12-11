package static

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"flexsupport/internal/config"
	// "fmt"
	// "strings"
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	// "path/filepath"
)

//go:embed all:assets
var Static embed.FS

//go:embed all:public
var Public embed.FS

func AssetRouter(cfg *config.Config) http.Handler {
	if cfg.Environment == config.DEV {

		fmt.Println("ASSETS DEV")
		return http.FileServer(http.Dir("./static/assets"))
	}
	st, err := fs.Sub(Static, "assets")
	if err != nil {
		panic(err)
	}
	handler := http.FileServer(http.FS(st))
	return handler
}

func PublicRouter(cfg *config.Config) http.Handler {
	if cfg.Environment == config.DEV {
		fmt.Println("PUBLIC DEV")
		return http.FileServer(http.Dir("./static/public"))
	}
	st, err := fs.Sub(Public, "public")
	if err != nil {
		panic(err)
	}
	handler := http.FileServer(http.FS(st))
	return handler
}

// func Handler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Get the static filesystem
// 		fmt.Println("ASSETS")
// 		st, err := fs.Sub(Static, "assets")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			panic(err)
// 		}
// 		static := http.FileServer(http.FS(st))

// 		// Try to open the requested file path directly
// 		requestedPath := strings.TrimPrefix(r.URL.Path, "/")
// 		fmt.Println(requestedPath)

// 		f, err := st.Open(requestedPath)
// 		fileExists := err == nil
// 		if f != nil {
// 			f.Close()
// 		}

// 		ext := filepath.Ext(r.URL.Path)
// 		fmt.Println(ext)
// 		// Handle all assets
// 			switch ext {
// 			case ".json":
// 				w.Header().Set("Cache-Control", "no-cache")
// 			case ".txt":
// 				w.Header().Set("Cache-Control", "no-cache")
// 			default:

// 				w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
// 			}

// 			if !fileExists {
// 				http.NotFound(w, r)
// 				return
// 			}

// 		w.Header().Set("Vary", "Accept-Encoding")
// 		static.ServeHTTP(w, r)
// 	}
// }
