-- Add message_count column to chats table
ALTER TABLE chats ADD COLUMN message_count INTEGER DEFAULT 0;

-- Create index for better performance on assistant_id queries
CREATE INDEX IF NOT EXISTS idx_chats_assistant_id_created_at ON chats(assistant_id, created_at DESC);

-- Create index for user search functionality
CREATE INDEX IF NOT EXISTS idx_chats_user_id_assistant_id ON chats(user_id, assistant_id);
