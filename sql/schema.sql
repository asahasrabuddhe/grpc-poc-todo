CREATE TABLE users
(
  id         int PRIMARY KEY AUTO_INCREMENT,
  name       varchar(64)            NOT NULL,
  email      varchar(64)            NOT NULL,
  password   varchar(64)            NOT NULL,
  created_on datetime DEFAULT now() NOT NULL,
  updated_on datetime DEFAULT now() NOT NULL,
  deleted_on datetime        DEFAULT null
);

CREATE TABLE todos
(
  id           int PRIMARY KEY AUTO_INCREMENT,
  title        varchar(64)            NOT NULL,
  description  text                   NOT NULL,
  is_completed bool DEFAULT false     NOT NULL,
  user_id      int                    NOT NULL,
  created_on   datetime DEFAULT now() NOT NULL,
  updated_on   datetime DEFAULT now() NOT NULL,
  deleted_on   datetime        DEFAULT null,
  CONSTRAINT todos_users_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);