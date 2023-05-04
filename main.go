package main

import (
	"net/http"
	"fmt"
	"math/rand"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for {
			var waterStatus, windStatus string
			water := rand.Intn(100)
			wind := rand.Intn(100)

			if water <= 5 {
				waterStatus = "aman"
			} else if water <= 8 {
				waterStatus = "siaga"
			} else {
				waterStatus = "bahaya"
			}

			if wind <= 6 {
				windStatus = "aman"
			} else if wind <= 15 {
				windStatus = "siaga"
			} else {
				windStatus = "bahaya"
			}

			data := map[string]interface{}{
				"water" : water,
				"wind" : wind,
			}


			// requestJson, err := json.Marshal(data)
			requestJson, err := json.MarshalIndent(data, "", "    ")

			client := &http.Client{}
			if err != nil{
				log.Fatalln(err)
			}

			url := "https://jsonplaceholder.typicode.com/posts"
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJson))
			req.Header.Set("Content-type", "application/json")
			if err != nil {
				log.Fatalln(err)
			}
			res, err := client.Do(req)
			if err != nil{
				log.Fatalln(err)
			}
			defer res.Body.Close()

			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatalln(err)
			}
			
			fmt.Println(string(requestJson))
			fmt.Printf("status water : %s\n", waterStatus)
			fmt.Printf("status wind : %s\n", windStatus)
			time.Sleep(15 * time.Second)
		}
	})

	server := new(http.Server)
	server.Addr = ":8000"
	server.ListenAndServe()
}