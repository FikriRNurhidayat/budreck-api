-- Create new wallets table
-- Store the amount of money user has, owned by one user
CREATE TABLE IF NOT EXISTS wallets (
  id UUID DEFAULT uuid_generate_v4(),
  balance FLOAT NOT NULL DEFAULT 0,
  user_id UUID NOT NULL,

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,

  PRIMARY KEY (id),
  CONSTRAINT fk_wallets_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  CONSTRAINT wallets_positive_balance_check CHECK (balance > 0)
);
