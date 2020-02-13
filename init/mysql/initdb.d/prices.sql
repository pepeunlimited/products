CREATE DATABASE IF NOT EXISTS products CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE products;

#  --------
# |products|
#  --------
CREATE TABLE products (
    id                  BIGINT      NOT NULL AUTO_INCREMENT,
    type                CHAR(40)    UNIQUE NOT NULL,
    list_of_places_id   BIGINT      UNIQUE NULL,
    PRIMARY KEY (id)
);
#  -------------
# |subscriptions|
#  -------------
CREATE TABLE plans (
   id            BIGINT      NOT NULL AUTO_INCREMENT,
   title_i18n_id BIGINT      NOT NULL,
   length        TINYINT     UNSIGNED NOT NULL,
   unit          CHAR(7)     NOT NULL, #days, weeks, months
   PRIMARY KEY (id)
);
# NOTE: not possible set same subscriptions at same time sequence
CREATE TABLE subscriptions (
   id                    BIGINT      NOT NULL AUTO_INCREMENT,
   user_id               BIGINT      NOT NULL,
   plan_subscriptions    BIGINT,
   start_at              DATETIME(3) NOT NULL,
   end_at                DATETIME(3) NOT NULL,
   FOREIGN KEY (plan_subscriptions) REFERENCES plans (id),
   PRIMARY KEY (id)
);
#  ------
# |prices|
#  ------
CREATE TABLE in_app_purchases (
    id                          TINYINT     NOT NULL AUTO_INCREMENT,
    apple_sku                   CHAR(7)     UNIQUE NULL,
    google_sku                  CHAR(7)     UNIQUE NULL,
    start_at                    DATETIME(3) NOT NULL,
    end_at                      DATETIME(3) NOT NULL,
    PRIMARY KEY (id)
);
# NOTE: not possible set multiple prices at same time sequence
CREATE TABLE prices (
    id                                       BIGINT      NOT NULL AUTO_INCREMENT,
    start_at                                 DATETIME(3) NOT NULL,
    end_at                                   DATETIME(3) NOT NULL,
    price                                    MEDIUMINT   UNSIGNED DEFAULT 0,
    discount                                 MEDIUMINT   UNSIGNED DEFAULT 0,
    product_id                               BIGINT      NULL, # actual product (required)
    plans_id                                 BIGINT      NULL, # if the price has option to subscriptions plans
    in_app_purchases_id                      TINYINT     NULL, # if the price has option for additional details of the in app purchases
    FOREIGN KEY (plans_id)                   REFERENCES  plans (id),
    FOREIGN KEY (product_id)                 REFERENCES  products (id),
    FOREIGN KEY (in_app_purchases_id)        REFERENCES  in_app_purchases (id),
    PRIMARY KEY (id)
);