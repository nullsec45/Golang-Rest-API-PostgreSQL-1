CREATE TABLE "public"."books" (
    "id"  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "title" character varying(255) NOT NULL,
    "description" text,
    "isbn" character varying(100) NOT NULL,
    "created_at" timestamp(6),
    "updated_at" timestamp(6),
    "deleted_at" timestamp(6),
    "cover_id" character varying(255),
    CONSTRAINT "books_pk" PRIMARY KEY ("id")
)