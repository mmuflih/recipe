
-- +migrate Up
CREATE TABLE public.recipes (
	id serial NOT NULL,
	"name" varchar(30) NOT NULL,
	slug varchar(30) NOT NULL,
	category_id smallint NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
	PRIMARY KEY (id),
    CONSTRAINT recipes_category_id_foreign FOREIGN KEY (category_id) REFERENCES public.categories(id)
);


-- +migrate Down
DROP TABLE public.recipes;