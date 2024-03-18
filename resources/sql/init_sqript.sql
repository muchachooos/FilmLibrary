CREATE TABLE IF NOT EXISTS actor
(
    id            CHAR(36) PRIMARY KEY,
    name          VARCHAR(255) NOT NULL,
    gender        VARCHAR(6),
    date_of_birth DATE
);

CREATE TABLE IF NOT EXISTS movie
(
    id     CHAR(36) PRIMARY KEY,
    title  VARCHAR(255) NOT NULL,
    rating REAL,
    release_year SMALLSERIAL
);

CREATE TABLE IF NOT EXISTS cast_record
(
    actor_id CHAR(36) NOT NULL,
    movie_id CHAR(36) NOT NULL,
    PRIMARY KEY (actor_id, movie_id),
    FOREIGN KEY (actor_id) REFERENCES actor (id) ON DELETE CASCADE,
    FOREIGN KEY (movie_id) REFERENCES movie (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_auth
(
    id          CHAR(36) PRIMARY KEY,
    login       VARCHAR(25) UNIQUE  NOT NULL,
    hashed_pass VARCHAR(300) UNIQUE NOT NULL,
    token       CHAR(36) UNIQUE,
    is_admin    BOOL                NOT NULL
);

-- создаём администратора с паролем my_pass
INSERT INTO user_auth (id, login, hashed_pass, token, is_admin)
VALUES ('8a1a10a8-554f-4d60-99e0-7d409084dc45', 'Igor',
        '$2a$10$Rv1GBrZsk.DxE/KhICjl2OhCs8YwzbHrOORqw5fSYXICNDcQDzWDa', NULL, TRUE)
