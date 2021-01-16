package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Handler struct {
	template string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("tmpl").Parse(h.template)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("err: %s", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, nil); err != nil {
		w.Write([]byte(fmt.Sprintf("err: %s", err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	t := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>App</title>
</head>
<body>
	Hello!
</body>
</html>`
	h := &Handler{template: t}
	http.Handle("/", h)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
