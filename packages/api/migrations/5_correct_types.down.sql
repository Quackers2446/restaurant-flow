alter table `OpeningPeriod` modify column `open_time` char(4) not null;
alter table `OpeningPeriod` modify column `close_time` char(4) null;

alter table `Review` modify column `rating` int not null;

alter table `GoogleRestaurant` modify column `name` varchar(64) not null;

alter table `GoogleRestaurant` modify column `avg_rating` tinyint;
