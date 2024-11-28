CREATE TABLE records (user_id TEXT NOT NULL, id TEXT NOT NULL, data bytea NOT NULL, revision INTEGER NOT NULL, schema_version TEXT NOT NULL, PRIMARY KEY (user_id, id));
CREATE TABLE user_revisions (user_id TEXT NOT NULL, revision INTEGER NOT NULL, PRIMARY KEY (user_id, revision));
