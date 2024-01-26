CREATE TABLE
    IF NOT EXISTS "url_lookup" (
        "hash" varchar(24) NOT NULL,
        "url" text NOT NULL CONSTRAINT "uq__url_look_up__url" UNIQUE,
        "created_at" timestamptz NOT NULL DEFAULT current_timestamp,
        PRIMARY KEY ("hash")
    );