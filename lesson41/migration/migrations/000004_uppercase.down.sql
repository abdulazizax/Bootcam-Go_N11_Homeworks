UPDATE users
SET username = LOWER(SUBSTRING(username, 1, 1)) || SUBSTRING(username, 2);
