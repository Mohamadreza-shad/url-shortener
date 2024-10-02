--liquibase formatted sql

-- changeset reza:1
CREATE TABLE IF NOT EXISTS urls (
  id bigserial PRIMARY KEY,
  long_url TEXT NOT NULL UNIQUE,
  short_url TEXT NOT NULL UNIQUE,
  created_at timestamp NOT NULL DEFAULT now(),
  expired_at timestamp NOT NULL
);
--rollback DROP TABLE IF EXISTS "urls";
