-- Remove user_id indexes and column from chats table
DROP INDEX IF EXISTS idx_chats_user_id;
DROP INDEX IF EXISTS idx_chats_user_id_assistant_id;
ALTER TABLE chats DROP COLUMN IF EXISTS user_id;
