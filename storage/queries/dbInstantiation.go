package queries

const (
	//CreateSchema- query to create sschema
	CreateSchema = `CREATE SCHEMA IF NOT EXISTS rtArchive 
    AUTHORIZATION postgres;`
	//CreateDialogSourceEnum - query to create enum fo dialog source platform name
	CreateDialogSourceEnum = `
	DO $$
	BEGIN
		CREATE TYPE rtArchive.dialog_source AS ENUM ('telegram', 'fb_messenger', 'whatsapp', 'viber', 'vk');
	EXCEPTION
    WHEN duplicate_object THEN null;
	END $$;
	ALTER TYPE rtArchive.dialog_source
    OWNER TO postgres;`

	//CreateTableRoundTrips - query to create table for roundtrips
	CreateTableRoundTrips = `CREATE TABLE IF NOT EXISTS rtArchive.roundtrips
	(
    id SERIAL PRIMARY KEY,
    source rtArchive.dialog_source NOT NULL,
    chat_id character varying NOT NULL,
    username character varying NOT NULL,
    request text,
    response text,
    created_at timestamp without time zone default current_timestamp,
    updated_at timestamp without time zone
	)
	WITH (
    OIDS = FALSE
	);

	ALTER TABLE rtArchive.roundtrips
    OWNER to postgres;`

	//CreateUpdatedAtFunction - create func to update timestamp column updated_at
	CreateUpdatedAtFunction = `CREATE OR REPLACE FUNCTION update_created_at_column() 
	RETURNS TRIGGER AS $$
	BEGIN
    NEW.updated_at = now();
    RETURN NEW; 
	END;
	$$ language 'plpgsql';`

	//CreateUpdatedAtTrigger  - create trigger that run func before updated any row
	CreateUpdatedAtTrigger = `
	DO $$
	BEGIN	
	CREATE TRIGGER update_created_at BEFORE UPDATE ON rtArchive.roundtrips FOR EACH ROW EXECUTE PROCEDURE  update_created_at_column();
	EXCEPTION
    WHEN duplicate_object THEN null;
	END $$;`
)
