CREATE TABLE "follows"(
    following_id INTEGER NOT NULL,
    followed_by_id INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (following_id) REFERENCES users(id),
    FOREIGN KEY (followed_by_id) REFERENCES users(id)
);
