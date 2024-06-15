-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 15, 2024 at 01:52 PM
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

--
-- Dumping data for table `activitate`
--

INSERT INTO `activitate` (`id_nota`, `id_scoala`, `nume_disciplina`, `id_clasa`, `id_elev`, `valoare`, `data`) VALUES
(1, 1, 'Istorie', '9A', 0, 0, '2024-05-01');

-- --------------------------------------------------------

--
-- Table structure for table `clasa`
--

CREATE TABLE `clasa` (
  `id_scoala` int(11) NOT NULL,
  `nume` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `clasa`
--

INSERT INTO `clasa` (`id_scoala`, `nume`) VALUES
(1, '5A'),
(1, '5B'),
(1, '5C'),
(1, '6A'),
(1, '6B'),
(1, '6C'),
(1, '7A'),
(1, '7B'),
(1, '7C'),
(1, '8A'),
(1, '8B'),
(1, '8C'),
(1, '9A');

-- --------------------------------------------------------

--
-- Table structure for table `cont`
--

CREATE TABLE `cont` (
  `id` int(11) NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  `parola` varchar(100) DEFAULT NULL,
  `nume` varchar(50) DEFAULT NULL,
  `prenume` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `cont`
--

INSERT INTO `cont` (`id`, `email`, `parola`, `nume`, `prenume`) VALUES
(7, 'stefan.mocanu@sss.tv', '$2a$10$4kDqWxUdm6XV4QluXCXBReqTntgbnFP6r0d59FESFUROH2SBaGRZm', 'Mocanu', 'Stefan');

-- --------------------------------------------------------

--
-- Table structure for table `cont_rol`
--

CREATE TABLE `cont_rol` (
  `id_cont` int(11) NOT NULL,
  `id_rol` varchar(30) NOT NULL,
  `id_scoala` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `cont_rol`
--

INSERT INTO `cont_rol` (`id_cont`, `id_rol`, `id_scoala`) VALUES
(7, 'Administrator', 1),
(7, 'Elev', 1),
(7, 'Parinte', 1);

-- --------------------------------------------------------

--
-- Table structure for table `discipline`
--

CREATE TABLE `discipline` (
  `id_scoala` int(11) NOT NULL,
  `nume` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `discipline`
--

INSERT INTO `discipline` (`id_scoala`, `nume`) VALUES
(1, 'Biologie'),
(1, 'Chimie'),
(1, 'Desen'),
(1, 'Engleza'),
(1, 'Fizica'),
(1, 'Franceza'),
(1, 'Geografie'),
(1, 'Istorie'),
(1, 'Limba Romana'),
(1, 'Matematica'),
(1, 'Muzica'),
(1, 'Romana'),
(1, 'Sport');

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

--
-- Dumping data for table `elev`
--

INSERT INTO `elev` (`id_scoala`, `id_clasa`, `id_elev`, `nume`, `prenume`, `gen`, `etnie`, `token_elev`, `token_parinte`, `id_cont_elev`, `id_cont_parinte`) VALUES
(1, '5A', 1, 'Smith', 'John', 'M', 'Rrom', 'Y7bA1RZVU2', '6y58GBNYkk', NULL, NULL),
(1, '5A', 2, 'White', 'Melissa', 'F', 'Roman', 'YHgkUTTXv5', 'oucaHEPqTd', NULL, NULL),
(1, '5A', 3, 'Hall', 'William', 'M', 'Rrom', 'LgbXVEIAi3', 'rg19u1uQfX', NULL, NULL),
(1, '5A', 4, 'Gonzalez', 'Rachel', 'F', 'Ucrainian', '56hNPIr7H8', 'Q04olzvBrt', NULL, NULL),
(1, '5B', 1, 'Jones', 'David', 'M', 'Rrom', 'jJQiiyHYx4', '2ASGbKNYtA', NULL, NULL),
(1, '5B', 2, 'Garcia', 'Michael', 'M', 'Roman', 'Vsnc2z1fxZ', 'gnuWE6pGwg', NULL, NULL),
(1, '5B', 3, 'King', 'Kimberly', 'F', 'Ucrainian', 'PuGxd2SKLd', 'd6Kp6ILd8a', NULL, NULL),
(1, '5B', 4, 'Perez', 'Jessica', 'F', 'Maghiar', '68uxD0mTuX', 'EX9L1ZB5xj', NULL, NULL),
(1, '5C', 1, 'Taylor', 'Matthew', 'M', 'Maghiar', '0UgNwkQKjU', 'TWcwYnMj5h', NULL, NULL),
(1, '5C', 2, 'Rodriguez', 'Amanda', 'F', 'Rrom', 'jjJc1ZMHjl', '70V5AM6KSi', NULL, NULL),
(1, '5C', 3, 'Scott', 'Amanda', 'F', 'Ucrainian', 'uGunsXuU6i', 'uEs6zYlG6w', NULL, NULL),
(1, '5C', 4, 'Campbell', 'Heather', 'F', 'Maghiar', 'XdqC6wc67S', 'tBVwAXysf8', NULL, NULL),
(1, '6A', 1, 'Anderson', 'Jennifer', 'F', 'Roman', 'xiv2eoFGKE', 'pldsPNgoFr', NULL, NULL),
(1, '6A', 2, 'Lewis', 'Melissa', 'F', 'Maghiar', 'PXJHJyYSVU', 'uMZgXLGFJv', NULL, NULL),
(1, '6A', 3, 'Green', 'Megan', 'F', 'Maghiar', 'raEcqT5YQe', '3ImoEIsMGX', NULL, NULL),
(1, '6A', 4, 'Parker', 'Amanda', 'F', 'Roman', 'uSMv7LtzlS', 'HvVDakmY7W', NULL, NULL),
(1, '6B', 1, 'Johnson', 'Emily', 'F', 'Roman', '657dNc6RaM', 'qm8A0XXKJp', NULL, NULL),
(1, '6B', 2, 'Harris', 'Daniel', 'M', 'Maghiar', 'C5kcVgvf6P', 'GgLr2SccFb', NULL, NULL),
(1, '6B', 3, 'Allen', 'Laura', 'F', 'Ucrainian', 'iBXbv6eQJ3', 'yAsjKKueO5', NULL, NULL),
(1, '6B', 4, 'Nelson', 'Brandon', 'M', 'Maghiar', 'nEFZbU62wc', '4raROBl1tB', NULL, NULL),
(1, '6C', 1, 'Davis', 'Sarah', 'F', 'Roman', 'a1ySJK6Kkb', 'tATe0UGtwW', NULL, NULL),
(1, '6C', 2, 'Martinez', 'Michelle', 'F', 'Rrom', 'cEhEHCgv6o', 'NrNkfSFcQD', NULL, NULL),
(1, '6C', 3, 'Wright', 'Jason', 'M', 'Rrom', 'EhkkOPyjtn', '3JbLy8vIrV', NULL, NULL),
(1, '6C', 4, 'Roberts', 'David', 'M', 'Roman', 'VdTuME8jn0', 'WSjJDd0rYd', NULL, NULL),
(1, '7A', 1, 'Miller', 'James', 'M', 'Maghiar', 'BrnvqkrKAg', 'UtKU2iAi3J', NULL, NULL),
(1, '7A', 2, 'Robinson', 'Matthew', 'M', 'Maghiar', 'WQQ7oiJJDO', 'pO1EJhnZOp', NULL, NULL),
(1, '7A', 3, 'Lopez', 'Emily', 'F', 'Maghiar', '3A0ZifOK2Q', 'uWvqQDHhcC', NULL, NULL),
(1, '7A', 4, 'Turner', 'Samantha', 'F', 'Rrom', 'g0IsTTc27O', 'gSXoLZxgNu', NULL, NULL),
(1, '7B', 1, 'Thomas', 'Christopher', 'M', 'Rrom', 'dHbAwSaLhX', 'xmIFKke6tj', NULL, NULL),
(1, '7B', 2, 'Lee', 'Joshua', 'M', 'Ucrainian', 'vbdekSNWrS', 'nv9lkOf9dc', NULL, NULL),
(1, '7B', 3, 'Adams', 'Nicholas', 'M', 'Roman', 'QxLhA9jufa', 'XpAOnvktDr', NULL, NULL),
(1, '7B', 4, 'Evans', 'Daniel', 'M', 'Rrom', 'BOsY1DAfUK', '0Sd346rGle', NULL, NULL),
(1, '7C', 1, 'Williams', 'Michael', 'M', 'Ucrainian', 'LlPkA93nDk', 'ZM94iiLV4S', NULL, NULL),
(1, '7C', 2, 'Martin', 'Lisa', 'F', 'Rrom', 'zDq7Q7C4wn', 'T1JCRBdy5o', NULL, NULL),
(1, '7C', 3, 'Young', 'Daniel', 'M', 'Maghiar', 'qJwsa3QEvv', 'qEoeX2k4XS', NULL, NULL),
(1, '7C', 4, 'Carter', 'Stephanie', 'F', 'Roman', 'cPWXWx0Q0a', 'MWC477ADxn', NULL, NULL),
(1, '8A', 1, 'Brown', 'Jessica', 'F', 'Maghiar', 'IrnhawLNUJ', 'jc3pJzVmuJ', NULL, NULL),
(1, '8A', 2, 'Thompson', 'Kimberly', 'F', 'Ucrainian', 'SdWpI4dq3N', 'U2fYZMTGZn', NULL, NULL),
(1, '8A', 3, 'Hernandez', 'Ashley', 'F', 'Roman', 'mSCzOHQhYN', 'EpMQa232oE', NULL, NULL),
(1, '8A', 4, 'Mitchell', 'Andrew', 'M', 'Rrom', 'EWPBGEB3ZO', '8UFzmA31uj', NULL, NULL),
(1, '8B', 1, 'Wilson', 'Ashley', 'F', 'Ucrainian', 'bfJUQp7qAk', 'dxR372GhOC', NULL, NULL),
(1, '8B', 2, 'Clark', 'Angela', 'F', 'Ucrainian', 'byS55aF4uG', 'BPC7drN9uI', NULL, NULL),
(1, '8B', 3, 'Hill', 'Christopher', 'M', 'Roman', 'IjQhyBwbGJ', 'LQTZ3ilXH0', NULL, NULL),
(1, '8B', 4, 'Phillips', 'Joshua', 'M', 'Ucrainian', 'wn5AVeLpzJ', '0wXSX0eD8r', NULL, NULL),
(1, '8C', 1, 'Jackson', 'Amanda', 'F', 'Ucrainian', 'dglX4zJIrg', '1vKhDlDsrB', NULL, NULL),
(1, '8C', 2, 'Walker', 'Sarah', 'F', 'Roman', 'NT2c7mPNcn', '5nUiODXZNw', NULL, NULL),
(1, '8C', 3, 'Baker', 'Sarah', 'F', 'Rrom', 'cP0H5MXvNp', 'eouQVRHliL', NULL, NULL),
(1, '8C', 4, 'Edwards', 'Melissa', 'F', 'Ucrainian', '7qhIxdFqMv', 'ysLdtQxZlO', NULL, NULL),
(1, '9A', 0, 'Mocanu', 'Stefan', 'M', 'Roman', NULL, NULL, 7, 7);

-- --------------------------------------------------------

--
-- Table structure for table `feedback`
--

CREATE TABLE `feedback` (
  `id_feedback` int(11) NOT NULL,
  `id_scoala` int(11) DEFAULT NULL,
  `nume_disciplina` varchar(100) DEFAULT NULL,
  `id_clasa` varchar(50) DEFAULT NULL,
  `id_elev` int(11) DEFAULT NULL,
  `content` mediumtext DEFAULT NULL,
  `data` date DEFAULT NULL,
  `tip` tinyint(1) DEFAULT NULL,
  `directie` tinyint(1) DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `feedback`
--

INSERT INTO `feedback` (`id_feedback`, `id_scoala`, `nume_disciplina`, `id_clasa`, `id_elev`, `content`, `data`, `tip`, `directie`) VALUES
(1, 1, 'Istorie', '9A', 0, 'A venit nepregatit la ora.', '2024-05-03', NULL, 0);

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

--
-- Dumping data for table `incadrare`
--

INSERT INTO `incadrare` (`id_scoala`, `id_clasa`, `id_profesor`, `nume_disciplina`) VALUES
(1, '9A', 3, 'Istorie'),
(1, '5A', 4, 'Matematica'),
(1, '5A', 5, 'Franceza'),
(1, '5A', 6, 'Limba Romana'),
(1, '5A', 7, 'Engleza'),
(1, '6A', 7, 'Engleza'),
(1, '6C', 7, 'Engleza'),
(1, '7B', 7, 'Engleza'),
(1, '8B', 7, 'Engleza'),
(1, '5A', 8, 'Istorie'),
(1, '6A', 8, 'Istorie'),
(1, '6C', 8, 'Istorie'),
(1, '7C', 8, 'Istorie'),
(1, '8B', 8, 'Istorie'),
(1, '5B', 9, 'Geografie'),
(1, '6A', 9, 'Geografie'),
(1, '6C', 9, 'Geografie'),
(1, '7C', 9, 'Geografie'),
(1, '8B', 9, 'Geografie'),
(1, '5B', 10, 'Sport'),
(1, '6A', 10, 'Sport'),
(1, '6C', 10, 'Sport'),
(1, '7C', 10, 'Sport'),
(1, '8B', 10, 'Sport'),
(1, '5B', 11, 'Muzica'),
(1, '6A', 11, 'Muzica'),
(1, '7A', 11, 'Muzica'),
(1, '7C', 11, 'Muzica'),
(1, '8C', 11, 'Muzica'),
(1, '5B', 12, 'Desen'),
(1, '6B', 12, 'Desen'),
(1, '7A', 12, 'Desen'),
(1, '7C', 12, 'Desen'),
(1, '8C', 12, 'Desen'),
(1, '5B', 13, 'Fizica'),
(1, '6B', 13, 'Fizica'),
(1, '7A', 13, 'Fizica'),
(1, '7C', 13, 'Fizica'),
(1, '8A', 13, 'Fizica'),
(1, '8C', 13, 'Fizica'),
(1, '5C', 14, 'Chimie'),
(1, '6B', 14, 'Chimie'),
(1, '7A', 14, 'Chimie'),
(1, '8A', 14, 'Chimie'),
(1, '5C', 15, 'Biologie'),
(1, '6B', 15, 'Biologie'),
(1, '7B', 15, 'Biologie'),
(1, '8A', 15, 'Biologie'),
(1, '5C', 16, 'Istorie'),
(1, '7B', 16, 'Limba Romana'),
(1, '8A', 16, 'Limba Romana'),
(1, '5C', 17, 'Matematica'),
(1, '6B', 17, 'Franceza'),
(1, '7B', 17, 'Franceza'),
(1, '8A', 17, 'Franceza'),
(1, '5C', 18, 'Limba Romana'),
(1, '6C', 18, 'Limba Romana'),
(1, '7B', 18, 'Matematica'),
(1, '8B', 18, 'Limba Romana');

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

--
-- Dumping data for table `note`
--

INSERT INTO `note` (`id_nota`, `id_scoala`, `nume_disciplina`, `id_clasa`, `id_elev`, `nota`, `data`) VALUES
(2, 1, 'Istorie', '9A', 0, 10, '2024-05-09'),
(4, 1, 'Engleza', '5A', 1, 2, '2024-05-16'),
(5, 1, 'Engleza', '5A', 1, 4, '2023-09-01'),
(6, 1, 'Franceza', '5A', 1, 9, '2024-05-25'),
(7, 1, 'Franceza', '5A', 1, 7, '2023-06-28'),
(8, 1, 'Franceza', '5A', 1, 7, '2023-08-15'),
(9, 1, 'Franceza', '5A', 1, 4, '2023-06-17'),
(10, 1, 'Franceza', '5A', 1, 4, '2024-01-18'),
(11, 1, 'Istorie', '5A', 1, 1, '2023-11-24'),
(12, 1, 'Istorie', '5A', 1, 3, '2023-09-20'),
(13, 1, 'Istorie', '5A', 1, 4, '2024-02-19'),
(14, 1, 'Istorie', '5A', 1, 5, '2023-11-24'),
(15, 1, 'Limba Romana', '5A', 1, 3, '2024-05-16'),
(16, 1, 'Limba Romana', '5A', 1, 7, '2023-11-09'),
(17, 1, 'Limba Romana', '5A', 1, 8, '2024-05-10'),
(18, 1, 'Matematica', '5A', 1, 7, '2024-02-01'),
(19, 1, 'Matematica', '5A', 1, 10, '2023-11-05'),
(20, 1, 'Matematica', '5A', 1, 8, '2023-09-24'),
(21, 1, 'Matematica', '5A', 1, 6, '2023-06-27'),
(22, 1, 'Matematica', '5A', 1, 5, '2023-06-04'),
(23, 1, 'Engleza', '5A', 2, 9, '2024-03-25'),
(24, 1, 'Engleza', '5A', 2, 5, '2024-05-07'),
(25, 1, 'Engleza', '5A', 2, 9, '2023-07-01'),
(26, 1, 'Engleza', '5A', 2, 6, '2023-09-12'),
(27, 1, 'Engleza', '5A', 2, 10, '2023-10-19'),
(28, 1, 'Franceza', '5A', 2, 5, '2023-12-09'),
(29, 1, 'Franceza', '5A', 2, 4, '2023-08-16'),
(30, 1, 'Franceza', '5A', 2, 10, '2024-01-09'),
(31, 1, 'Franceza', '5A', 2, 9, '2023-10-27'),
(32, 1, 'Istorie', '5A', 2, 1, '2024-01-06'),
(33, 1, 'Istorie', '5A', 2, 4, '2023-09-12'),
(34, 1, 'Istorie', '5A', 2, 3, '2023-11-26'),
(35, 1, 'Istorie', '5A', 2, 4, '2024-02-19'),
(36, 1, 'Limba Romana', '5A', 2, 7, '2023-11-19'),
(37, 1, 'Limba Romana', '5A', 2, 5, '2023-11-07'),
(38, 1, 'Matematica', '5A', 2, 4, '2024-03-08'),
(39, 1, 'Matematica', '5A', 2, 5, '2023-12-19'),
(40, 1, 'Matematica', '5A', 2, 5, '2024-01-07'),
(41, 1, 'Matematica', '5A', 2, 8, '2023-08-15'),
(42, 1, 'Engleza', '5A', 3, 1, '2023-10-12'),
(43, 1, 'Engleza', '5A', 3, 9, '2023-06-04'),
(44, 1, 'Engleza', '5A', 3, 6, '2023-07-02'),
(45, 1, 'Engleza', '5A', 3, 3, '2023-09-13'),
(46, 1, 'Franceza', '5A', 3, 6, '2023-08-01'),
(47, 1, 'Franceza', '5A', 3, 2, '2023-12-07'),
(48, 1, 'Franceza', '5A', 3, 4, '2023-11-23'),
(49, 1, 'Istorie', '5A', 3, 7, '2024-03-18'),
(50, 1, 'Istorie', '5A', 3, 1, '2024-04-24'),
(51, 1, 'Istorie', '5A', 3, 9, '2023-09-24'),
(52, 1, 'Istorie', '5A', 3, 9, '2023-10-20'),
(53, 1, 'Limba Romana', '5A', 3, 10, '2024-01-26'),
(54, 1, 'Limba Romana', '5A', 3, 3, '2024-04-27'),
(55, 1, 'Limba Romana', '5A', 3, 10, '2024-03-25'),
(56, 1, 'Limba Romana', '5A', 3, 4, '2023-12-09'),
(57, 1, 'Limba Romana', '5A', 3, 10, '2023-06-19'),
(58, 1, 'Matematica', '5A', 3, 4, '2023-07-05'),
(59, 1, 'Matematica', '5A', 3, 9, '2023-10-28'),
(60, 1, 'Matematica', '5A', 3, 7, '2024-02-07'),
(61, 1, 'Matematica', '5A', 3, 4, '2024-01-02'),
(62, 1, 'Engleza', '5A', 4, 10, '2024-05-21'),
(63, 1, 'Engleza', '5A', 4, 8, '2024-05-02'),
(64, 1, 'Franceza', '5A', 4, 6, '2023-12-22'),
(65, 1, 'Franceza', '5A', 4, 2, '2023-11-07'),
(66, 1, 'Franceza', '5A', 4, 8, '2024-05-16'),
(67, 1, 'Franceza', '5A', 4, 10, '2024-02-12'),
(68, 1, 'Istorie', '5A', 4, 3, '2023-09-06'),
(69, 1, 'Istorie', '5A', 4, 8, '2024-01-28'),
(70, 1, 'Limba Romana', '5A', 4, 9, '2023-12-15'),
(71, 1, 'Limba Romana', '5A', 4, 8, '2023-06-09'),
(72, 1, 'Limba Romana', '5A', 4, 10, '2023-09-23'),
(73, 1, 'Matematica', '5A', 4, 3, '2024-01-23'),
(74, 1, 'Matematica', '5A', 4, 10, '2024-01-28'),
(75, 1, 'Matematica', '5A', 4, 6, '2023-10-02'),
(76, 1, 'Matematica', '5A', 4, 8, '2023-06-04'),
(77, 1, 'Desen', '5B', 4, 8, '2023-11-01'),
(78, 1, 'Desen', '5B', 4, 4, '2024-04-10'),
(79, 1, 'Fizica', '5B', 4, 4, '2023-10-27'),
(80, 1, 'Fizica', '5B', 4, 8, '2024-01-18'),
(81, 1, 'Fizica', '5B', 4, 8, '2023-09-04'),
(82, 1, 'Geografie', '5B', 4, 7, '2024-03-14'),
(83, 1, 'Geografie', '5B', 4, 8, '2023-06-20'),
(84, 1, 'Muzica', '5B', 4, 7, '2023-10-06'),
(85, 1, 'Muzica', '5B', 4, 2, '2024-03-24'),
(86, 1, 'Muzica', '5B', 4, 1, '2023-11-21'),
(87, 1, 'Sport', '5B', 4, 3, '2023-12-21'),
(88, 1, 'Sport', '5B', 4, 5, '2023-12-01'),
(89, 1, 'Sport', '5B', 4, 5, '2023-09-22'),
(90, 1, 'Sport', '5B', 4, 7, '2024-04-08'),
(91, 1, 'Desen', '5B', 3, 1, '2023-10-19'),
(92, 1, 'Desen', '5B', 3, 5, '2024-03-25'),
(93, 1, 'Desen', '5B', 3, 3, '2023-07-19'),
(94, 1, 'Desen', '5B', 3, 9, '2023-10-03'),
(95, 1, 'Fizica', '5B', 3, 1, '2023-08-26'),
(96, 1, 'Fizica', '5B', 3, 9, '2023-07-18'),
(97, 1, 'Fizica', '5B', 3, 9, '2024-05-04'),
(98, 1, 'Geografie', '5B', 3, 8, '2024-05-01'),
(99, 1, 'Geografie', '5B', 3, 8, '2023-06-07'),
(100, 1, 'Geografie', '5B', 3, 5, '2024-04-07'),
(101, 1, 'Muzica', '5B', 3, 4, '2024-04-15'),
(102, 1, 'Muzica', '5B', 3, 8, '2023-12-19'),
(103, 1, 'Muzica', '5B', 3, 6, '2023-10-12'),
(104, 1, 'Muzica', '5B', 3, 2, '2023-09-10'),
(105, 1, 'Sport', '5B', 3, 2, '2024-03-09'),
(106, 1, 'Sport', '5B', 3, 4, '2023-09-23'),
(107, 1, 'Desen', '5B', 2, 10, '2024-03-28'),
(108, 1, 'Desen', '5B', 2, 6, '2023-06-26'),
(109, 1, 'Desen', '5B', 2, 4, '2024-05-15'),
(110, 1, 'Desen', '5B', 2, 7, '2023-06-15'),
(111, 1, 'Desen', '5B', 2, 7, '2023-10-05'),
(112, 1, 'Fizica', '5B', 2, 7, '2023-11-22'),
(113, 1, 'Fizica', '5B', 2, 10, '2024-05-01'),
(114, 1, 'Fizica', '5B', 2, 2, '2023-06-03'),
(115, 1, 'Geografie', '5B', 2, 10, '2024-04-14'),
(116, 1, 'Geografie', '5B', 2, 10, '2023-12-24'),
(117, 1, 'Geografie', '5B', 2, 9, '2024-02-23'),
(118, 1, 'Geografie', '5B', 2, 10, '2024-01-22'),
(119, 1, 'Geografie', '5B', 2, 8, '2024-03-15'),
(120, 1, 'Muzica', '5B', 2, 8, '2023-10-16'),
(121, 1, 'Muzica', '5B', 2, 5, '2023-10-20'),
(122, 1, 'Muzica', '5B', 2, 4, '2024-03-12'),
(123, 1, 'Muzica', '5B', 2, 6, '2023-11-06'),
(124, 1, 'Muzica', '5B', 2, 5, '2023-06-09'),
(125, 1, 'Sport', '5B', 2, 1, '2024-01-04'),
(126, 1, 'Sport', '5B', 2, 10, '2024-04-07'),
(127, 1, 'Sport', '5B', 2, 3, '2023-12-09'),
(128, 1, 'Desen', '5B', 1, 10, '2024-02-25'),
(129, 1, 'Desen', '5B', 1, 1, '2024-05-01'),
(130, 1, 'Desen', '5B', 1, 10, '2023-12-07'),
(131, 1, 'Desen', '5B', 1, 3, '2023-11-14'),
(132, 1, 'Desen', '5B', 1, 6, '2024-03-21'),
(133, 1, 'Fizica', '5B', 1, 8, '2024-04-21'),
(134, 1, 'Fizica', '5B', 1, 7, '2023-08-02'),
(135, 1, 'Fizica', '5B', 1, 7, '2024-03-18'),
(136, 1, 'Geografie', '5B', 1, 7, '2024-05-11'),
(137, 1, 'Geografie', '5B', 1, 9, '2024-04-03'),
(138, 1, 'Geografie', '5B', 1, 6, '2023-08-28'),
(139, 1, 'Geografie', '5B', 1, 4, '2024-02-21'),
(140, 1, 'Muzica', '5B', 1, 4, '2023-10-02'),
(141, 1, 'Muzica', '5B', 1, 7, '2023-09-07'),
(142, 1, 'Muzica', '5B', 1, 8, '2023-08-28'),
(143, 1, 'Muzica', '5B', 1, 1, '2023-06-10'),
(144, 1, 'Sport', '5B', 1, 5, '2023-10-10'),
(145, 1, 'Sport', '5B', 1, 10, '2023-06-22'),
(146, 1, 'Sport', '5B', 1, 3, '2024-04-14'),
(147, 1, 'Sport', '5B', 1, 9, '2023-08-10'),
(148, 1, 'Biologie', '5C', 1, 3, '2024-03-26'),
(149, 1, 'Biologie', '5C', 1, 9, '2023-09-12'),
(150, 1, 'Biologie', '5C', 1, 10, '2023-10-28'),
(151, 1, 'Biologie', '5C', 1, 3, '2024-02-02'),
(152, 1, 'Chimie', '5C', 1, 3, '2023-10-04'),
(153, 1, 'Chimie', '5C', 1, 5, '2023-08-05'),
(154, 1, 'Chimie', '5C', 1, 6, '2023-09-27'),
(155, 1, 'Istorie', '5C', 1, 7, '2024-01-18'),
(156, 1, 'Istorie', '5C', 1, 4, '2024-02-21'),
(157, 1, 'Istorie', '5C', 1, 9, '2024-03-23'),
(158, 1, 'Istorie', '5C', 1, 9, '2024-05-24'),
(159, 1, 'Limba Romana', '5C', 1, 5, '2023-12-01'),
(160, 1, 'Limba Romana', '5C', 1, 4, '2023-10-23'),
(161, 1, 'Limba Romana', '5C', 1, 6, '2023-11-08'),
(162, 1, 'Matematica', '5C', 1, 7, '2023-09-16'),
(163, 1, 'Matematica', '5C', 1, 3, '2023-08-01'),
(164, 1, 'Matematica', '5C', 1, 1, '2023-12-08'),
(165, 1, 'Biologie', '5C', 2, 5, '2023-08-10'),
(166, 1, 'Biologie', '5C', 2, 1, '2023-09-12'),
(167, 1, 'Biologie', '5C', 2, 4, '2023-06-17'),
(168, 1, 'Biologie', '5C', 2, 1, '2024-03-05'),
(169, 1, 'Biologie', '5C', 2, 10, '2023-09-04'),
(170, 1, 'Chimie', '5C', 2, 1, '2023-12-21'),
(171, 1, 'Chimie', '5C', 2, 2, '2023-12-18'),
(172, 1, 'Chimie', '5C', 2, 4, '2023-10-03'),
(173, 1, 'Chimie', '5C', 2, 8, '2024-02-08'),
(174, 1, 'Chimie', '5C', 2, 4, '2023-10-01'),
(175, 1, 'Istorie', '5C', 2, 7, '2023-12-12'),
(176, 1, 'Istorie', '5C', 2, 2, '2023-07-26'),
(177, 1, 'Istorie', '5C', 2, 6, '2023-06-15'),
(178, 1, 'Istorie', '5C', 2, 9, '2024-01-12'),
(179, 1, 'Limba Romana', '5C', 2, 6, '2023-12-25'),
(180, 1, 'Limba Romana', '5C', 2, 1, '2023-07-01'),
(181, 1, 'Limba Romana', '5C', 2, 7, '2024-01-10'),
(182, 1, 'Limba Romana', '5C', 2, 10, '2023-09-06'),
(183, 1, 'Matematica', '5C', 2, 4, '2023-08-15'),
(184, 1, 'Matematica', '5C', 2, 9, '2024-02-24'),
(185, 1, 'Matematica', '5C', 2, 8, '2023-10-05'),
(186, 1, 'Matematica', '5C', 2, 1, '2023-08-26'),
(187, 1, 'Biologie', '5C', 3, 10, '2024-05-16'),
(188, 1, 'Biologie', '5C', 3, 10, '2024-04-14'),
(189, 1, 'Chimie', '5C', 3, 10, '2023-08-22'),
(190, 1, 'Chimie', '5C', 3, 10, '2024-03-07'),
(191, 1, 'Istorie', '5C', 3, 6, '2024-03-16'),
(192, 1, 'Istorie', '5C', 3, 4, '2023-09-18'),
(193, 1, 'Istorie', '5C', 3, 9, '2023-11-16'),
(194, 1, 'Istorie', '5C', 3, 2, '2023-11-06'),
(195, 1, 'Limba Romana', '5C', 3, 3, '2023-12-10'),
(196, 1, 'Limba Romana', '5C', 3, 4, '2024-04-12'),
(197, 1, 'Limba Romana', '5C', 3, 9, '2023-10-04'),
(198, 1, 'Limba Romana', '5C', 3, 10, '2024-03-13'),
(199, 1, 'Limba Romana', '5C', 3, 10, '2023-09-11'),
(200, 1, 'Matematica', '5C', 3, 5, '2024-03-10'),
(201, 1, 'Matematica', '5C', 3, 1, '2024-01-04'),
(202, 1, 'Biologie', '5C', 4, 6, '2023-07-23'),
(203, 1, 'Biologie', '5C', 4, 7, '2023-09-08'),
(204, 1, 'Biologie', '5C', 4, 6, '2023-11-19'),
(205, 1, 'Biologie', '5C', 4, 6, '2023-12-20'),
(206, 1, 'Biologie', '5C', 4, 7, '2024-01-04'),
(207, 1, 'Chimie', '5C', 4, 1, '2024-01-12'),
(208, 1, 'Chimie', '5C', 4, 2, '2023-11-24'),
(209, 1, 'Istorie', '5C', 4, 1, '2024-05-24'),
(210, 1, 'Istorie', '5C', 4, 10, '2023-10-22'),
(211, 1, 'Istorie', '5C', 4, 9, '2024-02-09'),
(212, 1, 'Istorie', '5C', 4, 2, '2023-12-24'),
(213, 1, 'Limba Romana', '5C', 4, 6, '2024-05-13'),
(214, 1, 'Limba Romana', '5C', 4, 5, '2023-10-06'),
(215, 1, 'Limba Romana', '5C', 4, 6, '2024-01-26'),
(216, 1, 'Limba Romana', '5C', 4, 6, '2023-11-09'),
(217, 1, 'Matematica', '5C', 4, 9, '2023-06-15'),
(218, 1, 'Matematica', '5C', 4, 7, '2024-05-24'),
(219, 1, 'Matematica', '5C', 4, 3, '2023-06-24'),
(220, 1, 'Matematica', '5C', 4, 3, '2024-04-25'),
(221, 1, 'Matematica', '5C', 4, 8, '2023-12-23'),
(222, 1, 'Engleza', '6A', 4, 2, '2023-11-13'),
(223, 1, 'Engleza', '6A', 4, 9, '2023-06-14'),
(224, 1, 'Engleza', '6A', 4, 8, '2024-03-07'),
(225, 1, 'Geografie', '6A', 4, 2, '2024-03-08'),
(226, 1, 'Geografie', '6A', 4, 5, '2023-06-24'),
(227, 1, 'Geografie', '6A', 4, 7, '2023-12-04'),
(228, 1, 'Geografie', '6A', 4, 4, '2023-12-28'),
(229, 1, 'Istorie', '6A', 4, 10, '2023-07-25'),
(230, 1, 'Istorie', '6A', 4, 2, '2024-02-05'),
(231, 1, 'Muzica', '6A', 4, 5, '2023-10-13'),
(232, 1, 'Muzica', '6A', 4, 8, '2023-07-13'),
(233, 1, 'Sport', '6A', 4, 10, '2023-10-07'),
(234, 1, 'Sport', '6A', 4, 3, '2023-11-06'),
(235, 1, 'Sport', '6A', 4, 9, '2023-06-23'),
(236, 1, 'Sport', '6A', 4, 5, '2023-09-12'),
(237, 1, 'Sport', '6A', 4, 4, '2024-01-19'),
(238, 1, 'Engleza', '6A', 3, 3, '2023-06-27'),
(239, 1, 'Engleza', '6A', 3, 9, '2023-09-19'),
(240, 1, 'Engleza', '6A', 3, 7, '2023-10-25'),
(241, 1, 'Engleza', '6A', 3, 8, '2023-06-12'),
(242, 1, 'Engleza', '6A', 3, 6, '2024-03-25'),
(243, 1, 'Geografie', '6A', 3, 6, '2023-09-04'),
(244, 1, 'Geografie', '6A', 3, 10, '2024-05-05'),
(245, 1, 'Geografie', '6A', 3, 6, '2024-01-03'),
(246, 1, 'Istorie', '6A', 3, 9, '2023-10-23'),
(247, 1, 'Istorie', '6A', 3, 7, '2023-10-04'),
(248, 1, 'Istorie', '6A', 3, 6, '2024-01-28'),
(249, 1, 'Istorie', '6A', 3, 8, '2023-11-14'),
(250, 1, 'Muzica', '6A', 3, 3, '2023-11-07'),
(251, 1, 'Muzica', '6A', 3, 10, '2023-06-28'),
(252, 1, 'Sport', '6A', 3, 6, '2023-08-21'),
(253, 1, 'Sport', '6A', 3, 8, '2023-11-10'),
(254, 1, 'Sport', '6A', 3, 10, '2024-04-01'),
(255, 1, 'Sport', '6A', 3, 10, '2024-03-28'),
(256, 1, 'Engleza', '6A', 2, 4, '2023-08-28'),
(257, 1, 'Engleza', '6A', 2, 1, '2024-02-20'),
(258, 1, 'Geografie', '6A', 2, 10, '2023-11-16'),
(259, 1, 'Geografie', '6A', 2, 3, '2023-09-09'),
(260, 1, 'Geografie', '6A', 2, 9, '2024-05-22'),
(261, 1, 'Geografie', '6A', 2, 2, '2023-08-02'),
(262, 1, 'Istorie', '6A', 2, 9, '2024-04-25'),
(263, 1, 'Istorie', '6A', 2, 7, '2023-06-11'),
(264, 1, 'Istorie', '6A', 2, 1, '2024-04-25'),
(265, 1, 'Istorie', '6A', 2, 2, '2023-09-12'),
(266, 1, 'Muzica', '6A', 2, 7, '2024-05-26'),
(267, 1, 'Muzica', '6A', 2, 6, '2023-09-02'),
(268, 1, 'Muzica', '6A', 2, 8, '2023-10-06'),
(269, 1, 'Muzica', '6A', 2, 8, '2024-01-19'),
(270, 1, 'Muzica', '6A', 2, 6, '2023-08-27'),
(271, 1, 'Sport', '6A', 2, 2, '2024-05-19'),
(272, 1, 'Sport', '6A', 2, 4, '2024-04-19'),
(273, 1, 'Sport', '6A', 2, 5, '2023-08-25'),
(274, 1, 'Engleza', '6A', 1, 1, '2023-10-14'),
(275, 1, 'Engleza', '6A', 1, 4, '2023-07-21'),
(276, 1, 'Engleza', '6A', 1, 1, '2023-10-16'),
(277, 1, 'Engleza', '6A', 1, 10, '2024-05-05'),
(278, 1, 'Engleza', '6A', 1, 6, '2023-11-26'),
(279, 1, 'Geografie', '6A', 1, 5, '2024-04-10'),
(280, 1, 'Geografie', '6A', 1, 5, '2023-09-14'),
(281, 1, 'Geografie', '6A', 1, 8, '2024-01-23'),
(282, 1, 'Geografie', '6A', 1, 9, '2023-11-02'),
(283, 1, 'Istorie', '6A', 1, 9, '2023-10-03'),
(284, 1, 'Istorie', '6A', 1, 4, '2024-02-26'),
(285, 1, 'Istorie', '6A', 1, 5, '2023-09-28'),
(286, 1, 'Istorie', '6A', 1, 2, '2023-11-02'),
(287, 1, 'Istorie', '6A', 1, 6, '2023-06-03'),
(288, 1, 'Muzica', '6A', 1, 8, '2024-02-14'),
(289, 1, 'Muzica', '6A', 1, 10, '2024-03-14'),
(290, 1, 'Muzica', '6A', 1, 4, '2024-03-07'),
(291, 1, 'Muzica', '6A', 1, 7, '2024-01-06'),
(292, 1, 'Muzica', '6A', 1, 5, '2023-06-20'),
(293, 1, 'Sport', '6A', 1, 7, '2023-07-25'),
(294, 1, 'Sport', '6A', 1, 2, '2023-08-01'),
(295, 1, 'Biologie', '6B', 1, 2, '2023-09-07'),
(296, 1, 'Biologie', '6B', 1, 9, '2023-08-23'),
(297, 1, 'Biologie', '6B', 1, 2, '2023-10-27'),
(298, 1, 'Biologie', '6B', 1, 5, '2024-03-15'),
(299, 1, 'Chimie', '6B', 1, 3, '2023-10-13'),
(300, 1, 'Chimie', '6B', 1, 10, '2024-04-18'),
(301, 1, 'Chimie', '6B', 1, 6, '2023-07-16'),
(302, 1, 'Chimie', '6B', 1, 10, '2024-05-09'),
(303, 1, 'Chimie', '6B', 1, 2, '2023-06-06'),
(304, 1, 'Desen', '6B', 1, 5, '2023-08-05'),
(305, 1, 'Desen', '6B', 1, 1, '2023-08-14'),
(306, 1, 'Desen', '6B', 1, 3, '2024-05-16'),
(307, 1, 'Desen', '6B', 1, 2, '2023-06-11'),
(308, 1, 'Desen', '6B', 1, 5, '2023-11-01'),
(309, 1, 'Fizica', '6B', 1, 4, '2024-04-22'),
(310, 1, 'Fizica', '6B', 1, 9, '2024-04-27'),
(311, 1, 'Fizica', '6B', 1, 8, '2023-12-23'),
(312, 1, 'Franceza', '6B', 1, 9, '2024-04-18'),
(313, 1, 'Franceza', '6B', 1, 1, '2024-02-04'),
(314, 1, 'Franceza', '6B', 1, 3, '2023-10-01'),
(315, 1, 'Biologie', '6B', 2, 1, '2023-11-23'),
(316, 1, 'Biologie', '6B', 2, 7, '2024-04-11'),
(317, 1, 'Biologie', '6B', 2, 10, '2023-10-13'),
(318, 1, 'Chimie', '6B', 2, 7, '2023-12-06'),
(319, 1, 'Chimie', '6B', 2, 9, '2023-12-13'),
(320, 1, 'Desen', '6B', 2, 2, '2024-01-20'),
(321, 1, 'Desen', '6B', 2, 1, '2023-06-20'),
(322, 1, 'Fizica', '6B', 2, 2, '2023-07-07'),
(323, 1, 'Fizica', '6B', 2, 6, '2023-06-08'),
(324, 1, 'Fizica', '6B', 2, 3, '2023-06-23'),
(325, 1, 'Fizica', '6B', 2, 10, '2023-10-28'),
(326, 1, 'Franceza', '6B', 2, 8, '2023-11-22'),
(327, 1, 'Franceza', '6B', 2, 7, '2023-11-28'),
(328, 1, 'Biologie', '6B', 3, 9, '2023-07-13'),
(329, 1, 'Biologie', '6B', 3, 4, '2024-02-11'),
(330, 1, 'Biologie', '6B', 3, 2, '2023-11-14'),
(331, 1, 'Chimie', '6B', 3, 8, '2024-04-19'),
(332, 1, 'Chimie', '6B', 3, 8, '2024-02-05'),
(333, 1, 'Chimie', '6B', 3, 7, '2023-08-23'),
(334, 1, 'Chimie', '6B', 3, 10, '2023-12-18'),
(335, 1, 'Chimie', '6B', 3, 4, '2024-03-24'),
(336, 1, 'Desen', '6B', 3, 4, '2023-06-17'),
(337, 1, 'Desen', '6B', 3, 10, '2024-01-18'),
(338, 1, 'Fizica', '6B', 3, 8, '2024-01-27'),
(339, 1, 'Fizica', '6B', 3, 3, '2023-12-07'),
(340, 1, 'Fizica', '6B', 3, 2, '2024-01-26'),
(341, 1, 'Fizica', '6B', 3, 2, '2024-02-02'),
(342, 1, 'Franceza', '6B', 3, 4, '2024-04-13'),
(343, 1, 'Franceza', '6B', 3, 5, '2023-06-17'),
(344, 1, 'Franceza', '6B', 3, 4, '2023-10-13'),
(345, 1, 'Franceza', '6B', 3, 6, '2024-04-01'),
(346, 1, 'Biologie', '6B', 4, 6, '2023-11-12'),
(347, 1, 'Biologie', '6B', 4, 5, '2024-03-27'),
(348, 1, 'Biologie', '6B', 4, 9, '2023-11-03'),
(349, 1, 'Chimie', '6B', 4, 6, '2023-08-25'),
(350, 1, 'Chimie', '6B', 4, 7, '2023-06-25'),
(351, 1, 'Desen', '6B', 4, 6, '2024-01-09'),
(352, 1, 'Desen', '6B', 4, 4, '2024-04-10'),
(353, 1, 'Fizica', '6B', 4, 1, '2023-11-27'),
(354, 1, 'Fizica', '6B', 4, 8, '2024-01-03'),
(355, 1, 'Fizica', '6B', 4, 9, '2023-11-09'),
(356, 1, 'Fizica', '6B', 4, 10, '2023-10-07'),
(357, 1, 'Fizica', '6B', 4, 5, '2023-09-26'),
(358, 1, 'Franceza', '6B', 4, 6, '2024-02-26'),
(359, 1, 'Franceza', '6B', 4, 9, '2023-06-06'),
(360, 1, 'Engleza', '6C', 4, 4, '2024-05-27'),
(361, 1, 'Engleza', '6C', 4, 9, '2023-10-18'),
(362, 1, 'Engleza', '6C', 4, 10, '2023-11-06'),
(363, 1, 'Engleza', '6C', 4, 10, '2023-12-12'),
(364, 1, 'Geografie', '6C', 4, 1, '2024-05-17'),
(365, 1, 'Geografie', '6C', 4, 5, '2023-12-18'),
(366, 1, 'Geografie', '6C', 4, 2, '2024-04-04'),
(367, 1, 'Geografie', '6C', 4, 6, '2023-10-25'),
(368, 1, 'Istorie', '6C', 4, 4, '2024-04-05'),
(369, 1, 'Istorie', '6C', 4, 10, '2023-12-02'),
(370, 1, 'Istorie', '6C', 4, 2, '2023-09-11'),
(371, 1, 'Limba Romana', '6C', 4, 5, '2024-04-10'),
(372, 1, 'Limba Romana', '6C', 4, 8, '2023-11-24'),
(373, 1, 'Sport', '6C', 4, 9, '2023-06-04'),
(374, 1, 'Sport', '6C', 4, 2, '2024-02-25'),
(375, 1, 'Sport', '6C', 4, 2, '2023-12-23'),
(376, 1, 'Engleza', '6C', 3, 8, '2024-05-20'),
(377, 1, 'Engleza', '6C', 3, 7, '2023-12-14'),
(378, 1, 'Engleza', '6C', 3, 6, '2024-01-02'),
(379, 1, 'Engleza', '6C', 3, 7, '2023-12-09'),
(380, 1, 'Engleza', '6C', 3, 1, '2024-01-10'),
(381, 1, 'Geografie', '6C', 3, 2, '2023-12-07'),
(382, 1, 'Geografie', '6C', 3, 2, '2023-11-13'),
(383, 1, 'Geografie', '6C', 3, 3, '2023-06-16'),
(384, 1, 'Geografie', '6C', 3, 6, '2024-01-28'),
(385, 1, 'Geografie', '6C', 3, 6, '2023-10-05'),
(386, 1, 'Istorie', '6C', 3, 7, '2024-04-08'),
(387, 1, 'Istorie', '6C', 3, 1, '2023-09-08'),
(388, 1, 'Istorie', '6C', 3, 6, '2023-11-05'),
(389, 1, 'Limba Romana', '6C', 3, 3, '2024-03-19'),
(390, 1, 'Limba Romana', '6C', 3, 7, '2024-04-21'),
(391, 1, 'Limba Romana', '6C', 3, 8, '2023-08-26'),
(392, 1, 'Limba Romana', '6C', 3, 7, '2023-11-17'),
(393, 1, 'Limba Romana', '6C', 3, 4, '2023-09-11'),
(394, 1, 'Sport', '6C', 3, 2, '2024-04-05'),
(395, 1, 'Sport', '6C', 3, 2, '2024-05-13'),
(396, 1, 'Sport', '6C', 3, 10, '2023-08-28'),
(397, 1, 'Engleza', '6C', 2, 1, '2023-08-23'),
(398, 1, 'Engleza', '6C', 2, 1, '2024-05-15'),
(399, 1, 'Engleza', '6C', 2, 4, '2023-08-28'),
(400, 1, 'Engleza', '6C', 2, 1, '2023-11-01'),
(401, 1, 'Geografie', '6C', 2, 2, '2023-11-16'),
(402, 1, 'Geografie', '6C', 2, 8, '2023-07-03'),
(403, 1, 'Geografie', '6C', 2, 6, '2024-03-24'),
(404, 1, 'Geografie', '6C', 2, 6, '2023-11-28'),
(405, 1, 'Geografie', '6C', 2, 1, '2023-11-12'),
(406, 1, 'Istorie', '6C', 2, 6, '2024-01-15'),
(407, 1, 'Istorie', '6C', 2, 5, '2024-02-27'),
(408, 1, 'Istorie', '6C', 2, 3, '2023-11-05'),
(409, 1, 'Istorie', '6C', 2, 4, '2023-09-04'),
(410, 1, 'Istorie', '6C', 2, 4, '2023-10-01'),
(411, 1, 'Limba Romana', '6C', 2, 3, '2023-07-20'),
(412, 1, 'Limba Romana', '6C', 2, 4, '2023-11-23'),
(413, 1, 'Limba Romana', '6C', 2, 4, '2023-12-17'),
(414, 1, 'Limba Romana', '6C', 2, 8, '2023-12-07'),
(415, 1, 'Limba Romana', '6C', 2, 8, '2024-04-08'),
(416, 1, 'Sport', '6C', 2, 2, '2023-09-06'),
(417, 1, 'Sport', '6C', 2, 10, '2024-01-21'),
(418, 1, 'Sport', '6C', 2, 2, '2023-06-25'),
(419, 1, 'Sport', '6C', 2, 5, '2023-06-27'),
(420, 1, 'Engleza', '6C', 1, 9, '2023-09-27'),
(421, 1, 'Engleza', '6C', 1, 3, '2023-06-07'),
(422, 1, 'Engleza', '6C', 1, 5, '2024-01-13'),
(423, 1, 'Engleza', '6C', 1, 10, '2024-02-18'),
(424, 1, 'Geografie', '6C', 1, 10, '2024-02-01'),
(425, 1, 'Geografie', '6C', 1, 9, '2023-09-24'),
(426, 1, 'Geografie', '6C', 1, 2, '2024-03-08'),
(427, 1, 'Geografie', '6C', 1, 8, '2024-01-16'),
(428, 1, 'Istorie', '6C', 1, 9, '2024-04-19'),
(429, 1, 'Istorie', '6C', 1, 2, '2023-08-16'),
(430, 1, 'Limba Romana', '6C', 1, 2, '2023-09-16'),
(431, 1, 'Limba Romana', '6C', 1, 7, '2023-08-24'),
(432, 1, 'Limba Romana', '6C', 1, 1, '2023-12-19'),
(433, 1, 'Sport', '6C', 1, 9, '2023-12-04'),
(434, 1, 'Sport', '6C', 1, 1, '2024-01-22'),
(435, 1, 'Sport', '6C', 1, 8, '2024-05-17'),
(436, 1, 'Chimie', '7A', 4, 8, '2024-02-24'),
(437, 1, 'Chimie', '7A', 4, 6, '2023-07-13'),
(438, 1, 'Chimie', '7A', 4, 4, '2024-01-13'),
(439, 1, 'Desen', '7A', 4, 2, '2023-07-16'),
(440, 1, 'Desen', '7A', 4, 9, '2024-01-02'),
(441, 1, 'Fizica', '7A', 4, 3, '2023-08-15'),
(442, 1, 'Fizica', '7A', 4, 8, '2023-09-03'),
(443, 1, 'Muzica', '7A', 4, 9, '2024-02-16'),
(444, 1, 'Muzica', '7A', 4, 10, '2024-01-27'),
(445, 1, 'Muzica', '7A', 4, 1, '2024-01-24'),
(446, 1, 'Muzica', '7A', 4, 3, '2024-02-16'),
(447, 1, 'Chimie', '7A', 3, 10, '2023-10-27'),
(448, 1, 'Chimie', '7A', 3, 9, '2023-12-27'),
(449, 1, 'Chimie', '7A', 3, 6, '2024-02-26'),
(450, 1, 'Desen', '7A', 3, 4, '2023-10-23'),
(451, 1, 'Desen', '7A', 3, 1, '2024-04-11'),
(452, 1, 'Desen', '7A', 3, 2, '2023-12-08'),
(453, 1, 'Desen', '7A', 3, 3, '2024-03-02'),
(454, 1, 'Desen', '7A', 3, 9, '2023-12-16'),
(455, 1, 'Fizica', '7A', 3, 4, '2024-04-15'),
(456, 1, 'Fizica', '7A', 3, 10, '2023-09-09'),
(457, 1, 'Fizica', '7A', 3, 10, '2023-08-06'),
(458, 1, 'Fizica', '7A', 3, 9, '2023-09-14'),
(459, 1, 'Fizica', '7A', 3, 10, '2023-07-08'),
(460, 1, 'Muzica', '7A', 3, 7, '2023-10-14'),
(461, 1, 'Muzica', '7A', 3, 3, '2024-01-02'),
(462, 1, 'Chimie', '7A', 2, 8, '2024-04-17'),
(463, 1, 'Chimie', '7A', 2, 2, '2023-07-12'),
(464, 1, 'Desen', '7A', 2, 10, '2023-08-15'),
(465, 1, 'Desen', '7A', 2, 7, '2023-11-02'),
(466, 1, 'Desen', '7A', 2, 2, '2024-04-22'),
(467, 1, 'Desen', '7A', 2, 9, '2023-07-16'),
(468, 1, 'Desen', '7A', 2, 8, '2023-10-23'),
(469, 1, 'Fizica', '7A', 2, 5, '2024-03-21'),
(470, 1, 'Fizica', '7A', 2, 7, '2023-11-01'),
(471, 1, 'Fizica', '7A', 2, 1, '2024-03-16'),
(472, 1, 'Fizica', '7A', 2, 7, '2023-06-01'),
(473, 1, 'Fizica', '7A', 2, 4, '2024-05-02'),
(474, 1, 'Muzica', '7A', 2, 8, '2024-03-02'),
(475, 1, 'Muzica', '7A', 2, 4, '2023-10-14'),
(476, 1, 'Muzica', '7A', 2, 3, '2023-09-01'),
(477, 1, 'Chimie', '7A', 1, 7, '2024-05-09'),
(478, 1, 'Chimie', '7A', 1, 1, '2024-01-01'),
(479, 1, 'Chimie', '7A', 1, 8, '2023-08-24'),
(480, 1, 'Chimie', '7A', 1, 10, '2023-09-09'),
(481, 1, 'Desen', '7A', 1, 8, '2023-10-24'),
(482, 1, 'Desen', '7A', 1, 8, '2023-08-21'),
(483, 1, 'Desen', '7A', 1, 8, '2023-07-16'),
(484, 1, 'Desen', '7A', 1, 3, '2024-04-02'),
(485, 1, 'Desen', '7A', 1, 4, '2024-01-16'),
(486, 1, 'Fizica', '7A', 1, 8, '2023-06-26'),
(487, 1, 'Fizica', '7A', 1, 10, '2024-04-03'),
(488, 1, 'Muzica', '7A', 1, 3, '2023-08-11'),
(489, 1, 'Muzica', '7A', 1, 8, '2023-07-07'),
(490, 1, 'Muzica', '7A', 1, 8, '2023-09-05'),
(491, 1, 'Biologie', '7B', 1, 6, '2023-08-13'),
(492, 1, 'Biologie', '7B', 1, 8, '2024-03-19'),
(493, 1, 'Biologie', '7B', 1, 3, '2023-07-05'),
(494, 1, 'Engleza', '7B', 1, 7, '2023-06-17'),
(495, 1, 'Engleza', '7B', 1, 8, '2023-07-06'),
(496, 1, 'Engleza', '7B', 1, 3, '2023-06-18'),
(497, 1, 'Engleza', '7B', 1, 7, '2024-03-07'),
(498, 1, 'Engleza', '7B', 1, 6, '2024-02-01'),
(499, 1, 'Franceza', '7B', 1, 1, '2024-02-20'),
(500, 1, 'Franceza', '7B', 1, 8, '2023-06-02'),
(501, 1, 'Limba Romana', '7B', 1, 1, '2023-12-12'),
(502, 1, 'Limba Romana', '7B', 1, 8, '2024-01-28'),
(503, 1, 'Limba Romana', '7B', 1, 3, '2024-03-02'),
(504, 1, 'Matematica', '7B', 1, 10, '2024-01-09'),
(505, 1, 'Matematica', '7B', 1, 7, '2024-02-14'),
(506, 1, 'Matematica', '7B', 1, 5, '2024-03-12'),
(507, 1, 'Matematica', '7B', 1, 5, '2023-07-18'),
(508, 1, 'Biologie', '7B', 2, 9, '2023-10-07'),
(509, 1, 'Biologie', '7B', 2, 1, '2024-01-22'),
(510, 1, 'Biologie', '7B', 2, 5, '2024-03-28'),
(511, 1, 'Engleza', '7B', 2, 1, '2023-08-06'),
(512, 1, 'Engleza', '7B', 2, 7, '2023-10-21'),
(513, 1, 'Engleza', '7B', 2, 1, '2023-10-09'),
(514, 1, 'Engleza', '7B', 2, 9, '2023-06-13'),
(515, 1, 'Engleza', '7B', 2, 10, '2023-06-24'),
(516, 1, 'Franceza', '7B', 2, 6, '2024-01-22'),
(517, 1, 'Franceza', '7B', 2, 7, '2023-09-10'),
(518, 1, 'Franceza', '7B', 2, 5, '2024-03-05'),
(519, 1, 'Limba Romana', '7B', 2, 4, '2023-07-03'),
(520, 1, 'Limba Romana', '7B', 2, 6, '2023-06-27'),
(521, 1, 'Limba Romana', '7B', 2, 6, '2023-11-21'),
(522, 1, 'Limba Romana', '7B', 2, 6, '2023-09-16'),
(523, 1, 'Matematica', '7B', 2, 9, '2023-12-19'),
(524, 1, 'Matematica', '7B', 2, 8, '2023-07-18'),
(525, 1, 'Matematica', '7B', 2, 6, '2023-07-13'),
(526, 1, 'Biologie', '7B', 3, 9, '2024-03-02'),
(527, 1, 'Biologie', '7B', 3, 2, '2023-12-06'),
(528, 1, 'Biologie', '7B', 3, 1, '2024-04-23'),
(529, 1, 'Biologie', '7B', 3, 6, '2023-11-23'),
(530, 1, 'Engleza', '7B', 3, 7, '2023-06-06'),
(531, 1, 'Engleza', '7B', 3, 2, '2023-08-14'),
(532, 1, 'Engleza', '7B', 3, 9, '2024-03-20'),
(533, 1, 'Engleza', '7B', 3, 3, '2023-06-02'),
(534, 1, 'Franceza', '7B', 3, 3, '2023-08-11'),
(535, 1, 'Franceza', '7B', 3, 7, '2024-04-08'),
(536, 1, 'Franceza', '7B', 3, 3, '2023-12-21'),
(537, 1, 'Limba Romana', '7B', 3, 10, '2024-05-02'),
(538, 1, 'Limba Romana', '7B', 3, 3, '2024-03-17'),
(539, 1, 'Limba Romana', '7B', 3, 4, '2023-12-08'),
(540, 1, 'Limba Romana', '7B', 3, 7, '2024-05-14'),
(541, 1, 'Matematica', '7B', 3, 4, '2024-01-25'),
(542, 1, 'Matematica', '7B', 3, 2, '2024-03-20'),
(543, 1, 'Biologie', '7B', 4, 6, '2023-10-07'),
(544, 1, 'Biologie', '7B', 4, 1, '2023-11-04'),
(545, 1, 'Biologie', '7B', 4, 10, '2023-09-12'),
(546, 1, 'Biologie', '7B', 4, 3, '2024-05-28'),
(547, 1, 'Engleza', '7B', 4, 10, '2023-11-19'),
(548, 1, 'Engleza', '7B', 4, 7, '2023-12-05'),
(549, 1, 'Engleza', '7B', 4, 6, '2023-08-17'),
(550, 1, 'Engleza', '7B', 4, 1, '2023-09-02'),
(551, 1, 'Engleza', '7B', 4, 6, '2024-02-17'),
(552, 1, 'Franceza', '7B', 4, 8, '2023-09-10'),
(553, 1, 'Franceza', '7B', 4, 5, '2023-11-07'),
(554, 1, 'Franceza', '7B', 4, 4, '2024-03-26'),
(555, 1, 'Franceza', '7B', 4, 6, '2023-09-23'),
(556, 1, 'Limba Romana', '7B', 4, 4, '2024-02-13'),
(557, 1, 'Limba Romana', '7B', 4, 10, '2023-08-26'),
(558, 1, 'Limba Romana', '7B', 4, 2, '2023-07-15'),
(559, 1, 'Limba Romana', '7B', 4, 1, '2024-04-14'),
(560, 1, 'Limba Romana', '7B', 4, 10, '2024-01-17'),
(561, 1, 'Matematica', '7B', 4, 8, '2023-10-13'),
(562, 1, 'Matematica', '7B', 4, 8, '2023-08-16'),
(563, 1, 'Matematica', '7B', 4, 2, '2023-11-15'),
(564, 1, 'Desen', '7C', 4, 8, '2023-09-09'),
(565, 1, 'Desen', '7C', 4, 10, '2024-05-07'),
(566, 1, 'Desen', '7C', 4, 8, '2024-01-13'),
(567, 1, 'Desen', '7C', 4, 3, '2024-05-13'),
(568, 1, 'Desen', '7C', 4, 7, '2023-09-07'),
(569, 1, 'Fizica', '7C', 4, 6, '2023-09-24'),
(570, 1, 'Fizica', '7C', 4, 7, '2023-10-27'),
(571, 1, 'Fizica', '7C', 4, 9, '2024-02-09'),
(572, 1, 'Geografie', '7C', 4, 7, '2024-03-21'),
(573, 1, 'Geografie', '7C', 4, 9, '2024-05-11'),
(574, 1, 'Geografie', '7C', 4, 4, '2023-06-26'),
(575, 1, 'Geografie', '7C', 4, 3, '2023-10-01'),
(576, 1, 'Istorie', '7C', 4, 7, '2023-10-27'),
(577, 1, 'Istorie', '7C', 4, 2, '2024-05-04'),
(578, 1, 'Istorie', '7C', 4, 8, '2024-03-15'),
(579, 1, 'Muzica', '7C', 4, 6, '2023-09-19'),
(580, 1, 'Muzica', '7C', 4, 6, '2024-02-15'),
(581, 1, 'Muzica', '7C', 4, 1, '2024-04-05'),
(582, 1, 'Muzica', '7C', 4, 1, '2023-09-07'),
(583, 1, 'Muzica', '7C', 4, 2, '2023-11-17'),
(584, 1, 'Sport', '7C', 4, 5, '2023-06-25'),
(585, 1, 'Sport', '7C', 4, 10, '2023-06-26'),
(586, 1, 'Sport', '7C', 4, 4, '2024-05-09'),
(587, 1, 'Sport', '7C', 4, 9, '2024-02-28'),
(588, 1, 'Desen', '7C', 3, 10, '2024-03-07'),
(589, 1, 'Desen', '7C', 3, 10, '2024-03-02'),
(590, 1, 'Desen', '7C', 3, 10, '2024-04-27'),
(591, 1, 'Desen', '7C', 3, 2, '2023-09-11'),
(592, 1, 'Desen', '7C', 3, 8, '2023-07-18'),
(593, 1, 'Fizica', '7C', 3, 8, '2024-02-02'),
(594, 1, 'Fizica', '7C', 3, 3, '2023-09-09'),
(595, 1, 'Geografie', '7C', 3, 7, '2023-12-13'),
(596, 1, 'Geografie', '7C', 3, 1, '2024-01-26'),
(597, 1, 'Geografie', '7C', 3, 9, '2023-08-22'),
(598, 1, 'Geografie', '7C', 3, 4, '2023-12-15'),
(599, 1, 'Geografie', '7C', 3, 5, '2023-09-23'),
(600, 1, 'Istorie', '7C', 3, 5, '2023-12-07'),
(601, 1, 'Istorie', '7C', 3, 7, '2024-03-19'),
(602, 1, 'Istorie', '7C', 3, 4, '2023-09-02'),
(603, 1, 'Istorie', '7C', 3, 7, '2023-11-15'),
(604, 1, 'Istorie', '7C', 3, 4, '2024-01-02'),
(605, 1, 'Muzica', '7C', 3, 6, '2024-01-03'),
(606, 1, 'Muzica', '7C', 3, 4, '2024-02-04'),
(607, 1, 'Muzica', '7C', 3, 10, '2024-04-23'),
(608, 1, 'Muzica', '7C', 3, 8, '2023-07-24'),
(609, 1, 'Muzica', '7C', 3, 6, '2023-06-20'),
(610, 1, 'Sport', '7C', 3, 5, '2024-03-22'),
(611, 1, 'Sport', '7C', 3, 3, '2023-07-12'),
(612, 1, 'Sport', '7C', 3, 7, '2024-03-20'),
(613, 1, 'Desen', '7C', 2, 8, '2024-04-14'),
(614, 1, 'Desen', '7C', 2, 4, '2024-05-15'),
(615, 1, 'Desen', '7C', 2, 10, '2023-10-01'),
(616, 1, 'Desen', '7C', 2, 1, '2024-05-07'),
(617, 1, 'Fizica', '7C', 2, 10, '2024-01-15'),
(618, 1, 'Fizica', '7C', 2, 4, '2023-10-15'),
(619, 1, 'Fizica', '7C', 2, 2, '2024-04-05'),
(620, 1, 'Geografie', '7C', 2, 2, '2023-09-07'),
(621, 1, 'Geografie', '7C', 2, 10, '2023-12-13'),
(622, 1, 'Geografie', '7C', 2, 9, '2024-04-06'),
(623, 1, 'Geografie', '7C', 2, 6, '2024-05-13'),
(624, 1, 'Istorie', '7C', 2, 8, '2023-11-26'),
(625, 1, 'Istorie', '7C', 2, 5, '2024-02-24'),
(626, 1, 'Istorie', '7C', 2, 5, '2024-05-22'),
(627, 1, 'Istorie', '7C', 2, 4, '2024-05-14'),
(628, 1, 'Istorie', '7C', 2, 2, '2024-01-12'),
(629, 1, 'Muzica', '7C', 2, 6, '2023-07-28'),
(630, 1, 'Muzica', '7C', 2, 10, '2024-05-28'),
(631, 1, 'Muzica', '7C', 2, 1, '2023-10-06'),
(632, 1, 'Muzica', '7C', 2, 3, '2023-09-18'),
(633, 1, 'Muzica', '7C', 2, 8, '2023-08-14'),
(634, 1, 'Sport', '7C', 2, 9, '2024-01-06'),
(635, 1, 'Sport', '7C', 2, 8, '2023-11-18'),
(636, 1, 'Sport', '7C', 2, 7, '2023-06-07'),
(637, 1, 'Sport', '7C', 2, 8, '2023-07-17'),
(638, 1, 'Desen', '7C', 1, 2, '2023-07-23'),
(639, 1, 'Desen', '7C', 1, 5, '2023-12-07'),
(640, 1, 'Fizica', '7C', 1, 7, '2023-11-15'),
(641, 1, 'Fizica', '7C', 1, 2, '2023-12-21'),
(642, 1, 'Fizica', '7C', 1, 5, '2023-06-17'),
(643, 1, 'Fizica', '7C', 1, 2, '2024-03-07'),
(644, 1, 'Fizica', '7C', 1, 4, '2023-09-14'),
(645, 1, 'Geografie', '7C', 1, 10, '2024-02-21'),
(646, 1, 'Geografie', '7C', 1, 3, '2023-06-02'),
(647, 1, 'Geografie', '7C', 1, 7, '2024-04-06'),
(648, 1, 'Geografie', '7C', 1, 4, '2023-09-03'),
(649, 1, 'Geografie', '7C', 1, 5, '2023-08-17'),
(650, 1, 'Istorie', '7C', 1, 1, '2024-01-11'),
(651, 1, 'Istorie', '7C', 1, 4, '2024-01-14'),
(652, 1, 'Muzica', '7C', 1, 9, '2024-01-25'),
(653, 1, 'Muzica', '7C', 1, 3, '2023-10-23'),
(654, 1, 'Muzica', '7C', 1, 10, '2023-10-11'),
(655, 1, 'Muzica', '7C', 1, 8, '2023-11-10'),
(656, 1, 'Sport', '7C', 1, 3, '2023-10-06'),
(657, 1, 'Sport', '7C', 1, 6, '2023-11-18'),
(658, 1, 'Sport', '7C', 1, 1, '2023-10-16'),
(659, 1, 'Sport', '7C', 1, 5, '2023-10-14'),
(660, 1, 'Biologie', '8A', 4, 8, '2023-12-04'),
(661, 1, 'Biologie', '8A', 4, 6, '2024-01-17'),
(662, 1, 'Biologie', '8A', 4, 10, '2024-05-18'),
(663, 1, 'Biologie', '8A', 4, 7, '2023-09-07'),
(664, 1, 'Chimie', '8A', 4, 2, '2023-12-14'),
(665, 1, 'Chimie', '8A', 4, 10, '2024-05-05'),
(666, 1, 'Chimie', '8A', 4, 5, '2024-02-25'),
(667, 1, 'Fizica', '8A', 4, 1, '2023-07-04'),
(668, 1, 'Fizica', '8A', 4, 1, '2023-07-02'),
(669, 1, 'Fizica', '8A', 4, 4, '2023-09-06'),
(670, 1, 'Fizica', '8A', 4, 5, '2024-02-19'),
(671, 1, 'Fizica', '8A', 4, 8, '2024-04-24'),
(672, 1, 'Franceza', '8A', 4, 8, '2023-08-06'),
(673, 1, 'Franceza', '8A', 4, 4, '2023-10-05'),
(674, 1, 'Limba Romana', '8A', 4, 4, '2024-03-27'),
(675, 1, 'Limba Romana', '8A', 4, 4, '2024-01-06'),
(676, 1, 'Limba Romana', '8A', 4, 4, '2024-01-03'),
(677, 1, 'Limba Romana', '8A', 4, 10, '2023-10-13'),
(678, 1, 'Biologie', '8A', 3, 7, '2023-11-02'),
(679, 1, 'Biologie', '8A', 3, 2, '2024-03-25'),
(680, 1, 'Biologie', '8A', 3, 9, '2023-10-09'),
(681, 1, 'Chimie', '8A', 3, 6, '2023-09-09'),
(682, 1, 'Chimie', '8A', 3, 9, '2024-01-26'),
(683, 1, 'Chimie', '8A', 3, 7, '2024-02-16'),
(684, 1, 'Chimie', '8A', 3, 6, '2023-10-12'),
(685, 1, 'Chimie', '8A', 3, 3, '2024-03-11'),
(686, 1, 'Fizica', '8A', 3, 1, '2023-07-15'),
(687, 1, 'Fizica', '8A', 3, 6, '2023-07-03'),
(688, 1, 'Fizica', '8A', 3, 3, '2023-10-13'),
(689, 1, 'Franceza', '8A', 3, 8, '2023-06-12'),
(690, 1, 'Franceza', '8A', 3, 2, '2024-04-05'),
(691, 1, 'Limba Romana', '8A', 3, 3, '2023-06-04'),
(692, 1, 'Limba Romana', '8A', 3, 6, '2023-10-13'),
(693, 1, 'Limba Romana', '8A', 3, 3, '2023-09-17'),
(694, 1, 'Limba Romana', '8A', 3, 4, '2023-08-01'),
(695, 1, 'Biologie', '8A', 2, 5, '2024-04-26'),
(696, 1, 'Biologie', '8A', 2, 10, '2023-06-09'),
(697, 1, 'Biologie', '8A', 2, 6, '2024-01-04'),
(698, 1, 'Biologie', '8A', 2, 9, '2023-12-25'),
(699, 1, 'Biologie', '8A', 2, 2, '2024-05-23'),
(700, 1, 'Chimie', '8A', 2, 6, '2024-04-28'),
(701, 1, 'Chimie', '8A', 2, 4, '2023-08-05'),
(702, 1, 'Chimie', '8A', 2, 9, '2023-11-01'),
(703, 1, 'Chimie', '8A', 2, 6, '2023-06-18'),
(704, 1, 'Chimie', '8A', 2, 10, '2023-08-18'),
(705, 1, 'Fizica', '8A', 2, 5, '2023-06-17'),
(706, 1, 'Fizica', '8A', 2, 2, '2023-06-23'),
(707, 1, 'Fizica', '8A', 2, 4, '2024-04-17'),
(708, 1, 'Fizica', '8A', 2, 7, '2024-05-11'),
(709, 1, 'Franceza', '8A', 2, 7, '2023-09-24'),
(710, 1, 'Franceza', '8A', 2, 2, '2023-09-23'),
(711, 1, 'Franceza', '8A', 2, 9, '2023-08-04'),
(712, 1, 'Franceza', '8A', 2, 6, '2023-08-02'),
(713, 1, 'Franceza', '8A', 2, 8, '2023-11-08'),
(714, 1, 'Limba Romana', '8A', 2, 3, '2023-11-26'),
(715, 1, 'Limba Romana', '8A', 2, 8, '2024-01-09'),
(716, 1, 'Limba Romana', '8A', 2, 8, '2024-03-20'),
(717, 1, 'Biologie', '8A', 1, 3, '2023-10-13'),
(718, 1, 'Biologie', '8A', 1, 9, '2023-12-25'),
(719, 1, 'Biologie', '8A', 1, 7, '2023-06-27'),
(720, 1, 'Chimie', '8A', 1, 6, '2024-01-13'),
(721, 1, 'Chimie', '8A', 1, 9, '2023-10-07'),
(722, 1, 'Chimie', '8A', 1, 9, '2023-10-24'),
(723, 1, 'Chimie', '8A', 1, 2, '2023-08-17'),
(724, 1, 'Fizica', '8A', 1, 2, '2024-04-22'),
(725, 1, 'Fizica', '8A', 1, 6, '2023-08-23'),
(726, 1, 'Franceza', '8A', 1, 5, '2024-01-18'),
(727, 1, 'Franceza', '8A', 1, 6, '2024-03-09'),
(728, 1, 'Franceza', '8A', 1, 9, '2023-08-08'),
(729, 1, 'Franceza', '8A', 1, 8, '2023-12-28'),
(730, 1, 'Limba Romana', '8A', 1, 8, '2024-01-03'),
(731, 1, 'Limba Romana', '8A', 1, 2, '2023-06-24'),
(732, 1, 'Limba Romana', '8A', 1, 2, '2024-02-04'),
(733, 1, 'Limba Romana', '8A', 1, 1, '2023-10-11'),
(734, 1, 'Engleza', '8B', 4, 4, '2024-05-04'),
(735, 1, 'Engleza', '8B', 4, 9, '2024-01-16'),
(736, 1, 'Engleza', '8B', 4, 2, '2024-04-09'),
(737, 1, 'Geografie', '8B', 4, 1, '2024-02-21'),
(738, 1, 'Geografie', '8B', 4, 5, '2023-06-08'),
(739, 1, 'Geografie', '8B', 4, 3, '2024-04-28'),
(740, 1, 'Geografie', '8B', 4, 1, '2023-12-19'),
(741, 1, 'Istorie', '8B', 4, 10, '2024-01-28'),
(742, 1, 'Istorie', '8B', 4, 3, '2023-12-19'),
(743, 1, 'Istorie', '8B', 4, 7, '2023-07-14'),
(744, 1, 'Istorie', '8B', 4, 9, '2023-11-20'),
(745, 1, 'Limba Romana', '8B', 4, 3, '2023-09-01'),
(746, 1, 'Limba Romana', '8B', 4, 8, '2024-02-04'),
(747, 1, 'Sport', '8B', 4, 6, '2023-11-16'),
(748, 1, 'Sport', '8B', 4, 10, '2023-08-13'),
(749, 1, 'Engleza', '8B', 3, 10, '2024-03-06'),
(750, 1, 'Engleza', '8B', 3, 8, '2023-11-19'),
(751, 1, 'Geografie', '8B', 3, 1, '2023-11-28'),
(752, 1, 'Geografie', '8B', 3, 4, '2023-08-28'),
(753, 1, 'Geografie', '8B', 3, 4, '2024-05-02'),
(754, 1, 'Istorie', '8B', 3, 8, '2023-10-20'),
(755, 1, 'Istorie', '8B', 3, 5, '2023-12-20'),
(756, 1, 'Limba Romana', '8B', 3, 7, '2024-01-27'),
(757, 1, 'Limba Romana', '8B', 3, 9, '2023-06-11'),
(758, 1, 'Limba Romana', '8B', 3, 8, '2024-01-03'),
(759, 1, 'Sport', '8B', 3, 2, '2024-05-26'),
(760, 1, 'Sport', '8B', 3, 9, '2023-08-09'),
(761, 1, 'Sport', '8B', 3, 3, '2024-04-08'),
(762, 1, 'Sport', '8B', 3, 7, '2023-09-16'),
(763, 1, 'Sport', '8B', 3, 5, '2023-12-14'),
(764, 1, 'Engleza', '8B', 2, 1, '2024-01-15'),
(765, 1, 'Engleza', '8B', 2, 3, '2024-02-27'),
(766, 1, 'Engleza', '8B', 2, 1, '2023-10-26'),
(767, 1, 'Engleza', '8B', 2, 7, '2023-10-11'),
(768, 1, 'Geografie', '8B', 2, 3, '2024-04-01'),
(769, 1, 'Geografie', '8B', 2, 2, '2023-06-26'),
(770, 1, 'Geografie', '8B', 2, 8, '2023-08-01'),
(771, 1, 'Geografie', '8B', 2, 4, '2024-04-10'),
(772, 1, 'Geografie', '8B', 2, 2, '2024-03-05'),
(773, 1, 'Istorie', '8B', 2, 1, '2024-04-02'),
(774, 1, 'Istorie', '8B', 2, 4, '2024-03-22'),
(775, 1, 'Istorie', '8B', 2, 8, '2023-09-07'),
(776, 1, 'Istorie', '8B', 2, 2, '2023-12-12'),
(777, 1, 'Istorie', '8B', 2, 3, '2024-04-07'),
(778, 1, 'Limba Romana', '8B', 2, 9, '2024-05-13'),
(779, 1, 'Limba Romana', '8B', 2, 6, '2023-09-14'),
(780, 1, 'Limba Romana', '8B', 2, 3, '2023-08-10'),
(781, 1, 'Sport', '8B', 2, 9, '2024-04-11'),
(782, 1, 'Sport', '8B', 2, 4, '2023-06-16'),
(783, 1, 'Sport', '8B', 2, 4, '2024-03-24'),
(784, 1, 'Engleza', '8B', 1, 10, '2024-02-10'),
(785, 1, 'Engleza', '8B', 1, 1, '2023-09-25'),
(786, 1, 'Geografie', '8B', 1, 6, '2023-09-22'),
(787, 1, 'Geografie', '8B', 1, 4, '2023-10-27'),
(788, 1, 'Geografie', '8B', 1, 5, '2023-11-27'),
(789, 1, 'Geografie', '8B', 1, 3, '2023-08-28'),
(790, 1, 'Istorie', '8B', 1, 8, '2023-07-02'),
(791, 1, 'Istorie', '8B', 1, 6, '2023-08-23'),
(792, 1, 'Istorie', '8B', 1, 3, '2023-09-02'),
(793, 1, 'Limba Romana', '8B', 1, 2, '2024-03-15'),
(794, 1, 'Limba Romana', '8B', 1, 1, '2023-09-28'),
(795, 1, 'Limba Romana', '8B', 1, 10, '2023-11-08'),
(796, 1, 'Limba Romana', '8B', 1, 7, '2023-07-05'),
(797, 1, 'Sport', '8B', 1, 3, '2023-10-28'),
(798, 1, 'Sport', '8B', 1, 5, '2023-06-24'),
(799, 1, 'Sport', '8B', 1, 6, '2023-08-22'),
(800, 1, 'Sport', '8B', 1, 8, '2024-02-25'),
(801, 1, 'Desen', '8C', 1, 10, '2023-08-16'),
(802, 1, 'Desen', '8C', 1, 1, '2024-05-11'),
(803, 1, 'Desen', '8C', 1, 1, '2023-11-06'),
(804, 1, 'Desen', '8C', 1, 1, '2023-08-18'),
(805, 1, 'Fizica', '8C', 1, 5, '2023-07-22'),
(806, 1, 'Fizica', '8C', 1, 9, '2023-09-21'),
(807, 1, 'Fizica', '8C', 1, 9, '2024-02-08'),
(808, 1, 'Fizica', '8C', 1, 10, '2023-07-19'),
(809, 1, 'Fizica', '8C', 1, 5, '2023-06-14'),
(810, 1, 'Muzica', '8C', 1, 4, '2024-02-07'),
(811, 1, 'Muzica', '8C', 1, 9, '2023-07-28'),
(812, 1, 'Muzica', '8C', 1, 8, '2023-09-14'),
(813, 1, 'Muzica', '8C', 1, 1, '2023-10-25'),
(814, 1, 'Desen', '8C', 2, 8, '2024-05-23'),
(815, 1, 'Desen', '8C', 2, 2, '2023-10-03'),
(816, 1, 'Desen', '8C', 2, 4, '2023-10-03'),
(817, 1, 'Desen', '8C', 2, 3, '2024-01-16'),
(818, 1, 'Desen', '8C', 2, 8, '2024-05-10'),
(819, 1, 'Fizica', '8C', 2, 1, '2024-03-28'),
(820, 1, 'Fizica', '8C', 2, 2, '2023-07-24'),
(821, 1, 'Fizica', '8C', 2, 10, '2024-04-01'),
(822, 1, 'Fizica', '8C', 2, 2, '2024-01-28'),
(823, 1, 'Fizica', '8C', 2, 10, '2023-10-13'),
(824, 1, 'Muzica', '8C', 2, 2, '2023-11-08'),
(825, 1, 'Muzica', '8C', 2, 1, '2024-04-09'),
(826, 1, 'Muzica', '8C', 2, 4, '2024-03-09'),
(827, 1, 'Muzica', '8C', 2, 5, '2023-11-22'),
(828, 1, 'Muzica', '8C', 2, 2, '2024-05-06'),
(829, 1, 'Desen', '8C', 3, 10, '2023-12-13'),
(830, 1, 'Desen', '8C', 3, 5, '2023-06-25'),
(831, 1, 'Desen', '8C', 3, 4, '2023-11-05'),
(832, 1, 'Desen', '8C', 3, 3, '2024-05-18'),
(833, 1, 'Desen', '8C', 3, 9, '2023-12-02'),
(834, 1, 'Fizica', '8C', 3, 2, '2023-10-26'),
(835, 1, 'Fizica', '8C', 3, 2, '2023-06-27'),
(836, 1, 'Muzica', '8C', 3, 9, '2023-10-26'),
(837, 1, 'Muzica', '8C', 3, 6, '2023-10-15'),
(838, 1, 'Muzica', '8C', 3, 1, '2023-06-23'),
(839, 1, 'Desen', '8C', 4, 2, '2024-05-27'),
(840, 1, 'Desen', '8C', 4, 2, '2023-06-19'),
(841, 1, 'Desen', '8C', 4, 2, '2023-08-28'),
(842, 1, 'Fizica', '8C', 4, 6, '2023-08-11'),
(843, 1, 'Fizica', '8C', 4, 7, '2024-05-27'),
(844, 1, 'Fizica', '8C', 4, 4, '2024-02-24'),
(845, 1, 'Fizica', '8C', 4, 10, '2024-01-17'),
(846, 1, 'Fizica', '8C', 4, 9, '2024-05-23'),
(847, 1, 'Muzica', '8C', 4, 10, '2024-02-10'),
(848, 1, 'Muzica', '8C', 4, 7, '2023-07-01');

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

--
-- Dumping data for table `profesor`
--

INSERT INTO `profesor` (`id_scoala`, `id`, `nume`, `prenume`, `token`, `id_cont`) VALUES
(1, 1, 'Mocanu', 'Stefan', 'CkxZ5EDeZp', 7),
(1, 2, 'Tilica', 'Gabi', 'IpG7mMyUNL', NULL),
(1, 3, 'Ciobanu', 'Paul', 'DlEjFMMilP', NULL),
(1, 4, 'Smith', 'John', 'oifU6LUyYq', NULL),
(1, 5, 'Johnson', 'Emily', 't0uTXKxzep', NULL),
(1, 6, 'Williams', 'Michael', '1OYo2GL0p0', NULL),
(1, 7, 'Brown', 'Jessica', 'mNK424S2gD', NULL),
(1, 8, 'Jones', 'David', 'hAKYGZojJc', NULL),
(1, 9, 'Davis', 'Sarah', 'UPG9lCGSrZ', NULL),
(1, 10, 'Miller', 'James', 'jwElLjPHTP', NULL),
(1, 11, 'Wilson', 'Ashley', 'Q88WX4MFYT', NULL),
(1, 12, 'Taylor', 'Matthew', 'kFj3UWuNWJ', NULL),
(1, 13, 'Anderson', 'Jennifer', 'wDzbuN578I', NULL),
(1, 14, 'Thomas', 'Christopher', 'OK4qS9uhrY', NULL),
(1, 15, 'Jackson', 'Amanda', 'qrOA3GOcjW', NULL),
(1, 16, 'White', 'Melissa', 'yXk2cbFFaF', NULL),
(1, 17, 'Harris', 'Daniel', 'Onnfy2jiAD', NULL),
(1, 18, 'Martin', 'Lisa', 'gPzG20YVXX', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `roluri`
--

CREATE TABLE `roluri` (
  `nume_rol` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `roluri`
--

INSERT INTO `roluri` (`nume_rol`) VALUES
('Administrator'),
('Elev'),
('Parinte'),
('Profesor');

-- --------------------------------------------------------

--
-- Table structure for table `scoala`
--

CREATE TABLE `scoala` (
  `id` int(11) NOT NULL,
  `nume` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `scoala`
--

INSERT INTO `scoala` (`id`, `nume`) VALUES
(1, 'Scoala Gimnaziala C-tin Platon Bacau');

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
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

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
-- Indexes for table `feedback`
--
ALTER TABLE `feedback`
  ADD PRIMARY KEY (`id_feedback`),
  ADD KEY `id_scoala` (`id_scoala`,`id_clasa`,`id_elev`),
  ADD KEY `id_scoala_2` (`id_scoala`,`id_clasa`,`nume_disciplina`);

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
  MODIFY `id_nota` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `cont`
--
ALTER TABLE `cont`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `feedback`
--
ALTER TABLE `feedback`
  MODIFY `id_feedback` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `note`
--
ALTER TABLE `note`
  MODIFY `id_nota` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=849;

--
-- AUTO_INCREMENT for table `scoala`
--
ALTER TABLE `scoala`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

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
