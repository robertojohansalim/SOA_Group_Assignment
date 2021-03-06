CREATE SCHEMA "payment";

CREATE TABLE "payment".payment_record (
	id serial4 NOT NULL,
	external_id text NULL,
	"method" text NULL,
	status text NULL,
	expired_at timestamp NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	amount numeric NULL,
	user_id text NULL,
	CONSTRAINT payment_record_pk PRIMARY KEY (id)
);
