create table "ShopFlower"
(
    id           serial,
    name_fl      varchar,
    count_fl     bigint,
    date_fl_shop date
);

alter table "ShopFlower"
    owner to postgres;

INSERT INTO "Flower"."ShopFlower" (id, name_fl, count_fl, date_fl_shop) VALUES (1, 'Алое', 454, '2022-01-12');

