-- +migrate Up
CREATE TABLE "users" (
  "id" text PRIMARY KEY,
  "full_name" text,
  "first_name" text,
  "middle_name" text,
  "last_name" text,
  "gender" text NOT NULL,
  "date_of_birth" datetime,
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

CREATE TABLE "trashBins" (
  "id" text PRIMARY KEY,
  "trash_level" decimal(5,2),
  "location" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE reports (
  "id" text PRIMARY KEY,
  "description" text,
  "transactionIDs" text[] NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "transactions" (
  "id" text PRIMARY KEY,
  "userID" text,
  "trashBinID" text,
  "reportID" text,
  FOREIGN KEY ("trashBinID") REFERENCES "trashBins" ("id"),
  FOREIGN KEY ("userID") REFERENCES "users" ("id"),
  FOREIGN KEY ("reportID") REFERENCES "reports" ("id"),
  "updated_at" TIMESTAMPTZ NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL
);

-- bo sung sau
CREATE TABLE permission (
    "id" text PRIMARY KEY,
);

CREATE TABLE roles (
  "id" text PRIMARY KEY,
  "permissionID" text,
  "roleName" text UNIQUE,
  FOREIGN KEY ("permissionID") REFERENCES "permission" ("id")
);

CREATE TABLE userRoles (
  "userID" text,
  "roleID" text,
  PRIMARY KEY ("userID", "roleID"),
  FOREIGN KEY ("userID") REFERENCES "users" ("id"),
  FOREIGN KEY ("roleID") REFERENCES "roles" ("id")
);

-- +migrate Down
DROP TABLE userRoles;
DROP TABLE roles;
DROP TABLE permission;
DROP TABLE transactions;
DROP TABLE users;
DROP TABLE trashBins;