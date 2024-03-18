package storage

import (
	"errors"
	"filmlib/model"
	"fmt"
)

func (s *Storage) GetMovies(sortBy, sortType string) ([]model.Movie, error) {
	if sortBy == "" {
		sortBy = "rating"
	}

	if sortType == "" {
		sortType = "DESC"
	}

	if sortBy != "title" && sortBy != "rating" && sortBy != "release_year" {
		return nil, errors.New("no such sorting")
	}

	if sortType != "ASC" && sortType != "DESC" {
		return nil, errors.New("no such sorting order")
	}

	query := fmt.Sprintf("SELECT * FROM movie ORDER BY %s %s ", sortBy, sortType)

	var movies []model.Movie

	err := s.DB.Select(&movies, query)

	return movies, err
}

func (s *Storage) SearchMoviesByTitle(partOfMovieTitle string) ([]model.Movie, error) {
	query := `SELECT * FROM movie WHERE title LIKE %` + partOfMovieTitle + `%')`

	var movies []model.Movie

	err := s.DB.Select(&movies, query)

	return movies, err
}

func (s *Storage) SearchMoviesByActor(partOfActorName string) ([]model.Movie, error) {
	query := `WITH actor_ids AS (SELECT id FROM actor WHERE name LIKE '%` + partOfActorName + `%')
			SELECT *
			FROM movie
			WHERE id IN (SELECT movie_id FROM cast_record WHERE actor_id IN (SELECT id FROM actor_ids))`

	var movies []model.Movie

	err := s.DB.Select(&movies, query)

	return movies, err
}

func (s *Storage) CreateMovie(movie model.Movie) error {
	_, err := s.DB.NamedExec(
		`INSERT INTO movie (id, title,  rating, release_year) VALUES (:id, :title, :rating, :release_year)`,
		movie)
	return err
}

func (s *Storage) UpdateMovie(movie model.Movie) error {
	_, err := s.DB.NamedExec(
		"UPDATE movie SET title = :title, rating = :rating,release_year = :release_year WHERE id = :id",
		movie)
	return err
}

func (s *Storage) DeleteMovie(id string) error {
	_, err := s.DB.Exec("DELETE FROM movie WHERE id = $1", id)
	return err
}
