-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jan 15, 2022 at 04:59 PM
-- Server version: 10.4.22-MariaDB
-- PHP Version: 8.1.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `rellitel.ink`
--

-- --------------------------------------------------------

--
-- Table structure for table `linka_forgot_code`
--

CREATE TABLE `linka_forgot_code` (
  `email` varchar(200) NOT NULL,
  `code` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `linka_links`
--

CREATE TABLE `linka_links` (
  `id` int(11) NOT NULL,
  `slug` text NOT NULL,
  `link` text NOT NULL,
  `name` text NOT NULL,
  `pageid` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `time` datetime NOT NULL,
  `type` int(11) NOT NULL DEFAULT 1,
  `des` text NOT NULL,
  `thuimg` text NOT NULL,
  `preimg` text NOT NULL,
  `active` int(11) NOT NULL DEFAULT 1,
  `con` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `linka_pages`
--

CREATE TABLE `linka_pages` (
  `id` int(11) NOT NULL,
  `name` text NOT NULL,
  `slug` varchar(500) NOT NULL,
  `username` varchar(50) NOT NULL,
  `time` datetime NOT NULL,
  `cat` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '{}' CHECK (json_valid(`cat`)),
  `info` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '{}' CHECK (json_valid(`info`)),
  `theme` int(11) NOT NULL DEFAULT 1,
  `adult` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `linka_report`
--

CREATE TABLE `linka_report` (
  `id` int(11) NOT NULL,
  `type` text NOT NULL,
  `pageid` text NOT NULL,
  `slug` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `linka_users`
--

CREATE TABLE `linka_users` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(200) NOT NULL,
  `password` text NOT NULL,
  `fullname` text NOT NULL,
  `verify` int(11) NOT NULL DEFAULT 0,
  `phone` text DEFAULT NULL,
  `photo` text DEFAULT NULL,
  `time` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `linka_user_token`
--

CREATE TABLE `linka_user_token` (
  `token` text NOT NULL,
  `username` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `linka_verify_code`
--

CREATE TABLE `linka_verify_code` (
  `email` varchar(200) NOT NULL,
  `code` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `linka_wallet`
--

CREATE TABLE `linka_wallet` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `balance` float NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `linka_links`
--
ALTER TABLE `linka_links`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `linka_pages`
--
ALTER TABLE `linka_pages`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `linka_report`
--
ALTER TABLE `linka_report`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `linka_users`
--
ALTER TABLE `linka_users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `linka_wallet`
--
ALTER TABLE `linka_wallet`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `linka_links`
--
ALTER TABLE `linka_links`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `linka_pages`
--
ALTER TABLE `linka_pages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `linka_report`
--
ALTER TABLE `linka_report`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `linka_users`
--
ALTER TABLE `linka_users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `linka_wallet`
--
ALTER TABLE `linka_wallet`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
