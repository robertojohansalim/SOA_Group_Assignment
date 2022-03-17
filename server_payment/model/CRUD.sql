CREATE TABLE payment.payment_record (
	id serial NOT NULL,
	external_id text NULL,
	"method" text NULL,
	status text NULL,
	expired_at date NULL,
	created_at date NULL,
	updated_at date NULL,
	CONSTRAINT payment_record_pk PRIMARY KEY (id)
);
