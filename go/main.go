package main

import (
	"github.com/rs/xid"
	"path/filepath"
	"fmt"
	"log"	
	"net/http"
	"io/ioutil"
	"os"
	// "github.com/wcharczuk/go-chart"
	"errors"
	"encoding/json"
)

// var Ydata = make([]float64, 256)
var Xdata = make([]float64, 256)

type fileData struct{
	ID xid.ID
	Name string
	Path string
	Frequency []int
}

// func drawChart(res http.ResponseWriter, req *http.Request) {
// 	readdata(filepath.Clean(os.Args[1]))
// 	graph := chart.Chart{
// 		XAxis: chart.XAxis{
// 			Style: chart.StyleShow(), //enables / displays the x-axis
// 		},
// 		YAxis: chart.YAxis{
// 			Style: chart.StyleShow(), //enables / displays the y-axis
// 		},
// 		Title: "Value distribution " + os.Args[1],
// 		Width: 1920,
// 		TitleStyle: chart.StyleShow(),
// 		Background: chart.Style{
// 			Padding: chart.Box{
// 				Top: 40,
// 			},
// 		},
// 		Series: []chart.Series{
// 			chart.ContinuousSeries{
// 			 	// XValues: Xdata,
// 			 	// YValues: Ydata,
// 			},
// 		},
// 	}

// 	res.Header().Set("Content-Type", "image/png")
// 	graph.Render(chart.PNG, res)
// }

// func drawChartWide(res http.ResponseWriter, req *http.Request) {
// 	graph := chart.Chart{
// 		 //this overrides the default.
// 		Series: []chart.Series{
// 			chart.ContinuousSeries{
// 				XValues: []float64{1.0, 2.0, 3.0, 4.0},
// 				YValues: []float64{1.0, 2.0, 3.0, 4.0},
// 			},
// 		},
// 	}

// 	res.Header().Set("Content-Type", "image/png")
// 	graph.Render(chart.PNG, res)
// }

func sendJson(res http.ResponseWriter, req *http.Request){
	request_data := fileData{
		ID : xid.New(),
		Name: filepath.Base(os.Args[1]),
		Path: filepath.Clean(os.Args[1]),
		Frequency: readdata(filepath.Clean(os.Args[1])),
	}

	json, err := json.Marshal(request_data)
	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Fprintf(res,"%s", json)
}

func readdata(file string) []int{
	ydata := make([]int, 65536)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	file_lenght := len(data)
	var is_even bool

	if file_lenght % 2 == 0{
		is_even = true
	} else{
		is_even = false
	}

	var index uint16

	for i := 0; i < file_lenght; i+=2 {
		// word := data[i:i+2]
		index = uint16(data[i]) << 8
		index += uint16(data[i+1])
		// fmt.Printf("%#x\n", index)
		// used to debug
		// if i == 10000{
		// 	break
		// }
		ydata[int(index)]++
	}
	fmt.Print(is_even)

	// var temp_byte int
	// var is_qword bool

	// is_qword = false
	// for _,byte := range data {

	// 	if is_qword{
	// 		ydata[int(temp_byte+byte)]++
	// 	}

	// 	if !is_qword {
	// 		temp_byte = byte >> 8
	// 		is_qword = true
	// 	}
	// }

	fmt.Print(ydata)
	return ydata
}

func serveRoot(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "html/root.html")
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

	// http.HandleFunc("/", drawChartWide)
	http.HandleFunc("/favicon.ico", http.NotFound)
	// http.HandleFunc("/wide", drawChart)
	http.HandleFunc("/json", sendJson)
	http.HandleFunc("/view", serveRoot)
	// http.Handle("/", http.ServeFile("./html/"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
