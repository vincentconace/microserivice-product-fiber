CREATE DATABASE products;

CREATE TABLE products (
    id int not null primary key auto_increment,
    name varchar(255) not null,
    description text not null,
    price int not null,
    stock int not null,
    status bool not null,
)