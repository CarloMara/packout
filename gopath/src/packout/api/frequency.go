package api

import (
	"github.com/labstack/echo"
	"net/http"
	"packout/model"
)

// txt sample /home/carlo/Documents/packout/gopath/src/packout/samples/shakespeare.txt 5MB

// elf sample /home/carlo/Documents/chromecast/phillips/sections/UTDchrome/cast_shell 40MB

// elf sample /home/carlo/Documents/packout/gopath/src/packout/samples/a.out 17KB

// so sample /home/carlo/Documents/packout/gopath/src/packout/samples/libvirglrenderer.so 459 KB

// so sample /home/carlo/Documents/packout/gopath/src/packout/samples/libwireshark.so.10.1.6



func GetFrequencyTuple() echo.HandlerFunc{
	return func(context echo.Context) error {

		//echo.Logger.Debug("GetFrequencyTuple")

		fd := model.InitFile("/home/carlo/Documents/packout/gopath/src/packout/samples/libvirglrenderer.so")
		context.Logger().Debug(fd)
		fd.SampleTuple()
		context.Logger().Debug(fd.Data)

		//json, err := json.Marshal(request_data)
		//if err != nil{
		//	//log.Fatal(err)
		//	os.Exit(1)
		//}

		return context.JSON(http.StatusOK, fd.Data)
	}
}


func GetFrequencyByte() echo.HandlerFunc{
	return func(context echo.Context) error {
		fd := model.InitFile("/home/carlo/Documents/packout/gopath/src/packout/samples/a.out")
		context.Logger().Debug(fd)
		fd.SampleByte()
		context.Logger().Debug(fd.Frequency)

		//json, err := json.Marshal(request_data)
		//if err != nil{
		//	//log.Fatal(err)
		//	os.Exit(1)
		//}
		return context.JSON(http.StatusOK, fd)
	}
}