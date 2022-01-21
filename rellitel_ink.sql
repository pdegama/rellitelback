-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jan 21, 2022 at 02:55 PM
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
-- Table structure for table `linka_analytics`
--

CREATE TABLE `linka_analytics` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `pageid` int(11) NOT NULL,
  `linkid` int(11) NOT NULL,
  `time` datetime NOT NULL DEFAULT current_timestamp(),
  `cpc` float NOT NULL,
  `por` float NOT NULL,
  `device` int(11) NOT NULL DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `linka_analytics`
--

INSERT INTO `linka_analytics` (`id`, `username`, `pageid`, `linkid`, `time`, `cpc`, `por`, `device`) VALUES
(4, 'parthkao', 2, 30, '2022-01-16 12:20:53', 0.3, 0.196552, 1),
(5, '78612386', 3, 71, '2022-01-16 15:17:29', 0.3, 0.27931, 1),
(6, '78612386', 3, 71, '2022-01-16 15:18:36', 0.3, 0.196552, 1),
(7, '78612386', 3, 71, '2022-01-16 15:20:12', 0.003, 0.00124138, 1),
(10, 'parthka', 1, 65, '2021-12-16 11:34:30', 0.003, 0.00144828, 1),
(11, 'parthka', 1, 70, '2022-01-10 11:35:05', 0.003, 0.00174828, 1),
(12, 'parthka', 1, 69, '2022-01-10 17:29:13', 0.003, 0.003, 1),
(13, 'parthka', 1, 72, '2022-01-10 19:05:06', 0.003, 0.00155172, 1),
(14, 'parthka', 1, 73, '2022-01-10 19:37:10', 0.003, 0.00124138, 1),
(15, 'parthka', 1, 74, '2022-01-10 20:03:29', 0.003, 0.00124138, 1),
(16, 'parthka', 1, 82, '2022-01-10 20:43:18', 0.003, 0.003, 1),
(17, 'parthka', 1, 79, '2022-01-10 20:43:50', 0.003, 0.00134483, 1),
(18, 'parthka', 1, 79, '2022-01-10 20:51:17', 0.003, 0.00144828, 1),
(19, 'parthka', 1, 79, '2022-01-10 20:51:40', 0.003, 0.00175862, 1),
(20, 'parthka', 1, 82, '2022-01-10 20:54:59', 0.003, 0.003, 1),
(21, 'parthka', 1, 83, '2022-01-10 20:55:30', 0.003, 0.00144828, 1),
(22, 'parthka', 1, 82, '2022-01-17 20:58:48', 0.003, 0.00124138, 1),
(23, 'parthka', 1, 82, '2022-01-10 21:01:05', 0.003, 0.00134483, 1),
(24, 'parthka', 1, 69, '2022-01-10 21:01:40', 0.003, 0.00144828, 1),
(25, 'parthka', 1, 82, '2022-01-10 21:09:18', 0.003, 0.00155172, 1),
(26, 'parthka', 1, 82, '2022-01-10 21:10:19', 0.003, 0.00175862, 1),
(27, 'parthka2', 4, 84, '2022-01-10 21:19:59', 0.003, 0.003, 1),
(28, 'parthka2', 4, 84, '2022-01-10 21:23:37', 0.003, 0.003, 1),
(29, 'parthka', 1, 83, '2022-01-10 09:51:14', 0.003, 0.00124138, 1),
(30, 'parthka', 1, 65, '2022-01-10 10:20:24', 0.003, 0.00175862, 1),
(31, 'parthka5', 6, 85, '2022-01-18 12:08:06', 0.003, 0.00165517, 1),
(32, 'parthka', 1, 83, '2022-01-18 12:21:15', 0.003, 0.00175862, 1),
(33, 'parthka', 1, 69, '2022-01-18 16:44:29', 0.003, 0.00258621, 1),
(34, 'parthka', 1, 65, '2022-01-18 16:45:34', 0.003, 0.00134483, 1),
(35, 'parthka', 1, 83, '2022-01-10 16:46:57', 0.003, 0.00268966, 1),
(36, 'parthka', 1, 65, '2022-01-18 16:50:52', 0.003, 0.00124138, 1),
(37, 'parthka', 1, 82, '2022-01-18 17:12:28', 0.003, 0.00165517, 1),
(38, 'parthka', 1, 82, '2022-01-18 17:13:00', 0.003, 0.00124138, 1),
(39, 'parthka', 1, 65, '2022-01-18 20:42:17', 0.003, 0.00134483, 1),
(40, 'parthka', 1, 65, '2022-01-18 20:42:44', 0.003, 0.00144828, 1),
(41, 'parthka', 1, 5, '2022-01-18 20:43:39', 0.003, 0.00124138, 1),
(42, 'parthka', 1, 78, '2022-01-18 20:44:06', 0.003, 0.003, 1),
(43, 'parthka', 1, 58, '2022-01-18 20:44:22', 0.003, 0.003, 1),
(44, 'parthka', 1, 75, '2022-01-18 20:45:39', 0.003, 0.003, 1),
(45, 'parthka6', 7, 86, '2021-12-10 09:33:58', 0.003, 0.00206897, 1),
(46, 'parthka', 1, 94, '2022-01-10 09:39:32', 0.003, 0.003, 1),
(47, 'parthka', 1, 94, '2022-01-19 10:49:07', 0.003, 0.003, 1),
(48, 'parthka', 1, 94, '2022-01-19 10:49:41', 0.003, 0.00217241, 1),
(49, 'parthka', 1, 93, '2022-01-19 10:49:56', 0.003, 0.003, 1),
(50, 'parthka7', 8, 96, '2022-01-18 10:51:12', 0.003, 0.003, 1),
(51, 'parthka6', 7, 89, '2022-01-18 11:10:28', 0.003, 0.00113793, 1),
(52, 'parthka', 1, 94, '2022-01-19 11:18:13', 0.003, 0.00237931, 1),
(53, 'parthka', 1, 93, '2022-01-19 11:18:32', 0.003, 0.003, 1),
(54, 'parthka', 1, 80, '2022-01-19 11:18:45', 0.003, 0.003, 1),
(55, 'parthka', 1, 65, '2022-01-19 11:19:00', 0.003, 0.003, 1),
(56, 'parthka', 1, 65, '2022-01-19 11:19:14', 0.003, 0.003, 1),
(57, 'parthka', 1, 82, '2022-01-19 11:19:41', 0.003, 0.00175862, 1),
(58, 'parthka', 1, 65, '2022-01-19 11:19:57', 0.003, 0.003, 1),
(59, 'parthka', 1, 79, '2022-01-19 11:36:34', 0.003, 0.00134483, 1),
(60, 'parthka', 1, 93, '2022-01-20 11:50:37', 0.003, 0.00217241, 1),
(61, 'parthka', 1, 65, '2022-01-21 10:30:56', 5, 2.06897, 1);

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

--
-- Dumping data for table `linka_links`
--

INSERT INTO `linka_links` (`id`, `slug`, `link`, `name`, `pageid`, `username`, `time`, `type`, `des`, `thuimg`, `preimg`, `active`, `con`) VALUES
(1, 'c2WD8F2q', 'http://app.rellitel.ink/dashboard', 'Parthka\'s First link', 1, 'parthka', '2022-01-14 14:35:58', 1, '', 'thu_1_c2WD8F2q.png', 'pre_1_c2WD8F2q.png', 0, ''),
(2, 'NfHK5a84', 'http://app.rellitel.ink/dashboard', '123', 1, 'parthka', '2022-01-14 14:37:36', 1, '123', 'thu_1_NfHK5a84.png', '', 1, ''),
(3, 'jjJkwzDk', 'http://app.rellitel.ink/dashboard/links', '123', 1, 'parthka', '2022-01-14 14:46:53', 1, '', '', '', 1, ''),
(4, 'h9h2fhfU', 'http://app.rellitel.ink/dashboard/links', 'untitled', 1, 'parthka', '2022-01-14 14:47:13', 1, '', '', '', 1, ''),
(5, 'BpLnfgDs', 'http://app.rellitel.ink/dashboard/new', 'asd', 1, 'parthka', '2022-01-14 18:19:53', 1, '', '', '', 1, ''),
(6, 'VuS9jZ8u', 'http://app.rellitel.ink/dashboard/new', '1', 1, 'parthka', '2022-01-14 18:21:03', 1, '', '', '', 1, ''),
(7, 'VbhV3vC5', 'http://app.rellitel.ink/dashboard/new', '2', 1, 'parthka', '2022-01-14 18:21:07', 1, '', '', '', 1, ''),
(8, 'AWX39IVU', 'http://app.rellitel.ink/dashboard/new', '3', 1, 'parthka', '2022-01-14 18:21:09', 1, '', '', '', 1, ''),
(9, 'WSP2NcHc', 'http://app.rellitel.ink/dashboard/new', '4', 1, 'parthka', '2022-01-14 18:21:10', 1, '', '', '', 1, ''),
(10, 'iWvqZTa2', 'http://app.rellitel.ink/dashboard/new', '5', 1, 'parthka', '2022-01-14 18:21:11', 1, '', '', '', 1, ''),
(11, 'N95RxRTZ', 'http://app.rellitel.ink/dashboard/new', '6', 1, 'parthka', '2022-01-14 18:21:12', 1, '', '', '', 1, ''),
(12, 'HWUsaD6H', 'http://app.rellitel.ink/dashboard/new', '7', 1, 'parthka', '2022-01-14 18:21:14', 1, '', '', '', 1, ''),
(13, 'Edz0ThbX', 'http://app.rellitel.ink/dashboard/new', '8', 1, 'parthka', '2022-01-14 18:21:15', 1, '', '', '', 1, ''),
(14, 'fQ6pYSQ3', 'http://app.rellitel.ink/dashboard/new', '8', 1, 'parthka', '2022-01-14 18:21:15', 1, '', '', '', 1, ''),
(15, 'n267l1VQ', 'http://app.rellitel.ink/dashboard/new', '9', 1, 'parthka', '2022-01-14 18:21:21', 1, '', '', '', 1, ''),
(16, 'KGNbSuJE', 'http://app.rellitel.ink/dashboard/new', '10', 1, 'parthka', '2022-01-14 18:21:22', 1, '', '', '', 1, ''),
(17, '9fQbzONJ', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 18:26:33', 1, '', '', '', 1, ''),
(18, 'AAwdCxmM', '', 'untitled', 1, 'parthka', '2022-01-14 18:26:40', 1, '', '', '', 1, ''),
(19, '8BIabKER', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 18:26:41', 1, '', '', '', 1, ''),
(20, 'sUhPNmMm', '', 'untitled', 1, 'parthka', '2022-01-14 18:26:46', 1, '', '', '', 1, ''),
(21, 'df2eSJyY', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 18:26:48', 1, '', '', '', 1, ''),
(22, 'tqwcFiUI', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 18:26:50', 1, '', '', '', 1, ''),
(23, 'LzXv2fcN', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 18:26:51', 1, '', '', '', 1, ''),
(24, 'IrWO7sTo', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 18:26:53', 1, '', '', '', 1, ''),
(25, 'FgoilA0U', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 18:26:54', 1, '', '', '', 1, ''),
(26, '1WxNeW1g', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 18:26:56', 1, '', '', '', 1, ''),
(27, 'dgUVDsEW', 'http://app.rellitel.ink/dashboard/new', 'untitled', 2, 'parthkao', '2022-01-14 19:45:23', 1, '', '', '', 1, ''),
(28, 'J77aX7tL', 'http://app.rellitel.ink/dashboard/new', 'untitled', 2, 'parthkao', '2022-01-14 19:45:24', 1, '', '', '', 1, ''),
(29, 'FJ84qYU6', 'http://app.rellitel.ink/dashboard/new', 'untitled', 2, 'parthkao', '2022-01-14 19:45:25', 1, '', '', '', 1, ''),
(30, 'UrN8ctec', 'http://app.rellitel.ink/dashboard/new', 'untitled', 2, 'parthkao', '2022-01-14 19:45:27', 1, '', '', '', 1, ''),
(31, 'wZt5S4zj', 'http://app.rellitel.ink/dashboard/new', 'untitled', 2, 'parthkao', '2022-01-14 19:45:27', 1, '', '', '', 1, ''),
(32, 'hD0tXRTm', 'http://app.rellitel.ink/dashboard/new', 'untitled', 2, 'parthkao', '2022-01-14 19:45:28', 1, '', '', '', 1, ''),
(33, 'kYKQoN91', 'http://app.rellitel.ink/dashboard/new', 'untitled', 2, 'parthkao', '2022-01-14 19:45:29', 1, '', '', '', 1, ''),
(34, 'FmWnQSK2', 'http://app.rellitel.ink/dashboard/new', 'untitled', 2, 'parthkao', '2022-01-14 19:45:30', 1, '', '', '', 1, ''),
(35, 'wRC5UHK2', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 19:46:34', 1, '', '', '', 1, ''),
(36, 'KqAtxjP2', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 19:46:35', 1, '', '', '', 1, ''),
(37, 'ZmD1jtt3', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 19:46:36', 1, '', '', '', 1, ''),
(38, 'zgr5MeUj', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 19:46:37', 1, '', '', '', 1, ''),
(39, 'oAjcO9az', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 19:46:38', 1, '', '', '', 1, ''),
(40, 'MmtU3Ytv', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 19:46:38', 1, '', '', '', 1, ''),
(41, '0P7OPmmS', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 19:46:39', 1, '', '', '', 1, ''),
(42, 'Na87d6ts', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 19:46:40', 1, '', '', '', 1, ''),
(43, 'taxy5nac', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-14 19:46:40', 1, '', 'thu_1_taxy5nac.jpg', 'pre_1_taxy5nac.jpg', 1, ''),
(44, 'nJBSuFpO', 'http://app.rellitel.ink/dashboard/new', 'angel has fallen 300mb | 500mb | full HD | hindi - English', 1, 'parthka', '2022-01-14 19:46:41', 1, 'Elephants are herbivorous and feed on leaves, plants, grains, fruits and more. They are mostly found in Africa and Asia. Most of the elephants are grey in color, however, in Thailand, they have white elephants.', 'thu_1_nJBSuFpO.jpg', '', 1, 'Elephants are herbivorous and feed on leaves, plants, grains, fruits and more. They are mostly found in Africa and Asia. Most of the elephants are grey in color, however, in Thailand, they have white elephants.\n\nIn addition, elephants are one of the longest-lived animals with an average lifespan of around 5-70 years. But, the oldest elephant to ever live passed away at the age of 86 years.\n\nFurthermore, they mostly inhabit jungles but humans have forced them to work in zoos and circuses. Elephants are considered to be one of the most intelligent animals.\n\nSimilarly, they are quite obedient too. Usually, the female elephants live in groups but the male ones prefer solitary living. Additionally, this wild animal has great learning capacity. Humans use them for transport and entertainment purposes. Elephants are of great importance to the earth and mankind. Thus, we must protect them to not create an imbalance in nature’s cycle.\n\nGet the huge list of more than 500 Essay Topics and Ideas\n\nImportance of Elephants\nElephants come in the group of most intelligent creatures. They are capable of quite strong emotions. These creatures have earned the respect of people of Africa that share the landscape with them. This gives them a great cultural significance. Elephants are tourism magnets for mankind. In addition, they also play a great role in maintaining the biodiversity of the ecosystems.\n\nMost importantly, elephants are also significant for wildlife. They dig for water in the dry season with their tusks. It helps them survive the dry environment and droughts and also helps other animals to survive.\n\nIn addition, the elephants of the forest create gaps in the vegetation while eating. The gaps created enables the growth of new plants as well as pathways for smaller animals. This method also helps in dispersal of seeds by trees.\n\nFurthermore, even elephant dung is beneficial. The dung they leave contains seeds of plants they have consumed. This, in turn, helps the birth of new grasses, bushes, and even trees. Thus, they also boost the health of the savannah ecosystem.\n\nEndangerment of Elephants\nElephants have found their way on the list of endangered species. Selfish human activities have caused this endangerment. One of the biggest reasons for their endangerment is the illegal killing of elephants. As their body parts are very profitable, humans kill them off for their skin, bones, tusks, and more.\n\nMoreover, humans are wiping out the natural habitat of elephants i.e. the forests. This results in a lack of food, area to live, and resources to survive. Similarly, hunting and poaching just for the thrill of it also cause the death of elephants.\n\nTherefore, we see how humans are the main reason behind their endangerment. In other words, we must educate the public about the importance of elephants. Conservation efforts must be taken aggressively to protect them. In addition, poachers must be arrested to stop killing of the endangered species.\n\nFAQs on Essay on Elephant\nQ.1 Why are Elephants important?\n\nA.1 Elephants are important not only to humans but wildlife and vegetation too. They provide sources of water for other animals in the dry season. Their eating method helps in the growth of new plants. They maintain the balance of the savannah ecosystem.\n\nQ.2 Why is endangerment of elephants harmful?\n\nA.2 Human activities have caused endangerment of elephants. Extinction of these animals will create an imbalance in the ecosystem gravely. We must take steps to stop this endangerment so they can be protected from extinction.'),
(45, 'hyIXE6PJ', 'http://app.rellitel.ink/dashboard/new', 'untitled', 2, 'parthkao', '2022-01-14 21:00:32', 1, '', '', '', 1, ''),
(46, 'dgUVDsEW', 'http://static.rellitel.ink/images/thu_1_nJBSuFpO.jpg', 'untitled', 1, 'parthka', '2022-01-15 11:13:20', 4, '', '', '', 1, ''),
(47, 'J77aX7tL', 'https://file-examples-com.github.io/uploads/2018/04/file_example_AVI_480_750kB.avi', 'untitled', 1, 'parthka', '2022-01-15 11:14:34', 1, '', 'thu_1_J77aX7tL.png', '', 1, ''),
(48, 'FJ84qYU6', 'https://dr.sf-converter.com/download?id=751257acdf7c923fe160558a16c1a4846b638dc342a93354883404b448f9dc0a&payload=1*eJzVlNtu4zYQhl%2FFCBCiBSpHIqmDAxALJVYcnxJnffaNQVuyTEuiaJ1sZ9F3Lymn3fRyixbo3nwckSA1%2Fwx%2FfrvJ0zLbBtMsvrm%2F2ReFyO%2Fv7k6nU%2FOSlkW5CZrbNLk70WK7%2F1KRYHm051SU7cH05rePnV3%2FhzfS0mfp3%2F%2BYZVDTtJxrlIWxxfOoGaZpGAcV84O0PqqOREwvG7qNvgRnwbKAGBaGEOGW7oCAkepxzZYrjQ78pDvvH5ds6w4AEwQaVtNARtPQHTk6gPkk1dzXxPUmrui%2FxJ3OqVhhlJyzrb0fx29DN9QPj%2BMH21locxHxKWAFDYmBdXDVTD40giw4ljKPPI%2FJJchBsidWBZKEIOMWPkILJJx80iTn5Je550HEbbmYE1rKuZRnIKlIIsEIBCImEAPGWbE9cX8jcmK1oKnroBJZRQyQsCQgdQlv4VMiMOA58YYXz43d2QPtv%2FhepPEjXZ6Y0wEhy%2BrMtnHACW5hiKR%2Bv8wI0s0mwhDESSGrCB0DQmg4BjZtywb1nKysgXSEwa5SNQRREAgasyqoD9zJDsg8dd1ANlJiZWhDbIEtmXsPoJCrpokglvs5KbPsyXJ1vc%2BtYOGCXNCMSvXXJsrNAZNgQsFXkNVWtaprLYPvVZYfqgpyUFWQA88lpEhJJVEOUpyklAVyFhL39ai%2F9ZzT17C7dwcdeBxOp7Mlfcb55Z1RIz8NToyuW24vXqLw8BxFx3Qodi%2BZyzyX6e4ufp3sNmI1i5ZPU7QOaHfkpdUiX9iroUBR1MbpooKzt1vUBvGfspK9yi9RUAklKsWkzrlSOoVS8am9cqNKtIMe12f3I9F5dvCRv%2BzqXFuLjectn9wVHfnUWYhVmT%2BO1nyU8n2cap5KtByPWpO2N8nx2H2Oo2HRmQQPrH0Wx1Ab7Fg%2Bj1ujJ3%2FOWq1QJioNWHvpJzQgsgBVUS5DdekMhGuaNa2atqKl11cS1rxeT1zTruko2rVHbRWjFq5p1rRq2jWvqy1Jeb%2F%2FN%2B6vW%2FGj7keGbZnIxn%2FZHzr6J%2FtjbJiGbdgQ%2Fnv2N%2F%2BR%2Fa8t%2Fi8egJe%2BoLY%2FW8SO827D1szk5%2BlXQc2tjQ9RaPVLh25QL9SfzyPlqzW%2BRJnvHN7bO6d6HIWTqsv0Q8jPJ7PnHY%2Bd3luX97zxu9vTf7IHoGBFHEj3t%2FdB0NgVzYabsapsaA2PH9KLIqUJa%2FwyylK%2F2RhTXuzTfN94kfoulFP%2Bqzzj2r7cLW7uv78Cv%2F8BrS6k%2FQ%3D%3D*1642225781*8750998606027e08', 'untitled', 1, 'parthka', '2022-01-15 11:20:02', 2, '', '', '', 1, ''),
(49, 'UrN8ctec', 'http://app.rellitel.ink/dashboard/new', 'After', 1, 'parthka', '2022-01-15 11:28:54', 1, 'Tessa Young is a dedicated student, dutiful daughter and loyal girlfriend to her high school sweetheart.', 'thu_1_UrN8ctec.jpg', 'pre_1_UrN8ctec.jpg', 1, 'Tessa Young is a dedicated student, dutiful daughter and loyal girlfriend to her high school sweetheart. Entering her first semester of college, Tessa\'s guarded world opens up when she meets Hardin Scott, a mysterious and brooding rebel who makes her question all she thought she knew about herself -- and what she wants out of life.\nRelease date: 8 April 2019 (Los Angeles)\nDirector: <a href=\"https://www.imdb.com/name/nm1788310/\">Jenny Gage</a>\nSequel: After We Collided\nMusic by: <a href=\"https://www.imdb.com/name/nm0007212/\">Justin Burnett</a>\nAdapted from: After\nProduced by: Jennifer Gibgot; Courtney Solomon; Mark Canton; Aron Levitz; Anna Todd; Meadow Williams; Dennis Pelino'),
(50, 'wZt5S4zj', 'https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4', 'untitled', 1, 'parthka', '2022-01-15 14:18:16', 2, '', '', '', 1, ''),
(51, 'hD0tXRTm', 'https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4', 'untitled', 1, 'parthka', '2022-01-15 14:50:32', 1, '', '', '', 1, ''),
(52, 'kYKQoN91', 'https://file-examples-com.github.io/uploads/2017/04/file_example_MP4_1920_18MG.mp4', 'untitled', 1, 'parthka', '2022-01-15 14:50:46', 3, '', '', '', 1, ''),
(53, 'FmWnQSK2', 'http://rellitel.ink/assets/img/torange.png', '123', 1, 'parthka', '2022-01-15 14:59:17', 4, '', '', '', 1, ''),
(54, 'hyIXE6PJ', 'http://app.rellitel.ink/dashboard/new', '123123', 1, 'parthka', '2022-01-15 15:19:37', 1, '', '', '', 1, ''),
(56, '0DhUkLXq', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-15 18:16:16', 1, '', '', '', 1, ''),
(57, 'YYnENunq', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-15 18:16:43', 1, '', '', '', 1, ''),
(58, 'dswpTw4U', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-15 18:17:13', 1, '', '', '', 1, ''),
(59, 'DGDS23mS', 'http://app.rellitel.ink/dashboard/newd', 'untitled', 1, 'parthka', '2022-01-15 18:17:34', 1, '', '', '', 1, ''),
(60, 'PYK7vmub', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-15 18:18:25', 1, '', '', '', 1, ''),
(61, '2X8uXIu6', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-15 18:20:36', 1, '', '', '', 1, '1'),
(62, 'FNcJjARQ', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-15 18:21:36', 1, '', '', '', 1, ''),
(63, 'T2RVhVrt', 'http://google.com', 'No No Fool', 1, 'parthka', '2022-01-15 18:22:23', 1, 'A.2 Human Activities Have Caused Endangerment Of Elephants. Extinction Of These Animals Will Create An Imbalance In The Ecosystem Gravely. We Must Take Steps To Stop This Endangerment So They Can Be Protected From Extinction.', '', '', 1, 'Ho Content'),
(64, '13P6i5xC', 'https://www.youtube.com/watch?v=vQHbgxyWzJo', 'Theme Alpha by Verbana Team', 1, 'parthka', '2022-01-15 19:20:06', 1, 'Theme Alpha by Verbana Team', 'thu_1_13P6i5xC.jpg', 'pre_1_13P6i5xC.jpg', 1, 'Theme Alpha By Verbana Team'),
(65, 'link1', 'https://rr3---sn-cvh7knez.googlevideo.com/videoplayback?expire=1642267884&ei=jLDiYffTA4LmgQeF6YLICg&ip=77.243.191.120&id=o-AFVS5isDn3aijwyJ77kJE1rpJEyTPc4yawMKuMOYhzhj&itag=22&source=youtube&requiressl=yes&vprv=1&mime=video%2Fmp4&ns=9SzoLEQNFPd0xnD69k7ZQ30G&cnr=14&ratebypass=yes&dur=156.061&lmt=1604775030213267&fexp=24001373,24007246&c=WEB&txp=5432432&n=gQYjd88AoQfyOQ&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cvprv%2Cmime%2Cns%2Ccnr%2Cratebypass%2Cdur%2Clmt&sig=AOq0QJ8wRgIhALmrscbt1yy4yDrgvG042T2nW890Ad0I_47teQ5ZMloNAiEAl9bLD0nvccJr0Ch2QZYJb2V5ltz8vG3NK6x9k5zAXPg%3D&title=NO%20TIME%20TO%20DIE%20%E2%80%93%20Hindi%20Trailer&rm=sn-25gdl7s&req_id=10015acba32ea3ee&redirect_counter=2&cm2rm=sn-bu2a-nu8s7s&cms_redirect=yes&ipbypass=yes&mh=Jp&mip=103.251.59.52&mm=29&mn=sn-cvh7knez&ms=rdu&mt=1642254564&mv=m&mvi=3&pl=24&lsparams=ipbypass,mh,mip,mm,mn,ms,mv,mvi,pl&lsig=AG3C_xAwRQIgb5ghfFZRBf5JvuxTTUsDqzvV-UAsj96jcxUs193oT00CIQDESsuomoTPpl6M4nH61OQEYnK5dxlspWGb6UnqTqtmrA%3D%3D', 'NO TIME TO DIE – Hindi - Full HD', 1, 'parthka', '2022-01-15 19:40:12', 3, 'James Bond Is Enjoying A Tranquil Life In Jamaica After Leaving Active Service.', 'thu_1_rL5Fc3Gc.jpg', 'pre_1_rL5Fc3Gc.jpg', 1, 'James Bond is enjoying a tranquil life in Jamaica after leaving active service. However, his peace is short-lived as his old CIA friend, Felix Leiter, shows up and asks for help. The mission to rescue a kidnapped scientist turns out to be far more treacherous than expected, leading Bond on the trail… MORE\nRelease date: 30 September 2021 (India)\nDirector: Cary Joji Fukunaga\nBox office: 77.4 crores USD\n'),
(66, 'uHC03kau', 'https://wallpapercave.com/wp/wp2579649.png', 'untitled', 1, 'parthka', '2022-01-15 19:42:45', 4, 'Fudd Hd Wallaper', 'thu_1_uHC03kau.jpeg', '', 1, 'Thuder Wallaper'),
(67, 'nATUPRHj', 'https://wallpapercave.com/wp/wp2579649.png', 'HD Wallpaper', 1, 'parthka', '2022-01-15 20:07:41', 1, '', '', '', 0, ''),
(68, 'GmuVmxHs', 'http://app.rellitel.ink/dashboard/new', '123', 1, 'parthka', '2022-01-15 21:38:30', 1, '', '', '', 0, ''),
(69, 'yzzBbyOo', 'http://app.rellitel.ink/dashboard/new', 'http://app.rellitel.ink/dashboard/', 1, 'parthka', '2022-01-15 21:39:17', 1, '', '', '', 1, ''),
(70, 'rL5Fc3Gc', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-16 11:20:22', 1, '', '', '', 1, ''),
(71, 'uHC03kau', 'https://www.youtube.com/watch?v=LnfKij7ezr0', 'Trone Legacy', 3, '78612386', '2022-01-16 15:16:56', 1, 'Trone Legacy Full Orchesta', '', '', 1, 'Trone Legacy Full OrchestaTrone Legacy Full OrchestaTrone Legacy Full OrchestaTrone Legacy Full OrchestaTrone Legacy Full Orchesta\n'),
(72, 'nqVUSDK0', 'magnet:?xt=urn:btih:dd8255ecdc7ca55fb0bbf81323d87062db1f6d1c&dn=Big+Buck+Bunny&tr=udp%3A%2F%2Fexplodie.org%3A6969&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.empire-js.us%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=wss%3A%2F%2Ftracker.btorrent.xyz&tr=wss%3A%2F%2Ftracker.fastcast.nz&tr=wss%3A%2F%2Ftracker.openwebtorrent.com&ws=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2F&xs=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2Fbig-buck-bunny.torrent', 'untitled', 1, 'parthka', '2022-01-17 19:04:46', 1, 'adsd', 'thu_1_nqVUSDK0.png', '', 2, 'asd'),
(73, 'V6fFMTuP', 'http://app.rellitel.ink/dashboard/ne', 'Invalid link', 1, 'parthka', '2022-01-17 19:36:51', 1, '', 'thu_1_V6fFMTuP.png', '', 1, ''),
(74, 'CdWp9Qfu', 'magnet:?xt=urn:btih:dd8255ecdc7ca55fb0bbf81323d87062db1f6d1c&dn=Big+Buck+Bunny&tr=udp%3A%2F%2Fexplodie.org%3A6969&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.empire-js.us%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=wss%3A%2F%2Ftracker.btorrent.xyz&tr=wss%3A%2F%2Ftracker.fastcast.nz&tr=wss%3A%2F%2Ftracker.openwebtorrent.com&ws=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2F&xs=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2Fbig-buck-bunny.torrent', '123', 1, 'parthka', '2022-01-17 19:37:41', 1, '', '', '', 1, ''),
(75, 'WWig3e5g', 'http://localhost:3000', 'untitled', 1, 'parthka', '2022-01-17 19:43:13', 1, 'local redirect', '', '', 1, 'Local Redirect'),
(76, 'XH72r8NE', 'http://localhost:8000', 'Local Redirect', 1, 'parthka', '2022-01-17 19:43:40', 1, 'Local Redirect', '', '', 1, ''),
(77, 'lCnqsGwV', 'magnet:?xt=urn:btih:dd8255ecdc7ca55fb0bbf81323d87062db1f6d1c&dn=Big+Buck+Bunny&tr=udp%3A%2F%2Fexplodie.org%3A6969&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.empire-js.us%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=wss%3A%2F%2Ftracker.btorrent.xyz&tr=wss%3A%2F%2Ftracker.fastcast.nz&tr=wss%3A%2F%2Ftracker.openwebtorrent.com&ws=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2F&xs=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2Fbig-buck-bunny.torrent', 'untitled', 1, 'parthka', '2022-01-17 19:44:27', 1, '', '', '', 1, ''),
(78, 'qyTeCtdR', 'http://localhost:8000/movie-poster-template-design-21a1c803fe4ff4b858de24f5c91ec57f_screen.jpg', 'untitled', 1, 'parthka', '2022-01-17 19:44:56', 2, '', '', '', 1, ''),
(79, 'lF28hLZ4', 'http://localhost:8000/movie-poster-template-design-21a1c803fe4ff4b858de24f5c91ec57f_screen.jpg', 'untitled', 1, 'parthka', '2022-01-17 19:57:31', 4, '', '', '', 1, ''),
(80, 'ZS6U9NlU', 'http://localhost:8000/file_example_AVI_480_750kB%20%283%29.avi', 'Qwe', 1, 'parthka', '2022-01-17 19:59:04', 3, '', '', '', 1, ''),
(81, 'A9Wq8J58', 'magnet:?xt=urn:btih:dd8255ecdc7ca55fb0bbf81323d87062db1f6d1c&dn=Big+Buck+Bunny&tr=udp%3A%2F%2Fexplodie.org%3A6969&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.empire-js.us%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=wss%3A%2F%2Ftracker.btorrent.xyz&tr=wss%3A%2F%2Ftracker.fastcast.nz&tr=wss%3A%2F%2Ftracker.openwebtorrent.com&ws=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2F&xs=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2Fbig-buck-bunny.torrent', '1234567890', 1, 'parthka', '2022-01-17 20:06:46', 1, '123', '', '', 1, '123'),
(82, 'ONeDbgGb', 'magnet:?xt=urn:btih:dd8255ecdc7ca55fb0bbf81323d87062db1f6d1c&dn=Big+Buck+Bunny&tr=udp%3A%2F%2Fexplodie.org%3A6969&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969&tr=udp%3A%2F%2Ftracker.empire-js.us%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337&tr=wss%3A%2F%2Ftracker.btorrent.xyz&tr=wss%3A%2F%2Ftracker.fastcast.nz&tr=wss%3A%2F%2Ftracker.openwebtorrent.com&ws=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2F&xs=https%3A%2F%2Fwebtorrent.io%2Ftorrents%2Fbig-buck-bunny.torrent', '123qwe', 1, 'parthka', '2022-01-17 20:07:32', 1, '12', 'thu_1_ONeDbgGb.png', '', 1, '34'),
(83, 'V7RejsGN', 'http://localhost:8000', '123', 1, 'parthka', '2022-01-17 20:27:54', 1, '', '', '', 1, ''),
(84, 'BpLnfgDs', 'http://app.rellitel.ink/dashboard/new', '1234567890', 4, 'parthka2', '2022-01-17 21:19:19', 1, '', '', '', 1, ''),
(85, 'BpLnfgDs', 'http://app.rellitel.ink/dashboard/new', 'untitled', 6, 'parthka5', '2022-01-18 12:07:47', 1, '', '', '', 1, ''),
(86, 'BpLnfgDs', 'https://www.youtube.com/watch?v=NvD76VM3lbI', '123 link', 7, 'parthka6', '2022-01-19 09:33:34', 1, '', '', '', 1, ''),
(87, 'c2WD8F2q', 'https://www.youtube.com/watch?v=NvD76VM3lbI', 'untitled', 7, 'parthka6', '2022-01-19 09:35:16', 1, '', '', '', 1, ''),
(88, 'NfHK5a84', 'https://www.youtube.com/watch?v=NvD76VM3lbI', 'untitled', 7, 'parthka6', '2022-01-19 09:35:28', 1, '', '', '', 1, ''),
(89, 'jjJkwzDk', 'https://www.youtube.com/watch?v=NvD76VM3lbI', 'untitled', 7, 'parthka6', '2022-01-19 09:35:49', 1, '', '', '', 1, ''),
(90, 'hi3E6olI', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-19 09:36:28', 1, '', '', '', 1, ''),
(91, 'd9IL2Ghx', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-19 09:37:09', 1, '', '', '', 1, ''),
(92, '7sSGGHgm', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-19 09:37:32', 1, '', '', '', 1, ''),
(93, 'KwUMoa5S', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-19 09:38:08', 1, '', '', '', 1, ''),
(94, '4VZlr9bU', 'http://app.rellitel.ink/', 'untitled', 1, 'parthka', '2022-01-19 09:38:32', 1, '', '', '', 1, ''),
(95, 'F2lCCzFc', 'http://app.rellitel.ink/dashboard/new', 'untitled', 8, 'parthka7', '2022-01-19 10:19:54', 1, '', '', '', 0, ''),
(96, '48bzlmW3', 'http://app.rellitel.ink/dashboard/links', 'untitled', 8, 'parthka7', '2022-01-19 10:50:16', 1, '', '', '', 1, ''),
(97, 'F2lCCzFc', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-20 12:54:21', 1, 'fdfsdf', '', '', 1, 'sdfsdfsdfsdf'),
(98, '48bzlmW3', 'http://app.rellitel.ink/dashboard/new', 'untitled', 1, 'parthka', '2022-01-20 14:00:09', 1, '', 'thu_1_48bzlmW3.png', '', 1, '');

-- --------------------------------------------------------

--
-- Table structure for table `linka_news`
--

CREATE TABLE `linka_news` (
  `name` text NOT NULL,
  `link` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `linka_news`
--

INSERT INTO `linka_news` (`name`, `link`) VALUES
('Change CPC 5 to 6', ''),
('new interface', 'http://rellitel.ink/');

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

--
-- Dumping data for table `linka_pages`
--

INSERT INTO `linka_pages` (`id`, `name`, `slug`, `username`, `time`, `cat`, `info`, `theme`, `adult`) VALUES
(1, 'parthka\'s Page', 'movieflix', 'parthka', '2022-01-15 21:34:28', '{}', '{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"1_2\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}', 3, 0),
(2, 'parthkao\'s Page', 'parthkaos-page', 'parthkao', '2022-01-15 21:36:01', '{}', '{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"/assets/img/brand/g61.png\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}', 1, 0),
(3, '78612386\'s Page', '78612386s-page', '78612386', '2022-01-16 15:14:22', '{}', '{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"/assets/img/brand/g61.png\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}', 1, 0),
(4, 'parthka2\'s Page', 'parthka2s-page', 'parthka2', '2022-01-17 21:17:10', '{}', '{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"/assets/img/brand/g61.png\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}', 1, 0),
(5, 'pka', 'pka', 'parthka3', '2022-01-17 21:18:30', '{}', '{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"/assets/img/brand/g61.png\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}', 1, 0),
(6, 'parthka5\'s Page', 'parthka5s-page', 'parthka5', '2022-01-18 09:29:30', '{}', '{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"/assets/img/brand/g61.png\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}', 1, 0),
(7, 'parthka6\'s Page', 'parthka6s-page', 'parthka6', '2022-01-19 09:31:53', '{}', '{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"/assets/img/brand/g61.png\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}', 1, 0),
(8, 'parthka7\'s Page', 'parthka7s-page', 'parthka7', '2022-01-19 10:07:49', '{}', '{\"logo\":\"/assets/img/brand/g61.png\",\"logo2\":\"/assets/img/brand/g61.png\",\"des\":\"\",\"seo\":true,\"social\":{\"yt\":\"#\",\"tg\":\"#\",\"fb\":\"#\",\"ig\":\"#\"}}', 1, 0);

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

--
-- Dumping data for table `linka_users`
--

INSERT INTO `linka_users` (`id`, `username`, `email`, `password`, `fullname`, `verify`, `phone`, `photo`, `time`) VALUES
(1, 'parthka', 'parthka.2005@gmail.com', 'axb8fBIFALZJMViRD0LjCw', 'parthka boss', 1, '8866881066', NULL, '2022-01-15 21:34:28'),
(2, 'parthkao', 'parthka2@gmail.com', '1SxKpmIJab9bOxU1uHeFLw', 'parthkao', 0, NULL, NULL, '2022-01-15 21:36:01'),
(3, '78612386', 'parthka.2005@gmail.co', 'RhPlDZCvTrAJXqJXW0jWJQ', '78612386', 1, NULL, NULL, '2022-01-16 15:14:22'),
(4, 'parthka2', 'parth@123.com', 'axb8fBIFALZJMViRD0LjCw', 'parthka2', 0, NULL, NULL, '2022-01-17 21:17:10'),
(5, 'parthka3', 'parthkj@123.com', '1SxKpmIJab9bOxU1uHeFLw', 'parthka3', 0, NULL, NULL, '2022-01-17 21:18:30'),
(6, 'parthka5', 'parth@4267.com', '1SxKpmIJab9bOxU1uHeFLw', 'parthka5', 0, NULL, NULL, '2022-01-18 09:29:30'),
(7, 'parthka6', 'parth@123456.com', '1SxKpmIJab9bOxU1uHeFLw', 'parthka6', 0, NULL, NULL, '2022-01-19 09:31:53'),
(8, 'parthka7', '7@7.com', '1SxKpmIJab9bOxU1uHeFLw', 'parthka7', 0, NULL, NULL, '2022-01-19 10:07:49');

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
('MTY0MjI2MjgxMnxCZFNXM3ZzeGREN0VpeEtyQm0tX1hGcHBQeUZMWFFsQ19hdjhQQTRWRC15TktJNUo0NGR0VjJFeWZ1ZmpZTXVZT3A4OEVTaUh2TUpGRjNkTEZpVjg3OV9PN3Q3UFdqdFV80JBNLxc27_H3mtMqKmY9sIh5-tq2t7jo7iHxBjGXbw4=', 'parthka'),
('MTY0MjMyNjM3MXx1Rk5vVVJmODZJaURXLTZRUTlQdUZHcjhTZnFBOXNwdmFVUFZxRm1Ybk5FUkFDV2tTODhhLWNOVlBYQy1NRzQ0cndxR0FzY0RXV1dWSTRUWXhINVlRSm9BT3FoUjJPdUM4QT09fEqRJR5vGXe12kcrLHfsHVmZkkLRxmwpqF3Db7q5NsfC', '78612386'),
('MTY0MjMyODQ2NHxqV0lCeVhwa2lFZGsycENYUmpHTDBWbjRtNmg4ZG9XblJPQU5yU2lJa2VET2hNaTVTd2RrYkZhdVZsMk52ekRUNG1EMld2dkNvYTI0VTlXUnNxYWVEX1dxZnVlRmpqWkJ8pcIGdPlP-qw-rXdDmOTPqz3ixz-E29UXzCSCl7LZYIo=', 'parthka'),
('MTY0MjMzMTA2OXw0Mk1WZUVybWQyZTlreUxqZXhfMUNyMTIyNElXSlZrSmpLbXNNVXlkUFliSFZKbGgyVmRyRnJneFQ2MUJoN3pLSm4wYW13ZEpkSHZIaFRCRGd0dXZINUg1RW9yNGRRbmx8GL-8DkuFTcpCJcsgB9XBQDBQjQRLhNMdLVjJg7aEK7s=', 'parthka'),
('MTY0MjMzNzEwM3wtaW1YZHUxZ3NzcDJFd1NsTmJKTXVWMldXeF83NEZ3Sk9jcFFHSFRFX09iY0FBM25TcjYwYVQ0MDlUSFdySkpjdGEyZ2FDTG1EdmxYdDJZZGN6MXNEM1hOTDlaRDlzcVp8CGlOsX-NlZrgBPWeSPjB7Y2jaw8YvrBG2lQIwGJ3AHk=', 'parthka'),
('MTY0MjQwNDI1NnxoR29wZ3VvQVdNSl9rUE9WMFVjZHQwUW5zRG1mWUlycGRfTS1FT0EwT2VOOWJETTB2TGcwQ29XY0w2QTYxU3BOa2hYWm1FT2Y2Zkt3WTNEWjRQWThPZXhmOUxVdU1YWmJ8vBY-UL7zXigKkk6yJm6fXzdHd-k4wSTPbqlpTVR_zYc=', 'parthka'),
('MTY0MjQwODcxNXx3NnNzUEJDbk01c19VWUlfY1Q3MlhKNGxpY0prNmRvZUJtOU5qenRoWTg2bnFlQWpyRS1CeXE2WGdER200SlB5MG5kX2NGT0VUdkVuc2NPLTRjWkZXQjN2MzdkV1FHN1h8mgToQai2SNh1H_zQ-cz2rY0DbPCMkzjRyZQgFfeOMoQ=', 'parthka'),
('MTY0MjQxNjQ4OHxtd0FTVEQzX1RwN1Vmc3NreDFtVTdVWUZqVl91dkJla0VBUVc1NlJSVXVmejhfXzlFODFyeEE3WEFWYlo1d1lhejhqa1lYbl9DVDJCaGZsaDBTTzVZTm9vRXZVamlHZXN8suY6ZZYQ8xeP1cbYpxQ6RhV5QEXYBmYRolNUEwdp_3U=', 'parthka'),
('MTY0MjQxOTk0MXxsMEJNLUFVR2pjUmRGWmQ4dEhDdU00emlCZmk2NVJzQzlCWDVyazQ1dkV3NThxZF85R01iWmZ0WDBHUUNmd0hQUmtuRUNfUVN2TEloajJZQmRPODNiZUtFV2JJbXdSQnl8J4CjHfhHCAjt0q-Ncn5679oBhQMED5XrsyhzGACgs7Y=', 'parthka'),
('MTY0MjQyNjgyM3xRcWV2bkw5ZW1DdGlEamt1WHNfeWVfbU9fQjh3TkE0M0pfenVtSXJOZDk4VWdaVERLOGNTa2g4NmY4bEQ1d29fcEtZeFpVd25LUEh4dVNJd0NQOFBUUUJFLVowLWNPN3J8ju5URFRXZaXZNlFgwo7tV3KHkG5XXpbxUMOBywSW-fk=', 'parthka'),
('MTY0MjQzNDg3NXxDSnlBZVhrNmhFc09DZE8tZ0NNSnFNa1hpSGtQQXNRRUZ1UDd1V01yQU5ETHM5MXNZRWxJb2VNdXVweHIwZ3hyang3eFlqMHdIQmxMd0MwN3NUZDJFNjdyd25BTG5fTXV84ZlF9Z0E-LKHr1Mg418Dm6EaX7lp72RMRWSRAC_ULBo=', 'parthka'),
('MTY0MjQzNDkzNXxfekFTR3hSeU5pY0JmQWduYW5MYmxKVGF6Qm9zRzNzMnRNVUVIVVpfWmk2WkhBREYxYjFVRmdOdmdMM0RiSTM5QkpqOFliTTZQbkhLM0FqMmJ6SC1RSFlmUllJN2phVFR82Y9oG9w5rIDbDwCuHXK0J163XNT6wY2UzslLdNqNjKY=', 'parthka'),
('MTY0MjQ3NTA5NXxoS2lyVDRnTFktbkc2VHpDbVU1NjJXZC1lSDlsZkR1NEtFUHdiM1dUZzlqRjBCbXFSREpLY3VudUxWV1p4b3A2UElucXJzemhGYVhwYUg4SlNfcEw4M1c3M09Jby1nNUl8-0Jm8cfmCW-et6EyaXg1Hx_69palqoXqH-tQzyftEQI=', 'parthka'),
('MTY0MjQ4NzkxOXxXZXNXem0xR3BDNExPZl94MUppN0c1TWxOdndaZHh2X0ZoM2lfOFNaMDlzclRKZTc2T2d3VTFKYlVCdEZHSG5EMF9xR1lkandlckh1a2NVdy1oMHBqal9ybVlqYUxvZ2N81NXuQEKh3NCdUFbrmBARE40r3Iyb-qiRELRj_JWYOEw=', 'parthka'),
('MTY0MjUxNjYyNXwzUXM1X1BabWdPRTdIdEx0VkZrdlVsM19VNnBFUEItREJGQ2NIX0IwVUlMeFJXYVJ0c2wyMl9mVm5SSzRpYUZiMlBiaEcxdTZBQlBiUmpxemJxbXBQZ1l0SC05SGl0eWZ8F-WQhdmB7o4kjb_icQaT64HCvO3sZB-ZfrbKob6FOzo=', 'parthka'),
('MTY0MjU2NTE3MHxXcGFMTDJ2UmN4eGhENS1wVHowV2ZTalhnQmlQYUpDMEFETmtxVVVCbmRpblNBcDhoeTdCN29OZGF1WFNFRWEzcEZQTjg3dllvLWJJU3FEblctQ25OWjZyME45NV9tZnB8lKhYY9Qj8TJg-sG-mYK64mMd84ZVs8hyWJJhNtBCDwc=', 'parthka'),
('MTY0MjU2Njg1M3xHOU1TQXRpUEREajNzRWtfUENaUXh1ZWk0aTZOYXpUMnNwaU8zWmFZNXNoQTJ5N3pCRDQ1bUpiSDFVNU12QjNKVS1ILXI2emFqc3lFeVNUM0prRm9ndVE4WWp4M3VrRzJ8dIQ1eKiCGa2qmmi9UJIR6a40auMZDt8xTVuvxzO5v8c=', 'parthka'),
('MTY0MjU2NzA4MXw2SHpMTFIwVkFPQVhkeWNzSWJ2bnpadVdRb1RSdHVhVURNdEdRdFo4NUdZMTZldjVpeGpCZ0NtV050T3ktZTJFS1VPSzBvNkh0VFdJdXdHTnZJQkRrTnBzT19reUsxTVZ2QT09fLavTB_UzbzPzbw2jm_HHwz54Eai_-duEw8Zhj-mG06X', 'parthka7'),
('MTY0MjU3MDc1OHxyVkFSeTBYd0dIU0stWDYzdXlGaldfSjEyZ0h6SF85U2VBOG5LRzdzZzdWYzA1Qm5NN0NlMktoQ29SY1RTeHBaX2NiTndGQk5uRUVNc1dPdTJLczZZOWNHLXg4cTlxcFFnUT09fCN5CjrSudizy8D2pSBJGlCOM-yhFrkjcGlTEPrgKvVh', 'parthka6'),
('MTY0MjU3MjM0NHx5Vy1UMlZ2ekhjWGs4M21CMXhqSHJINUlKSFJqYXcyUnJIR1g3ZnkybnYtZlY2aDAxV1dtYWF2b0NxMGtmS0kwMlJ1UTJ1ZHJtWE9TWlB6dmZ3cnZqS3dVZFBPMVlqek5VUT09fEfJMzC3seAcUMeSe7-_2oKL2ECap59VfN_AUWox39uy', 'parthka6'),
('MTY0MjU3NjY5M3w0c3RTS3hmaEJOR0tzM0wyOEJaVXpjRXdESTk0bHJReVJHOTRJcS16WmtEVlJDUnA5N00tX2diZklCNlg1MWlxQXhDMzNIRzIyU0ZjWHNHMlJfVjdqNFRqNzBBMWJEWFJ8jsCOmqdZPaJ6XPgi0rlb7u-erq4ruv56FhL0RMfRvFo=', 'parthka'),
('MTY0MjY2NDQ5MnxpXzFSeE1CaFhQcFVUR29NWWNkX3dWcndZOFRNTU1hcXZ2TFhzendGaVlNRFVVRHVBOTlvdHIxaU5fUnZrUUhETVJJYmpSdkVaSEFYNVpOUXhQZm92RVdpR1F5NEpvSlV87SuMYCUBTD7nkAWO-SEcu2sGRZklkz8me1nDeZcT2xE=', 'parthka'),
('MTY0MjY5MDIyOXxmMHVEcThGdXZVTzlIekxCMzRGeGtPYWVsU1lrMHFzZlFGLXZpXy1MSXlwc3FnOWNLU0x1YWRranlpekstcE43cnVUMUdPaldnMmdSNTZnUk52VndoYWRoNVM3NUNHQnF8PqSMrwiYl0QD92DLf5sWtS_qtrCMPIhOarpsS_L6zF8=', 'parthka'),
('MTY0MjY5MjYwMXwwbHBiRHc1dloyQ296QXVEYVZiVW1WX04zRkZUcmFtMHdleWtPU2VuQWREMVJJX0ZXME5INXNYVjlOQXdROWRKUS1fME9CTllnOGtaVHYyWGVobTZXNUtXcXEtVzd6Q0F8kcbtEWEh3k2MQOyEbHG-r6UwV9ctqWJN16mhJq9Vx4c=', 'parthka'),
('MTY0MjY5NDIxOHxJQ2N5Z0M3bHBzR0l3QXVSTWdPY19pRTc1aURqWEtsX2o3SHlUZkVSSkkwbzB5bjdfMWlucTJpSXFNaVVBSF9uLXF5eXYxVnJzZ2RFdEhqUndpTWY2UnlYUHktTVhlSWZ8hru-Cza1L0OVt5-ZeCycbCXObxbDz6x8uADI0i5CEsw=', 'parthka'),
('MTY0MjczNjU4N3xBSmluTl9kRWl1SXlNUXREcllTRHMxNkVDckpUOWN6V1J0SGhFcVpHSVJZLWF6V2VsQ3ppNEtsWUF3NnhQM2RCUjZJSVg2WnlIeUVSSjNfcWI1OVZTeVJ5cUluZE1jaWF8bpag5DfIt9stJ6xFMaP48oMfWtRv6kiLB9xHbDa76nU=', 'parthka'),
('MTY0MjczNjkzMXxUZW14WVJXV2k0S3g5OVB6OUliS0lLd0xjbFlyRmwxcVJlby16NFVCSmFjZTBDM0taM0tncnd2NXNUSnpSOUFGMloyTkp4Yk5PNTZBd0VJLXhQTnQwNER2eEMtSDhhbjR8rTPrHmpYjIQrMJbvdSH226_7uvED92_eRqrxwAdUJe0=', 'parthka'),
('MTY0Mjc0MTE4Nnw1cU1fYmZEbWM2elZpMGsxellTM3hTM2NrTzkwbE9xV1lueDdJbkJid0NndlRWOXdaNkk5bTNwcHd4dVpkUTM5dkpzRG5qVTg5Rms3OVR0OHk1R0hGWWh4Y1VjempaeFV8K_w6YiWzKBOcET2cNTMr3khBc5ilhygZTO4pOTZ-s6g=', 'parthka'),
('MTY0Mjc0ODU1OXxHQ2Y3YkExWDcyeG8wSGgzQnBxa3lCOUVxVF84TTdpZGVqSGs3am4yS3pIOF9JaFhYRWVKYnNnZXhmZ0Q2ZFE4Nkdyd3VEeFlfV3lpZTBITko1LURRYjE3NmdmV2JKRU18bPgJFrp_25KxkqtG3P7-MmppkFWFheqzZzoyys77i8U=', 'parthka'),
('MTY0Mjc0OTg4NnxIUnRVcTIwTXVkamtCMUllTVpWbERmQ3N5OFN4MmI0SW41WHZJc21HTHlRVFg2cDRXSlBmamdEakpod003V3dHN1BlRjdYVE1kUjR4ekpDbE1IdWwtRm56andLYzZVdVdnQT09fIdwNG-yhsO1xgIPaydh0T9vvcGYIq51ljuBIsxnzv2z', 'parthka2'),
('MTY0Mjc3MzI2N3xrck1xUnhuNXhUMldxYUJzZ0l0UGNocnhrQVQzblR6SHA1M19oclZqQWYtUzJGRlB0X0VhNlFUT1hqTko5MlRfc0hOd0NEazEycnRHcDNaVVNlWUp4TXNLVGVtMEppa2lPdz09fKau76N4bKdXIycE1NuhYganAOo3V7cS4wgwVvQVpxyy', 'parthka2');

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
('parthka2@gmail.com', 'RTZMRDWQ4SZKRFECWRSGFUGCV3XOVMTEOLU'),
('parth@123.com', 'R2GPFSEO2CBJVKES6SMNVHHMYSRHU6Q===='),
('parthkj@123.com', 'UTPMEZHAQKNKREXUTDNJZ3GEUJ5HU======'),
('parth@4267.com', 'SCCNBRW4UTPKFCFC6KONJRXKWJSHF2A===='),
('parth@123456.com', 'RTZMRDWQQKNKREXUTSEKUZEY3KOOZRFCPJ5'),
('7@7.com', 'TSEJUYU2RDHGVHDAQJTJRWU45TCKE6T2===');

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
-- Dumping data for table `linka_wallet`
--

INSERT INTO `linka_wallet` (`id`, `username`, `balance`) VALUES
(1, 'parthka', 317),
(2, 'parthkao', 0.196552),
(3, '78612386', 0.477103),
(4, 'parthka2', 33),
(5, 'parthka3', 0),
(6, 'parthka5', 0.00165517),
(7, 'parthka6', 0.0032069),
(8, 'parthka7', 0.003);

-- --------------------------------------------------------

--
-- Table structure for table `linka_withdraw`
--

CREATE TABLE `linka_withdraw` (
  `id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `nowbal` float NOT NULL,
  `por` int(11) NOT NULL,
  `pom` text NOT NULL,
  `poa` text NOT NULL,
  `date` datetime NOT NULL DEFAULT current_timestamp(),
  `status` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `linka_withdraw`
--

INSERT INTO `linka_withdraw` (`id`, `username`, `nowbal`, `por`, `pom`, `poa`, `date`, `status`) VALUES
(1, 'parthka', 5.222, 4, 'paypal', 'parthka.2005@gmail.com', '2022-01-21 12:46:51', 1),
(2, 'parthka', 503, 5, 'phonepe', '8866881066', '2022-01-21 12:47:24', 0),
(3, 'parthka', 500, 5, 'paypal', 'parthka.2005@gmail.com', '2022-01-21 12:51:40', 1),
(4, 'parthka', 495, 4, 'paypal', 'parthka.2005@gmail.com', '2022-01-21 12:53:08', 1),
(5, 'parthka', 491, 7, 'paypal', 'parthka.2005@gmail.com', '2022-01-21 12:53:51', 1),
(6, 'parthka2', 50, 17, 'google', 'parthka.2005@gmail.com', '2022-01-21 12:55:13', 1),
(7, 'parthka', 484, 3, 'paypal', 'parthka.2005@gmail.com', '2022-01-21 12:56:07', 1),
(8, 'parthka', 481, 10, 'upi', 'parthka.2005@gmail.com', '2022-01-21 13:01:00', 1),
(9, 'parthka', 471, 8, 'upi', '8866881066@122', '2022-01-21 13:44:27', 1),
(10, 'parthka', 463, 123, 'google', 'parthka.2005@gmail.com', '2022-01-21 18:00:43', 1),
(11, 'parthka', 340, 12, 'upi', 'parthka.2005@gmail.com', '2022-01-21 18:04:54', 1),
(12, 'parthka', 328, 11, 'amazon', 'parthka@helka.com', '2022-01-21 18:47:36', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `linka_analytics`
--
ALTER TABLE `linka_analytics`
  ADD PRIMARY KEY (`id`);

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
-- Indexes for table `linka_withdraw`
--
ALTER TABLE `linka_withdraw`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `linka_analytics`
--
ALTER TABLE `linka_analytics`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=62;

--
-- AUTO_INCREMENT for table `linka_links`
--
ALTER TABLE `linka_links`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=99;

--
-- AUTO_INCREMENT for table `linka_pages`
--
ALTER TABLE `linka_pages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `linka_report`
--
ALTER TABLE `linka_report`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `linka_users`
--
ALTER TABLE `linka_users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `linka_wallet`
--
ALTER TABLE `linka_wallet`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `linka_withdraw`
--
ALTER TABLE `linka_withdraw`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
