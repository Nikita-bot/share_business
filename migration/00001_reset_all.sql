-- +goose Up  
;
-- +goose Down
DROP TABLE IF EXISTS role_permissions;

DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS tokens;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS wash_servers;
DROP TABLE IF EXISTS wash_sessions;