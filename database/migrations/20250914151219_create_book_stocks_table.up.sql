CREATE TABLE "public"."book_stocks" (
    "book_id"  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "code" character varying(50) NOT NULL,
    "status" character varying(50) NOT NULL,
    "borrower_id" character varying(36),
    "borrowed_at" timestamp(6),
    CONSTRAINT "book_stocks_pk" PRIMARY KEY ("code")
)