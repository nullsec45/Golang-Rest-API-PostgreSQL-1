CREATE TABLE "public"."charges" (
    "id"  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "journal_id" character varying(36) NOT NULL,
    "days_late" integer DEFAULT '1' NOT NULL,
    "daily_late_fee" integer NOT NULL,
    "total" integer NOT NULL,
    "user_id" character varying(36) NOT NULL,
    "created_at" timestamp(6)
)