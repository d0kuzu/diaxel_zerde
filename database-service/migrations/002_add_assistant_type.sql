-- Migration: Add type column to assistants table
-- Distinguishes between 'api' and 'telegram' assistant types

ALTER TABLE assistants ADD COLUMN IF NOT EXISTS type VARCHAR(20) NOT NULL DEFAULT 'api';
