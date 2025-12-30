-- run as root init script, switch to appdb
CREATE USER debezium WITH REPLICATION PASSWORD 'debezium';
\connect appdb
CREATE PUBLICATION dbz_publication FOR ALL TABLES;
