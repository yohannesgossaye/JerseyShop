-- Drop is_merchant column from customers
ALTER TABLE customers
DROP COLUMN IF EXISTS is_merchant;


