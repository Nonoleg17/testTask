CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS "users" (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    firstname text not null,
    surname text not null,
    middlename text not null,
    fio text GENERATED ALWAYS AS (surname || ' ' || firstname || ' ' ||
                                  middlename) STORED,
    sex text not null,
    age int not null
    );
CREATE TABLE IF NOT EXISTS "friendships" (
    first_user_id uuid,
    second_user_id uuid,
    friendship_status text,
    PRIMARY KEY (first_user_id, second_user_id),
    CONSTRAINT fk_user_1_id FOREIGN KEY (first_user_id) REFERENCES "users"(id),
    CONSTRAINT fk_user_2_id FOREIGN KEY (second_user_id) REFERENCES "users"(id)
);
CREATE TABLE IF NOT EXISTS "orders"
(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id uuid NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES "users"(id)
);

CREATE TABLE IF NOT EXISTS "products"
(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    description text,
    price int,
    currency text,
    left_in_stock int

);
CREATE TABLE IF NOT EXISTS "orderProducts"
(
    order_id BIGINT,
    product_id BIGINT,
    PRIMARY KEY (order_id, product_id),
    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES "orders"(id),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES "products"(id)
);

CREATE FUNCTION check_left_in_stock() RETURNS trigger as $$
BEGIN
        If (NEW.left_in_stock < 0) then
            raise exception 'The number of products cannot be less than 0';
end if;
return NEW;
END;
    $$ language plpgsql;

CREATE TRIGGER check_left_in_stock_trigger
    before insert or update on "products"
    for each row
    execute function check_left_in_stock();

INSERT INTO "users"(firstname,surname,middlename,sex,age) values
                                                             ('John','Snow','Eddard','M',20),
                                                             ('Mike','Lobo', 'Dekkard','M',21),
                                                             ('Lina','Rola', 'Ridic','W',45),
                                                             ('Alex','Carsa', 'Robin','M',30),
                                                             ('Peter','Polo', 'Line','M',60);


INSERT INTO "products"(description,price,currency, left_in_stock) values
                                                           ('Car', 1000, 'dollars', 10),
                                                           ('Book', 500, 'rubbles', 1),
                                                           ('Radio', 15, 'euro', 0),
                                                           ('Clothes', 300, 'dollars', 19),
                                                           ('Clock', 800, 'rubbles', 2);

