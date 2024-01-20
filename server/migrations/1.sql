-- +migrate Up
CREATE TABLE "users" (
  "id" text PRIMARY KEY,
  "first_name" text,
  "middle_name" text,
  "last_name" text,
  "gender" text NOT NULL,
  "date_of_birth" date,
  "nationality" text,
  "cin" char(12),
  "poo" text,
  "por" text,
  "email" text UNIQUE,
  "username" text UNIQUE, 
  "phone" text UNIQUE NOT NULL,
  "password" text,
  "category" float,
  "token" text UNIQUE,
  "refresh_token" text UNIQUE,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "trash_bins" (
  "id" text PRIMARY KEY,
  "trash_level" decimal(5,2),
  "location" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE reports (
  "id" text PRIMARY KEY,
  "description" text,
  "transaction_ids" text[] NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "transactions" (
  "id" text PRIMARY KEY,
  "user_id" text,
  "trash_bin_id" text,
  "report_id" text,
  FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
  FOREIGN KEY ("trash_bin_id") REFERENCES "trash_bins" ("id"),
  FOREIGN KEY ("report_id") REFERENCES "reports" ("id"),
  "updated_at" TIMESTAMPTZ NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL
);

-- bo sung sau
CREATE TABLE permission (
    "id" text PRIMARY KEY,
    "permission_name" text
);

CREATE TABLE roles (
  "id" text PRIMARY KEY,
  "permission_id" text,
  "role_name" text UNIQUE,
  FOREIGN KEY ("permission_id") REFERENCES "permission" ("id")
);

CREATE TABLE user_roles (
  "user_id" text,
  "role_id" text,
  PRIMARY KEY ("user_id", "role_id"),
  FOREIGN KEY ("user_id") REFERENCES "users" ("id"),
  FOREIGN KEY ("role_id") REFERENCES "roles" ("id")
);

-- +migrate Down
DROP TABLE users;
DROP TABLE trash_bins;
DROP TABLE reports;
DROP TABLE transactions;
DROP TABLE permission;
DROP TABLE roles;
DROP TABLE user_roles;
