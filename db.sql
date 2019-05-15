CREATE TABLE IF NOT EXISTS users (
    ID INT NOT NULL AUTO_INCREMENT,
    Username VARCHAR(24) CHARACTER SET utf8 NOT NULL,
    Password VARCHAR(255) CHARACTER SET utf8 NOT NULL,
    Token VARCHAR(255) CHARACTER SET utf8 NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE IF NOT EXISTS accounts (
    ID INT NOT NULL AUTO_INCREMENT,
    Region VARCHAR(4) CHARACTER SET utf8 NOT NULL,
    Tag VARCHAR(255) CHARACTER SET utf8 NOT NULL,
    Ign VARCHAR(16) CHARACTER SET utf8 NOT NULL,
    Username VARCHAR(24) CHARACTER SET utf8 NOT NULL,
    Password VARCHAR(255) CHARACTER SET utf8 NOT NULL,
    User VARCHAR(24) CHARACTER SET utf8 NOT NULL,
    Leaverbuster INT NOT NULL,
    Ban DATETIME,
    Perma BOOLEAN NOT NULL,
    PasswordChanged BOOLEAN NOT NULL,
    Pre30 BOOLEAN NOT NULL,
    Elo VARCHAR(24) CHARACTER SET utf8 NOT NULL DEFAULT "Not parsed",
    PRIMARY KEY (ID)
);
