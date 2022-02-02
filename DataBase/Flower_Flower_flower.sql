create table flower
(
    id       serial
        constraint flower_pk
            primary key,
    name_fl  varchar default 255 not null,
    count_fl integer             not null,
    money_fl double precision    not null,
    date_fl  date                not null
);

alter table flower
    owner to postgres;

create unique index flower_id_uindex
    on flower (id);

INSERT INTO "Flower".flower (id, name_fl, count_fl, money_fl, date_fl) VALUES (1, 'Незабудка', 27, 37.24, '2022-01-19');

