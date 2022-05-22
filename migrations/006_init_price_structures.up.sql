CREATE TYPE price_type AS ENUM('Estimated', 'Fixed');

CREATE TABLE price_structures
(
    id                    serial,
    topic_number          integer,
    laboriousness         integer,
    basic_salary          integer,
    additional_salary     integer,
    payroll_fund          integer,
    materials             integer,
    special_equipment     integer,
    social_payments       integer,
    overhead_costs        integer,
    business_trips        integer,
    other_direct_expenses integer,
    work_own_cost         integer,
    profit                integer,
    work_cost             integer,
    value_added_tax       integer,
    work_price            integer,
    price_type            price_type,

    primary key (id)
)