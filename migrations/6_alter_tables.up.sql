ALTER TABLE respondent
        ADD COLUMN inn varchar(255) default '';

ALTER TABLE respondent ALTER COLUMN rating_communication TYPE float USING rating_communication::double precision;

ALTER TABLE respondent ALTER COLUMN rating_experience TYPE float USING rating_experience::double precision;

ALTER TABLE respondent ALTER COLUMN rating_punctuality TYPE float USING rating_punctuality::double precision;
