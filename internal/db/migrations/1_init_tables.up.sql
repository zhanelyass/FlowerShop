-- Таблица пользователей
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT NOT NULL UNIQUE,
                       password TEXT NOT NULL,
                       role TEXT DEFAULT 'user'
);

-- Таблица цветов
CREATE TABLE flowers (
                         id SERIAL PRIMARY KEY,
                         name TEXT NOT NULL,
                         description TEXT,
                         price NUMERIC(10, 2) NOT NULL,
                         user_id INTEGER NOT NULL,
                         FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
