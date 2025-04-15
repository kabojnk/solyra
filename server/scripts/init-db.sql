-- Initialize database schema for etrenank API

-- Create applications table
CREATE TABLE IF NOT EXISTS applications (
    id UUID PRIMARY KEY,
    client_id VARCHAR(255) NOT NULL UNIQUE,
    client_secret VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create index on client_id for faster lookups
CREATE INDEX IF NOT EXISTS idx_applications_client_id ON applications(client_id);

-- Insert a sample application for testing
INSERT INTO applications (id, client_id, client_secret)
VALUES 
    ('00000000-0000-0000-0000-000000000001', 'test_client', 'test_secret')
ON CONFLICT (id) DO NOTHING;
