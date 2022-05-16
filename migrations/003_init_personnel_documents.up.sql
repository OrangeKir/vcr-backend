CREATE TYPE document_type AS ENUM('Order', 'Writ');

CREATE TABLE personnel_documents
(
    document_type   document_type,
    document_number integer,
    document_date   date,

    PRIMARY KEY (document_number)
)