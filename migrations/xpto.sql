use xpto;

DROP TABLE IF EXISTS `transactions`;
DROP TABLE IF EXISTS `accounts`;

CREATE TABLE IF NOT EXISTS `accounts` (
    `id` VARCHAR(36) NOT NULL,
    `document_number` VARCHAR(11) NOT NULL UNIQUE,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)  DEFAULT CHARSET=UTF8;


CREATE TABLE IF NOT EXISTS `transactions` (
    `id` VARCHAR(36) NOT NULL,
    `account_id` VARCHAR(36) NOT NULL,
    `operation_type` ENUM('1', '2', '3', '4') NOT NULL,
    `amount` VARCHAR(11) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`account_id`)
        REFERENCES accounts (`id`)
        ON DELETE CASCADE
)  DEFAULT CHARSET=UTF8;

COMMIT;