CREATE TABLE "spaces" (
    "id" bigserial PRIMARY KEY,
    "name" varchar,
    "description" varchar,
    "location" varchar,
    "link" varchar,
    "host_id" bigint NOT NULL,
    "image_id" bigint,
    "date_id" bigint,
    "updated_at" timestamptz,
    "created_at" timestamptz DEFAULT (now())
);
CREATE TABLE "dates" (
    "id" bigserial PRIMARY KEY,
    "available_on" timestamptz
);
CREATE TABLE "images" ("id" bigserial PRIMARY KEY, "link" varchar);
CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "user_name" varchar NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "last_login" timestamptz
);
ALTER TABLE "spaces"
ADD FOREIGN KEY ("host_id") REFERENCES "users" ("id");
ALTER TABLE "spaces"
ADD FOREIGN KEY ("image_id") REFERENCES "images" ("id");
ALTER TABLE "spaces"
ADD FOREIGN KEY ("date_id") REFERENCES "dates" ("id");
CREATE INDEX ON "spaces" ("host_id");
CREATE INDEX ON "spaces" ("name");
CREATE INDEX ON "users" ("user_name");
CREATE INDEX ON "users" ("email");