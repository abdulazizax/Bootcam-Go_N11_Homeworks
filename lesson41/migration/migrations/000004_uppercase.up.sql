UPDATE users
SET username = UPPER(SUBSTRING(username, 1, 1)) || LOWER(SUBSTRING(username, 2));