-- +goose Up  
;
-- +goose Down
DROP TABLE IF EXISTS role_permissions;

DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS wash_servers;