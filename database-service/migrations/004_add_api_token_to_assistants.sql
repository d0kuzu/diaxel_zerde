-- Add api_token column to assistants table
ALTER TABLE assistants ADD COLUMN IF NOT EXISTS api_token VARCHAR(255) NOT NULL DEFAULT '';

-- Create index for api_token for better lookup performance
CREATE INDEX IF NOT EXISTS idx_assistants_api_token ON assistants(api_token);

-- Add trigger to update updated_at timestamp for assistants table
CREATE TRIGGER update_assistants_updated_at BEFORE UPDATE ON assistants
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
