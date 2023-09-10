CREATE DATABASE alta_online_shop;

create table users (
  `id` varchar(36) primary key not null,
  `name` varchar(255) not null,
  `birthdate` date,
  `gender` enum('male', 'female'),
  `status` enum('active', 'inactive'),
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

create table addresses (
  `id` varchar(36) primary key not null,
  `user_id` varchar(36) not null,
  `line_one` varchar(255),
  `line_two` varchar(255),
  `city` varchar(63),
  `province` varchar(63),
  `country` varchar(31)
);

create table product_descriptions (
  `id` int primary key not null,
  `name` varchar(255) not null,
  `price` float not null,
  `description` text
);

create table product_types (
  `id` int not null primary key,
  `name` varchar(127) not null
);

create table products (
  `id` int primary key not null,
  `product_description_id` int not null unique,
  `product_type_id` int not null
);

create table operators (
  `id` varchar(36) not null primary key,
  `name` varchar(255) not null
);

create table payment_methods (
  `id` int not null primary key,
  `payment_method_description_id` int unique not null
);

create table payment_method_descriptions (
  `id` int not null primary key,
  `name` varchar(63) not null,
  `description` text
);

create table user_payment_method_details (
  `user_id` varchar(36) not null,
  `payment_method_id` int not null
);

create table transactions (
  `id` varchar(36) not null primary key,
  `product_id` int not null,
  `user_id` varchar(36) not null,
  `operator_id` varchar(36) not null,
  `transaction_detail_id` varchar(36) not null unique,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

create table transaction_details (
  `id` varchar(36) not null primary key,
  `amount` int not null,
  `costs` float not null
);

create table couriers (
  `id` varchar(36) primary key not null,
  `name` varchar(127) not null,
  `created_at` timestamp default current_timestamp,
  `updated_at` datetime default current_timestamp on update current_timestamp
);

alter table couriers
  add column `basic_cost` float;

alter table couriers
  rename shippings;

alter table addresses
  add key `address_with_user` (`user_id`);

alter table products
  add key `product_with_product_type`(`product_type_id`),
  add key `product_meets_product_description`(`product_description_id`);

alter table payment_methods
  add key `payment_method_meets_payment_description_method`(`payment_method_description_id`);

alter table transactions
  add key `transaction_with_product`(`product_id`),
  add key `transaction_with_operator`(`operator_id`),
  add key `transaction_with_user`(`user_id`),
  add key `transaction_meets_transaction_detail`(`transaction_detail_id`);

alter table user_payment_method_details
  add key `payment_method_with_user`(`payment_method_id`),
  add key `user_with_payment_method`(`user_id`);


alter table addresses
  add constraint `address_with_user` foreign key (`user_id`) references users(`id`);

alter table products
  add constraint `product_with_product_type` foreign key (`product_type_id`) references product_types(`id`),
  add constraint `product_meets_product_description` foreign key (`product_description_id`) references product_descriptions(`id`);

alter table payment_methods
  add constraint `payment_method_meets_payment_method_description` foreign key (`payment_method_description_id`) references payment_methods(`id`);

alter table transactions
  add constraint `transaction_with_product` foreign key (`product_id`) references products(`id`),
  add constraint `transaction_with_operator` foreign key (`operator_id`) references operators(`id`),
  add constraint `transaction_with_user` foreign key (`user_id`) references users(`id`),
  add constraint `transaction_meets_transaction_detail` foreign key (`transaction_detail_id`) references transaction_details(`id`);

alter table user_payment_method_details
  add constraint `user_with_payment_method` foreign key (`user_id`) references users(`id`),
  add constraint `payment_method_with_user` foreign key (`payment_method_id`) references payment_methods(`id`);

drop table shippings;   