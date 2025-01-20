package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	/*"context"
	  "os"
	  "encoding/json"
	  "time"

	  "github.com/lrstanley/go-ytdlp" */)

func main()  {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Access-Control-Allow-Origin", "*")
    res.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if req.Method == "OPTIONS" {
        res.WriteHeader(http.StatusOK)
        return
    }

    if req.Method != "POST" {
        http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var data struct {
        Link string `json:"link"` 
    }

    decoder := json.NewDecoder(req.Body)
    if err := decoder.Decode(&data); err != nil {
        http.Error(res, err.Error(), http.StatusBadRequest)
        return
    }
    
    fmt.Printf("Received link: %s", data.Link)

    res.Header().Set("Content-Type", "application/json")
    json.NewEncoder(res).Encode(map[string]string{
        "message": "Link received succesfully",
    })
}
