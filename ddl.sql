CREATE TABLE Users (
    EmailUsername VARCHAR(255) PRIMARY KEY,
    Password VARCHAR(255),
    DepositAmount DECIMAL(10, 2) DEFAULT 0
);

CREATE TABLE Equipment (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(255),
    Availability BOOLEAN,
    RentalCosts DECIMAL(10, 2),
    Category VARCHAR(50)
);

CREATE TABLE EquipmentRentalHistory (
    RentalID INT PRIMARY KEY AUTO_INCREMENT,
    UserID VARCHAR(255 ),
    EquipmentID INT,
    RentalDate DATE,
    ReturnDate DATE,
    RentalCost DECIMAL(10, 2),
    Status VARCHAR(50),
    FOREIGN KEY (UserID) REFERENCES Users(EmailUsername),
    FOREIGN KEY (EquipmentID) REFERENCES Equipment(ID)
);

