-- +migrate Up
CREATE TABLE "users" (
  "id" text PRIMARY KEY,
  "full_name" text,
  "first_name" text,
  "middle_name" text,
  "last_name" text,
  "date_of_birth" datetime,
  "email" text UNIQUE,
  "phone" text UNIQUE,
  "password" text,
  "total_points" float,
  "token" text UNIQUE,
  "refresh_token" text UNIQUE,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "trashBins" (
  "id" text PRIMARY KEY,
  "location" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

CREATE TABLE "transactions" (
  "id" text PRIMARY KEY,
  "userID" text,
  "trashBinID" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL,
  FOREIGN KEY ("userID") REFERENCES "users" ("id"),
  FOREIGN KEY ("trashBinID") REFERENCES "trashBins" ("id")
);

CREATE TABLE 
-- bo sung sau
CREATE TABLE permission (
    "id" text PRIMARY KEY,
)

CREATE TABLE roles (
  "id" text PRIMARY KEY,
  "permissionID" text,
  "roleName" text UNIQUE
  FOREIGN KEY ("permissionID") REFERENCES "permissions" ("id")
),

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