DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'crypto_currency') THEN
        CREATE TYPE crypto_currency AS ENUM ('Bitcoin', 'USDT', 'Ethereum');
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS wallets (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  balance NUMERIC NOT NULL DEFAULT 0,
  currency_type crypto_currency
);
