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
    photo_url VARCHAR(100) NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE icons
(
    id        VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    ic_url    VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL    DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE learning_material_categories
(
    id          VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    title       VARCHAR(100) NOT NULL,
    description TEXT         NOT NULL,
    created_at  TIMESTAMP    NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL    DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE files(
    id VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    pdf_url VARCHAR(255) NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE learning_materials(
    id VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    title VARCHAR(100) NOT NULL ,
    description TEXT NOT NULL,
    icon_id VARCHAR(100) UNIQUE ,
    learning_material_category_id VARCHAR(100) UNIQUE ,
    file_id VARCHAR(100) UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_learning_material_icons FOREIGN KEY (icon_id) REFERENCES icons (id),
    CONSTRAINT fk_learning_material_categories FOREIGN KEY (learning_material_category_id) REFERENCES learning_material_categories (id),
    CONSTRAINT fk_learning_material_files FOREIGN KEY (file_id) REFERENCES files(id)
);


CREATE TABLE quiz_categories(
    id VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE quizzes(
    id VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    quiz_category_id VARCHAR(100) UNIQUE,
    name VARCHAR(100) NOT NULL ,
    description VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_quizzes_category_id FOREIGN KEY (quiz_category_id) REFERENCES quiz_categories(id)
);

CREATE TABLE questions(
    id VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    quiz_id VARCHAR(100),
    question TEXT NOT NULL ,
    correct_answer INTEGER NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT question_quiz_id FOREIGN KEY (quiz_id) REFERENCES quizzes(id)
);

CREATE TABLE answers(
    id VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    question_id VARCHAR(100) NOT NULL,
    value INTEGER NOT NULL ,
    text TEXT NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_answer_question_id FOREIGN KEY (question_id) REFERENCES questions (id)
)

insert into quiz_categories (id, name) values ('qt1','contoh')
insert into quizzes (id, quiz_category_id, name, description) values ('qz1','qt1', 'modalitas', 'quiz modalitas')
insert into
questions
(id, quiz_id, question, correct_answer)
values
('q1', 'qz1', 'Why are you gay?', 1),
('q2', 'qz1', 'Why are you ugly af?', 2)

insert into
answers
(id, question_id, value, text)
values
('a2-1','q2', 0,'YNTKTS' ),
('a2-2','q2', 1,'Terserah saya' ),
('a2-3','q2', 3,'Wong saya suka kok' ),
('a2-4','q2', 4,'No comment' )