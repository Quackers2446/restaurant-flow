create table `dummyTable` (
    `id` int primary key auto_increment,
    `description` text
);

insert into dummyTable (`description`) values ("test1"), ("test2"), ("test3");
