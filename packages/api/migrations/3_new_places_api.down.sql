alter table `GoogleRestaurant` rename column `viewport_high` to `northeast`;
alter table `GoogleRestaurant` rename column `viewport_low` to `southwest`;

alter table `GoogleRestaurant` drop column `supports_takeout`;

alter table `GoogleRestaurant` modify column `serves_breakfast` boolean;
alter table `GoogleRestaurant` modify column `serves_brunch` boolean;
alter table `GoogleRestaurant` modify column `serves_dinner` boolean;
alter table `GoogleRestaurant` modify column `serves_lunch` boolean;
alter table `GoogleRestaurant` modify column `serves_vegetarian_food` boolean;
alter table `GoogleRestaurant` modify column `serves_wine` boolean;
alter table `GoogleRestaurant` drop column `serves_beer`;
alter table `GoogleRestaurant` drop column `serves_cocktails`;
alter table `GoogleRestaurant` drop column `serves_coffee`;
alter table `GoogleRestaurant` drop column `serves_dessert`;

alter table `GoogleRestaurant` drop column `good_for_groups`;
alter table `GoogleRestaurant` drop column `good_for_watching_sports`;
alter table `GoogleRestaurant` drop column `has_outdoor_seating`;
alter table `GoogleRestaurant` drop column `has_restroom`;

alter table `GoogleRestaurant` drop column `accepts_credit_cards`;
alter table `GoogleRestaurant` drop column `accepts_debit_cards`;
alter table `GoogleRestaurant` drop column `accepts_cash_only`;
alter table `GoogleRestaurant` drop column `accepts_nfc`;

alter table `GoogleRestaurant` modify column `wheelchair_accessible_entrance` boolean;
alter table `GoogleRestaurant` drop column `wheelchair_accessible_seating`;
