CREATE TABLE standards
(
    id                serial,
    standard_type     varchar(256),
    basic_salary      integer,
    additional_salary integer,
    payroll_fund      integer,
    social_payments   integer,
    overhead_cost     integer,
    profitability     integer
)