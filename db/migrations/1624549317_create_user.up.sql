-- Create new users table
-- Store authentication related data of a user
CREATE TABLE IF NOT EXISTS users (
  id uuid DEFAULT uuid_generate_v4(),
  email VARCHAR,
  encrypted_password VARCHAR,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,

  PRIMARY KEY (id),
  CONSTRAINT users_email_check CHECK (
    email IS NULL OR
    TRIM(email) <> '' AND
    email ~ '^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$'
  ),
  CONSTRAINT users_encrypted_password_check CHECK (
    encrypted_password IS NULL OR
    TRIM(encrypted_password) <> ''
  )
);

-- Create unique index, ensure email is always unique when it is not deleted
CREATE UNIQUE INDEX users_email_key ON users(email) WHERE (deleted_at IS NULL); 

COMMENT ON COLUMN users.id IS 'unique identifier';
COMMENT ON COLUMN users.email IS 'can be null, cannot be empty, and must be valid email format';
COMMENT ON COLUMN users.encrypted_password IS 'a generated hash, can be null, and cannot be empty.';
COMMENT ON COLUMN users.deleted_at IS 'can be null, tell if the record been soft deleted';
