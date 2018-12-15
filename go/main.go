package main

import (
	"github.com/rs/xid"
	"path/filepath"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"github.com/wcharczuk/go-chart"
	"errors"
	"encoding/json"
)

// var Ydata = make([]float64, 256)
var Xdata = make([]float64, 256)

type fileData struct{
	ID xid.ID
	Name string
	Path string
	Ordered_frequency []int
}

func drawChart(res http.ResponseWriter, req *http.Request) {
	readdata()
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.StyleShow(), //enables / displays the x-axis
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(), //enables / displays the y-axis
		},
		Title: "Value distribution " + os.Args[1],
		Width: 1920,
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
			 	// XValues: Xdata,
			 	// YValues: Ydata,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func drawChartWide(res http.ResponseWriter, req *http.Request) {
	graph := chart.Chart{
		 //this overrides the default.
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func sendJson(res http.ResponseWriter, req *http.Request){
	request_data := fileData{
		ID : xid.New(),
		Name: filepath.Base(os.Args[1]),
		Path: filepath.Clean(os.Args[1]),
		Ordered_frequency: readdata(),
	}

	json, err := json.Marshal(request_data)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Fprintf(res,"%s", json)
}

func readdata() []int{
	ydata := make([]int, 256)

	data, err := ioutil.ReadFile(filepath.Clean(os.Args[1]))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}


	for _,byte := range data {
		ydata[int(byte)]++
	}

	fmt.Print(ydata)
	return ydata
}

func main() {
	if len(os.Args)< 2 {
		err := errors.New("Not enough arguments, missing input filename")
		log.Fatal(err)
		os.Exit(1)
	}

	// readdata()

	for i := 0; i < 256; i++ {
		Xdata[i] = float64(i)
	}

	http.HandleFunc("/", drawChartWide)
	http.HandleFunc("/wide", drawChart)
	http.HandleFunc("/json", sendJson)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
