package repositories

import "github.com/Ferriem/Iris/code/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	//Database query here
	movie := &datamodels.Movie{Name: "test"}
	return movie.Name
}
