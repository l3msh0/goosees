
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `companies` ADD `address` VARCHAR(200) NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `companies` DROP `address`;
