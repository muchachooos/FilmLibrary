package storage

import (
	"FilmLibrary/model"
	"errors"
	"fmt"
)

func (s *Storage) CreateActor(actor model.Actor) error {
	result, err := s.DB.Exec("INSERT INTO actor (id, name,  gender, date_of_birth) VALUES ($1, $2, $3, $4)", actor.ID, actor.Name, actor.Gender, actor.Birth)
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

func (s *Storage) GetActor(actorId string) (model.Actor, error) {
	var actor model.Actor

	err := s.DB.Get(&actor, "SELECT * FROM actor WHERE id = $1", actorId)
	if err != nil {
		return model.Actor{}, err
	}

	return actor, nil
}

func (s *Storage) UpdateActor(actor model.Actor) (model.Actor, error) {
	result, err := s.DB.NamedExec("UPDATE actor SET name = :name, gender = :gender, date_of_birth = :date_of_birth WHERE id = :id", actor)
	if err != nil {
		fmt.Println(111111111111)
		return model.Actor{}, err
	}

	countOfModifiedRows, err := result.RowsAffected()
	if err != nil {
		fmt.Println(222222222222)
		return model.Actor{}, err
	}

	if countOfModifiedRows == 0 {
		fmt.Println(33333333333)
		return model.Actor{}, errors.New("Nothing has been changed")
	}

	var response model.Actor

	err = s.DB.Get(&response, "SELECT * FROM actor WHERE id = $1", actor.ID)
	if err != nil {
		fmt.Println(44444444444444)
		return model.Actor{}, err
	}

	return actor, nil
}
