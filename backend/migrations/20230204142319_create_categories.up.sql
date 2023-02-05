SET statement_timeout = 0;

--bun:split

CREATE TABLE categories (
    "id" uuid,
    "name" varchar NOT NULL,
    PRIMARY KEY ("id"),
    UNIQUE ("name")
);

--bun:split

INSERT INTO categories VALUES
    ('8ff16cd9-e1fd-4a8a-a9ce-876e65e3bf96', 'Meat'),
    ('ee1e73bc-ae9b-4b7a-ae96-a714f074b524', 'Sweets'),
    ('2ac002c9-4fae-4dbc-afb1-05c197ef1427', 'Fruits'),
    ('584a2e6a-4691-4443-9f6f-67bbcfae8d10', 'Vegetables')
