CREATE TABLE merchants (
  id SERIAL PRIMARY KEY,
  customer_id BIGINT NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
  store_name VARCHAR(150) NOT NULL,
  business_license VARCHAR(100),
  business_address VARCHAR(255),
  business_type VARCHAR(100),
  business_email VARCHAR(150),
  business_phone VARCHAR(30),
  owner_first_name VARCHAR(100),
  owner_last_name VARCHAR(100),
  is_verified BOOLEAN DEFAULT FALSE,  
  is_active BOOLEAN DEFAULT TRUE,     
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

-- Indexes
CREATE INDEX idx_merchants_customer_id ON merchants(customer_id);
CREATE INDEX idx_merchants_store_name ON merchants(store_name);
CREATE INDEX idx_merchants_is_verified ON merchants(is_verified);
