ALTER TABLE respondent 
    ADD COLUMN account_number int default 0;

ALTER TABLE company 
    ADD COLUMN account_number int default 0;