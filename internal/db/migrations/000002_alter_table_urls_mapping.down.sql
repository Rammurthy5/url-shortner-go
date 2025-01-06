-- Add new identity column
ALTER TABLE url_mapping ADD COLUMN new_id BIGINT;

-- Drop old primary key constraint
ALTER TABLE url_mapping DROP CONSTRAINT url_mapping_pkey;

-- Add primary key to the new column
ALTER TABLE url_mapping ADD PRIMARY KEY (new_id);

-- Drop the old column
ALTER TABLE url_mapping DROP COLUMN id;

-- Rename new column to 'id' (optional)
ALTER TABLE url_mapping RENAME COLUMN new_id TO id;
