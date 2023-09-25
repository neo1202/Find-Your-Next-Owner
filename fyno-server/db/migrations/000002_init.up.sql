CREATE TABLE "users" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "name" text,
  "email" text,
  "avatar_url" text,
  "bio" text DEFAULT '',
  "signature" text DEFAULT ''
);

CREATE TABLE "posts" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "user_id" uuid,
  "kind" text,
  "name" text,
  "age" int,
  "gender" text,
  "content" text,
  "location_id" uuid,
  "category_id" uuid,
  "created_at" timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE "locations" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "name" text
);

CREATE TABLE "categories" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "name" text
);

ALTER TABLE "messages" ADD FOREIGN KEY ("sender") REFERENCES "users" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("receiver") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "images" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "saved" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "saved" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
