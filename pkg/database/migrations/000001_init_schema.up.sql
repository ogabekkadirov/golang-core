CREATE TABLE statuses (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL
);

INSERT INTO statuses (id, name)
VALUES (1,'Active'),(2,'Blocked');

CREATE TABLE users (
    "id" SERIAL PRIMARY KEY,
    "username" VARCHAR(50)  NOT NULL,
    "fullname" VARCHAR(100) NOT NULL,
    "status_id" INTEGER DEFAULT 1,
    "password" TEXT NOT NULL,
    "created_at" TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX statuses_id_uindex ON "statuses" ("id");

CREATE UNIQUE INDEX users_id_uindex ON "users" ("id");
CREATE UNIQUE INDEX users_username_uindex ON "users" ("username");

ALTER TABLE users ADD CONSTRAINT users_statuses_id_fk
     FOREIGN KEY ("status_id") REFERENCES statuses ("id");
