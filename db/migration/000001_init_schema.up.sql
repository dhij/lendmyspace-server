CREATE TABLE "spaces" (
    "id" bigserial PRIMARY KEY,
    "name" varchar,
    "description" varchar,
    "location" varchar,
    "link" varchar,
    "host_id" bigint NOT NULL,
    "images" text [],
    "dates" text [],
    "updated_at" timestamptz,
    "created_at" timestamptz DEFAULT (now())
);
CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "user_name" varchar NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "password" varchar NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "created_at" timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE "spaces"
ADD FOREIGN KEY ("host_id") REFERENCES "users" ("id");