-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 14, 2021 at 02:11 AM
-- Server version: 10.4.21-MariaDB
-- PHP Version: 8.0.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pd_backend`
--

-- --------------------------------------------------------

--
-- Table structure for table `customer`
--

CREATE TABLE `customer` (
  `email` varchar(255) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `no_telp` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `customer`
--

INSERT INTO `customer` (`email`, `nama`, `no_telp`) VALUES
('', '', ''),
('asdfasdfs@gmail.com', 'christ', '0818763984'),
('christ@gmail.com', 'christ', '0818763984'),
('pembeli@gmail.com', 'pembeli', '081758934'),
('plutoo@gmail.com', 'christ', '0818763984'),
('test@gmail.com', 'test', '0814872384'),
('testt@gmail.com', 'testt', '081876398412');

-- --------------------------------------------------------

--
-- Table structure for table `order`
--

CREATE TABLE `order` (
  `id` int(50) NOT NULL,
  `customer_email` varchar(255) NOT NULL,
  `waktu` datetime NOT NULL,
  `alamat` varchar(255) NOT NULL,
  `status` int(50) NOT NULL,
  `rating` int(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `order`
--

INSERT INTO `order` (`id`, `customer_email`, `waktu`, `alamat`, `status`, `rating`) VALUES
(1, 'pembeli@gmail.com', '2021-12-02 13:53:07', 'Pluto', 0, 8),
(2, 'test@gmail.com', '2021-12-02 13:57:34', 'Milky Way', 1, 8),
(18, 'testt@gmail.com', '2021-12-07 01:51:29', 'bumi', 0, 5),
(19, 'testt@gmail.com', '2021-12-07 01:55:16', 'bumi', 0, 5),
(50, '', '2021-12-14 08:06:42', '', 0, 0),
(51, '', '2021-12-14 08:06:53', '', 0, 0);

-- --------------------------------------------------------

--
-- Table structure for table `orderdetail`
--

CREATE TABLE `orderdetail` (
  `id` int(50) NOT NULL,
  `pizza_id` int(50) NOT NULL,
  `order_id` int(50) NOT NULL,
  `quantity` int(50) NOT NULL,
  `total_harga` float(50,0) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orderdetail`
--

INSERT INTO `orderdetail` (`id`, `pizza_id`, `order_id`, `quantity`, `total_harga`) VALUES
(1, 1, 1, 1, 100000),
(44, 2, 18, 2, 200000),
(45, 2, 19, 2, 200000);

-- --------------------------------------------------------

--
-- Table structure for table `payment`
--

CREATE TABLE `payment` (
  `id` int(50) NOT NULL,
  `order_id` int(50) NOT NULL,
  `status_pembayaran` int(50) NOT NULL,
  `total_pembayaran` float(50,0) NOT NULL,
  `waktu_pembayaran` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `payment`
--

INSERT INTO `payment` (`id`, `order_id`, `status_pembayaran`, `total_pembayaran`, `waktu_pembayaran`) VALUES
(1, 1, 1, 100000, '2021-12-02 13:58:34'),
(2, 2, 0, 200000, '2021-12-02 13:58:48'),
(3, 19, 0, 200000, '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `pizza`
--

CREATE TABLE `pizza` (
  `id` int(50) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `deskripsi` text NOT NULL,
  `harga` decimal(50,0) NOT NULL,
  `gambar` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `pizza`
--

INSERT INTO `pizza` (`id`, `nama`, `deskripsi`, `harga`, `gambar`) VALUES
(1, 'Beef Delight', 'Lorem Ipsum Dolor Sit Amet', '100000', 'beef-delight.png'),
(2, 'Beef Pepperoni Feast', 'Lorem Ipsum Dolor Sit Amet', '100000', 'beef-pepperoni-feast.png'),
(3, 'Beef Rasher', 'Lorem Ipsum Dolor Sit Amet', '150000', 'beef-rasher.png'),
(4, 'Cheesy Sausage', 'Lorem Ipsum Dolor Sit Amet', '135000', 'cheesy-sausage.png'),
(5, 'Chicken Delight', 'Lorem Ipsum Dolor Sit Amet', '110000', 'chicken-delight.png'),
(6, 'Chicken Pepperoni Feast', 'Lorem Ipsum Dolor Sit Amet', '125000', 'chicken-pepperoni-feast.png'),
(7, 'Chicken Sausage', 'Lorem Ipsum Dolor Sit Amet', '115000', 'chicken-sausage.png'),
(8, 'Chili Chicken', 'Lorem Ipsum Dolor Sit Amet', '110000', 'chili-chicken.png'),
(9, 'Double Beef Burger', 'Lorem Ipsum Dolor Sit Amet', '95000', 'double-beef-burger.png'),
(10, 'Margherita', 'Lorem Ipsum Dolor Sit Amet', '100000', 'margherita.png'),
(11, 'Meat Meat', 'Lorem Ipsum Dolor Sit Amet', '155000', 'meat-meat.png'),
(12, 'Tuna Delight', 'Lorem Ipsum Dolor Sit Amet', '175000', 'tuna-delight.png');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(50) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `no_telp` varchar(255) NOT NULL,
  `position` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama`, `email`, `password`, `no_telp`, `position`) VALUES
(1, 'owner', 'owner@gmail.com', 'owner', '123', 'owner'),
(2, 'staff', 'staff@gmail.com', 'staff', '123', 'staff'),
(3, 'chef', 'chef@gmail.com', 'chef', '123', 'chef'),
(4, 'delivery', 'delivery@gmail.com', 'delivery', '123', 'delivery');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`email`);

--
-- Indexes for table `order`
--
ALTER TABLE `order`
  ADD PRIMARY KEY (`id`),
  ADD KEY `customer_email` (`customer_email`);

--
-- Indexes for table `orderdetail`
--
ALTER TABLE `orderdetail`
  ADD PRIMARY KEY (`id`),
  ADD KEY `pizza_id` (`pizza_id`),
  ADD KEY `order_id` (`order_id`);

--
-- Indexes for table `payment`
--
ALTER TABLE `payment`
  ADD PRIMARY KEY (`id`),
  ADD KEY `order` (`order_id`);

--
-- Indexes for table `pizza`
--
ALTER TABLE `pizza`
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
-- AUTO_INCREMENT for table `order`
--
ALTER TABLE `order`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=52;

--
-- AUTO_INCREMENT for table `orderdetail`
--
ALTER TABLE `orderdetail`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=106;

--
-- AUTO_INCREMENT for table `payment`
--
ALTER TABLE `payment`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;

--
-- AUTO_INCREMENT for table `pizza`
--
ALTER TABLE `pizza`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `order`
--
ALTER TABLE `order`
  ADD CONSTRAINT `customer_email` FOREIGN KEY (`customer_email`) REFERENCES `customer` (`email`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `orderdetail`
--
ALTER TABLE `orderdetail`
  ADD CONSTRAINT `order_id` FOREIGN KEY (`order_id`) REFERENCES `order` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `pizza_id` FOREIGN KEY (`pizza_id`) REFERENCES `pizza` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;

--
-- Constraints for table `payment`
--
ALTER TABLE `payment`
  ADD CONSTRAINT `order` FOREIGN KEY (`order_id`) REFERENCES `order` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
