-- Targeted Migration: Cleanup chats table
-- This migration only removes the obsolete user_id column and related indexes.

-- Drop indexes that depend on user_id
DROP INDEX IF EXISTS idx_chats_user_id;
DROP INDEX IF EXISTS idx_chats_user_id_assistant_id;

-- Drop the user_id column
ALTER TABLE chats DROP COLUMN IF EXISTS user_id;
