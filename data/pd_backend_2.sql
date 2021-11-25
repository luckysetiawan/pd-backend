-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 25, 2021 at 08:49 AM
-- Server version: 10.4.13-MariaDB
-- PHP Version: 7.4.8

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
-- Table structure for table `order`
--

CREATE TABLE `order` (
  `id` int(50) NOT NULL,
  `idCustomer` int(50) NOT NULL,
  `idOrderDetail` int(50) NOT NULL,
  `invoice` varchar(255) NOT NULL,
  `waktu` datetime NOT NULL,
  `alamat` varchar(255) NOT NULL,
  `status` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `order`
--

INSERT INTO `order` (`id`, `idCustomer`, `idOrderDetail`, `invoice`, `waktu`, `alamat`, `status`) VALUES
(0, 1, 1, 'P-2021-11-25 14:35:42', '2021-11-25 14:35:42', 'Pluto', 0),
(40, 2, 2, 'P-2021-11-24 19:03:15', '2021-11-24 19:03:15', 'MARSSSSSSS', 2),
(41, 2, 2, 'P-2021-11-24 20:39:19', '2021-11-24 20:39:19', 'Moon', 1);

-- --------------------------------------------------------

--
-- Table structure for table `orderdetail`
--

CREATE TABLE `orderdetail` (
  `id` int(50) NOT NULL,
  `pizza` int(50) NOT NULL,
  `rating` int(50) NOT NULL,
  `quantity` int(50) NOT NULL,
  `totalPesanan` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orderdetail`
--

INSERT INTO `orderdetail` (`id`, `pizza`, `rating`, `quantity`, `totalPesanan`) VALUES
(1, 1, 5, 2, 100000);

-- --------------------------------------------------------

--
-- Table structure for table `payment`
--

CREATE TABLE `payment` (
  `invoice` varchar(255) NOT NULL,
  `statusPembayaran` int(10) NOT NULL,
  `totalHarga` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `payment`
--

INSERT INTO `payment` (`invoice`, `statusPembayaran`, `totalHarga`) VALUES
('pay01', 1, 100000);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `notelp` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama`, `email`, `password`, `notelp`) VALUES
(1, 'Dave', 'dave@gmail.com', 'dave123', '089123123123'),
(2, 'Christian', 'christian@gmail.com', 'christian123', '089098098098'),
(3, 'Lucky', 'lucky@gmail.com', 'lucky123', '089234234234'),
(4, 'Gilbert', 'gilbert@gmail.com', 'gilbert123', '089987987987'),
(5, 'MelvinSebastian', 'melvin@gmail.com', 'melvin123', '087785705296');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `order`
--
ALTER TABLE `order`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idOrderDetail` (`idOrderDetail`),
  ADD KEY `invoice` (`invoice`),
  ADD KEY `idCustomer` (`idCustomer`);

--
-- Indexes for table `orderdetail`
--
ALTER TABLE `orderdetail`
  ADD PRIMARY KEY (`id`),
  ADD KEY `pizza` (`pizza`);

--
-- Indexes for table `payment`
--
ALTER TABLE `payment`
  ADD PRIMARY KEY (`invoice`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
