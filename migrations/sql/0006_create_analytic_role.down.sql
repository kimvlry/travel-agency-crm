DO $$
    DECLARE
        rec record;
    BEGIN
        FOR rec IN
            SELECT table_schema, table_name
            FROM information_schema.role_table_grants
            WHERE grantee = 'analytic'
            LOOP
                EXECUTE format('REVOKE ALL PRIVILEGES ON TABLE %I.%I FROM analytic', rec.table_schema, rec.table_name);
            END LOOP;
    END $$;

REVOKE USAGE ON SCHEMA public FROM analytic;

ALTER DEFAULT PRIVILEGES IN SCHEMA public
    REVOKE SELECT ON TABLES FROM analytic;

DROP ROLE IF EXISTS analytic;
