DROP TABLE movie, actor, cast_record;

CREATE TABLE actor
(
    id            CHAR(36) PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    gender        VARCHAR(6),
    date_of_birth DATE
);

CREATE TABLE movie
(
    id           CHAR(36) PRIMARY KEY,
    title        VARCHAR(150) NOT NULL,
    rating       REAL,
    release_year SMALLSERIAL
);

CREATE TABLE cast_record
(
    actor_id CHAR(36),
    movie_id CHAR(36),
    PRIMARY KEY (actor_id, movie_id),
    FOREIGN KEY (actor_id) REFERENCES actor (id) ON DELETE CASCADE,
    FOREIGN KEY (movie_id) REFERENCES movie (id) ON DELETE CASCADE
);



INSERT INTO actor (id, name, gender, date_of_birth)
VALUES ('bradpittIDb8actorID1', 'Brad Pit', 'male', '18-12-1963');

INSERT INTO actor (id, name, gender, date_of_birth)
VALUES ('ayacash0IDb8actorID1', 'Aya Cash', 'female', '13-07-1982');

UPDATE actor
SET name = 'Brad Pitt'
WHERE id = 'bradpittIDb8actorID1';

UPDATE actor
SET gender = 'male'
WHERE id = 'bradpittIDb8actorID1';

UPDATE actor
SET date_of_birth = '13-07-1982'
WHERE id = 'bradpittIDb8actorID1';

DELETE
FROM actor
WHERE id = 'bradpittIDb8actorID1';



INSERT INTO movie (id, title, rating, release_year)
VALUES ('fightclubID84filmID1', 'Fight Club', 8.7, 1999);

INSERT INTO movie (id, title, rating, release_year)
VALUES ('mrmssmithID84filmID2', 'Mr and Mrs Smith', 7.5, 2005);

UPDATE movie
SET title = 'Mr and Mrs Smith'
WHERE id = 'mrmssmithID84filmID2';

UPDATE movie
SET rating = '7.5'
WHERE id = 'mrmssmithID84filmID2';

UPDATE movie
SET release_year = '2005'
WHERE id = 'mrmssmithID84filmID2';

DELETE
FROM movie
WHERE id = 'mrmssmithID84filmID2';



INSERT INTO cast_record (actor_id, movie_id)
VALUES ('ayacash0IDb8actorID1', 'mrmssmithID84filmID2');

INSERT INTO cast_record (actor_id, movie_id)
VALUES ('bradpittIDb8actorID1', 'fightclubID84filmID1');

DELETE
FROM cast_record
WHERE cast_record.actor_id = 'ayacash0IDb8actorID1'
  AND cast_record.movie_id = 'fightclubID84filmID1';
