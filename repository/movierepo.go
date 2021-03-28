package repository

import "iris/model"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}
func (m *MovieManager) GetMovieName() string  {
	movie:=&model.Movie{Name:"模拟的movie"}
	return movie.Name
}