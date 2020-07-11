CREATE TABLE IF NOT EXISTS users
(
uid        SERIAL       NOT NULL,
name       VARCHAR(50)  NOT NULL,
nickname   VARCHAR(50)  NOT NULL,
birthday   DATE         NOT NULL,
phone      INT          NOT NULL,
email      VARCHAR(50)  NOT NULL,
password   VARCHAR(40)  NOT NULL,
state      BOOLEAN      NOT NULL,
created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
update_at  TIMESTAMP,
delete_at  TIMESTAMP,
CONSTRAINT users_pkey PRIMARY KEY(uid)
)