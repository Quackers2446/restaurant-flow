delete from `user` where `email`="email";

alter table `review` drop column `is_anonymous`;
alter table `user` modify column `email` varchar(320) not null;
alter table `user` drop column `user_id_text`;

alter table `review` drop index `unique_user_restaurant_review`;
