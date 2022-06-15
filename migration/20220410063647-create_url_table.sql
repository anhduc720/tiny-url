
-- +migrate Up
CREATE TABLE IF NOT EXISTS `urls` (
    `hash` VARCHAR(16) NOT NULL,
    `original_url` VARCHAR(512) NOT NULL,
    `creation_date` DATETIME,
    `expiration_date` DATETIME,
    `user_id` int
    PRIMARY KEY (`hash`)
    ) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS `urls`;