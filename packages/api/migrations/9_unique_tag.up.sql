alter table `tag` add constraint unique_rid_name unique (`name`, `restaurant_id`);
