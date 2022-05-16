CREATE TABLE works_states
(
    id                 integer,
    work_name          varchar(256),
    state_name         varchar(256),
    state_number       integer,
    start_date         date,
    end_date           date,
    price_structure_id integer,

    PRIMARY KEY (id),
    FOREIGN KEY (price_structure_id) REFERENCES price_structures (id)
)