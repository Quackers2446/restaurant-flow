alter table `OpeningPeriod` modify column `open_time` smallint unsigned not null;
alter table `OpeningPeriod` modify column `close_time` smallint unsigned null;

alter table `Review` modify column `rating` tinyint unsigned not null check (`rating` >= 0 and `rating` <= 10);

alter table `GoogleRestaurant` modify column `name` varchar(256) not null; -- I can't believe some restaurants have such a long fucking name
