alter table `opening_period` modify column `open_time` char(4) not null;
alter table `opening_period` modify column `close_time` char(4) null;

alter table `review` modify column `rating` int not null;

alter table `google_restaurant` modify column `name` varchar(64) not null;

alter table `google_restaurant` modify column `avg_rating` tinyint;
