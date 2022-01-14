-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jan 14, 2022 at 10:14 AM
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
  `des` longtext NOT NULL,
  `thuimg` text NOT NULL,
  `preimg` text NOT NULL,
  `active` int(11) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `linka_links`
--

INSERT INTO `linka_links` (`id`, `slug`, `link`, `name`, `pageid`, `username`, `time`, `type`, `des`, `thuimg`, `preimg`, `active`) VALUES
(1, 'c2WD8F2q', 'http://app.rellitel.ink/dashboard', 'Parthka\'s First link', 1, 'parthka', '2022-01-14 14:35:58', 1, '', 'thu_1_c2WD8F2q.png', 'pre_1_c2WD8F2q.png', 0),
(2, 'NfHK5a84', 'http://app.rellitel.ink/dashboard', '123', 1, 'parthka', '2022-01-14 14:37:36', 1, '123', 'thu_1_NfHK5a84.png', '', 1);

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
  `theme` int(11) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `linka_pages`
--

INSERT INTO `linka_pages` (`id`, `name`, `slug`, `username`, `time`, `cat`, `info`, `theme`) VALUES
(1, 'Parthka Page', 'parthka-page2', 'parthka', '2022-01-14 14:15:12', '{}', '{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"1_2\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}', 3),
(2, 'Parthkao\'s Page', 'parthkaos-page', 'parthkao', '2022-01-14 14:23:12', '{}', '{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"/assets/img/brand/g61.png\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}', 1);

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

--
-- Dumping data for table `linka_users`
--

INSERT INTO `linka_users` (`id`, `username`, `email`, `password`, `fullname`, `verify`, `phone`, `photo`, `time`) VALUES
(1, 'parthka', 'pathka.2005@gmail.com', '1SxKpmIJab9bOxU1uHeFLw', 'hello world', 0, '1234567895', NULL, '2022-01-14 14:15:12'),
(2, 'parthkao', 'parthka.2005@gmail.com', '1SxKpmIJab9bOxU1uHeFLw', 'parthka offical', 1, '132123123123', NULL, '2022-01-14 14:23:12');

-- --------------------------------------------------------

--
-- Table structure for table `linka_user_token`
--

CREATE TABLE `linka_user_token` (
  `token` text NOT NULL,
  `username` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `linka_user_token`
--

INSERT INTO `linka_user_token` (`token`, `username`) VALUES
('MTY0MjE1MTE5OHwtS2dfaGlabGpFakM5RE14UkNSUkxzYkduSEd1cVcwamlSVXdzbUtFSm44WnNTTEgwdnZ2cG94MElGUVlXdXBVSktuUHhmY0JZakgySXNZcE1ia0NQX2JWY2tEYk9kNTl86StBBnmW3efZdHk7Bo7oQjwSmrPmw3OymmtVaUQ_RLQ=', 'parthka'),
('MTY0MjE1MTIxMnxmUXA3OWM3ZV9BdGFpdjByYzdUb2E2eTR0bmVnbnRtQU54UlFiN2d0MWV4bG5qU3FETlVURDZ1ZjVUdHE1cHd0dTZ6TU5tZTJUUjVrQUtKTThVOWxaM21fTjYxaXF4N3A3Zz09fMxZu4xr9E4P2oG4D6H9-cdeS3FNDFutcw55uLIdEL5p', 'parthkao');

-- --------------------------------------------------------

--
-- Table structure for table `linka_verify_code`
--

CREATE TABLE `linka_verify_code` (
  `email` varchar(200) NOT NULL,
  `code` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `linka_verify_code`
--

INSERT INTO `linka_verify_code` (`email`, `code`) VALUES
('pathka.2005@gmail.com', '2CMNJEXOTKEKZAVUMRRNBQVO53VLEZDS5A=');

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
-- Indexes for table `linka_users`
--
ALTER TABLE `linka_users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `linka_links`
--
ALTER TABLE `linka_links`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `linka_pages`
--
ALTER TABLE `linka_pages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `linka_users`
--
ALTER TABLE `linka_users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
