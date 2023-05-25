CREATE TABLE couriers
(
    id            serial       not null unique,
    type          varchar(255) not null,
    region     integer[] not null ,
    working_hours varchar(255)[] not null
);

CREATE TABLE orders
(
    id          serial       not null unique,
    weight      float not null,
    cost int not null,
    region     int not null ,
    delivery_hours varchar(255) not null,
    complete_time varchar(255) 
);




CREATE TABLE completeOrderdate
(
    id      serial                                           not null unique,
    couriers_id int references couriers (id) on delete cascade      not null,
    orders_id int[] not null,
    deliveryDate date not null
);