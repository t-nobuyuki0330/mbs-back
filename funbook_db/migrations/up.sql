CREATE TABLE IF NOT EXISTS cache (
    id SERIAL PRIMARY KEY,
    req_lang VARCHAR(32) NOT NULL,
    req_func VARCHAR(128) NOT NULL,
    req_resp VARCHAR(128) NOT NULL,
    req_count INT NOT NULL,
    ans_json JSON NOT NULL
);