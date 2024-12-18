ALTER TABLE wallets ADD user_id VARCHAR(36),
ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id);