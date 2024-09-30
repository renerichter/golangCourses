package main

import (
"net/http"
)

func respondWithJSON(w htt.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)
	if err!= nil{
		log.Printf("Faild to marshal JSON response: %v",payload)
		w.WriteHeader(500)
		return 
	}
	w.Header().Add("Content-Type","application/json")
}