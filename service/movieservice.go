package service

import (
	"fmt"
	"iris/repository"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repository.MovieRepository
}

func NewMovieServiceManager(repo repository.MovieRepository) MovieService  {
	return &MovieServiceManager{ repo: repo }
}

func (m * MovieServiceManager)ShowMovieName() string{
	fmt.Println("我们获取到的视频名称为："+m.repo.GetMovieName())
	return "我们获取到的视频名称为："+m.repo.GetMovieName()
}
