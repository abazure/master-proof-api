CREATE TYPE user_role AS ENUM(
    'STUDENT',
    'TEACHER'
    );

CREATE TABLE users
(
    id         VARCHAR(100) PRIMARY KEY,
    nim        VARCHAR(100) UNIQUE NOT NULL,
    role        user_role DEFAULT 'STUDENT',
    name       VARCHAR(100)        NOT NULL,
    email      VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

);

insert into users (id, nim, name, email) VALUES ('123','1212121', 'roni', 'efef@gmail.com');


CREATE TABLE files(
    id VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    url VARCHAR(100) NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE learning_materials(
    id VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    title VARCHAR(100) NOT NULL ,
    description TEXT NOT NULL,
    file_id VARCHAR(100) UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_learning_materials_id FOREIGN KEY (file_id) REFERENCES files(id)
);


select lm.title, lm.description,f.url from learning_materials as lm join files as f on lm.file_id = f.id