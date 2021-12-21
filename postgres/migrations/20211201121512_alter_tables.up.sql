ALTER TABLE "users" 
  ADD COLUMN "deleted" bool NOT NULL DEFAULT false, 
  ADD COLUMN "nickname" varchar(255) NOT NULL,
  ADD CONSTRAINT "email" UNIQUE ("email"),
  ADD CONSTRAINT "nickname" UNIQUE ("nickname"),
  ADD created_at timestamp NOT NULL DEFAULT NOW()
  ;

  