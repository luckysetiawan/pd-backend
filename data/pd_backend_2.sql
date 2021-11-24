-- phpMyAdmin SQL Dump
-- version 5.0.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 24, 2021 at 08:26 AM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pd_backend_2`
--

-- --------------------------------------------------------

--
-- Table structure for table `menus`
--

CREATE TABLE `menus` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `deskripsi` text NOT NULL,
  `harga` int(11) NOT NULL,
  `gambar` varchar(255) NOT NULL,
  `varian` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `menus`
--

INSERT INTO `menus` (`id`, `nama`, `deskripsi`, `harga`, `gambar`, `varian`) VALUES
(1, 'Beef Delight', 'Lorem Ipsum Dolor Sit Amet', 115000, 'beef-delight.png', 'BD'),
(2, 'Beef Pepperoni Feast', 'Lorem Ipsum Dolor Sit Amet', 125000, 'beef-pepperoni-feast.png', 'BPF'),
(3, 'Beef Rasher', 'Lorem Ipsum Dolor Sit Amet', 150000, 'beef-rasher.png', 'BR'),
(4, 'Cheesy Sausage', 'Lorem Ipsum Dolor Sit Amet', 135000, 'cheesy-sausage.png', 'CSS'),
(5, 'Chicken Delight', 'Lorem Ipsum Dolor Sit Amet', 110000, 'chicken-delight.png', 'CD'),
(6, 'Chicken Pepperoni Feast', 'Lorem Ipsum Dolor Sit Amet', 125000, 'chicken-pepperoni-feast.png', 'CPF'),
(7, 'Chicken Sausage', 'Lorem Ipsum Dolor Sit Amet', 115000, 'chicken-sausage.png', 'CHS'),
(8, 'Chili Chicken', 'Lorem Ipsum Dolor Sit Amet', 110000, 'chili-chicken.png', 'CC'),
(9, 'Double Beef Burger', 'Lorem Ipsum Dolor Sit Amet', 95000, 'double-beef-burger.png', 'DBB'),
(10, 'Margherita', 'Lorem Ipsum Dolor Sit Amet', 100000, 'margherita.png', 'MR'),
(11, 'Meat Meat', 'Lorem Ipsum Dolor Sit Amet', 155000, 'meat-meat.png', 'MM'),
(12, 'Tuna Delight', 'Lorem Ipsum Dolor Sit Amet', 175000, 'tuna-delight.png', 'TD');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `notelp` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama`, `email`, `notelp`) VALUES
(1, 'Dave', 'dave@gmail.com', '089123123123'),
(2, 'Christian', 'christian@gmail.com', '089098098098'),
(3, 'Lucky', 'lucky@gmail.com', '089234234234'),
(4, 'Gilbert', 'gilbert@gmail.com', '089987987987'),
(5, 'MelvinSebastian', 'melvin@gmail.com', '012345678901');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `menus`
--
ALTER TABLE `menus`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `menus`
--
ALTER TABLE `menus`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
