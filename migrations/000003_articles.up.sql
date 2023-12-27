CREATE TABLE "articles"(
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "body" TEXT,
    "author_id" INT REFERENCES users(id),
    "category_id" INT REFERENCES categories(id),
    "is_draft" BOOLEAN,
    "published_at" TIMESTAMP,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
