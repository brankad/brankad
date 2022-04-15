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
		responseData.Result = "not ok"
		responseData.Text = "problem with user json data"
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}

	if data.From == "" || data.To == ""{
		var responseData DataOutput
		responseData.RespCode = 0
		responseData.Result = "not ok"
		responseData.Text = "all fields must be filled"
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}

	url:= SwissTransportURL + DistanceAPI + "?from="+data.From+"&to="+data.To

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	respData, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		log.Fatal(err1)
	}

	var responseObject Response
	json.Unmarshal(respData, &responseObject)


	var arr1 RespToArray
	var fixedValue = responseObject.Connections[0].ToC.ArrivalTimestamp + 1
	for i := 0; i < len(responseObject.Connections); i++ {
		timestampDiff := (responseObject.Connections[i].ToC.ArrivalTimestamp - responseObject.Connections[i].FromC.DepartureTimestamp)/60

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

	responseData := responseOk(arr1)
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}

func responseOk(arr1 RespToArray) (data DataOutput) {

	var responseData DataOutput
	responseData.RespCode = 1
	responseData.Result = "ok"
	responseData.Text = "it is working"
	responseData.Time = arr1.DepartureTime
	responseData.Duration = arr1.DiffMin

	return responseData
}

func responseNotOk(arr1 RespToArray) (data DataOutput) {

	var responseData DataOutput
	responseData.RespCode = 0
	responseData.Result = "error occured"
	responseData.Text = "check the logs"
	responseData.Time = " / "
	responseData.Duration = 0

	return responseData
}