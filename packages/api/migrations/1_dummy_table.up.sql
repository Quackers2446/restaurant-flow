create table `dummy_table` (
    `id` int primary key auto_increment,
    `description` text
);

insert into dummy_table (`description`) values ("test1"), ("test2"), ("test3");
