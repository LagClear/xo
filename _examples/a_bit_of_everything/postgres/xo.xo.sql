-- Generated by xo for the public schema.

-- enum a_enum
CREATE TYPE a_enum AS ENUM (
  'ONE',
  'TWO'
);

-- table a_bit_of_everything
CREATE TABLE a_bit_of_everything (
  a_enum A_ENUM NOT NULL,
  a_enum_nullable A_ENUM,
  a_bigint BIGINT NOT NULL,
  a_bigint_nullable BIGINT,
  a_bigserial BIGSERIAL,
  a_bigserial_nullable BIGSERIAL,
  a_bit BIT(1) NOT NULL,
  a_bit_nullable BIT(1),
  a_bit_varying BIT VARYING NOT NULL,
  a_bit_varying_nullable BIT VARYING,
  a_bool BOOLEAN NOT NULL,
  a_bool_nullable BOOLEAN,
  a_boolean BOOLEAN NOT NULL,
  a_boolean_nullable BOOLEAN,
  a_bpchar BPCHAR NOT NULL,
  a_bpchar_nullable BPCHAR,
  a_bytea BYTEA NOT NULL,
  a_bytea_nullable BYTEA,
  a_char CHAR(1) NOT NULL,
  a_char_nullable CHAR(1),
  a_character CHAR(1) NOT NULL,
  a_character_nullable CHAR(1),
  a_character_varying VARCHAR NOT NULL,
  a_character_varying_nullable VARCHAR,
  a_date DATE NOT NULL,
  a_date_nullable DATE,
  a_decimal NUMERIC NOT NULL,
  a_decimal_nullable NUMERIC,
  a_double_precision DOUBLE PRECISION NOT NULL,
  a_double_precision_nullable DOUBLE PRECISION,
  a_inet INET NOT NULL,
  a_inet_nullable INET,
  a_int INTEGER NOT NULL,
  a_int_nullable INTEGER,
  a_integer INTEGER NOT NULL,
  a_integer_nullable INTEGER,
  a_interval INTERVAL NOT NULL,
  a_interval_nullable INTERVAL,
  a_json JSON NOT NULL,
  a_json_nullable JSON,
  a_jsonb JSONB NOT NULL,
  a_jsonb_nullable JSONB,
  a_money MONEY NOT NULL,
  a_money_nullable MONEY,
  a_numeric NUMERIC NOT NULL,
  a_numeric_nullable NUMERIC,
  a_real REAL NOT NULL,
  a_real_nullable REAL,
  a_serial SERIAL,
  a_serial_nullable SERIAL,
  a_smallint SMALLINT NOT NULL,
  a_smallint_nullable SMALLINT,
  a_smallserial SMALLSERIAL,
  a_smallserial_nullable SMALLSERIAL,
  a_text TEXT NOT NULL,
  a_text_nullable TEXT,
  a_time TIME NOT NULL,
  a_time_nullable TIME,
  a_timestamp TIMESTAMP NOT NULL,
  a_timestamp_nullable TIMESTAMP,
  a_timestamptz TIMESTAMPTZ NOT NULL,
  a_timestamptz_nullable TIMESTAMPTZ,
  a_timetz TIMETZ NOT NULL,
  a_timetz_nullable TIMETZ,
  a_uuid UUID NOT NULL,
  a_uuid_nullable UUID,
  a_varchar VARCHAR NOT NULL,
  a_varchar_nullable VARCHAR,
  a_xml XML NOT NULL,
  a_xml_nullable XML
);

-- table a_foreign_key
CREATE TABLE a_foreign_key (
  a_key INTEGER REFERENCES a_primary (a_key)
);

-- table a_foreign_key_composite
CREATE TABLE a_foreign_key_composite (
  a_key1 INTEGER,
  a_key2 INTEGER,
  FOREIGN KEY (a_key1, a_key2) REFERENCES a_primary_composite (a_key1, a_key2)
);

-- table a_index
CREATE TABLE a_index (
  a_key INTEGER
);

-- index a_index_idx
CREATE INDEX a_index_idx ON a_index (a_key);

-- table a_index_composite
CREATE TABLE a_index_composite (
  a_key1 INTEGER,
  a_key2 INTEGER
);

-- index a_index_composite_idx
CREATE INDEX a_index_composite_idx ON a_index_composite (a_key1, a_key2);

-- table a_manual_table
CREATE TABLE a_manual_table (
  a_text VARCHAR(255)
);

-- table a_primary
CREATE TABLE a_primary (
  a_key INTEGER NOT NULL,
  PRIMARY KEY (a_key)
);

-- table a_primary_composite
CREATE TABLE a_primary_composite (
  a_key1 INTEGER NOT NULL,
  a_key2 INTEGER NOT NULL,
  PRIMARY KEY (a_key1, a_key2)
);

-- table a_primary_multi
CREATE TABLE a_primary_multi (
  a_key INTEGER NOT NULL,
  a_text VARCHAR(255),
  PRIMARY KEY (a_key)
);

-- table a_sequence
CREATE TABLE a_sequence (
  a_seq SERIAL,
  PRIMARY KEY (a_seq)
);

-- table a_sequence_multi
CREATE TABLE a_sequence_multi (
  a_seq SERIAL,
  a_text VARCHAR(255),
  PRIMARY KEY (a_seq)
);

-- table a_unique_index
CREATE TABLE a_unique_index (
  a_key INTEGER,
  UNIQUE (a_key)
);

-- table a_unique_index_composite
CREATE TABLE a_unique_index_composite (
  a_key1 INTEGER,
  a_key2 INTEGER,
  UNIQUE (a_key1, a_key2)
);

-- view a_view_of_everything
CREATE VIEW a_view_of_everything AS
SELECT a_bit_of_everything.a_enum,
    a_bit_of_everything.a_enum_nullable,
    a_bit_of_everything.a_bigint,
    a_bit_of_everything.a_bigint_nullable,
    a_bit_of_everything.a_bigserial,
    a_bit_of_everything.a_bigserial_nullable,
    a_bit_of_everything.a_bit,
    a_bit_of_everything.a_bit_nullable,
    a_bit_of_everything.a_bit_varying,
    a_bit_of_everything.a_bit_varying_nullable,
    a_bit_of_everything.a_bool,
    a_bit_of_everything.a_bool_nullable,
    a_bit_of_everything.a_boolean,
    a_bit_of_everything.a_boolean_nullable,
    a_bit_of_everything.a_bpchar,
    a_bit_of_everything.a_bpchar_nullable,
    a_bit_of_everything.a_bytea,
    a_bit_of_everything.a_bytea_nullable,
    a_bit_of_everything.a_char,
    a_bit_of_everything.a_char_nullable,
    a_bit_of_everything.a_character,
    a_bit_of_everything.a_character_nullable,
    a_bit_of_everything.a_character_varying,
    a_bit_of_everything.a_character_varying_nullable,
    a_bit_of_everything.a_date,
    a_bit_of_everything.a_date_nullable,
    a_bit_of_everything.a_decimal,
    a_bit_of_everything.a_decimal_nullable,
    a_bit_of_everything.a_double_precision,
    a_bit_of_everything.a_double_precision_nullable,
    a_bit_of_everything.a_inet,
    a_bit_of_everything.a_inet_nullable,
    a_bit_of_everything.a_int,
    a_bit_of_everything.a_int_nullable,
    a_bit_of_everything.a_integer,
    a_bit_of_everything.a_integer_nullable,
    a_bit_of_everything.a_interval,
    a_bit_of_everything.a_interval_nullable,
    a_bit_of_everything.a_json,
    a_bit_of_everything.a_json_nullable,
    a_bit_of_everything.a_jsonb,
    a_bit_of_everything.a_jsonb_nullable,
    a_bit_of_everything.a_money,
    a_bit_of_everything.a_money_nullable,
    a_bit_of_everything.a_numeric,
    a_bit_of_everything.a_numeric_nullable,
    a_bit_of_everything.a_real,
    a_bit_of_everything.a_real_nullable,
    a_bit_of_everything.a_serial,
    a_bit_of_everything.a_serial_nullable,
    a_bit_of_everything.a_smallint,
    a_bit_of_everything.a_smallint_nullable,
    a_bit_of_everything.a_smallserial,
    a_bit_of_everything.a_smallserial_nullable,
    a_bit_of_everything.a_text,
    a_bit_of_everything.a_text_nullable,
    a_bit_of_everything.a_time,
    a_bit_of_everything.a_time_nullable,
    a_bit_of_everything.a_timestamp,
    a_bit_of_everything.a_timestamp_nullable,
    a_bit_of_everything.a_timestamptz,
    a_bit_of_everything.a_timestamptz_nullable,
    a_bit_of_everything.a_timetz,
    a_bit_of_everything.a_timetz_nullable,
    a_bit_of_everything.a_uuid,
    a_bit_of_everything.a_uuid_nullable,
    a_bit_of_everything.a_varchar,
    a_bit_of_everything.a_varchar_nullable,
    a_bit_of_everything.a_xml,
    a_bit_of_everything.a_xml_nullable
   FROM a_bit_of_everything;

-- view a_view_of_everything_some
CREATE VIEW a_view_of_everything_some AS
SELECT a_bit_of_everything.a_bool,
    a_bit_of_everything.a_text
   FROM a_bit_of_everything;

-- procedure a_0_in_0_out
CREATE PROCEDURE a_0_in_0_out() AS $$
BEGIN
END;
$$ LANGUAGE plpgsql;

-- function a_0_in_1_out
CREATE FUNCTION a_0_in_1_out() RETURNS INTEGER AS $$
BEGIN
  a_return := 10;
END;
$$ LANGUAGE plpgsql;

-- procedure a_1_in_0_out
CREATE PROCEDURE a_1_in_0_out(a_param INTEGER, OUT a_param INTEGER) AS $$
BEGIN
END;
$$ LANGUAGE plpgsql;

-- function a_1_in_1_out
CREATE FUNCTION a_1_in_1_out(a_param INTEGER) RETURNS INTEGER AS $$
BEGIN
  a_return := a_param;
END;
$$ LANGUAGE plpgsql;

-- function a_2_in_2_out
CREATE FUNCTION a_2_in_2_out(param_one INTEGER, param_two INTEGER, OUT param_one INTEGER, OUT param_two INTEGER) AS $$
BEGIN
  return_one := param_one;
  return_two := param_two;
END;
$$ LANGUAGE plpgsql;

-- function a_func_0_in
CREATE FUNCTION a_func_0_in() RETURNS INTEGER AS $$
BEGIN
  RETURN 10;
END;
$$ LANGUAGE plpgsql;

-- function a_func_1_in
CREATE FUNCTION a_func_1_in(a_param INTEGER) RETURNS INTEGER AS $$
BEGIN
  RETURN a_param;
END;
$$ LANGUAGE plpgsql;

-- function a_func_2_in
CREATE FUNCTION a_func_2_in(param_one INTEGER, param_two INTEGER) RETURNS INTEGER AS $$
BEGIN
  RETURN param_one + param_two;
END;
$$ LANGUAGE plpgsql;

-- procedure a_overloaded
CREATE PROCEDURE a_overloaded(param_one INTEGER, OUT param_one INTEGER) AS $$
BEGIN
END;
$$ LANGUAGE plpgsql;

-- procedure a_overloaded
CREATE PROCEDURE a_overloaded(param_one INTEGER, param_two INTEGER, OUT param_one INTEGER, OUT param_two INTEGER) AS $$
BEGIN
END;
$$ LANGUAGE plpgsql;
