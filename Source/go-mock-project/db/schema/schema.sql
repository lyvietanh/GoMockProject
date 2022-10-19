CREATE TABLE "flights" (
  "id" BIGSERIAL PRIMARY KEY,
  "origin_airport_code" varchar(3) NOT NULL,
  "destination_airport_code" varchar(3) NOT NULL,
  "departure_date_time" timestamp NOT NULL,
  "flight_number" varchar(8) NOT NULL,
  "booking_class" varchar(16) NOT NULL,
  "flight_duration" int NOT NULL,
  "enabled" boolean NOT NULL DEFAULT true,
  "seat_available" int NOT NULL DEFAULT 10,
  "currency_code" varchar(3) NOT NULL,
  "price" float NOT NULL DEFAULT 1000000,
  "create_date" timestamp DEFAULT (now()),
  "modify_date" timestamp
);

CREATE TABLE "booking_transactions" (
  "id" BIGSERIAL PRIMARY KEY,
  "reservation_code" varchar(8),
  "status" varchar(32),
  "error_code" int,
  "error_message" varchar(256),
  "contact_email" varchar(64) NOT NULL,
  "contact_phone" varchar(16) NOT NULL,
  "currency_code" varchar(3),
  "total_price" float,
  "create_date" timestamp DEFAULT (now()),
  "modify_date" timestamp
);

CREATE TABLE "booking_flights" (
  "id" BIGSERIAL PRIMARY KEY,
  "booking_transaction_id" bigint,
  "flight_id" bigint,
  "currency_code" varchar(3) NOT NULL,
  "price" float NOT NULL
);

CREATE TABLE "booking_passengers" (
  "id" BIGSERIAL PRIMARY KEY,
  "booking_transaction_id" bigint,
  "title" varchar(8),
  "last_name" varchar(25),
  "first_name" varchar(25),
  "date_of_birth" date,
  "travel_doc_type" varchar(1),
  "travel_doc_number" varchar(16),
  "ffp_number" varchar(16)
);

CREATE TABLE "roles" (
  "id" varchar(32) PRIMARY KEY,
  "name" varchar(32),
  "enabled" boolean NOT NULL DEFAULT true,
  "create_date" timestamp DEFAULT (now()),
  "modify_date" timestamp
);

CREATE TABLE "users" (
  "email" varchar(64) PRIMARY KEY,
  "role_id" varchar(32),
  "title" varchar(8),
  "last_name" varchar(25),
  "first_name" varchar(25),
  "date_of_birth" date,
  "phone" varchar(16),
  "travel_doc_type" varchar(1),
  "travel_doc_number" varchar(16),
  "ffp_number" varchar(16) UNIQUE,
  "password_salt" varchar(16),
  "secure_password" varchar(256),
  "enabled" boolean NOT NULL DEFAULT true,
  "last_login_date" timestamp,
  "create_date" timestamp DEFAULT (now()),
  "modify_date" timestamp
);

ALTER TABLE "booking_flights" ADD FOREIGN KEY ("booking_transaction_id") REFERENCES "booking_transactions" ("id");

ALTER TABLE "booking_flights" ADD FOREIGN KEY ("flight_id") REFERENCES "flights" ("id");

ALTER TABLE "booking_passengers" ADD FOREIGN KEY ("booking_transaction_id") REFERENCES "booking_transactions" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
