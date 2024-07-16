CREATE TABLE IF NOT EXISTS `db_tasks`.`tasks` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `completed` BOOLEAN NOT NULL,
    PRIMARY KEY (`id`)
);