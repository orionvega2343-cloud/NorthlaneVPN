CREATE TABLE IF NOT EXISTS subscription_servers (
    -- Соединение первичных ключей таблицы подписок и серверов
    subscription_id INT NOT NULL,
    server_id INT NOT NULL,
    -- создание составного первичного ключа
    PRIMARY KEY(subscription_id, server_id),
    -- связывание с таблицей подписок
    CONSTRAINT fk_subscriptions
        FOREIGN KEY (subscription_id)
        REFERENCES subscriptions(id)
        ON DELETE CASCADE,

    -- связывание с таблицей серверов
    CONSTRAINT fk_servers
        FOREIGN KEY (server_id)
        REFERENCES servers(id)
        ON DELETE CASCADE
)