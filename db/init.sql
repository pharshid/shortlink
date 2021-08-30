CREATE TABLE IF NOT EXISTS shortlink (
  id INT AUTO_INCREMENT PRIMARY KEY,
  token VARCHAR(32) NOT NULL UNIQUE,
  target_url VARCHAR(500) NOT NULL,
  created_at DATE NOT NULL,
  visits INT NOT NULL
);