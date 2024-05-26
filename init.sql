-- Define DB schema and seed here
-- We will setup migrations later if needed, but this will suffice

-- CreateTable
CREATE TABLE `dummyTable` (
    `id` INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    `description` TEXT,
    PRIMARY KEY (`id`)
);

INSERT INTO `dummyTable` (`description`) VALUES ("test1"), ("test2"), ("test3");
