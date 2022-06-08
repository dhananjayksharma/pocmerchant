-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jun 08, 2022 at 09:15 PM
-- Server version: 8.0.29-0ubuntu0.20.04.3
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `Pacenow_DemoDb`
--

-- --------------------------------------------------------

--
-- Table structure for table `merchants`
--

CREATE TABLE `merchants` (
  `id` bigint UNSIGNED NOT NULL,
  `code` char(24) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `name` varchar(55) NOT NULL,
  `address` varchar(465) NOT NULL,
  `status` tinyint DEFAULT '0' COMMENT '1-Not Active, 2-Active, 3-Deactivated',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `merchants`
--

INSERT INTO `merchants` (`id`, `code`, `name`, `address`, `status`, `created_at`, `updated_at`) VALUES
(9, 'cadjq02gqpmvra1scb0g', 'Cinestaan', 'Mumbai', 1, '2022-06-04 16:40:28', '2022-06-04 16:40:28'),
(10, 'cadjq02gqpmvra18971', 'Rediff', 'Mumbai, Ville Parle', 1, '2022-06-04 16:41:34', '2022-06-05 08:44:45'),
(108, 'cadjq02gqpmvljdra98', 'Sony ltd', 'Mumbai', 1, '2022-06-08 19:11:12', '2022-06-08 19:11:12'),
(111, 'cadjq02gqpmvljdrad98', 'Sony d ltd', 'Mumbai', 1, '2022-06-08 19:17:00', '2022-06-08 19:17:00');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int UNSIGNED NOT NULL,
  `fk_code` char(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `first_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `last_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `is_active` tinyint UNSIGNED DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `fk_code`, `first_name`, `last_name`, `email`, `mobile`, `is_active`, `created_at`, `updated_at`) VALUES
(200, 'cadjq02gqpmvra18971', 'dhananjay', 'sharma', 'dhananjay@gmail.com', '', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(202, 'cadjq02gqpmvra18971', 'name 1', 'last 1', 'name1@gmail.com', '9855555478', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(204, 'cadjq02gqpmvra18971', 'name 2', 'last 2', 'name2@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(205, 'cadjq02gqpmvra18971', 'name 3', 'last 3', 'name3@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(206, 'cadjq02gqpmvra18971', 'name 4', 'last 4', 'name4@gmail.com', '9855555423', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(207, 'cadjq02gqpmvra18971', 'name 1a', 'last 1a', 'name1a@gmail.com', '9855555433', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(208, 'cadjq02gqpmvra18971', 'name 2a', 'last 2a', 'name2a@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(209, 'cadjq02gqpmvra18971', 'name 3a', 'last 3a', 'name3a@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(210, 'cadjq02gqpmvra18971', 'name 4a', 'last 4a', 'name4a@gmail.com', '9855555423', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(211, 'cadjq02gqpmvra18971', 'name 1ba', 'last 1ba', 'name1ba@gmail.com', '9855555433', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(212, 'cadjq02gqpmvra18971', 'name 2ba', 'last 2ba', 'name2ba@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(213, 'cadjq02gqpmvra18971', 'name 3ba', 'last 3ba', 'name3ba@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(214, 'cadjq02gqpmvra18971', 'name 4ba', 'last 4ba', 'name4ba@gmail.com', '9855555423', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(215, 'cadjq02gqpmvra1scb0g', 'dhananjay', 'sharma', 'dhananjay@gmail.com', '', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(216, 'cadjq02gqpmvra1scb0g', 'm2mname 1', 'last 1', 'm2mname1@gmail.com', '9855555478', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(217, 'cadjq02gqpmvra1scb0g', 'm2mname 2', 'last 2', 'm2mname2@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(218, 'cadjq02gqpmvra1scb0g', 'm2mname 3', 'last 3', 'm2mname3@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(219, 'cadjq02gqpmvra1scb0g', 'm2mname 4', 'last 4', 'm2mname4@gmail.com', '9855555423', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(220, 'cadjq02gqpmvra1scb0g', 'm2mname 1a', 'last 1a', 'm2mname1a@gmail.com', '9855555433', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(221, 'cadjq02gqpmvra1scb0g', 'm2mname 2a', 'last 2a', 'm2mname2a@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(222, 'cadjq02gqpmvra1scb0g', 'm2mname 3a', 'last 3a', 'm2mname3a@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(223, 'cadjq02gqpmvra1scb0g', 'm2mname 4a', 'last 4a', 'm2mname4a@gmail.com', '9855555423', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(224, 'cadjq02gqpmvra1scb0g', 'm2mname 1ba', 'last 1ba', 'm2mname1ba@gmail.com', '9855555433', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(225, 'cadjq02gqpmvra1scb0g', 'm2mname 2ba', 'last 2ba', 'm2mname2ba@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(226, 'cadjq02gqpmvra1scb0g', 'm2mname 3ba', 'last 3ba', 'm2mname3ba@gmail.com', '985555542', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41'),
(227, 'cadjq02gqpmvra1scb0g', 'm2mname 4ba', 'last 4ba', 'm2mname4ba@gmail.com', '9855555423', 1, '2022-06-04 22:32:41', '2022-06-04 22:32:41');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `merchants`
--
ALTER TABLE `merchants`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `Code_UniqueIndex` (`code`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `Code_Email_UniqueIndex` (`fk_code`,`email`) USING BTREE;

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `merchants`
--
ALTER TABLE `merchants`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=112;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=228;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
