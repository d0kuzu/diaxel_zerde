-- Remove bot_token column from assistants table
ALTER TABLE assistants DROP COLUMN IF EXISTS bot_token;

-- Remove index for bot_token if it exists
DROP INDEX IF EXISTS idx_assistants_bot_token;
