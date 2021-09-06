
-- +migrate Up
CREATE TABLE public.recipe_details (
	id serial NOT NULL,
    recipe_id integer not null,
    ingredient_id integer not null,
	"qty" varchar(30) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
	PRIMARY KEY (id),
    CONSTRAINT recipe_details_recipe_id_foreign FOREIGN KEY (recipe_id) REFERENCES public.recipes(id),
    CONSTRAINT recipe_details_ingredient_id_foreign FOREIGN KEY (ingredient_id) REFERENCES public.ingredients(id)
);


-- +migrate Down
DROP TABLE public.recipe_details;