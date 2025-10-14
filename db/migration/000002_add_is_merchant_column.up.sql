-- Add is_merchant column to customers
ALTER TABLE customers
ADD COLUMN IF NOT EXISTS is_merchant BOOLEAN NOT NULL DEFAULT FALSE;


