CREATE TABLE conferences
(
    id   serial,
    name varchar(256),

    primary key (id)
);

CREATE TABLE conferences_to_persons
(
    id            serial,
    person_login  varchar(256),
    conference_id integer,
    topic_name    varchar(256),

    PRIMARY KEY (id),
    FOREIGN KEY (person_login) REFERENCES persons (login),
    FOREIGN KEY (conference_id) REFERENCES conferences (id)
)