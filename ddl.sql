-- Create Users table
CREATE TABLE Users (
    UserID SERIAL PRIMARY KEY,
    EmailUsername VARCHAR(255) UNIQUE NOT NULL,
    Password VARCHAR(255) NOT NULL,
    DepositAmount DECIMAL DEFAULT 0
);

-- Create EquipmentTypes table
CREATE TABLE EquipmentTypes (
    TypeID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Availability BOOLEAN NOT NULL,
    RentalCosts DECIMAL NOT NULL,
    Category VARCHAR(255) NOT NULL
);

-- Create RentalHistory table
CREATE TABLE RentalHistory (
    RentalID SERIAL PRIMARY KEY,
    UserID INT REFERENCES Users(UserID),
    EquipmentID INT REFERENCES EquipmentTypes(TypeID),
    RentalDate DATE NOT NULL,
    ReturnDate DATE,
    CONSTRAINT fk_user_rental FOREIGN KEY (UserID) REFERENCES Users(UserID),
    CONSTRAINT fk_equipment_rental FOREIGN KEY (EquipmentID) REFERENCES EquipmentTypes(TypeID)
);

-- Create EquipmentRentals table
CREATE TABLE Orders (
    RentalID SERIAL PRIMARY KEY,
    UserID INT REFERENCES Users(UserID),
    EquipmentID INT REFERENCES EquipmentTypes(TypeID),
    RentalDate DATE NOT NULL,
    ReturnDate DATE,
    CONSTRAINT fk_user_rental_current FOREIGN KEY (UserID) REFERENCES Users(UserID),
    CONSTRAINT fk_equipment_rental_current FOREIGN KEY (EquipmentID) REFERENCES EquipmentTypes(TypeID)
);
