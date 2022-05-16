CREATE TABLE persons
(
    login           varchar(256),
    first_name      varchar(256),
    second_name     varchar(256),
    third_name      varchar(256),
    date_of_birth   date,
    education       varchar(256),
    academic_degree varchar(256),
    academic_rank   varchar(256),
    contact_number  varchar(32),
    contact_email   varchar(256),

    PRIMARY KEY (login)
)