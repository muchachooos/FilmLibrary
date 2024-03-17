package storage

import (
	"FilmLibrary/model"
)

func (s *Storage) ActorsHandlerInDB(actor model.Actor) error {
	result, err := s.DB.Exec("INSERT INTO actors (id, actor_name,  gender, date_of_birth) VALUES ($1, $2, $3, $4)", actor.ID, actor.Name, actor.Gender, actor.Birth)
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
