CREATE TABLE "public"."media" (
    "id" character varying(36) DEFAULT 'gen_random_uuid()' NOT NULL,
    "path" text,
    "created_at" timestamp(6) NOT NULL,
    CONSTRAINT "media_pk" PRIMARY KEY ("id")
)