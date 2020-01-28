package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shushutochako/pkg/signinapple"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "error")
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error")
		return
	}
	vtp := signinapple.ValidateTokensParams{}
	err = json.Unmarshal(body, &vtp)

	apClient := signinapple.NewClient()
	apClient.ValidateTokens(signinapple.ValidateTokensParams{
		AuthorizationCode: vtp.AuthorizationCode,
	})
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
