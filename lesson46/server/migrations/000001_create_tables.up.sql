-- Transport
CREATE TABLE IF NOT EXISTS schedule (
    id SERIAL PRIMARY KEY,
    route_number VARCHAR(50) NOT NULL,
    stop_name VARCHAR(100) NOT NULL,
    arrival_time TIME NOT NULL
);

CREATE TABLE IF NOT EXISTS bus_location (
    id SERIAL PRIMARY KEY,
    bus_id VARCHAR(50) NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS traffic_jam_report (
    id SERIAL PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    severity VARCHAR(50) NOT NULL,
    description TEXT
);

-- Weather 

CREATE TABLE IF NOT EXISTS current_weather (
    id SERIAL PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    temperature FLOAT NOT NULL,
    humidity FLOAT NOT NULL,
    speed_of_wind FLOAT NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS weather_forecast (
    id SERIAL PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    temperature_high FLOAT NOT NULL,
    temperature_low FLOAT NOT NULL,
    condition VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS weather_condition_report (
    id SERIAL PRIMARY KEY,
    location VARCHAR(255) NOT NULL,
    temperature FLOAT NOT NULL,
    humidity FLOAT NOT NULL,
    condition VARCHAR(255) NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);