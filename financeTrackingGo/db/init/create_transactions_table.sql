CREATE TABLE IF NOT EXISTS tb_transactions (
  id TEXT PRIMARY KEY,
  agency TEXT NOT NULL,
  account_id TEXT NOT NULL,
  release_date DATE NOT NULL,
  accounting_date DATE NOT NULL,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  income NUMERIC(18,2) NOT NULL DEFAULT 0,
  expense NUMERIC(18,2) NOT NULL DEFAULT 0,
  daily_balance NUMERIC(18,2) NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_tb_transactions_account_date
  ON tb_transactions (account_id, release_date);