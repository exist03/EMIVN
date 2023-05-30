CREATE TABLE `Card`
(
    `ID`        varchar(40) PRIMARY KEY,
    `Owner`     varchar(40) COMMENT 'Daimyo.Nickname',
    `BankInfo`  varchar(40),
    `LimitInfo` float
);

CREATE TABLE `Samurai`
(
    `Nickname`         varchar(255),
    `TelegramUsername` varchar(40) PRIMARY KEY,
    `Owner`            varchar(40) COMMENT 'Daimyo.Nickname'
);

CREATE TABLE `Daimyo`
(
    `Nickname`         varchar(40),
    `Owner`            varchar(40),
    `TelegramUsername` varchar(40) PRIMARY KEY
);

CREATE TABLE `Shogun`
(
    `TelegramUsername` varchar(40) PRIMARY KEY,
    `Nickname`         varchar(40)
);

CREATE TABLE `Collector`
(
    `TelegramUsername` varchar(40) PRIMARY KEY,
    `Nickname`         varchar(40)
);

CREATE TABLE `Application`
(
    `Daimyo` varchar(40),
    `Sum`    float,
    `ID`     varchar(40) PRIMARY KEY COMMENT 'card_id'
);
CREATE TABLE `Admin`
(
    `TelegramUsername` varchar(40) PRIMARY KEY,
    `Nickname`         varchar(40)
);

CREATE TABLE `Transaction`
(
    `OperationType` bool,
    `Amount`        float,
    `Date`          DATETIME,
    `CardID`        varchar(40),
    FOREIGN KEY (`CardID`) REFERENCES Card (`ID`),
    PRIMARY KEY (`Date`, `CardID`)
);

CREATE TABLE `SamuraiTurnover`
(
    `Amount`          float,
    `Date`            DATE,
    `SamuraiUsername` varchar(40),
    `Bank`            varchar(40),
    FOREIGN KEY (`SamuraiUsername`) REFERENCES Samurai (`TelegramUsername`),
    PRIMARY KEY (`SamuraiUsername`, `Date`, `Bank`)
);

create table TurnoverController
(
    Amount          float       null,
    Date            date        not null,
    SamuraiUsername varchar(40) not null,
    Bank            varchar(40) not null,
    primary key (SamuraiUsername, Date, Bank),
    foreign key (SamuraiUsername) references Samurai (TelegramUsername)
);

create table CardBalance
(
    CardID  varchar(40),
    Balance varchar(40),
    primary key (CardID),
    foreign key (CardID) references Card (ID)
);
create table CardBalanceDaily
(
    CardID varchar(40),
    Amount varchar(40),
    `Time` DATE,
    primary key (CardID, Time),
    foreign key (CardID) references Card (ID)
);
create table ControllerBalance
(
    Amount varchar(40),
    `Time` DATE,
    primary key (Time)
);

ALTER TABLE `Card`
    ADD FOREIGN KEY (`Owner`) REFERENCES `Daimyo` (`TelegramUsername`);

ALTER TABLE `Samurai`
    ADD FOREIGN KEY (`Owner`) REFERENCES `Daimyo` (`TelegramUsername`);

ALTER TABLE `Daimyo`
    ADD FOREIGN KEY (`Owner`) REFERENCES `Shogun` (`TelegramUsername`);

ALTER TABLE `Application`
    ADD FOREIGN KEY (`Daimyo`) REFERENCES `Daimyo` (`TelegramUsername`);
#сумма всех карт на начало дня
# SELECT TelegramUsername, SUM(CBD.Amount) FROM Daimyo
# JOIN  Card ON Daimyo.TelegramUsername = Card.Owner
# JOIN CardBalanceDaily CBD ON Card.ID = CBD.CardID
# WHERE CBD.Time="2023-05-16" AND Daimyo.Owner="exist03" GROUP BY TelegramUsername;

#сумма всех пополнений за день
# SELECT TelegramUsername, SUM(T.Amount) FROM Daimyo
# JOIN Card ON Daimyo.TelegramUsername = Card.Owner
# JOIN Transaction T on Card.ID = T.CardID WHERE T.OperationType=true AND Daimyo.Owner="exist03" AND T.Date BETWEEN "2023-05-16 8:00" AND "2023-05-17 8:00" GROUP BY TelegramUsername;
#оборот всех самураев
#
# SELECT Daimyo.TelegramUsername, SUM(TE.Amount) FROM Daimyo
# JOIN Samurai S ON Daimyo.TelegramUsername = S.Owner
# JOIN TurnoverEnd TE ON TE.SamuraiUsername = S.TelegramUsername
# WHERE TE.Date = "2023-05-16" AND Daimyo.Owner = "exist03" GROUP BY Daimyo.TelegramUsername;
# -
# SELECT Daimyo.TelegramUsername, SUM(TB.Amount) FROM Daimyo
# JOIN Samurai S ON Daimyo.TelegramUsername = S.Owner
# JOIN TurnoverBegin TB ON TB.SamuraiUsername = S.TelegramUsername
# WHERE TB.Date = "2023-05-16" AND Daimyo.Owner = "exist03" GROUP BY Daimyo.TelegramUsername;
# SELECT Daimyo.TelegramUsername,TC.Amount FROM Daimyo
# JOIN Samurai S on Daimyo.TelegramUsername = S.Owner
# JOIN TurnoverController TC on S.TelegramUsername = TC.SamuraiUsername
# WHERE TC.Date = "2023-05-09" AND Daimyo.Owner = "exist03";
# SELECT Amount FROM ControllerBalance WHERE `Time` = "2023-05-27"