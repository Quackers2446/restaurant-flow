alter table `opening_period` modify column `open_time` smallint unsigned not null;
alter table `opening_period` modify column `close_time` smallint unsigned null;

alter table `review` modify column `rating` tinyint unsigned not null check (`rating` >= 0 and `rating` <= 10);

alter table `google_restaurant` modify column `name` varchar(256) not null; -- I can't believe some restaurants have such a long fucking name

alter table `google_restaurant` modify column `avg_rating` float(2, 1);
