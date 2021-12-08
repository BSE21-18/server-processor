package main


import (
    "fmt"
    "log"
	"flag"
	"strings"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type ProcessedResult struct {
  Date string `json:"date"`
  Time string `json:"time"`
  Sniffer string `json:"sniffer"` //the topic ie sniffer device unique ID
  Disease string `json:"disease" default:"Late blight"`
  PlantStatus string `json:"plantStatus"` //healthy, mild +ve, moderate +ve, severe +ve
  Recommendation string `json:"recommendation"`
}

//data: {"mq1":"345PH","mq2":"345PH","mq3":"345PH","mq4":"345PH","mq5":"345PH","mq135":"345PH"}
type SnifferMessage struct {
    Data string `json:"data"`
}

func processing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var msg SnifferMessage
	_ = json.NewDecoder(r.Body).Decode(&msg)
	
	fmt.Println("Processor endpoint: Processing ...")
	fmt.Println("Sniffer Message: => msg.Data = ", msg.Data)
	
	//TODO: process and attach processed data to the result
	
	result := &ProcessedResult{
	    Date: "2021-12-26",
        Time: "09:40:04",
        Sniffer: "DV00023",
        Disease: "Late blight",
        PlantStatus: "mild +ve", //healthy, mild +ve, moderate +ve, severe +ve
        Recommendation: "Spray with the recommende chemicals, or contact experts for help.",
	}
	json.NewEncoder(w).Encode(result)
}

func getRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/processor", processing).Methods("POST")
	return router
}

func main() {
    //++++| os.Args |+++++
    wsEndPoint := ":7500" 
    addr := flag.String("addr", wsEndPoint, "Processor endpoint address") 
    flag.Parse()
    //++++++++++++++++++++
    
    fmt.Println("DATAVOC Processor endpoint listening on port: "+(strings.Split(wsEndPoint,":")[1])) 
    log.Fatal(http.ListenAndServe(*addr, getRouter()))
}








