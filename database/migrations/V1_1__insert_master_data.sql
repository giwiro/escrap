INSERT INTO escrap.scrap_provider (scrap_provider_id, name)
VALUES (1, 'Amazon'),
       (2, 'Ebay');

INSERT INTO escrap.scrap_result_state (scrap_result_state_id, name)
VALUES (1, 'SUCCESS'),
       (2, 'ERROR');

INSERT INTO escrap.scrap_provider_domain (scrap_provider_id, domain)
VALUES (1, 'www.amazon.com');
