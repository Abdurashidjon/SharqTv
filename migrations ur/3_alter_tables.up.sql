ALTER TABLE researcher 
    ADD COLUMN company_id uuid REFERENCES company(id);

ALTER TABLE respondent
    ADD COLUMN company varchar(255) default '';

ALTER TABLE respondent 
    ADD COLUMN position varchar(255) default '';

ALTER TABLE respondent 
    ADD COLUMN description varchar(255) default '';

ALTER TABLE respondent
    ADD COLUMN photo varchar(255) default '';