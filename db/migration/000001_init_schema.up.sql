CREATE TABLE "topics" (
    "id" bigserial PRIMARY KEY,
    "room_id" bigint
);
CREATE TABLE "rooms" (
    "id" bigserial PRIMARY KEY,
    "host_id" bigint NOT NULL,
    "name" varchar,
    "description" varchar,
    "participant_id" bigint,
    "message_id" bigint,
    "updated_at" timestamptz,
    "created_at" timestamptz DEFAULT (now())
);
CREATE TABLE "messages" (
    "id" bigserial PRIMARY KEY,
    "body" varchar NOT NULL,
    "updated_at" timestamptz,
    "created_at" timestamptz,
    "user_id" bigint NOT NULL
);
CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "user_name" varchar NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "last_login" timestamptz
);
ALTER TABLE "topics"
ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");
ALTER TABLE "rooms"
ADD FOREIGN KEY ("host_id") REFERENCES "users" ("id");
ALTER TABLE "rooms"
ADD FOREIGN KEY ("participant_id") REFERENCES "users" ("id");
ALTER TABLE "rooms"
ADD FOREIGN KEY ("message_id") REFERENCES "messages" ("id");
ALTER TABLE "messages"
ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
CREATE INDEX ON "rooms" ("host_id");
CREATE INDEX ON "rooms" ("name");
CREATE INDEX ON "users" ("user_name");
CREATE INDEX ON "users" ("email");