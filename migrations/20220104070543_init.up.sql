-- users --
CREATE TABLE IF NOT EXISTS users
(
    user_id    BIGSERIAL PRIMARY KEY,
    nickname   TEXT UNIQUE NOT NULL,
    first_name TEXT        NOT NULL,
    last_name  TEXT        NOT NULL,

    created_at TIMESTAMP   NOT NULL,
    updated_at TIMESTAMP   NOT NULL,
    deleted_at TIMESTAMP
);

INSERT INTO users (user_id, nickname, first_name, last_name, created_at, updated_at)
VALUES (1, 'admin', 'Администратор', 'Главный', NOW(), NOW());
ALTER SEQUENCE users_user_id_seq RESTART WITH 2;

-- roles and permissions --
CREATE TABLE IF NOT EXISTS permissions
(
    permission_id SMALLINT PRIMARY KEY,
    name          TEXT NOT NULL UNIQUE
);

INSERT INTO permissions (permission_id, name)
VALUES (1, '-read'),
       (2, '-write'),
       (3, 'edit'),
       (4, 'delete'),
       (5, 'administrate'),
       (6, 'write');

CREATE TABLE IF NOT EXISTS roles
(
    role_id BIGSERIAL PRIMARY KEY,
    name    TEXT NOT NULL
);

INSERT INTO roles (role_id, name)
VALUES (1, 'admins'),
       (2, 'banned'),
       (3, 'read only'),
       (4, 'editors'),
       (5, 'members');
ALTER SEQUENCE roles_role_id_seq RESTART WITH 6;

CREATE TABLE IF NOT EXISTS roles_permissions
(
    role_id       BIGINT REFERENCES roles (role_id) ON DELETE CASCADE             NOT NULL,
    permission_id BIGINT REFERENCES permissions (permission_id) ON DELETE CASCADE NOT NULL,
    UNIQUE (role_id, permission_id)
);

INSERT INTO roles_permissions (permission_id, role_id)
VALUES (3, 1), -- admins: edit
       (4, 1), -- admins: delete
       (5, 1), -- admins: administrate (change roles)
       (6, 1), -- admins: write
       (1, 2), -- banned: -read
       (2, 2), -- banned: -write
       (2, 3), -- readonly: -write
       (3, 4), -- editors: edit
       (4, 4), -- editors: delete
       (6, 5);
-- members: write

-- tags --
CREATE TABLE IF NOT EXISTS tags
(
    tag_id BIGSERIAL PRIMARY KEY,
    name   TEXT UNIQUE NOT NULL
);

--  groups --
CREATE TABLE IF NOT EXISTS groups
(
    group_id    BIGSERIAL PRIMARY KEY,
    name        TEXT                              NOT NULL,
    description TEXT                              NOT NULL,
    slug        TEXT UNIQUE                       NOT NULL,
    user_id     BIGINT REFERENCES users (user_id) NOT NULL,

    created_at  TIMESTAMP                         NOT NULL,
    updated_at  TIMESTAMP                         NOT NULL,
    deleted_at  TIMESTAMP
);

CREATE TABLE IF NOT EXISTS groups_tags
(
    group_id BIGINT REFERENCES groups (group_id) NOT NULL,
    tag_id   BIGINT REFERENCES tags (tag_id)     NOT NULL,
    UNIQUE (group_id, tag_id)
);

CREATE TABLE IF NOT EXISTS groups_users_roles
(
    user_id    BIGINT REFERENCES users (user_id) ON DELETE CASCADE NOT NULL,
    group_id   BIGINT REFERENCES groups (group_id) ON DELETE CASCADE,
    role_id    BIGINT REFERENCES roles (role_id) ON DELETE CASCADE NOT NULL,
    expires_at TIMESTAMP,
    UNIQUE (user_id, group_id)
);

INSERT INTO groups_users_roles (user_id, group_id, role_id)
VALUES (1, NULL, 1);


-- posts --
CREATE TABLE IF NOT EXISTS posts
(
    post_id    BIGSERIAL PRIMARY KEY,
    name       TEXT                                NOT NULL,
    text       TEXT                                NOT NULL,
    user_id    BIGINT REFERENCES users (user_id)   NOT NULL,
    group_id   BIGINT REFERENCES groups (group_id),

    created_at TIMESTAMP                           NOT NULL,
    updated_at TIMESTAMP                           NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS posts_tags
(
    post_id BIGINT REFERENCES posts (post_id) NOT NULL,
    tag_id  BIGINT REFERENCES tags (tag_id)   NOT NULL,
    UNIQUE (post_id, tag_id)
);

--  subscriptions --
CREATE TABLE IF NOT EXISTS groups_subs
(
    user_id  BIGINT REFERENCES users (user_id)   NOT NULL,
    group_id BIGINT REFERENCES groups (group_id) NOT NULL,
    UNIQUE (user_id, group_id)
);

CREATE TABLE IF NOT EXISTS users_subs
(
    user_id    BIGINT REFERENCES users (user_id) NOT NULL,
    to_user_id BIGINT REFERENCES users (user_id) NOT NULL,
    UNIQUE (user_id, to_user_id)
);

CREATE TABLE IF NOT EXISTS tags_subs
(
    user_id BIGINT REFERENCES users (user_id) NOT NULL,
    tag_id  BIGINT REFERENCES tags (tag_id)   NOT NULL,
    UNIQUE (user_id, tag_id)
);