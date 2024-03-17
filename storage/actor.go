package storage

import (
	"FilmLibrary/model"
	"errors"
)

func (s *Storage) CreateActor(actor model.Actor) error {
	result, err := s.DB.NamedExec(`INSERT INTO actor (idgen, name,  gender, date_of_birth) VALUES (:idgen, :name, :gender, :date_of_birt, :date_of_birth)`, actor)
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

func (s *Storage) UpdateActor(actor model.Actor) (model.Actor, error) {
	result, err := s.DB.NamedExec(`UPDATE actor SET name = :name, gender = :gender,date_of_birth = :date_of_birth WHERE idgen = :idgen`, actor)
	if err != nil {
		return model.Actor{}, err
	}

	countOfModifiedRows, err := result.RowsAffected()
	if err != nil {
		return model.Actor{}, err
	}

	if countOfModifiedRows == 0 {
		return model.Actor{}, errors.New("Nothing has been changed")
	}

	var response model.Actor

	err = s.DB.Get(&response, "SELECT * FROM actor WHERE idgen = $1", actor.ID)
	if err != nil {
		return model.Actor{}, err
	}

	return actor, nil
}

func (s *Storage) DeleteActor(id string) error {
	result, err := s.DB.Exec(`DELETE FROM actor WHERE idgen = $1`, id)

	countOfModifiedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if countOfModifiedRows == 0 {
		errors.New("Nothing has been changed")
	}

	return nil
}

func (s *Storage) AddInCast(actorID, movieID string) error {
	_, err := s.DB.Exec(`INSERT INTO cast_record (actor_id, movie_id) VALUES ($1, $2)`, actorID, movieID)

	return err
}

func (s *Storage) DeleteFromCast(actorID, movieID string) error {
	result, err := s.DB.Exec(`DELETE FROM actor WHERE actor_id = $1 AND movie_id= $2`, actorID, movieID)

	countOfModifiedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if countOfModifiedRows == 0 {
		errors.New("Nothing has been changed")
	}

	return err
}
