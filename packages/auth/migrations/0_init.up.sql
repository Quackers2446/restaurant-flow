create table `user` (
    `user_id` binary(16) primary key,
    `user_id_text` varchar(36) generated always as (
        insert(
            insert(
                insert(
                    insert(hex(user.user_id),9,0,'-'),
                    14,0,'-'
                ),
                19,0,'-'
            ),
            24,0,'-'
        )
    ) virtual,
    `email` varchar(320) not null unique
);

create table `user_auth` (
    `user_id` binary(16) primary key,
    `user_id_text` varchar(36) generated always as (
        insert(
            insert(
                insert(
                    insert(hex(user_auth.user_id),9,0,'-'),
                    14,0,'-'
                ),
                19,0,'-'
            ),
            24,0,'-'
        )
    ) virtual,
    `password_hash` char(60) character set latin1 collate latin1_bin not null
);

alter table `user_auth` add foreign key (`user_id`) references `user`(`user_id`);

create table `session` (
    `session_id` int primary key auto_increment,
    `user_id` binary(16) not null,
    `user_id_text` varchar(36) generated always as (
        insert(
            insert(
                insert(
                    insert(hex(session.user_id),9,0,'-'),
                    14,0,'-'
                ),
                19,0,'-'
            ),
            24,0,'-'
        )
    ) virtual,
    `ip_addr` char(45) not null,
    `user_agent` text not null,
    `created_at` timestamp not null default current_timestamp,
    `last_used_at` timestamp not null default current_timestamp on update current_timestamp,
    `expires_at` timestamp not null,
    `refresh_key` binary(64) not null unique,
    `valid` boolean not null default true,

    index `session_user_id` (`user_id`)
);

alter table `session` add foreign key (`user_id`) references `user`(`user_id`);
