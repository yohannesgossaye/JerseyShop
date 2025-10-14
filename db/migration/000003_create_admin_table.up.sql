CREATE TABLE admins (
  id SERIAL PRIMARY KEY,
  customer_id BIGINT REFERENCES customers(id) ON DELETE CASCADE,
  role VARCHAR(50) DEFAULT 'superadmin',
  permissions JSONB DEFAULT '{}'::jsonb,
  current_merchant_numbers VARCHAR(50),
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);


-- Indexes
CREATE INDEX idx_admins_customer_id ON admins(customer_id);
CREATE INDEX idx_admins_role ON admins(role);
