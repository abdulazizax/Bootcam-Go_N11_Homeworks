CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE books (
    book_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year_published INT NOT NULL
);

CREATE TABLE rent (
    rental_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    book_id UUID NOT NULL,
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);
