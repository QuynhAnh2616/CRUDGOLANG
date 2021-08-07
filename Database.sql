USE master
go

DROP DATABASE IF EXISTS BaiThiNetCore
CREATE DATABASE BaiThiNetCore
GO



USE BaiThiNetCore
GO

CREATE TABLE Customer(
Id int primary key identity,
[Name] nvarchar(250),
[Address] nvarchar(250),
Birthday date,
Phone varchar(20)
)
go

Create table Orders(
Id int primary key identity,
[Name] nvarchar(250),
DateCreation date,
[Status] bit,
Payment nvarchar(250),
CustomerId int constraint fk_Orders_Customer foreign key (CustomerId) references Customer(Id)
)

go

insert into Customer
values('Cus1','cus12','2000/12/03','123456'),
('Cus2','cus22','2000/12/25','123456'),
('Cus3','cus33','2000/06/03','123456'),
('Cus4','cus44','2000/07/03','123456')
go

insert into Orders
values ('order1','2020/12/20',1,'Cash',1),
('order2','2020/12/20',0,'Cash',1),
('order3','2020/12/20',1, 'Visa Card',1),
('order4','2020/12/20',1,'Master Card ',4),
('order5','2020/12/20',1,'Master Card ',3),
('order6','2020/12/20',1,'Paypal',3)
go
select * from Orders