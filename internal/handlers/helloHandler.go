package handlers

import(
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request){


	if r.Method != http.MethodGet{
		http.Error(w, "only Get is allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _  = w.Write([]byte("Hello, from Go net/http server"))
} 