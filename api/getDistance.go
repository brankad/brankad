package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func GetDistance(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data DataInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		var responseData DataOutput
		responseData.Result = "nok"
		responseData.Text = "problem with user json data"
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	//fmt.Println("kako hvatam params")
	//fmt.Println(data.From)
	//fmt.Println(data.To)


	url:= "http://transport.opendata.ch/v1/connections?from="+data.From+"&to="+data.To


	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	respData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(respData, &responseObject)


	var arr1 RespToArray
	var fixedValue = responseObject.Connections[0].ToC.ArrivalTimestamp + 1
	for i := 0; i < len(responseObject.Connections); i++ {
		timestampDiff := (responseObject.Connections[i].ToC.ArrivalTimestamp - responseObject.Connections[i].FromC.DepartureTimestamp)/60

		fmt.Println ("Minutes diff ")
		fmt.Println (timestampDiff)
		fmt.Println (responseObject.Connections[i].ToC.ArrivalTimestamp)


		fmt.Println("sta mi je i")
		fmt.Println(i)

		fmt.Println("sta mi je fixed value")
		fmt.Println(fixedValue)


		if (fixedValue > timestampDiff){
			unixTimeUTC:=time.Unix(int64(responseObject.Connections[i].FromC.DepartureTimestamp), 0) //gives unix time stamp in utc

			unitTimeInRFC3339 := unixTimeUTC.Format(time.RFC1123)

			arr1.DiffMin = timestampDiff
			arr1.DepartureTime = unitTimeInRFC3339 //responseObject.Connections[i].FromC.Departure
		}

		fixedValue = timestampDiff

	}


	timer := time.Now()
	time.Sleep(1 * time.Second)
	end := time.Since(timer)
	fmt.Println("processing takes: " + end.String())
	var responseData DataOutput
	responseData.Result = "ok"
	responseData.Text = "everything went smooth"
	responseData.Time = arr1.DepartureTime
	responseData.Duration = arr1.DiffMin
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}
