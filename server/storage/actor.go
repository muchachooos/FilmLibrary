package storage

import (
	"filmlib/model"
)

func (s *Storage) CreateActor(actor model.Actor) error {
	_, err := s.DB.NamedExec(
		`INSERT INTO actor (id, name,  gender, date_of_birth) 
				VALUES (:id, :name, :gender, :date_of_birth)`,
		actor)
	return err
}

func (s *Storage) GetActor(actorId string) (model.Actor, error) {
	var actor model.Actor

	err := s.DB.Get(&actor, "SELECT * FROM actor WHERE id = $1", actorId)
	if err != nil {
		return model.Actor{}, err
	}

	return actor, nil
}

func (s *Storage) UpdateActor(actor model.Actor) error {
	_, err := s.DB.NamedExec(
		`UPDATE actor SET 
				name = :name, 
				gender = :gender, 
				date_of_birth = :date_of_birth 
           	    WHERE id = :id`,
		actor)
	return err
}

func (s *Storage) DeleteActor(id string) error {
	_, err := s.DB.Exec(`DELETE FROM actor WHERE id = $1`, id)
	return err
}

func (s *Storage) AddCast(actorID, movieID string) error {
	_, err := s.DB.Exec(`INSERT INTO cast_record (actor_id, movie_id) VALUES ($1, $2)`, actorID, movieID)
	return err
}

func (s *Storage) DeleteCast(actorID, movieID string) error {
	_, err := s.DB.Exec(`DELETE FROM actor WHERE actor_id = $1 AND movie_id= $2`, actorID, movieID)
	return err
}

func (s *Storage) GetActorsWithMovies() ([]model.Actor, error) {
	var actors []model.Actor
	err := s.DB.Select(&actors, "SELECT * FROM actor")
	if err != nil {
		return nil, err
	}

	var casts []model.CastRecord
	err = s.DB.Select(&casts, "SELECT * FROM cast_record")
	if err != nil {
		return nil, err
	}

	var movies []model.Movie
	err = s.DB.Select(&movies, "SELECT * FROM movie")
	if err != nil {
		return nil, err
	}

	// создаём мапу movie ID - movie
	movieMap := make(map[string]model.Movie, len(movies))
	for _, movie := range movies {
		movieMap[movie.ID] = movie
	}

	for i := range actors {
		for j := range casts {
			if actors[i].ID == casts[j].ActorID {
				actors[i].Movies = append(actors[i].Movies, movieMap[casts[j].MovieID])
			}
		}
	}

	return actors, nil
}
