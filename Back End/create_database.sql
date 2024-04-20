-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 20, 2024 at 10:31 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `catalog`
--

-- --------------------------------------------------------

--
-- Table structure for table `activitate`
--

CREATE TABLE `activitate` (
  `id_nota` int(11) NOT NULL,
  `id_scoala` int(11) DEFAULT NULL,
  `nume_disciplina` varchar(100) DEFAULT NULL,
  `id_clasa` varchar(50) DEFAULT NULL,
  `id_elev` int(11) DEFAULT NULL,
  `valoare` int(2) DEFAULT NULL,
  `data` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `clasa`
--

CREATE TABLE `clasa` (
  `id_scoala` int(11) NOT NULL,
  `nume` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `cont`
--

CREATE TABLE `cont` (
  `id` int(11) NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  `parola` varchar(50) DEFAULT NULL,
  `nume` varchar(50) DEFAULT NULL,
  `prenume` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `cont_rol`
--

CREATE TABLE `cont_rol` (
  `id_cont` int(11) NOT NULL,
  `id_rol` varchar(30) NOT NULL,
  `id_scoala` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `discipline`
--

CREATE TABLE `discipline` (
  `id_scoala` int(11) NOT NULL,
  `nume` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `elev`
--

CREATE TABLE `elev` (
  `id_scoala` int(11) NOT NULL,
  `id_clasa` varchar(50) NOT NULL,
  `id_elev` int(11) NOT NULL,
  `nume` varchar(50) DEFAULT NULL,
  `prenume` varchar(50) DEFAULT NULL,
  `gen` varchar(10) DEFAULT NULL,
  `etnie` varchar(20) DEFAULT NULL,
  `token_elev` varchar(10) DEFAULT NULL,
  `token_parinte` varchar(10) DEFAULT NULL,
  `id_cont_elev` int(11) DEFAULT NULL,
  `id_cont_parinte` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `incadrare`
--

CREATE TABLE `incadrare` (
  `id_scoala` int(11) NOT NULL,
  `id_clasa` varchar(50) NOT NULL,
  `id_profesor` int(11) NOT NULL,
  `nume_disciplina` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `note`
--

CREATE TABLE `note` (
  `id_nota` int(11) NOT NULL,
  `id_scoala` int(11) DEFAULT NULL,
  `nume_disciplina` varchar(100) DEFAULT NULL,
  `id_clasa` varchar(50) DEFAULT NULL,
  `id_elev` int(11) DEFAULT NULL,
  `nota` int(2) DEFAULT NULL,
  `data` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `profesor`
--

CREATE TABLE `profesor` (
  `id_scoala` int(11) NOT NULL,
  `id` int(11) NOT NULL,
  `nume` varchar(50) DEFAULT NULL,
  `prenume` varchar(50) DEFAULT NULL,
  `token` varchar(10) DEFAULT NULL,
  `id_cont` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `roluri`
--

CREATE TABLE `roluri` (
  `nume_rol` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `scoala`
--

CREATE TABLE `scoala` (
  `id` int(11) NOT NULL,
  `nume` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `activitate`
--
ALTER TABLE `activitate`
  ADD PRIMARY KEY (`id_nota`),
  ADD KEY `id_scoala` (`id_scoala`,`id_clasa`,`id_elev`),
  ADD KEY `id_scoala_2` (`id_scoala`,`id_clasa`,`nume_disciplina`);

--
-- Indexes for table `clasa`
--
ALTER TABLE `clasa`
  ADD PRIMARY KEY (`id_scoala`,`nume`);

--
-- Indexes for table `cont`
--
ALTER TABLE `cont`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `cont_rol`
--
ALTER TABLE `cont_rol`
  ADD PRIMARY KEY (`id_cont`,`id_rol`,`id_scoala`),
  ADD KEY `id_rol` (`id_rol`),
  ADD KEY `id_scoala` (`id_scoala`);

--
-- Indexes for table `discipline`
--
ALTER TABLE `discipline`
  ADD PRIMARY KEY (`id_scoala`,`nume`);

--
-- Indexes for table `elev`
--
ALTER TABLE `elev`
  ADD PRIMARY KEY (`id_scoala`,`id_clasa`,`id_elev`),
  ADD KEY `id_cont_elev` (`id_cont_elev`),
  ADD KEY `id_cont_parinte` (`id_cont_parinte`);

--
-- Indexes for table `incadrare`
--
ALTER TABLE `incadrare`
  ADD PRIMARY KEY (`id_scoala`,`id_clasa`,`nume_disciplina`),
  ADD KEY `id_scoala` (`id_scoala`,`nume_disciplina`),
  ADD KEY `id_scoala_2` (`id_scoala`,`id_profesor`);

--
-- Indexes for table `note`
--
ALTER TABLE `note`
  ADD PRIMARY KEY (`id_nota`),
  ADD KEY `id_scoala` (`id_scoala`,`id_clasa`,`id_elev`),
  ADD KEY `id_scoala_2` (`id_scoala`,`id_clasa`,`nume_disciplina`);

--
-- Indexes for table `profesor`
--
ALTER TABLE `profesor`
  ADD PRIMARY KEY (`id_scoala`,`id`),
  ADD KEY `id_cont` (`id_cont`);

--
-- Indexes for table `roluri`
--
ALTER TABLE `roluri`
  ADD PRIMARY KEY (`nume_rol`);

--
-- Indexes for table `scoala`
--
ALTER TABLE `scoala`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `activitate`
--
ALTER TABLE `activitate`
  MODIFY `id_nota` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `cont`
--
ALTER TABLE `cont`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `note`
--
ALTER TABLE `note`
  MODIFY `id_nota` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `scoala`
--
ALTER TABLE `scoala`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `activitate`
--
ALTER TABLE `activitate`
  ADD CONSTRAINT `activitate_ibfk_1` FOREIGN KEY (`id_scoala`,`id_clasa`,`id_elev`) REFERENCES `elev` (`id_scoala`, `id_clasa`, `id_elev`),
  ADD CONSTRAINT `activitate_ibfk_2` FOREIGN KEY (`id_scoala`,`id_clasa`,`nume_disciplina`) REFERENCES `incadrare` (`id_scoala`, `id_clasa`, `nume_disciplina`);

--
-- Constraints for table `clasa`
--
ALTER TABLE `clasa`
  ADD CONSTRAINT `clasa_ibfk_1` FOREIGN KEY (`id_scoala`) REFERENCES `scoala` (`id`);

--
-- Constraints for table `cont_rol`
--
ALTER TABLE `cont_rol`
  ADD CONSTRAINT `cont_rol_ibfk_1` FOREIGN KEY (`id_cont`) REFERENCES `cont` (`id`),
  ADD CONSTRAINT `cont_rol_ibfk_2` FOREIGN KEY (`id_rol`) REFERENCES `roluri` (`nume_rol`),
  ADD CONSTRAINT `cont_rol_ibfk_3` FOREIGN KEY (`id_scoala`) REFERENCES `scoala` (`id`);

--
-- Constraints for table `discipline`
--
ALTER TABLE `discipline`
  ADD CONSTRAINT `discipline_ibfk_1` FOREIGN KEY (`id_scoala`) REFERENCES `scoala` (`id`);

--
-- Constraints for table `elev`
--
ALTER TABLE `elev`
  ADD CONSTRAINT `elev_ibfk_1` FOREIGN KEY (`id_scoala`,`id_clasa`) REFERENCES `clasa` (`id_scoala`, `nume`),
  ADD CONSTRAINT `elev_ibfk_2` FOREIGN KEY (`id_cont_elev`) REFERENCES `cont` (`id`),
  ADD CONSTRAINT `elev_ibfk_3` FOREIGN KEY (`id_cont_parinte`) REFERENCES `cont` (`id`);

--
-- Constraints for table `incadrare`
--
ALTER TABLE `incadrare`
  ADD CONSTRAINT `incadrare_ibfk_1` FOREIGN KEY (`id_scoala`,`id_clasa`) REFERENCES `clasa` (`id_scoala`, `nume`),
  ADD CONSTRAINT `incadrare_ibfk_2` FOREIGN KEY (`id_scoala`,`nume_disciplina`) REFERENCES `discipline` (`id_scoala`, `nume`),
  ADD CONSTRAINT `incadrare_ibfk_3` FOREIGN KEY (`id_scoala`,`id_profesor`) REFERENCES `profesor` (`id_scoala`, `id`);

--
-- Constraints for table `note`
--
ALTER TABLE `note`
  ADD CONSTRAINT `note_ibfk_1` FOREIGN KEY (`id_scoala`,`id_clasa`,`id_elev`) REFERENCES `elev` (`id_scoala`, `id_clasa`, `id_elev`),
  ADD CONSTRAINT `note_ibfk_2` FOREIGN KEY (`id_scoala`,`id_clasa`,`nume_disciplina`) REFERENCES `incadrare` (`id_scoala`, `id_clasa`, `nume_disciplina`);

--
-- Constraints for table `profesor`
--
ALTER TABLE `profesor`
  ADD CONSTRAINT `profesor_ibfk_1` FOREIGN KEY (`id_scoala`) REFERENCES `scoala` (`id`),
  ADD CONSTRAINT `profesor_ibfk_2` FOREIGN KEY (`id_cont`) REFERENCES `cont` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
