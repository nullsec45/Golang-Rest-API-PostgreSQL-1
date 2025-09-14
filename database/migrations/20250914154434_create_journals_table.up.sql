CREATE TABLE "public"."journals" (
    "id" character varying(36) DEFAULT 'gen_random_uuid()' NOT NULL,
    "book_id" character varying(36) NOT NULL,
    "stock_code" character varying(255) NOT NULL,
    "customer_id" character varying(36) NOT NULL,
    "status" character varying(50) NOT NULL,
    "borrowed_at" timestamp(6) NOT NULL,
    "returned_at" timestamp(6),
    "due_at" timestamp
)