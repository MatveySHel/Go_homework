package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"task2/internal/app"
	"time"
	"errors"
)


func GetRequestVersion(){
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/version", nil)
	if err != nil{
		fmt.Println("Request error:", err)
		return 
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Do error:", err)
		return 
	}
	defer resp.Body.Close()
	resp_content,err := io.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("Reading error:", err)
		return
	}
	fmt.Println(string(resp_content))
}


func PostRequestBase64(inputString string){
	json_input, err := json.Marshal(app.EncodedBase64String{InputBase64String: inputString})
	if err != nil{
		fmt.Println("Json marshal error:", err)
		return
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/decode", bytes.NewReader(json_input))
	if err != nil{
		fmt.Println("Request error:", err)
		return 
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Do error:", err)
		return 
	}

	defer resp.Body.Close()
	resp_content, err := io.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("Reading error:", err)
		return
	}

	resp_object := app.DecodedBase64String{}
	if err := json.Unmarshal(resp_content, &resp_object); err != nil{
		fmt.Println("Json unmarshal error:", err)
		return
	}
	fmt.Println(resp_object.OutoutString)
}


func GetRequestHarOp(){
	ctx, cansel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cansel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/hard-op", nil)
	if err != nil{
		fmt.Println("Request error:", err)
		return 
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Do error:", err)
		if errors.Is(err, context.DeadlineExceeded){
			fmt.Println("false")
		}
		return 
	}
	defer resp.Body.Close()
	fmt.Println(true, resp.StatusCode)
}


func LaunchClient(){
	GetRequestVersion()
	PostRequestBase64("eHh4")
	GetRequestHarOp()
}