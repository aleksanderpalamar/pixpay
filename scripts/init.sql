CREATE TABLE payments (
  id SERIAL PRIMARY KEY,
  receiver_id VARCHAR(255) NOT NULL,
  sender_id VARCHAR(255) NOT NULL,
  amount DECIMAL(10, 2) NOT NULL,
  status VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);