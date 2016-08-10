package web

import (
	"html/template"
	"hyperagent/log"
	"net/http"
	"os"
	"runtime/debug"
)

var templates map[string]*template.Template

const (
	TEMPLATE_DIR = "./views"
	ASSETS_DIR   = "./assets"
)

func init() {
	templates = make(map[string]*template.Template)
}

func HandlerWebSite(mux *http.ServeMux) {
	log.Debug("HandlerRestServices")
	staticDirHandler(mux, "/assets/", ASSETS_DIR, 0)
}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func safeHandlerViews(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, "ViewError", http.StatusInternalServerError)
				w.WriteHeader(http.StatusInternalServerError)
				renderHtml(w, "error", nil)
				log.Warn("WARN: panic in %v. - %v", fn, e)
				log.Debug(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags) == 0 {
			if exists := isExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
