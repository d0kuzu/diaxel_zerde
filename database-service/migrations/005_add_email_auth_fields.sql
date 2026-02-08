-- Add email authentication fields to users table
ALTER TABLE users ADD COLUMN IF NOT EXISTS email VARCHAR(255) UNIQUE;
ALTER TABLE users ADD COLUMN IF NOT EXISTS password_hash VARCHAR(255);
ALTER TABLE users ADD COLUMN IF NOT EXISTS role VARCHAR(50) DEFAULT 'user';

-- Create index for email lookups
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Make telegram_id nullable since we'll have email users too
ALTER TABLE users ALTER COLUMN telegram_id DROP NOT NULL;

-- Add constraint to ensure at least one identifier exists
ALTER TABLE users ADD CONSTRAINT check_user_identifier 
  CHECK (email IS NOT NULL OR telegram_id IS NOT NULL);
