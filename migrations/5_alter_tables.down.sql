ALTER TABLE company
    DROP COLUMN owner_id uuid REFERENCES researcher(id);
