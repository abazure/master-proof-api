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
    quiz_category_id VARCHAR(100),
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
);

CREATE TABLE diagnostic_reports
(
    name        VARCHAR(100) NOT NULL PRIMARY KEY,
    description TEXT         NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_diagnostic_reports
(
    id                   VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    user_id              VARCHAR(100) NOT NULL,
    quiz_id              VARCHAR(100) NOT NULL,
    diagnostic_report_id VARCHAR(100) NOT NULL,
    created_at           TIMESTAMP    NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_diagnostic_reports_user_id FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT user_diagnostic_reports_quiz_id FOREIGN KEY (quiz_id) REFERENCES quizzes (id),
    CONSTRAINT user_diagnostic_reports_diagnostic_report_id FOREIGN KEY (diagnostic_report_id) references diagnostic_reports (name)
);


CREATE TABLE user_competence_reports
(
    id         VARCHAR(100) PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    user_id    VARCHAR(100) NOT NULL,
    quiz_name  VARCHAR(100) NOT NULL,
    score      int          NOT NULL,
    created_at TIMESTAMP    NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT user_diagnostic_report_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);