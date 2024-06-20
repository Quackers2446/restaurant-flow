delete from `User` where `email`="email";

alter table `Review` drop column `is_anonymous`;
alter table `User` modify column `email` varchar(320) not null;
alter table `User` drop column `user_id_text`;

alter table `Review` drop index `unique_user_restaurant_review`;
