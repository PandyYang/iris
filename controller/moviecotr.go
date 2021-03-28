package controller

import (
	"github.com/kataras/iris/v12/mvc"
	"iris/repository"
	"iris/service"
)

type MovieController struct {

}

func (c *MovieController) Get() mvc.View  {
	movieRepository := repository.NewMovieManager()
	movieService := service.NewMovieServiceManager(movieRepository)
	MovieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "index.html",
		Data: MovieResult,
	}
}

