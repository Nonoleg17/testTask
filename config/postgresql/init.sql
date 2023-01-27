CREATE TABLE IF NOT EXISTS "User" (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    firstname text,
    surname text,
    midlename text,
    fio text GENERATED ALWAYS AS (surname || ' ' || firstname || ' ' ||
                                  midlename) STORED,
    sex text,
    age int
    );

-- CREATE FUNCTION change_fio() RETURNS trigger as $$
--     BEGIN
--         update "User" SET fio = NEW.midlename where NEW.id == OLD.id;
--         return NEW
--     END;
--     $$ language plpgsql;
--
-- CREATE TRIGGER  IF NOT EXISTS add_fio
--     after insert or update on "User"
--     for each row
--     execute function change_fio();