-- +migrate Up
CREATE TABLE usrs
(
    id UUID DEFAULT UUID_GENERATE_V4()
        CONSTRAINT usrs_usr_id_pkey
            PRIMARY KEY,
    fio TEXT NOT NULL,
    post TEXT NOT NULL,
    department TEXT NOT NULL
);

-- +migrate Down
DROP TABLE usrs;