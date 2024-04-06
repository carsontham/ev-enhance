CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "address" varchar
);

CREATE TABLE "operators" (
  "id" bigserial PRIMARY KEY,
  "name" bigserial NOT NULL,
  "chargers_id" bigserial,
  "point_id" bigserial 
);

CREATE TABLE "chargers" (
  "id" bigserial PRIMARY KEY,
  "name" bigserial NOT NULL,
  "point_id" bigserial,
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigserial NOT NULL,
  "operator_id" bigserial NOT NULL,
  "point_id" bigserial NOT NULL,
  "charging_time" time,
  "charging_type" varchar,
  "payment_type" varchar,
  "status" varchar
);

COMMENT ON TABLE "users" IS 'Stores user date';

COMMENT ON TABLE "operators" IS 'Stores operators details';

COMMENT ON TABLE "chargers" IS 'Stores operators details';

COMMENT ON TABLE "transactions" IS 'Stores transactions';

ALTER TABLE "operators" ADD FOREIGN KEY ("name") REFERENCES "users" ("id");

ALTER TABLE "chargers" ADD FOREIGN KEY ("name") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("operator_id") REFERENCES "operators" ("id");
