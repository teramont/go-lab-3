START TRANSACTION;

DROP TABLE IF EXISTS Disks;
DROP TABLE IF EXISTS Machines;

CREATE TABLE Machines (
  id  serial PRIMARY KEY,
  name  varchar(64) NOT NULL
);

CREATE TABLE Disks (
  id serial PRIMARY KEY,
  space bigint NOT NULL,
  machineId integer NULL
);

ALTER TABLE Disks ADD CONSTRAINT fkDisksMachinesId FOREIGN KEY (machineId) REFERENCES Machines (Id);

COMMIT;