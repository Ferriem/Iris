package controllers

import (
	"github.com/Ferriem/Iris/code/repositories"
	"github.com/Ferriem/Iris/code/services"
	"github.com/kataras/iris/v12/mvc"
)

type MovieController struct {
}

func (m *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepository)
	MoiveResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: MoiveResult,
	}

}
