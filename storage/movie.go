package storage

import (
	"FilmLibrary/model"
	"errors"
)

func (s *Storage) CreateMovie(movie model.Movie) error {
	result, err := s.DB.NamedExec(`INSERT INTO movie (idgen, title,  rating, release_year) VALUES (:idgen, :title, :rating, :release_year)`, movie)
	if err != nil {
		return err
	}

	countOfModifiedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if countOfModifiedRows == 0 {
		return nil
	}

	return nil
}

func (s *Storage) GetMovie(movieId string) (model.Movie, error) {
	var movie model.Movie

	err := s.DB.Get(&movie, "SELECT * FROM movie WHERE idgen = $1", movieId)
	if err != nil {
		return model.Movie{}, err
	}

	return movie, nil
}

func (s *Storage) UpdateMovie(movie model.Movie) (model.Movie, error) {
	result, err := s.DB.NamedExec(`UPDATE movie SET title = :title, rating = :rating,release_year = :release_year WHERE idgen = :idgen`, movie)
	if err != nil {
		return model.Movie{}, err
	}

	countOfModifiedRows, err := result.RowsAffected()
	if err != nil {
		return model.Movie{}, err
	}

	if countOfModifiedRows == 0 {
		return model.Movie{}, errors.New("Nothing has been changed")
	}

	var response model.Movie

	err = s.DB.Get(&response, "SELECT * FROM movie WHERE idgen = $1", movie.ID)
	if err != nil {
		return model.Movie{}, err
	}

	return movie, nil
}

func (s *Storage) DeleteMovie(id string) error {
	_, err := s.DB.Exec(`DELETE FROM movie WHERE idgen = $1`, id)

	return err
}
