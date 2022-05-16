CREATE TABLE divisions
(
    division_code       integer,
    organisation_name   varchar(256),
    division_name       varchar(256),
    head_division_code  integer,
    division_supervisor varchar(256),

    PRIMARY KEY (division_code),
    FOREIGN KEY (division_supervisor) REFERENCES persons (login)
)
