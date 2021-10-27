ALTER TABLE respondent
    ADD COLUMN rating_communication varchar(255) default '0';
ALTER TABLE respondent
    ADD COLUMN rating_experience varchar(255) default '0';
ALTER TABLE respondent
    ADD COLUMN rating_punctuality varchar(255) default '0';
