CREATE DATABASE IF NOT EXISTS products CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE products;

#  --------
# |products|
#  --------
CREATE TABLE products (
    id                  BIGINT      NOT NULL AUTO_INCREMENT,
    sku                 CHAR(32)    UNIQUE NOT NULL,
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
# NOTE: not possible set same subscription's time sequence for plansId
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
CREATE TABLE iap_source (
    id                          TINYINT     NOT NULL AUTO_INCREMENT,
    in_app_purchase_sku         CHAR(32)    UNIQUE NOT NULL,
    google_billing_service_sku  CHAR(32)    UNIQUE NULL, # for the future
    start_at                    DATETIME(3) NOT NULL,
    end_at                      DATETIME(3) NOT NULL,
    PRIMARY KEY (id)
);
# NOTE: not possible set multiple price's time sequence for productId
CREATE TABLE prices (
    id                                       BIGINT      NOT NULL AUTO_INCREMENT,
    start_at                                 DATETIME(3) NOT NULL,
    end_at                                   DATETIME(3) NOT NULL,
    price                                    MEDIUMINT   UNSIGNED DEFAULT 0,
    discount                                 MEDIUMINT   UNSIGNED DEFAULT 0,
    product_prices                           BIGINT      NOT NULL, # actual product
    plan_prices                              BIGINT      NULL, # if the price has option to subscriptions plans
    iap_source_prices                        TINYINT     NULL, # if the price has option for additional details of the in app purchases
    FOREIGN KEY (plan_prices)                REFERENCES  plans (id),
    FOREIGN KEY (product_prices)             REFERENCES  products (id),
    FOREIGN KEY (iap_source_prices)          REFERENCES  iap_source (id),
    PRIMARY KEY (id)
);