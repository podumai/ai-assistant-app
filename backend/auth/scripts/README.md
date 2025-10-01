# Scripts

## user-storage

Example `init.sql` for Postgres:  

```sql
CREATE SCHEMA IF NOT EXISTS auth;

CREATE TYPE users_role AS ENUM ('user', 'admin');
CREATE TYPE users_status AS ENUM ('active', 'inactive');

CREATE TABLE IF NOT EXISTS auth.users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(40) NOT NULL UNIQUE,
  email VARCHAR(320) NOT NULL UNIQUE,
  passwd BYTEA NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_DATE,
  user_role USERS_ROLE NOT NULL DEFAULT('user'),
  user_status USERS_STATUS NOT NULL DEFAULT('active')
);

CREATE ROLE backend WITH NOSUPERUSER
                         NOCREATEDB
                         NOCREATEROLE
                         NOINHERIT
                         LOGIN
                         REPLICATION
                         PASSWORD 'your_super_secret_password_for_backend_user';

GRANT USAGE ON SCHEMA auth TO backend GRANTED BY postgres;

GRANT USAGE, SELECT ON SEQUENCE auth.users_id_seq TO backend;

GRANT SELECT, UPDATE, INSERT, DELETE ON auth.users TO backend GRANTED BY postgres;
```
