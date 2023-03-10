CREATE TABLE plants (
	id varchar(36) primary key not null,
	name varchar(50) not null,
	location varchar(50) not null,
	description text null,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP,
   	updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX name_plants ON plants (name);