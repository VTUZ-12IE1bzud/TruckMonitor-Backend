-- Schema
DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA public;

-- Вид ТС
CREATE TABLE track_type(
  id SERIAL PRIMARY KEY,
	name TEXT NOT NULL);

-- Марка ТС
CREATE TABLE track_brand(
  id SERIAL PRIMARY KEY,
  track_type_id INTEGER NOT NULL REFERENCES track_type,
	name TEXT NOT NULL);

-- Модель ТС
CREATE TABLE track_model(
  id SERIAL PRIMARY KEY,
  track_brand_id INTEGER NOT NULL REFERENCES track_brand,
	name TEXT NOT NULL,
	max_capacity REAL NOT NULL);

-- ТС
CREATE TABLE track(
  id SERIAL PRIMARY KEY,
  track_model_id INTEGER NOT NULL REFERENCES track_model,
	vin TEXT NOT NULL,
	licence_plate TEXT NOT NULL);

-- Роль пользователя
CREATE TYPE role AS ENUM ('admin', 'accountant', 'manager', 'driver', 'client');

-- Сотрудник
CREATE TABLE account(
  id SERIAL PRIMARY KEY,
  role role NOT NULL,
  surmane TEXT NOT NULL,
  name TEXT NOT NULL,
  patronymic TEXT,
  date_of_birth DATE NOT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  photo TEXT NOT NULL,
  phone VARCHAR(11) NOT NULL);

  -- Упаковка
CREATE TABLE packaging(
  id SERIAL PRIMARY KEY,
	name TEXT NOT NULL);

  -- Банк
CREATE TABLE bank(
  id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	itn VARCHAR(12) NOT NULL,
	bic VARCHAR(9) NOT NULL,
	address TEXT NOT NULL,
	сorrespondent_account VARCHAR(20) NOT NULL);

-- Контрагент
CREATE TABLE client(
  id SERIAL PRIMARY KEY,
  role role NOT NULL,
  bank_id INTEGER NOT NULL REFERENCES bank,
	name TEXT NOT NULL,
	itn VARCHAR(12) NOT NULL,
	iec VARCHAR(12) NOT NULL,
	address TEXT NOT NULL,
	current_account VARCHAR(20) NOT NULL,
	chief TEXT NOT NULL,
	accountant TEXT NOT NULL,
	email TEXT NOT NULL,
	password TEXT NOT NULL);

-- Договор
CREATE TABLE contract(
  id SERIAL PRIMARY KEY,
  account_id INTEGER NOT NULL REFERENCES account,
  client_id INTEGER NOT NULL REFERENCES client,
  packaging_id INTEGER NOT NULL REFERENCES packaging,
  number TEXT NOT NULL,
  price MONEY NOT NULL,
  price_check TEXT,
  customs_check TEXT,
  departure TEXT NOT NULL,
  destination TEXT NOT NULL,
  date DATE NOT NULL);

  -- Грузоперевозка
CREATE TABLE freight_traffic(
  id SERIAL PRIMARY KEY,
  track_id INTEGER NOT NULL REFERENCES track,
  driver_id INTEGER NOT NULL REFERENCES account);

-- График движения
CREATE TABLE track_route(
  id SERIAL PRIMARY KEY,
  freight_traffic_id INTEGER NOT NULL REFERENCES freight_traffic,
	coordinate POINT NOT NULL,
	timestamp_planned TIMESTAMP NOT NULL,
	timestamp_actual TIMESTAMP);

-- Состав груза
CREATE TABLE сargo_composition(
  id SERIAL PRIMARY KEY,
  freight_traffic_id INTEGER NOT NULL REFERENCES freight_traffic,
  contract_id INTEGER NOT NULL REFERENCES contract);

-- Единица измерения
CREATE TABLE measure(
  id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	abbreviation TEXT NOT NULL);

-- Классификатор опасности груза
CREATE TABLE hazard(
  id SERIAL PRIMARY KEY,
	name TEXT NOT NULL);

-- Номенклатура
CREATE TABLE nomenclature(
  id SERIAL PRIMARY KEY,
  measure_id INTEGER NOT NULL REFERENCES measure,
  hazard_id INTEGER NOT NULL REFERENCES hazard,
	name TEXT NOT NULL);

-- Предмет договора
CREATE TABLE contract_subject(
  id SERIAL PRIMARY KEY,
  contract_id INTEGER NOT NULL REFERENCES contract,
  nomenclature_id INTEGER NOT NULL REFERENCES nomenclature,
  hazard_id INTEGER NOT NULL REFERENCES hazard,
	amount REAL NOT NULL,
	price REAL NOT NULL);