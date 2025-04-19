DO
$$
    BEGIN
        IF NOT EXISTS (SELECT
                       FROM pg_roles
                       WHERE rolname = 'analytic') THEN
            CREATE ROLE analytic NOLOGIN;
        END IF;
    END
$$;
-- USAGE: viewing but no access for objects
-- grant select
GRANT USAGE ON SCHEMA public TO analytic;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO analytic;
-- for next objects in scheme
ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT SELECT ON TABLES TO analytic;