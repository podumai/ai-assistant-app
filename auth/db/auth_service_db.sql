DROP SCHEMA IF EXISTS auth CASCADE;

CREATE SCHEMA IF NOT EXISTS auth;

CREATE TABLE IF NOT EXISTS auth.users
(
  userId UUID PRIMARY KEY,
  login TEXT NOT NULL,
  password TEXT NOT NULL,
  CONSTRAINT uniqueLogin UNIQUE(login),
  CONSTRAINT uniquePassword UNIQUE(password),
  CONSTRAINT validLogin CHECK(LENGTH(login) > 0),
  CONSTRAINT validPassword CHECK(LENGTH(password) > 0)
);