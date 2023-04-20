CREATE TABLE "agent" (
  "id" integer PRIMARY KEY,
  "fullname" varchar,
  "role" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "house" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "price" varchar,
  "location" varchar,
  "description" text,
  "status" varchar,
  "agent_id" integer,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "house_gallery" (
  "id" integer PRIMARY KEY,
  "house_id" integer,
  "photos" text[]
);

CREATE TABLE "house_details" (
  "id" integer PRIMARY KEY,
  "house_id" integer,
  "rooms" integer,
  "bathrooms" varchar,
  "parking_lot" varchar,
  "sq_ft" double
);

COMMENT ON COLUMN "house"."description" IS 'House Description';

ALTER TABLE "house" ADD FOREIGN KEY ("agent_id") REFERENCES "agent" ("id");

ALTER TABLE "house_details" ADD FOREIGN KEY ("house_id") REFERENCES "house" ("id");

ALTER TABLE "house_gallery" ADD FOREIGN KEY ("house_id") REFERENCES "house" ("id");
