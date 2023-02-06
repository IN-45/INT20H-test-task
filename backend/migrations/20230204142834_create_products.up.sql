SET statement_timeout = 0;

--bun:split

CREATE TABLE products (
    "id" uuid,
    "category_id" uuid NOT NULL,
    "name" varchar(255) NOT NULL,
    "image_url" varchar(2000),
    "amount_type" amount_type,
    PRIMARY KEY ("id"),
    UNIQUE ("name"),
    CONSTRAINT category_id_fk FOREIGN KEY("category_id") REFERENCES categories("id") ON DELETE CASCADE
);

--bun:split

INSERT INTO products
VALUES (
        'f056513a-c64b-4f54-89e2-c23d26b55a97',
        '2ac002c9-4fae-4dbc-afb1-05c197ef1427',
        'Apple',
        NULL,
        'g'
    ),
    (
        'f725fb66-5719-481c-a484-386f5c27a084',
        '2ac002c9-4fae-4dbc-afb1-05c197ef1427',
        'Cherry',
        NULL,
        'g'
    ),
    (
        '61e3326f-6f5c-4c2b-ba42-6d95e56a2022',
        'ee1e73bc-ae9b-4b7a-ae96-a714f074b524',
        'Chocolate',
        NULL,
        'g'
    ),
    (
        '0578ae60-ca58-4130-97ff-a217acf45f43',
        '584a2e6a-4691-4443-9f6f-67bbcfae8d10',
        'Potato',
        NULL,
        'g'
    ),
    (
        'cf43becb-6772-464e-b58d-fb959aa43278',
        '8ff16cd9-e1fd-4a8a-a9ce-876e65e3bf96',
        'Pork',
        NULL,
        'g'
    )