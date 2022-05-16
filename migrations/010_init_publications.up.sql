CREATE TABLE publications
(
    id                  integer,
    publishing_house_id integer,
    topic_name          varchar(256),

    PRIMARY KEY (id),
    FOREIGN KEY (publishing_house_id) REFERENCES publishing_houses (id)
);

create table publications_to_persons
(
    id             integer,
    person_login   varchar(256),
    publication_id integer,

    PRIMARY KEY (id),
    FOREIGN KEY (person_login) REFERENCES persons (login),
    FOREIGN KEY (publication_id) REFERENCES publications (id)
)