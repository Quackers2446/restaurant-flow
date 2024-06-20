delete from `User` where `email`="email";

alter table `Review` drop column `is_anonymous`;
alter table `User` modify column `email` varchar(320) not null;
