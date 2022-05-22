CREATE TYPE position_type AS ENUM('Main', 'InternalCombination', 'ExternalCombination');
CREATE TYPE status_type AS ENUM('Active', 'Inactive');
CREATE TYPE status_change_reason_type AS ENUM('Hired', 'Fired', 'Transfer');

CREATE TABLE staff
(
    id                               serial,
    person_login                     varchar(256),
    position                         varchar(256),
    working_rate                     float,
    professional_qualification_group varchar(256),
    qualification_level              varchar(256),
    position_type                    position_type,
    division_id                      integer,
    supervisor_login                 varchar(256),
    is_supervisor                    bool,
    contract_end_date                date,
    status                           status_type,
    status_change_reason             status_change_reason_type,

    PRIMARY KEY (id),
    FOREIGN KEY (person_login) REFERENCES persons (login),
    FOREIGN KEY (supervisor_login) REFERENCES persons (login),
    FOREIGN KEY (division_id) REFERENCES divisions (division_code)
)