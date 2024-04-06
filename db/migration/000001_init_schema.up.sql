CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "address" varchar
);

CREATE TABLE "operators" (
  "operator_id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "chargers" (
  "charger_id" varchar PRIMARY KEY,
  "operator_id" bigserial,
  "location" bigserial
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

--ALTER TABLE "operators" ADD FOREIGN KEY ("name") REFERENCES "users" ("id");

--ALTER TABLE "chargers" ADD FOREIGN KEY ("operator_id") REFERENCES "operators" ("operator_id");

--ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

--ALTER TABLE "transactions" ADD FOREIGN KEY ("operator_id") REFERENCES "operators" ("operator_id");

INSERT INTO users (username, email, password_hash, address)
VALUES ('john_doe', 'john@example.com', 'hashed_password', '123 Main St');

INSERT INTO operators (operator_id, name)
VALUES (1, 'SP Group'), (2, 'Schneider Electric');

INSERT INTO chargers (charger_id, operator_id, location) 
VALUES ('A1234X', 1, 2412);
INSERT INTO chargers (charger_id, operator_id, location) 
VALUES ('A3285X', 1, 5245);
INSERT INTO chargers (charger_id, operator_id, location) 
VALUES ('B3434T', 2, 1245);
INSERT INTO chargers (charger_id, operator_id, location) 
VALUES ('B2728T', 2, 4829);