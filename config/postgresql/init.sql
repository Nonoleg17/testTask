CREATE TABLE IF NOT EXISTS "User" (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    firstname text not null,
    surname text not null,
    middlename text not null,
    fio text GENERATED ALWAYS AS (surname || ' ' || firstname || ' ' ||
                                  middlename) STORED,
    sex text not null,
    age int not null
    );
CREATE TABLE IF NOT EXISTS "Friendship" (
    user_1_id BIGINT,
    user_2_id BIGINT,
    friendship_status text,
    PRIMARY KEY (user_1_id, user_2_id),
    CONSTRAINT fk_user_1_id FOREIGN KEY (user_1_id) REFERENCES "User"(id),
    CONSTRAINT fk_user_2_id FOREIGN KEY (user_2_id) REFERENCES "User"(id)
);
CREATE TABLE IF NOT EXISTS "Order"
(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES "User"(id)
);

CREATE TABLE IF NOT EXISTS "Product"
(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    description text,
    price int,
    currency text,
    left_in_stock int

);
CREATE TABLE IF NOT EXISTS "Order_product"
(
    order_id BIGINT,
    product_id BIGINT,
    PRIMARY KEY (order_id, product_id),
    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES "Order"(id),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES "Product"(id)
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
    before insert or update on "Product"
    for each row
    execute function check_left_in_stock();

INSERT INTO "User"(firstname,surname,middlename,sex,age) values
                                                             ('John','Snow','Eddard','M',20),
                                                             ('Mike','Lobo', 'Dekkard','M',21),
                                                             ('Lina','Rola', 'Ridic','W',45),
                                                             ('Alex','Carsa', 'Robin','M',30),
                                                             ('Peter','Polo', 'Line','M',60);


INSERT INTO "Product"(description,price,currency, left_in_stock) values
                                                           ('Car', 1000, 'dollars', 10),
                                                           ('Book', 500, 'rubbles', 1),
                                                           ('Radio', 15, 'euro', 0),
                                                           ('Clothes', 300, 'dollars', 19),
                                                           ('Clock', 800, 'rubbles', 2);

