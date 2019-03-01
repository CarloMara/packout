package api

import (
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io"
	"net/http"
	"packout/db"
	"packout/model"
)

//{name:}

func Getfilelist() echo.HandlerFunc {
	return func(context echo.Context) error {

		mongo := db.Init()
		bucket, err := gridfs.NewBucket(mongo.Client.Database("testing"))
		if err != nil {
			context.Logger().Error("Newbucket fail")
		}

		cur, err := bucket.Find(bson.D{})
		if err != nil {
			context.Logger().Error("Bucket Find fail")
		}

		var Projects []model.Project

		for cur.Next(mongo.Contex) {

			project := &model.Project{}
			if err := cur.Decode(project); err != nil {
				context.Logger().Error(err)

			}

			Projects = append(Projects, *project)
		}

		return context.Render(http.StatusOK, "main_page.html", Projects)
	}
}

func Uploadfile() echo.HandlerFunc {
	return func(context echo.Context) error {

		form, err := context.MultipartForm()
		if err != nil {
			return err
		}
		files := form.File["file"]
		context.Logger().Debug("name ", files)

		mongo := db.Init()
		bucket, err := gridfs.NewBucket(mongo.Client.Database("testing"))

		if err != nil {
			context.Logger().Error("Newbucket Fail")
		}

		for _, file := range files {
			// Source
			src, err := file.Open()
			//Filename Size Content

			if err != nil {
				context.Logger().Error(err)
			}

			dst, err := bucket.OpenUploadStream(file.Filename)

			if err != nil {
				context.Logger().Error(err)
			}

			if _, err = io.Copy(dst, src); err != nil {
				return err
			}

			defer dst.Close()

			//context.Logger().Debug("name ", src)
			defer src.Close()

		}

		//mongo := db.Init()
		//collection := mongo.Client.Database("testing").Collection("numbers")
		//res, err := collection.InsertOne(mongo.Contex, bson.M{"name": "pi", "value": 3.14159})
		////id := res.InsertedID
		//if err != nil{
		//	panic(fmt.Sprint("panic on insertion"))
		//}
		return context.Redirect(http.StatusFound, "/")
	}
}
