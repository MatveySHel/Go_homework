package server

import (
	"net/http"
	"log"
	"time"
	"math/rand"
	"task2/internal/app"
	"encoding/json"
	"encoding/base64"
)

func GetVersion(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet:
		log.Println("GET API version: v1.0.0")
		w.Write([]byte("v1.0.0"))
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func JsonDecoder(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodPost:
		input := app.EncodedBase64String{}
		if err := json.NewDecoder(r.Body).Decode(&input); err!=nil{
			log.Println("JSON input string decode error: ", err)
			return
		}
		decoded, err := base64.StdEncoding.DecodeString(input.InputBase64String)
		if err != nil{
			log.Println("Base64 decode error: ", err)
			return
		}
		json_out, err := json.Marshal(app.DecodedBase64String{OutoutString: string(decoded)})
		if err != nil{
			log.Println("JSON encode error: ", err)
			return
		}
		w.Write(json_out)
		log.Printf("Output string is '%v'", string(decoded))
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func Sleeping(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case http.MethodGet:
		ctx := r.Context()
		select{
		case <- time.After(time.Duration(rand.Intn(11)+10)*time.Second):
			log.Println("Internal Server Error")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError+rand.Intn(11))
		case <- ctx.Done():
			log.Println("Proccess Interrupted")
		}
		
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}