insert into `User` set
    `user_id`=unhex(replace(uuid(),'-','')),
    `name`="Placeholder",
    `username`="Placeholder",
    `email`="email";

alter table `Review` add column `is_anonymous` boolean not null default false;
alter table `User` modify column `email` varchar(320) not null unique;
