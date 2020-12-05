START TRANSACTION;

CREATE TABLE Machines (
  id  serial PRIMARY KEY,
  name  varchar(64) NOT NULL
);

CREATE TABLE Disks (
  id serial PRIMARY KEY,
  name varchar(64) NOT NULL,
  space integer NOT NULL,
  machineId integer NULL
);

ALTER TABLE Disks ADD CONSTRAINT fkDisksMachinesId FOREIGN KEY (machineId) REFERENCES Machines (Id);

COMMIT;