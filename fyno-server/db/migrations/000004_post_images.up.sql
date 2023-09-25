-- INSERT INTO post_images (url, rank, post_id) VALUES ($1, $2, $3)
CREATE TABLE post_images (
    id uuid PRIMARY KEY,
    url TEXT NOT NULL,
    rank INTEGER NOT NULL,
    post_id uuid NOT NULL REFERENCES posts(id)
);