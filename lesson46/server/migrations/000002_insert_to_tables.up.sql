INSERT INTO schedule (route_number, stop_name, arrival_time) VALUES
('1', 'Mustaqillik maydoni', '08:00:00'),
('1', 'Amir Temur xiyoboni', '08:15:00'),
('2', 'Alisher Navoiy metro', '09:00:00'),
('2', 'Chorsu bozor', '09:20:00'),
('3', 'Minor masjid', '10:00:00'),
('3', 'Gafur Gulom metro', '10:30:00'),
('4', 'Oybek metro', '11:00:00'),
('4', 'Bunyodkor stadion', '11:45:00'),
('5', 'Buyuk Ipak Yoli metro', '12:30:00'),
('5', 'Xadra maydoni', '12:45:00');

INSERT INTO bus_location (bus_id, latitude, longitude, timestamp) VALUES
('B001', 41.311151, 69.279737, '2023-06-25 08:00:00'),
('B002', 41.326015, 69.250545, '2023-06-25 08:05:00'),
('B003', 41.331871, 69.244211, '2023-06-25 08:10:00'),
('B004', 41.299496, 69.240073, '2023-06-25 08:15:00'),
('B005', 41.314389, 69.287367, '2023-06-25 08:20:00'),
('B006', 41.316339, 69.282367, '2023-06-25 08:25:00'),
('B007', 41.318239, 69.278367, '2023-06-25 08:30:00'),
('B008', 41.320139, 69.274367, '2023-06-25 08:35:00'),
('B009', 41.322039, 69.270367, '2023-06-25 08:40:00'),
('B010', 41.324039, 69.266367, '2023-06-25 08:45:00');

INSERT INTO traffic_jam_report (location, severity, description) VALUES
('Shota Rustaveli ko''chasi', 'Yuqori', 'Transport vositalari ko''pligi sababli tirbandlik.'),
('Amir Temur xiyoboni', 'O''rtacha', 'Yengil avtomobillar sababli tirbandlik.'),
('Alisher Navoiy ko''chasi', 'Past', 'Engil tirbandlik.'),
('Buyuk Ipak Yo''li ko''chasi', 'Yuqori', 'Tuzatish ishlari olib borilmoqda.'),
('Beruniy ko''chasi', 'O''rtacha', 'Transport vositalari ko''pligi sababli tirbandlik.'),
('Navoiy shoh ko''chasi', 'Past', 'Yengil avtomobillar sababli tirbandlik.'),
('Mirzo Ulug''bek ko''chasi', 'Yuqori', 'Yuk avtomobillari sababli tirbandlik.'),
('G''alaba ko''chasi', 'O''rtacha', 'Tuzatish ishlari olib borilmoqda.'),
('Bunyodkor ko''chasi', 'Past', 'Transport vositalari ko''pligi sababli tirbandlik.'),
('Yunusobod ko''chasi', 'Yuqori', 'Transport vositalari ko''pligi sababli tirbandlik.');

INSERT INTO current_weather (location, temperature, humidity, speed_of_wind, timestamp) VALUES
('Toshkent', 35.5, 40, 5.5, '2023-06-25 08:00:00'),
('Samarqand', 33.0, 45, 6.0, '2023-06-25 08:00:00'),
('Buxoro', 37.2, 38, 4.5, '2023-06-25 08:00:00'),
('Andijon', 31.8, 50, 3.5, '2023-06-25 08:00:00'),
('Namangan', 32.5, 48, 4.0, '2023-06-25 08:00:00'),
('Farg''ona', 34.0, 42, 4.8, '2023-06-25 08:00:00'),
('Urganch', 36.0, 35, 5.0, '2023-06-25 08:00:00'),
('Qo''qon', 33.5, 47, 3.8, '2023-06-25 08:00:00'),
('Guliston', 35.2, 41, 5.2, '2023-06-25 08:00:00'),
('Nukus', 38.0, 33, 5.6, '2023-06-25 08:00:00');

INSERT INTO weather_forecast (location, date, temperature_high, temperature_low, condition) VALUES
('Toshkent', '2023-06-26', 36.0, 25.0, 'Quyoshli'),
('Toshkent', '2023-06-27', 37.0, 26.0, 'Quyoshli'),
('Toshkent', '2023-06-28', 36.5, 25.5, 'Bulutli'),
('Toshkent', '2023-06-29', 35.0, 24.0, 'Yomg''irli'),
('Toshkent', '2023-06-30', 34.0, 23.0, 'Quyoshli'),
('Toshkent', '2023-07-01', 33.0, 22.0, 'Quyoshli'),
('Toshkent', '2023-07-02', 32.0, 21.0, 'Bulutli'),
('Toshkent', '2023-07-03', 31.0, 20.0, 'Quyoshli'),
('Toshkent', '2023-07-04', 30.0, 19.0, 'Quyoshli'),
('Toshkent', '2023-07-05', 29.0, 18.0, 'Yomg''irli');

INSERT INTO weather_condition_report (location, temperature, humidity, condition, timestamp) VALUES
('Toshkent', 35.5, 40, 'Quyoshli', '2023-06-25 08:00:00'),
('Samarqand', 33.0, 45, 'Bulutli', '2023-06-25 08:00:00'),
('Buxoro', 37.2, 38, 'Quyoshli', '2023-06-25 08:00:00'),
('Andijon', 31.8, 50, 'Yomg''irli', '2023-06-25 08:00:00'),
('Namangan', 32.5, 48, 'Quyoshli', '2023-06-25 08:00:00'),
('Farg''ona', 34.0, 42, 'Quyoshli', '2023-06-25 08:00:00'),
('Urganch', 36.0, 35, 'Bulutli', '2023-06-25 08:00:00'),
('Qo''qon', 33.5, 47, 'Yomg''irli', '2023-06-25 08:00:00'),
('Guliston', 35.2, 41, 'Quyoshli', '2023-06-25 08:00:00'),
('Nukus', 38.0, 33, 'Quyoshli', '2023-06-25 08:00:00');
