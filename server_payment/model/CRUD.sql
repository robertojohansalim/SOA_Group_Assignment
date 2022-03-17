CREATE TABLE payment.payment_record (
	id serial NOT NULL,
	external_id text NOT NULL,
	amount numeric NOT NULL,
	"method" text NOT NULL,
	status text NOT NULL,
	expired_at date NULL,
	created_at date NULL,
	updated_at date NULL,
	CONSTRAINT payment_record_pk PRIMARY KEY (id)
);
