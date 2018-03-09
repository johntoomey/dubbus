package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const url string = "https://data.dublinked.ie/cgi-bin/rtpi"

func realtimebusinfo(stopid int) {

	// Result : comment
	type Result struct {
		ArrivalDateTime        string `json:"arrivaldatetime"`
		DueTime                string `json:"duetime"`
		DepartureDateTime      string `json:"departuretime"`
		DepartureDueTime       string `json:"departureduetime"`
		ScheduledArrivalTime   string `json:"scheduledarrivaltime"`
		ScheduledDepartureTime string `json:"scheduleddeparturetime"`
		Destination            string `json:"destination"`
		Origin                 string `json:"origin"`
		Direction              string `json:"direction"`
		Operator               string `json:"operator"`
		AdditionalInfo         string `json:"additionalinformation"`
		LowFloorStat           string `json:"lowfloorstatus"`
		Route                  string `json:"route"`
		SourceTimeStamp        string `json:"destination"`
	}
	// RealTimeBusInfo : comment
	type RealTimeBusInfo struct {
		ErrorCode       string `json:"errorcode"`
		ErrorMessage    string `json:"errormessage"`
		NumberOfResults int    `json:"numberofresults"`
		StopID          string `json:"stopid"`
		Timestamp       string `json:"timestamp"`
		Results         []*Result
	}

	rtpiURL := fmt.Sprintf("%srealtimebusinformation?stopid=%d&format=json", url, stopid)

	client := http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Get(rtpiURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	var rtbi RealTimeBusInfo

	if err := json.NewDecoder(resp.Body).Decode(&rtbi); err != nil {
		fmt.Println(err)
		return
	}

	if rtbi.ErrorCode != "0" {
		fmt.Println(rtbi.ErrorMessage)
		return
	}

	for i, r := range rtbi.Results {
		arrival := strings.Fields(r.ArrivalDateTime)
		fmt.Printf("%d Route %s due %s (%s mins) %s\n", i, r.Route, arrival[1], r.DueTime, r.Destination)
	}
}

func main() {
	realtimebusinfo(317)
	//realtimebusinfo(6089)
}
