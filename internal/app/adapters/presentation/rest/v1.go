package restv1

import (
	"log/slog"
	"net/http"
)

type V1API struct{}

func (rv *V1API) Start() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("webserver crashed", "error", err)
	}
}
