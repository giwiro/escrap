CREATE FUNCTION escrap.update_modified_column() RETURNS TRIGGER
    LANGUAGE PLPGSQL
AS
$$
BEGIN
    IF ROW (NEW.*) IS DISTINCT FROM ROW (OLD.*) THEN
        NEW.updated_at = NOW();
        RETURN NEW;
    ELSE
        RETURN OLD;
    END IF;
END;
$$;

ALTER FUNCTION escrap.update_modified_column() OWNER TO escrap_admin;

------------- Provider
CREATE TABLE escrap.scrap_provider
(
    scrap_provider_id BIGINT PRIMARY KEY,
    name              VARCHAR NOT NULL
);

ALTER TABLE escrap.scrap_provider
    OWNER TO escrap_admin;

GRANT SELECT, DELETE, INSERT, UPDATE ON escrap.scrap_provider TO escrap_admin;


------------- Provider Domain
CREATE TABLE escrap.scrap_provider_domain
(
    scrap_provider_domain_id BIGSERIAL PRIMARY KEY,
    scrap_provider_id        BIGINT REFERENCES escrap.scrap_provider,
    domain                   VARCHAR NOT NULL
);

CREATE UNIQUE INDEX uq_domain
    ON escrap.scrap_provider_domain (domain);

ALTER TABLE escrap.scrap_provider_domain
    OWNER TO escrap_admin;

GRANT SELECT, DELETE, INSERT, UPDATE ON escrap.scrap_provider_domain TO escrap_admin;

------------- Product
CREATE TABLE escrap.scrap_product
(
    scrap_product_id  BIGSERIAL PRIMARY KEY,
    scrap_provider_id BIGINT REFERENCES escrap.scrap_provider,
    name              VARCHAR                                NOT NULL,
    price             NUMERIC(12, 2),
    vendor_id         VARCHAR                                NOT NULL,
    description       VARCHAR,
    image_url         VARCHAR,
    last_scrapped_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at        TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL

);

CREATE UNIQUE INDEX uq_vendor_id
    ON escrap.scrap_product (vendor_id);

ALTER TABLE escrap.scrap_product
    OWNER TO escrap_admin;

GRANT SELECT, DELETE, INSERT, UPDATE ON escrap.scrap_product TO escrap_admin;

------------- Scrap Result State
CREATE TABLE escrap.scrap_result_state
(
    scrap_result_state_id BIGINT PRIMARY KEY,
    name                  VARCHAR NOT NULL
);

ALTER TABLE escrap.scrap_result_state
    OWNER TO escrap_admin;

GRANT SELECT, DELETE, INSERT, UPDATE ON escrap.scrap_result_state TO escrap_admin;

------------- Scrap Result
CREATE TABLE escrap.scrap_result
(
    scrap_result_id       BIGSERIAL PRIMARY KEY,
    scrap_product_id      BIGINT REFERENCES escrap.scrap_product,
    scrap_result_state_id BIGINT REFERENCES escrap.scrap_result_state,
    api_result            JSONB,
    created_at            TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

ALTER TABLE escrap.scrap_result
    OWNER TO escrap_admin;

GRANT SELECT, DELETE, INSERT, UPDATE ON escrap.scrap_result TO escrap_admin;