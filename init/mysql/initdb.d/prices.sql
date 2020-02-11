CREATE DATABASE IF NOT EXISTS prices CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE prices;

#  ------
# |prices|
#  ------
CREATE TABLE google_iap (
    sku        CHAR(100) NOT NULL,
    created_at DATETIME(3) NOT NULL,
    PRIMARY KEY (sku)
);
CREATE TABLE apple_iap (
    sku        CHAR(100) NOT NULL,
    created_at DATETIME(3) NOT NULL,
    PRIMARY KEY (sku)
);
CREATE TABLE prices (
    id                BIGINT NOT NULL AUTO_INCREMENT,
    google_iap_prices CHAR(100) NULL,
    apple_iap_prices  CHAR(100) NULL,
    start_at          DATETIME(3) NOT NULL,
    end_at            DATETIME(3) NOT NULL,
    cost              MEDIUMINT UNSIGNED DEFAULT 0,
    discount          MEDIUMINT UNSIGNED DEFAULT 0,
    FOREIGN KEY (google_iap_prices) REFERENCES google_iap (sku),
    FOREIGN KEY (apple_iap_prices) REFERENCES apple_iap (sku),
    PRIMARY KEY (id)
);