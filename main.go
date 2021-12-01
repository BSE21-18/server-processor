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

//{"gas1":"345PH","gas2":"345PH","gas3":"345PH","gas4":"345PH","gas5":"345PH","gas6":"345PH"}
type SnifferMessage struct {
    Gas1 string `json:"gas1"`
    Gas2 string `json:"gas2"`
    Gas3 string `json:"gas3"`
    Gas4 string `json:"gas4"`
    Gas5 string `json:"gas5"`
    Gas6 string `json:"gas6"`
}

func processing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var msg SnifferMessage
	_ = json.NewDecoder(r.Body).Decode(&msg)
	
	fmt.Println("Processor endpoint: Processing ...")
	fmt.Println("Sniffer Message: => ", msg)
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








