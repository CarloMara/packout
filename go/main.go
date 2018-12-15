package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
	"github.com/wcharczuk/go-chart"
	"errors"
)

var Ydata = make([]float64, 256)
var Xdata = make([]float64, 256)

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
			 	XValues: Xdata,
			 	YValues: Ydata,
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

func drawPage(res http.ResponseWriter, req *http.Request){}

func readdata() error{
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		return err
	}
	for _,byte := range data {
		Ydata[int(byte)]++
	}
	fmt.Print(Ydata)

	return nil
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
