CREATE TYPE crypto_currency AS ENUM ('Bitcoin', 'USDT', 'Ethereum');

  CREATE TABLE wallets (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    balance NUMERIC NOT NULL DEFAULT 0
  );