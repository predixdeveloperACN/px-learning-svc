package server

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/predixdeveloperACN/px-learning-svc/model"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!!")
}

func handleGetLanguages(w http.ResponseWriter, r *http.Request) {
	outJson := `["english","tagalog","french"]`

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, outJson)
}

func handleGetLanguage(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	lang := vars["language"]

	var outString string
	if lang == "english" {
		outString = "i love you Dhen"
	} else if lang == "tagalog" {
		outString = "mahal kita Dhen"
	} else if lang == "french" {
		outString = "je t'aime Dhen"
	} else {
		outString = "ambot sa imo!!!"
	}

	w.Header().Set("content-type","application/text")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, outString)
}

func handlePostLanguage(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var language model.LanguageType
	err := json.NewDecoder(r.Body).Decode(&language)
	if err != nil {
		http.Error(w, "wrong schema", http.StatusInternalServerError)
		return
	}

	err = language.IsValid()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outputJson, err := json.Marshal(language)
	if err != nil {
		http.Error(w, "error marshal", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(outputJson))
}