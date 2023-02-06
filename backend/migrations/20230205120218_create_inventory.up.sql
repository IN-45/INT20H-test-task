SET statement_timeout = 0;
--bun:split
CREATE TABLE inventories (
    "user_id" uuid NOT NULL,
    "product_id" uuid NOT NULL,
    "added_at" timestamp DEFAULT NOW(),
    "amount" int,
    "amount_type" amount_type,
    UNIQUE("user_id", "product_id"),
    CONSTRAINT user_id_fk FOREIGN KEY("user_id") REFERENCES users("id") ON DELETE CASCADE,
    CONSTRAINT product_id_fk FOREIGN KEY("product_id") REFERENCES products("id") ON DELETE CASCADE
);