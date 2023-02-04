SET statement_timeout = 0;

--bun:split

CREATE TABLE users (
    "id" UUID,
    "email" varchar NOT NULL,
    "password_hash" varchar NOT NULL,
    "created_at" timestamp DEFAULT NOW(),
    PRIMARY KEY ("id"),
    UNIQUE ("email")
);
