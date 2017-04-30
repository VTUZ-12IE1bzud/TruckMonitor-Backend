-- Schema
DROP SCHEMA IF EXISTS public CASCADE;
CREATE SCHEMA public;

-- Роль пользователя
CREATE TYPE ROLE AS ENUM ('admin', 'accountant', 'manager', 'driver');

-- Статус грузоперевозки
CREATE TYPE CARRIAGE_STATUS AS ENUM ('current', 'archive', 'onward');

-- Вид ТС
CREATE TABLE vehicle_type(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL);

-- Марка ТС
CREATE TABLE vehicle_brand(
    id SERIAL PRIMARY KEY,
    type_id INTEGER NOT NULL REFERENCES vehicle_type,
    name TEXT NOT NULL);

-- Модель ТС
CREATE TABLE vehicle_model(
    id SERIAL PRIMARY KEY,
    brand_id INTEGER NOT NULL REFERENCES vehicle_brand,
    name TEXT NOT NULL,
    capacity_max REAL NOT NULL);

-- ТС
CREATE TABLE vehicle(
    id SERIAL PRIMARY KEY,
    model_id INTEGER NOT NULL REFERENCES vehicle_model,
    vin TEXT UNIQUE NOT NULL,
    licence_plate TEXT NOT NULL,
    region_number INTEGER NOT NULL);

-- Упаковка
CREATE TABLE packaging(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL);

-- Сотрудник
CREATE TABLE employee(
    id SERIAL PRIMARY KEY,
    role ROLE NOT NULL,
    surmane TEXT NOT NULL,
    name TEXT NOT NULL,
    patronymic TEXT,
    birth_date DATE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    photo TEXT NOT NULL,
    phone VARCHAR(20) NOT NULL);

-- Контрагент
CREATE TABLE client(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    itn VARCHAR(12) NOT NULL,
    iec VARCHAR(12) NOT NULL,
    address TEXT NOT NULL);

-- Склад
CREATE TABLE store(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    latitude REAL NOT NULL,
    longitude REAL NOT NULL);

-- Договор
CREATE TABLE contract(
    id SERIAL PRIMARY KEY,
    manager_id INTEGER NOT NULL REFERENCES employee,
    client_id INTEGER NOT NULL REFERENCES client,
    packaging_id INTEGER NOT NULL REFERENCES packaging,
    form_store_id INTEGER NOT NULL REFERENCES store,
    before_store_id INTEGER NOT NULL REFERENCES store,
    number TEXT UNIQUE NOT NULL,
    price REAL NOT NULL,
    confirmation_payment_link TEXT,
    confirmation_customs_link TEXT,
    date_shipment DATE NOT NULL);

-- Грузоперевозка
CREATE TABLE carriage(
    id SERIAL PRIMARY KEY,
    status CARRIAGE_STATUS NOT NULL,
    vehicle_id INTEGER NOT NULL REFERENCES vehicle,
    driver_id INTEGER NOT NULL REFERENCES employee);

-- Состав груза
CREATE TABLE carriage_detail(
    id SERIAL PRIMARY KEY,
    carriage_id INTEGER NOT NULL REFERENCES carriage,
    contract_id INTEGER NOT NULL REFERENCES contract);

-- Единица измерения
CREATE TABLE measure(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    abbreviation TEXT NOT NULL);

-- Номенклатура
CREATE TABLE nomenclature(
    id SERIAL PRIMARY KEY,
    measure_id INTEGER NOT NULL REFERENCES measure,
    name TEXT NOT NULL);

  -- Предмет договора
CREATE TABLE contract_detail(
    id SERIAL PRIMARY KEY,
    contract_id INTEGER NOT NULL REFERENCES contract,
    nomenclature_id INTEGER NOT NULL REFERENCES nomenclature,
    amount REAL NOT NULL,
    price REAL NOT NULL);

-- Контрольная точка
CREATE TABLE check_point(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    latitude REAL NOT NULL,
    longitude REAL NOT NULL);

-- График движения
CREATE TABLE carriage_route(
    id SERIAL PRIMARY KEY,
    carriage_id INTEGER NOT NULL REFERENCES carriage,
    check_point_id INTEGER NOT NULL REFERENCES check_point,
    timestamp_planned TIMESTAMP NOT NULL,
    timestamp_fact TIMESTAMP);

INSERT INTO vehicle_type (id, name) VALUES (1, 'Грузовой автомобиль общего назначения');
INSERT INTO vehicle_type (id, name) VALUES (2, 'Специальный грузовой автомобиль');

INSERT INTO vehicle_brand (id, type_id, name) VALUES (1, 1, 'Volvo');
INSERT INTO vehicle_brand (id, type_id, name) VALUES (2, 1, 'Mercedes-Benz');
INSERT INTO vehicle_brand (id, type_id, name) VALUES (3, 1, 'Man');

INSERT INTO vehicle_model (id, brand_id, name, capacity_max) VALUES (1, 1, 'FH', 18000);
INSERT INTO vehicle_model (id, brand_id, name, capacity_max) VALUES (2, 1, 'FE', 26000);
INSERT INTO vehicle_model (id, brand_id, name, capacity_max) VALUES (3, 2, 'Atego', 7000);
INSERT INTO vehicle_model (id, brand_id, name, capacity_max) VALUES (4, 2, 'Axor', 26000);
INSERT INTO vehicle_model (id, brand_id, name, capacity_max) VALUES (5, 2, 'Actros', 50000);
INSERT INTO vehicle_model (id, brand_id, name, capacity_max) VALUES (6, 3, 'TXG', 41000);

INSERT INTO vehicle (id, model_id, vin, licence_plate, region_number)
    VALUES (1, 1, '1FDYU90L9KVA08444', 'В444МВ', 33);
INSERT INTO vehicle (id, model_id, vin, licence_plate, region_number)
    VALUES (2, 2, '3HSCHAPR8AN202147', 'Х777АМ', 72);
INSERT INTO vehicle (id, model_id, vin, licence_plate, region_number)
    VALUES (3, 3, 'WAUHFAFL1AA754008', 'С100СС', 45);
INSERT INTO vehicle (id, model_id, vin, licence_plate, region_number)
    VALUES (4, 4, 'JALE5B1U0R3098851', 'В595ОР', 777);
INSERT INTO vehicle (id, model_id, vin, licence_plate, region_number)
    VALUES (5, 5, '1HD1GM41XAC396026', 'Н777ВВ', 35);
INSERT INTO vehicle (id, model_id, vin, licence_plate, region_number)
    VALUES (6, 6, '1GDS7H4J2TJ535301', 'В567АО', 36);

INSERT INTO packaging (id, name) VALUES (1, 'Контейнер');
INSERT INTO packaging (id, name) VALUES (2, 'Коробка');

INSERT INTO employee (id, role, surmane, name, patronymic, birth_date, email, password, photo, phone)
    VALUES (1, 'admin', 'Аннин', 'Павел', 'Александрович', '1993-11-19', 'annin@truck.ru', md5('annin'), 'https://randus.ru/avatars/m/myAvatar19.png', '8(800)000-00-00');
INSERT INTO employee (id, role, surmane, name, patronymic, birth_date, email, password, photo, phone)
    VALUES (2, 'accountant', 'Сморчков', 'Мирослав', 'Борисович', '1975-07-4', 'smorchkov@truck.ru', md5('smorchkov'), 'https://randus.ru/avatars/m/myAvatar17.png', '8(967)393-27-29');
INSERT INTO employee (id, role, surmane, name, patronymic, birth_date, email, password, photo, phone)
    VALUES (3, 'manager', 'Антонова', 'Аполлинария', 'Артемовна', '1974-09-15', 'antonova@truck.ru', md5('antonova'), 'https://randus.ru/avatars/w/myAvatar4.png', '8(910)336-71-23');
INSERT INTO employee (id, role, surmane, name, patronymic, birth_date, email, password, photo, phone)
    VALUES (4, 'driver', 'Леонтьев', 'Захар', 'Николаевич', '1976-07-16', 'leontev@truck.ru', md5('leontev'), 'https://randus.ru/avatars/m/myAvatar18.png', '8(905)363-75-34');
INSERT INTO employee (id, role, surmane, name, patronymic, birth_date, email, password, photo, phone)
    VALUES (5, 'driver', 'Кузьмин', 'Архип', 'Ильич', '1989-02-16', 'kuzmin@truck.ru', md5('kuzmin'), 'https://randus.ru/avatars/m/myAvatar13.png', '8(938)962-78-37');

INSERT INTO client (id, name, itn, iec, address)
    VALUES (1, 'ООО Спортмастер', '7728551528', '772801001', '117437, Российская Федерация, Москва, ул. Миклухо-Маклая, д. 18, корп. 2, ком. 102');
INSERT INTO client (id, name, itn, iec, address)
    VALUES (2, 'ООО Exist', '7721531730', '772101001', '109428, Российская Федерация, Москва, ул. Рязанский проспект, д.65, ком. 2');
INSERT INTO client (id, name, itn, iec, address)
    VALUES (3, 'ООО Decatlon', '5029086747', '502901001', ': 141031, Российская Федерация, Московская область, Мытищинский район, 84км МКАД, ТПЗ «Алтуфьево», владение 3, строение 3');
INSERT INTO client (id, name, itn, iec, address)
    VALUES (4, 'ООО ИнтерАвто', '6229045919', '622901001', '390044, Российская Федерация, Рязань, ш. Московское, д. 24');

INSERT INTO store(id, name, address, latitude, longitude)
    VALUES (1, 'ООО МосТраст', '109518, Москва, Рязанский пр., 2с26', 55.729273, 37.742948);
INSERT INTO store(id, name, address, latitude, longitude)
    VALUES (2, 'ООО ТоргЕкспрес', '443030, Самарская обл., Самара, ул. Спортивная, 1А',53.186259, 50.123964);
INSERT INTO store(id, name, address, latitude, longitude)
    VALUES (3, 'ООО Верст', '620146, Свердловская обл., Екатеринбург, ул. Амундсена, 74', 56.793216, 60.574221);
INSERT INTO store(id, name, address, latitude, longitude)
    VALUES (4, 'ООО ТюменьТраст', '625046, Тюменская обл., Тюмень, ул. Широтная, 170 корпус 3', 57.101800, 65.598239);
INSERT INTO store(id, name, address, latitude, longitude)
    VALUES (5, 'ООО ВладТранст', '690109, Приморский край, Владивосток, ул. Нейбута, 81А', 43.118757, 131.967477);
INSERT INTO store(id, name, address, latitude, longitude)
    VALUES (6, '高道汽車', 'Гонконг, Kam Tin, Kam Sheung Rd, 106', 22.788984, 114.096886);

INSERT INTO contract(id, manager_id, client_id, packaging_id, form_store_id, before_store_id, number, price, confirmation_payment_link, confirmation_customs_link, date_shipment)
    VALUES (1, 3, 1, 1, 6, 1, 'N-001', 170000, 'http://ipipip.ru/blank/doc/sberbank1.gif', 'http://www.2blanka.ru/static/media/prim_declaracija1.png', '2017-02-01');
INSERT INTO contract(id, manager_id, client_id, packaging_id, form_store_id, before_store_id, number, price, confirmation_payment_link, confirmation_customs_link, date_shipment)
    VALUES (2, 3, 3, 1, 6, 3, 'N-002', 140000, 'http://ipipip.ru/blank/doc/sberbank1.gif', 'http://www.2blanka.ru/static/media/prim_declaracija1.png', '2017-02-01');
INSERT INTO contract(id, manager_id, client_id, packaging_id, form_store_id, before_store_id, number, price, confirmation_payment_link, confirmation_customs_link, date_shipment)
    VALUES (3, 3, 3, 1, 5, 4, 'N-003', 96000, 'http://ipipip.ru/blank/doc/sberbank1.gif', 'http://www.2blanka.ru/static/media/prim_declaracija1.png', '2017-04-22');
INSERT INTO contract(id, manager_id, client_id, packaging_id, form_store_id, before_store_id, number, price, confirmation_payment_link, confirmation_customs_link, date_shipment)
    VALUES (4, 3, 3, 1, 4, 2, 'N-004', 56000, 'http://ipipip.ru/blank/doc/sberbank1.gif', 'http://www.2blanka.ru/static/media/prim_declaracija1.png', '2017-04-22');
INSERT INTO contract(id, manager_id, client_id, packaging_id, form_store_id, before_store_id, number, price, confirmation_payment_link, confirmation_customs_link, date_shipment)
    VALUES (5, 3, 2, 1, 5, 4, 'N-005', 270000, 'http://ipipip.ru/blank/doc/sberbank1.gif', 'http://www.2blanka.ru/static/media/prim_declaracija1.png', '2017-04-28');
INSERT INTO contract(id, manager_id, client_id, packaging_id, form_store_id, before_store_id, number, price, confirmation_payment_link, confirmation_customs_link, date_shipment)
    VALUES (6, 3, 2, 1, 6, 4, 'N-006', 250000, 'http://ipipip.ru/blank/doc/sberbank1.gif', 'http://www.2blanka.ru/static/media/prim_declaracija1.png', '2017-06-08');
INSERT INTO contract(id, manager_id, client_id, packaging_id, form_store_id, before_store_id, number, price, confirmation_payment_link, confirmation_customs_link, date_shipment)
    VALUES (7, 3, 4, 1, 5, 4, 'N-007', 220000, 'http://ipipip.ru/blank/doc/sberbank1.gif', 'http://www.2blanka.ru/static/media/prim_declaracija1.png', '2017-06-08');

INSERT INTO carriage(id, status, vehicle_id, driver_id) VALUES (1, 'archive', 1, 4);
INSERT INTO carriage(id, status, vehicle_id, driver_id) VALUES (2, 'current', 1,4);
INSERT INTO carriage(id, status, vehicle_id, driver_id) VALUES (3, 'current', 2,5);
INSERT INTO carriage(id, status, vehicle_id, driver_id) VALUES (4, 'onward', 1, 4);

INSERT INTO carriage_detail(id, carriage_id, contract_id) VALUES (1, 1, 1);
INSERT INTO carriage_detail(id, carriage_id, contract_id) VALUES (2, 1, 2);
INSERT INTO carriage_detail(id, carriage_id, contract_id) VALUES (3, 2, 3);
INSERT INTO carriage_detail(id, carriage_id, contract_id) VALUES (4, 2, 4);
INSERT INTO carriage_detail(id, carriage_id, contract_id) VALUES (5, 3, 5);
INSERT INTO carriage_detail(id, carriage_id, contract_id) VALUES (6, 4, 6);
INSERT INTO carriage_detail(id, carriage_id, contract_id) VALUES (7, 4, 7);

INSERT INTO measure(id, name, abbreviation) VALUES (1, 'Штука', 'шт.');
INSERT INTO measure(id, name, abbreviation) VALUES (2, 'Набор', 'набор');
INSERT INTO measure(id, name, abbreviation) VALUES (3, 'Пара (2 шт.)', 'пар');
INSERT INTO measure(id, name, abbreviation) VALUES (4, 'Упаковка', 'упак.');

INSERT INTO nomenclature(id, measure_id, name) VALUES (1, 2, 'Втулка уплотнительная клапанной крышки');
INSERT INTO nomenclature(id, measure_id, name) VALUES (2, 1, 'Шестерня распредвала малой цепи');
INSERT INTO nomenclature(id, measure_id, name) VALUES (3, 2, 'Кольцо уплотнительное свечного колодца');
INSERT INTO nomenclature(id, measure_id, name) VALUES (4, 1, 'Сальник коленвала передний');
INSERT INTO nomenclature(id, measure_id, name) VALUES (5, 1, 'Шестерня вала коленчатого');
INSERT INTO nomenclature(id, measure_id, name) VALUES (6, 1, 'Шестерня распредвала малой цепи');
INSERT INTO nomenclature(id, measure_id, name) VALUES (7, 1, 'Цепь приводная малая ГРМ');
INSERT INTO nomenclature(id, measure_id, name) VALUES (8, 1, 'Натяжитель цепи ГРМ');
INSERT INTO nomenclature(id, measure_id, name) VALUES (9, 1, 'Палатка 4-местная Nordway Bergen 4');
INSERT INTO nomenclature(id, measure_id, name) VALUES (10, 1, 'Палатка 6-местная Nordway Camper 4+2');
INSERT INTO nomenclature(id, measure_id, name) VALUES (11, 1, 'Палатка 6-местная Nordway Camper 6');
INSERT INTO nomenclature(id, measure_id, name) VALUES (12, 1, 'Палатка 6-местная Nordway Dalen 6');
INSERT INTO nomenclature(id, measure_id, name) VALUES (13, 1, 'Спальный мешок для кемпинга Outventure Light +20');
INSERT INTO nomenclature(id, measure_id, name) VALUES (14, 1, 'Рюкзак Outventure NEW Discovery 15');

INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (1, 1, 9, 20, 6999);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (2, 1, 10, 20, 12999);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (3, 1, 11, 10, 14999);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (4, 1, 12, 5, 15999);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (5, 1, 13, 50, 699);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (6, 1, 14, 50, 699);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (7, 2, 10, 60, 12999);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (8, 2, 11, 50, 14999);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (9, 2, 12, 30, 15999);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (10, 3, 13, 500, 699);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (11, 3, 14, 500, 699);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (12, 4, 10, 60, 12999);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (13, 4, 11, 50, 14999);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (14, 4, 12, 30, 15999);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (15, 5, 1, 600, 294);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (16, 5, 2, 400, 2170);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (17, 5, 3, 700, 220);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (18, 5, 4, 200, 537);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (19, 6, 5, 50, 1838);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (20, 6, 6, 60, 2134);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (21, 6, 7, 40, 1397);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (22, 6, 8, 10, 3323);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (23, 7, 2, 40, 2170);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (24, 7, 3, 70, 220);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (25, 7, 4, 70, 537);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (26, 7, 5, 50, 1838);
INSERT INTO contract_detail(id, contract_id, nomenclature_id, amount, price) VALUES (27, 7, 6, 60, 2134);

INSERT INTO check_point(id, name, address, latitude, longitude)
    VALUES (1, 'ООО МосТраст', '109518, Москва, Рязанский пр., 2с26', 55.729273, 37.742948);
INSERT INTO check_point(id, name, address, latitude, longitude)
    VALUES (2, 'ООО ТоргЕкспрес', '443030, Самарская обл., Самара, ул. Спортивная, 1А', 53.186259, 50.123964);
INSERT INTO check_point(id, name, address, latitude, longitude)
    VALUES (3, 'ООО Верст', '620146, Свердловская обл., Екатеринбург, ул. Амундсена, 74', 56.793216, 60.574221);
INSERT INTO check_point(id, name, address, latitude, longitude)
    VALUES (4, 'ООО ТюменьТраст', '625046, Тюменская обл., Тюмень, ул. Широтная, 170 корпус 3', 57.101800, 65.598239);
INSERT INTO check_point(id, name, address, latitude, longitude)
    VALUES (5, 'ООО РосПуть', '627755, Тюменская обл., Ишим, Красноярская ул., 133', 56.120191, 69.499795);
INSERT INTO check_point(id, name, address, latitude, longitude)
    VALUES (6, 'ООО НовосибирскТраст', '630119, Новосибирская обл., Новосибирск, ул. Петухова, 79', 54.943100, 82.929316);
INSERT INTO check_point(id, name, address, latitude, longitude)
    VALUES (7, 'ООО Красноярская Транспортная Компания', '660079, Красноярский край, Красноярск, ул. Александра Матросова, 41А', 55.974116, 92.887662);
INSERT INTO check_point(id, name, address, latitude, longitude)
    VALUES (8, 'ООО ВладТранст', '690109, Приморский край, Владивосток, ул. Нейбута, 81А', 43.118757, 131.967477);
INSERT INTO check_point(id, name, address, latitude, longitude)
    VALUES (9, '高道汽車', 'Гонконг, Kam Tin, Kam Sheung Rd, 106', 22.788984, 114.096886);

INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned, timestamp_fact)
    VALUES (1, 1, 9, '2017-02-01 15:00:00', '2017-02-01 15:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned, timestamp_fact)
    VALUES (2, 1, 8, '2017-02-03 12:00:00', '2017-02-03 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned, timestamp_fact)
    VALUES (3, 1, 7, '2017-02-05 12:00:00', '2017-02-05 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned, timestamp_fact)
    VALUES (4, 1, 6, '2017-02-06 12:00:00', '2017-02-06 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned, timestamp_fact)
    VALUES (5, 1, 5, '2017-02-08 12:00:00', '2017-02-08 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned, timestamp_fact)
    VALUES (6, 1, 4, '2017-02-10 12:00:00', '2017-02-10 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned, timestamp_fact)
    VALUES (7, 1, 6, '2017-02-11 12:00:00', '2017-02-11 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned, timestamp_fact)
    VALUES (8, 1, 2, '2017-02-12 12:00:00', '2017-02-12 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned, timestamp_fact)
    VALUES (9, 1, 1, '2017-02-13 12:00:00', '2017-02-13 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (10, 2, 8, '2017-04-22 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (11, 2, 7, '2017-04-24 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (12, 2, 6, '2017-04-26 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (13, 2, 5, '2017-04-28 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (14, 2, 4, '2017-04-30 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (15, 2, 6, '2017-05-1 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (16, 2, 2, '2017-05-2 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (17, 3, 8, '2017-04-28 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (18, 3, 7, '2017-04-30 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (19, 3, 6, '2017-05-01 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (20, 3, 5, '2017-05-03 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (21, 3, 4, '2017-05-05 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (22, 4, 9, '2017-06-08 15:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (23, 4, 8, '2017-06-10 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (24, 4, 7, '2017-06-12 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (25, 4, 6, '2017-06-14 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (26, 4, 5, '2017-06-16 12:00:00');
INSERT INTO carriage_route(id, carriage_id, check_point_id, timestamp_planned)
    VALUES (27, 4, 4, '2017-06-18 12:00:00');