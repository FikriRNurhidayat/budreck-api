-- Create new transactions table
-- Store any information about user's transactions
CREATE TABLE IF NOT EXISTS transactions (
  id uuid DEFAULT uuid_generate_v4(),
  amount FLOAT NOT NULL,
  wallet_id UUID NOT NULL,

  created_at TIMESTAMP NOT NULL DEFAULT NOW(),

  PRIMARY KEY (id),
  CONSTRAINT fk_transactions_wallets FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE CASCADE,
  CONSTRAINT transactions_amount_not_zero_check CHECK (amount != 0)
);
