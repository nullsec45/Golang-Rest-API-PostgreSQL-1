CREATE SEQUENCE customers_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."customers" (
    "id"  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "code" character varying(255) NOT NULL,
    "name" character varying(255) NOT NULL,
    "created_at" timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp,
    CONSTRAINT "customers_pkey" PRIMARY KEY ("id")
)