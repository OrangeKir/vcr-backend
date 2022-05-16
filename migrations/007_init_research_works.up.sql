CREATE TYPE work_type AS ENUM('ResearchWork', 'ResearchWorkPart', 'DevelopmentWork', 'DevelopmentWorkPart', 'Service', 'Grant', 'Other');

CREATE TABLE research_works
(
    lead_customer              varchar(256),
    customer                   varchar(256),
    full_work_title            varchar(256),
    short_work_title           varchar(256),
    code                       varchar(32),
    contract_number            integer,
    contract_date              date,
    topic_number               integer,
    work_type                  work_type,
    work_price                 integer,
    start_date                 date,
    end_date                   date,
    stages_number              integer,
    work_price_structure_id    integer,
    responsible_executor_login varchar(256),

    PRIMARY KEY (code),
    FOREIGN KEY (work_price_structure_id) REFERENCES price_structures (id),
    FOREIGN KEY (responsible_executor_login) REFERENCES persons (login)
)