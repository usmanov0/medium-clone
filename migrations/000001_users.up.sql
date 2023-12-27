CREATE TYPE "user_roles" AS ENUM(
    'author',
    'editor',
    'admin'
);


CREATE TABLE "users"(
    "id" SERIAL PRIMARY KEY,
    "user_name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(55) NOT NULL,
    "password" VARCHAR(55) NOT NULL,
    "role" user_roles,
    "bio" VARCHAR(255) NOT NULL,
    "date_of_birth" DATE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


--    Author: Users who can create and publish articles.
--    Editor: Users who can review and edit articles, possibly with additional privileges.
--    Admin: Users with administrative rights, allowing them to manage users, settings, and the overall system.
