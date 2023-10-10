-- Psql Table Schema for task objects

DROP TABLE IF EXISTS tasks;

CREATE TABLE IF NOT EXISTS tasks (
    id int NOT NULL,
    name varchar(255) DEFAULT NULL,
    created timestamp DEFAULT NULL,
    done bit,
    PRIMARY KEY(id)
);

INSERT INTO tasks (id, name, created, done) VALUES 
    ( 1, 'Mow Lawn', current_timestamp, CAST('0' AS bit) ),
    ( 2, 'Buy Groceries', current_timestamp, CAST('0' AS bit) ),
    ( 3, 'Buy Groceries', current_timestamp, CAST('0' AS bit) )
;
