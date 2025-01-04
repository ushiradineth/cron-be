ALTER TABLE users
ADD COLUMN active BOOLEAN DEFAULT TRUE,
ADD COLUMN deleted_at TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE events
ADD COLUMN active BOOLEAN DEFAULT TRUE,
ADD COLUMN deleted_at TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
