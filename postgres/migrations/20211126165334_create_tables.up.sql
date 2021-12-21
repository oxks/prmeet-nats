CREATE TABLE "users" (
  "id" bigserial,
  "firstname" text NOT NULL DEFAULT '',
  "lastname" text NOT NULL DEFAULT '',
   "email" varchar(255) NOT NULL,
   "password" varchar(255) NOT NULL,
  PRIMARY KEY ("id")
)
;


