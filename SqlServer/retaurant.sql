create database restaurant

use restaurant 
go

create table restaurants
(
id INT IDENTITY(1,1) PRIMARY KEY,
ower_id INT not null,
[name] VARCHAR(50) not null,
addr VARCHAR(250) not null,
city_id INT default null,
lat FLOAT default null ,
lng FLOAT default null ,
cover NVARCHAR(max) not null,
logo NVARCHAR(max) not null,
shipping_fee_per_km FLOAT default 0,
[status] BIT not null default 1,
created_at DATETIME null default CURRENT_TIMESTAMP,
upated_at DATETIME null default CURRENT_TIMESTAMP
);
go
