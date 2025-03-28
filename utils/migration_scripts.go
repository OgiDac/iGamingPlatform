package utils

const CreatePlayers = `
CREATE TABLE IF NOT EXISTS players (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name TEXT NOT NULL,
	password TEXT NOT NULL,
	email TEXT NOT NULL,
	accountBalance DOUBLE PRECISION NOT NULL,
    role TEXT NOT NULL
);`

const CreateTournaments = `
CREATE TABLE IF NOT EXISTS tournaments (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name TEXT NOT NULL,
	prizePool DOUBLE PRECISION NOT NULL,
	startDate DATETIME NOT NULL,
	endDate DATETIME NOT NULL,
    chanceToWin INT NOT NULL
);`

const CreatePlayerTournaments = `
CREATE TABLE IF NOT EXISTS player_tournaments (
	id INT AUTO_INCREMENT PRIMARY KEY,
	player_id INT NOT NULL,
	tournament_id INT NOT NULL,
	score FLOAT NOT NULL,
    totalInvested FLOAT NOT NULL DEFAULT 0,
	UNIQUE KEY unique_player_tournament (player_id, tournament_id),
	FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE,
	FOREIGN KEY (tournament_id) REFERENCES tournaments(id) ON DELETE CASCADE
);`

const InsertPlayers = `
INSERT INTO players (name, password, email, accountBalance, role) VALUES
('Alice', 'pass1', 'alice@example.com', 100.0, 'user'),
('Bob', 'pass2', 'bob@example.com', 200.0, 'user'),
('Carol', 'pass3', 'carol@example.com', 150.0, 'user'),
('Dave', 'pass4', 'dave@example.com', 250.0, 'user'),
('Eve', 'pass5', 'eve@example.com', 300.0, 'user'),
('Frank', 'pass6', 'frank@example.com', 120.0, 'user'),
('Grace', 'pass7', 'grace@example.com', 180.0, 'user'),
('Heidi', 'pass8', 'heidi@example.com', 220.0, 'user'),
('Ivan', 'pass9', 'ivan@example.com', 140.0, 'user'),
('Judy', 'pass10', 'judy@example.com', 160.0, 'user'),
('Admin', 'pass11', 'admin@example.com', 160.0, 'admin');`

const InsertTournaments = `
INSERT INTO tournaments (name, prizePool, startDate, endDate, chanceToWin) VALUES
('Spring Cup', 10000.0, '2025-01-01 10:00:00', '2025-04-10 18:00:00', 30),
('Summer Showdown', 20000.0, '2025-01-05 09:00:00', '2025-07-15 20:00:00', 70),
('Autumn Arena', 15000.0, '2025-02-01 12:00:00', '2025-10-10 17:00:00', 5),
('Past Event', 15000.0, '2025-01-01 12:00:00', '2025-02-10 17:00:00', 5),
('Future Event', 15000.0, '2025-09-01 12:00:00', '2025-11-10 17:00:00', 5);
`

const CreatePrizeDistributionSP = `
CREATE PROCEDURE DistributePrizes(IN tournamentId INT)
BEGIN
    WITH ranked AS (
        SELECT
            player_id,
            RANK() OVER (ORDER BY totalInvested DESC) AS r
        FROM player_tournaments
        WHERE player_tournaments.tournament_id = tournament_id 
    )
    UPDATE players
    JOIN ranked ON players.id = ranked.player_id
    JOIN tournaments t ON t.id = tournamentId
    SET players.accountBalance = players.accountBalance +
        CASE
            WHEN ranked.r = 1 THEN t.prizePool * 0.5
            WHEN ranked.r = 2 THEN t.prizePool * 0.3 
            WHEN ranked.r = 3 THEN t.prizePool * 0.2 
            ELSE 0
        END
    WHERE ranked.r IN (1, 2, 3);
END
`
