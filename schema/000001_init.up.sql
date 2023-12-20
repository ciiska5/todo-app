CREATE TABLE users
(
    id              SERIAL NOT NULL UNIQUE,
    name            VARCHAR(250) NOT NULL,
    nickname        VARCHAR(250) NOT NULL UNIQUE,
    password_hash   VARCHAR(250) NOT NULL
);

CREATE TABLE lists
(
    id              SERIAL NOT NULL UNIQUE,
    title           VARCHAR(250) NOT NULL,
    description     VARCHAR(350),
    created_at      TIMESTAMP NOT NULL
);

CREATE TABLE users_lists
(
    user_id         INTEGER REFERENCES users (id) ON DELETE CASCADE,
    list_id         INTEGER REFERENCES lists (id) ON DELETE CASCADE    
);

CREATE TABLE tasks
(
    id              SERIAL NOT NULL UNIQUE,
    title           VARCHAR(250) NOT NULL,
    description     VARCHAR(500),
    is_done         BOOLEAN NOT NULL DEFAULT FALSE,
    created_at      TIMESTAMP NOT NULL
);


CREATE TABLE lists_tasks
(
    list_id         INTEGER REFERENCES lists (id) ON DELETE CASCADE,
    task_id         INTEGER REFERENCES tasks (id) ON DELETE CASCADE    
);