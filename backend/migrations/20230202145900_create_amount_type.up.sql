SET statement_timeout = 0;

--bun:split

CREATE TYPE amount_type AS ENUM (
    'g',
    'mg',
    'kg',
    'l',
    'ml',
    'tablesp',
    'teasp',
    'pcs'
);