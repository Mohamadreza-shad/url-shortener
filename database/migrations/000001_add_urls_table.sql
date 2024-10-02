--liquibase formatted sql

-- changeset majid:2
CREATE TABLE IF NOT EXISTS source_keys (
  id bigserial PRIMARY KEY,
  source_key TEXT NOT NULL UNIQUE,
  total BIGINT NOT NULL,
  frozen DECIMAL NOT NULL DEFAULT 0,
  spent DECIMAL NOT NULL DEFAULT 0,
  last_order_amount BIGINT NOT NULL,
  plan VARCHAR(64) NOT NULL,
  renewal_date timestamp NOT NULL DEFAULT now(),
  created_at timestamp NOT NULL DEFAULT now()
);
--rollback DROP TABLE IF EXISTS "source_keys";
