CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    age INTEGER NOT NULL CHECK ( age >= 18 AND age <= 100),
    gender VARCHAR(255) NOT NULL,
    interests TEXT[],
    status VARCHAR(255),
    bio TEXT,
    country VARCHAR(100),
    city VARCHAR(100),
    likedUsers INTEGER[],
    yourLikes INTEGER[],
    yourPartners INTEGER[]
);

CREATE TABLE IF NOT EXISTS user_photo
(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    photo VARCHAR(255),
    is_avatar BOOLEAN,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);