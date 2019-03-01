package route

import (
	"github.com/labstack/echo"
	"html/template"
	"packout/api"
	"packout/utils"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug = true

	//Serve webapp
	//e.Static("/", "pubblic/istogram.html")
	//e.Static("/", "pubblic/d3js_scatterplot.html")

	//build template
	t := &utils.Template{
		Templates: template.Must(template.ParseGlob("pubblic/main_page.html")),
	}
	//register template
	e.Renderer = t

	//e.Static("/", "pubblic/main_page.html")
	e.GET("/", api.Getfilelist())

	//v1 endpoints for json backend
	v1 := e.Group("/v1")
	{
		v1.GET("/hello", api.GetData())
		v1.GET("/touplejson", api.GetFrequencyTuple())
		v1.GET("/bytejson", api.GetFrequencyByte())

	}

	v15 := e.Group("/v2")
	{
		v15.GET("/getfilelist", api.Getfilelist())
		v15.POST("/uploadfile", api.Uploadfile())
		//v15.GET("/uploadfile", redirect)

	}

	return e
}
