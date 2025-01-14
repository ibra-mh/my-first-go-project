-- +goose Up
-- +goose StatementBegin
CREATE TABLE paths (
    id integer PRIMARY KEY,
    name varchar(50),
    description varchar(50),
    guid varchar(19),
    created_at date,
    updated_at date,
    deleted_at date
);
CREATE TABLE db (
    id integer PRIMARY KEY,
    name varchar(50),
    description varchar(50),
    guid varchar(19),
    created_at date,
    updated_at date,
    deleted_at date
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS paths;
-- +goose StatementEnd
