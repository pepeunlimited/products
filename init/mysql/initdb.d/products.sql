CREATE DATABASE IF NOT EXISTS products CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE products;

#  --------
# |products|
#  --------
CREATE TABLE products (
    id          BIGINT      NOT NULL AUTO_INCREMENT,
    sku         CHAR(32)    UNIQUE NOT NULL,
    PRIMARY KEY (id)
);
#  ------------------
# |third_party_prices|
#  ------------------
CREATE TABLE third_party_prices (
    id                          BIGINT      NOT NULL AUTO_INCREMENT,
    in_app_purchase_sku         CHAR(32)    UNIQUE NOT NULL,
    google_billing_service_sku  CHAR(32)    UNIQUE NULL, # for the future
    start_at                    DATETIME(3) NOT NULL,
    end_at                      DATETIME(3) NOT NULL,
    type                        CHAR(28)    NOT NULL, # CONSUMABLE, NON_CONSUMABLE, AUTO_RENEWABLE_SUBSCRIPTIONS, NON_RENEWING_SUBSCRIPTIONS
    PRIMARY KEY (id)
);
#  -------------
# |subscriptions|
#  -------------
CREATE TABLE plans (
   id                BIGINT      NOT NULL AUTO_INCREMENT,
   title_i18n_id     BIGINT      NOT NULL,
   start_at          DATETIME(3) NOT NULL,
   end_at            DATETIME(3) NOT NULL,
   price             MEDIUMINT   UNSIGNED DEFAULT 0,
   discount          MEDIUMINT   UNSIGNED DEFAULT 0,
   length            TINYINT     UNSIGNED NOT NULL,
   unit              CHAR(7)     NOT NULL, # days, weeks, months
   product_plans     BIGINT      NOT NULL, # product whose pricing this plan determines.
   third_party_price_plans      BIGINT     NULL,     # if plans has option to use in-app-purchase
   FOREIGN KEY (product_plans)  REFERENCES  products (id),
   FOREIGN KEY (third_party_price_plans)   REFERENCES  third_party_prices (id),
   PRIMARY KEY (id)
);
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
CREATE TABLE prices (
    id                           BIGINT      NOT NULL AUTO_INCREMENT,
    start_at                     DATETIME(3) NOT NULL,
    end_at                       DATETIME(3) NOT NULL,
    price                        MEDIUMINT   UNSIGNED DEFAULT 0,
    discount                     MEDIUMINT   UNSIGNED DEFAULT 0,
    product_prices               BIGINT      NOT NULL, # product whose pricing this determines.
    third_party_price_prices     BIGINT     NULL,     # if price has option to use in-app-purchase
    FOREIGN KEY (third_party_price_prices)  REFERENCES  third_party_prices (id),
    FOREIGN KEY (product_prices) REFERENCES  products (id),
    PRIMARY KEY (id)
);