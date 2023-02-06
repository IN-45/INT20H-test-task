SET statement_timeout = 0;
--bun:split
CREATE TABLE recipes (
    "id" uuid,
    "name" varchar(255) NOT NULL,
    "description" varchar(2000),
    "author_id" uuid,
    "cooking_time_minutes" int,
    "image_url" varchar(2000),
    "created_at" timestamp DEFAULT NOW(),
    PRIMARY KEY ("id"),
    CONSTRAINT author_id_fk FOREIGN KEY("author_id") REFERENCES users("id") ON DELETE CASCADE
);
--bun:split
CREATE TABLE instructions (
    "recipe_id" uuid,
    "description" varchar(2000),
    "priority" int,
    CONSTRAINT recipe_id_fk FOREIGN KEY("recipe_id") REFERENCES recipes("id") ON DELETE CASCADE
);
--bun:split
/* CREATE TYPE amount_type AS ENUM (
    'g',
    'mg',
    'kg',
    'l',
    'ml',
    'tablesp',
    'teasp',
    'pcs'
); */
CREATE TABLE recipes_products(
    "recipe_id" uuid,
    "product_id" uuid,
    "amount" int,
    "amount_type" amount_type,
    CONSTRAINT recipe_id_fk FOREIGN KEY("recipe_id") REFERENCES recipes("id") ON DELETE CASCADE,
    CONSTRAINT product_id_fk FOREIGN KEY("product_id") REFERENCES products("id") ON DELETE CASCADE
);