-- Create assistants table for Telegram bot registration
CREATE TABLE IF NOT EXISTS assistants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    bot_token VARCHAR(255) UNIQUE NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_assistants_user_id ON assistants(user_id);
CREATE INDEX IF NOT EXISTS idx_assistants_bot_token ON assistants(bot_token);

-- Create trigger to update updated_at timestamp
CREATE TRIGGER update_assistants_updated_at BEFORE UPDATE ON assistants
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
