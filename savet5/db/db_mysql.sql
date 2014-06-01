# Users table
 
# --- !Ups
 
CREATE TABLE Users (
    id bigint(20) NOT NULL AUTO_INCREMENT,
    email varchar(255) NOT NULL,
    password varchar(255) not null,
    firstName varchar(255) NOT NULL,
    lastName varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    role varchar(255) not null default 'OPERATOR',
    PRIMARY KEY (id)
);

insert into Users(email, password, firstName, lastName, description,role) values ('a@a.a','1','Test','Tester','Bogotac','GOD');

# DROP TABLE Users;

CREATE TABLE Saveti (
    id bigint(20) NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    address varchar(255) not null,
    created_at TIMESTAMP NOT NULL default current_timestamp,
    PRIMARY KEY (id)
);

# drop table Saveti;

insert into Saveti(name, address) values ('Savet 1','Testni savet br.1 na nekoj adresi');
insert into Saveti(name, address) values ('Savet 2','Testni savet br.2 na nekoj drugoj adresi');
insert into Saveti(name, address) values ('Savet 3','Testni savet br.3 opet na nekoj adresi');

# Users table
 
# --- !Ups
 
CREATE TABLE Stanari (
    id bigint(20) NOT NULL AUTO_INCREMENT,
    SAVET_ID bigint(20) NOT NULL,
    broj_Stana varchar(255) not null,
    redosled int ,
    name varchar(255) NOT NULL,
    last_name varchar(255),
	PRIMARY KEY (id)
);

insert into Stanari(id, SAVET_ID, broj_Stana, redosled, name, last_name) values (1, 1, 'S1', 1, 'Icabod', 'Crane');
insert into Stanari(id, SAVET_ID, broj_Stana, redosled, name, last_name) values (2, 1, 'S2', 2, 'Icabod2', 'Crane2');
insert into Stanari(id, SAVET_ID, broj_Stana, redosled, name, last_name) values (3, 1, 'S3', 3, 'Icabod3', 'Crane3');

insert into Stanari(id, SAVET_ID, broj_Stana, redosled, name, last_name) values (11, 2, 'S1', 1, 'X-Icabod', 'Crane');
insert into Stanari(id, SAVET_ID, broj_Stana, redosled, name, last_name) values (12, 2, 'S2', 2, 'X-Icabod2', 'Crane2');
insert into Stanari(id, SAVET_ID, broj_Stana, redosled, name, last_name) values (13, 2, 'S3', 3, 'X-Icabod3', 'Crane3');

