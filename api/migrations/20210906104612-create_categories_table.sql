
-- +migrate Up
CREATE TABLE public.categories (
	id smallserial NOT NULL,
	"name" varchar(30) NOT NULL,
	slug varchar(30) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
	PRIMARY KEY (id)
);


-- +migrate Down
DROP TABLE public.categories;