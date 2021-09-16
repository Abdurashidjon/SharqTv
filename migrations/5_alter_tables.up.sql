ALTER TABLE company
    ADD COLUMN owner_id uuid REFERENCES researcher(id);

