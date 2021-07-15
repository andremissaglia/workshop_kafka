-- Configuração para facilitar o desenvolvimento
SET 'auto.offset.reset' = 'earliest';
-- Criação do stream de entrada
CREATE STREAM votes (cat STRING KEY, vote STRING) WITH (
    KAFKA_TOPIC = 'votes',
    WRAP_SINGLE_VALUE = false,
    KEY_FORMAT = 'DELIMITED',
    VALUE_FORMAT = 'DELIMITED'
);

-- Criação de um stream intermediário
CREATE OR REPLACE STREAM parsed_votes
WITH (VALUE_FORMAT = 'JSON') AS
SELECT cat,
    CAST(vote AS INT) vote
FROM votes;

-- Debug: Ler o que tem no stream
SELECT *
FROM parsed_votes EMIT CHANGES;


-- Criação da tabela de médias
CREATE OR REPLACE TABLE ratings WITH (
    KAFKA_TOPIC = 'ratings-ksqldb'
) AS
SELECT cat,
    AVG(vote) rating
FROM parsed_votes
GROUP BY cat EMIT CHANGES;

-- Debug: Ler o que tem na tabela
SELECT *
FROM ratings EMIT CHANGES;