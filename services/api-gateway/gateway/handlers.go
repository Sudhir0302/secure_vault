package gateway

import (
	"net/http"
)

func (rev *Gateway) AuthHandler(w http.ResponseWriter, r *http.Request) {
	rev.ServeHTTP(w, r)
}

func (rev *Gateway) ShareHandler(w http.ResponseWriter, r *http.Request) {
	rev.ServeHTTP(w, r)
}

func (rev *Gateway) StoreHandler(w http.ResponseWriter, r *http.Request) {
	rev.ServeHTTP(w, r)
}
