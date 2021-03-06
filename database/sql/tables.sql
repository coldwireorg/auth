SELECT 'CREATE DATABASE auth ENCODING ''UTF8'''
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'auth')\gexec

CREATE TABLE IF NOT EXISTS users (
    username    VARCHAR(25) PRIMARY KEY NOT NULL UNIQUE,
    password    VARCHAR(256) NOT NULL,
    public_key  VARCHAR(256) NOT NULL,
    private_key VARCHAR(256) NOT NULL
);
