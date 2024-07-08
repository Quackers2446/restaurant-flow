insert into `user` set
    `user_id`=unhex(replace("00000000-0000-0000-0000-000000000000",'-','')),
    `name`="Placeholder",
    `username`="Placeholder",
    `email`="email";

alter table `review` add column `is_anonymous` boolean not null default false;
alter table `user` modify column `email` varchar(320) not null unique;
alter table `user` add column `user_id_text` varchar(36) generated always as
    (insert(
        insert(
            insert(
                insert(hex(user_id),9,0,'-'),
                14,0,'-'
            ),
            19,0,'-'
        ),
        24,0,'-')
    ) virtual;

alter table `review` add unique `unique_user_restaurant_review`(`user_id`, `restaurant_id`);
