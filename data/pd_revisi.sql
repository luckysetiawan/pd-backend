-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 29 Nov 2021 pada 17.19
-- Versi server: 10.4.21-MariaDB
-- Versi PHP: 8.0.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pd_revisi`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `customer`
--

CREATE TABLE `customer` (
  `email` varchar(255) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `no_telp` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `customer`
--

INSERT INTO `customer` (`email`, `nama`, `no_telp`) VALUES
('asdf', 'asdf', 'asdf'),
('test@gmail.com', 'testtt', '0818763984'),
('testt@gmail.com', 'testtt', '0818763984');

-- --------------------------------------------------------

--
-- Struktur dari tabel `order`
--

CREATE TABLE `order` (
  `id` int(50) NOT NULL,
  `customer_email` varchar(255) NOT NULL,
  `waktu` datetime NOT NULL,
  `alamat` varchar(255) NOT NULL,
  `status` int(11) NOT NULL,
  `rating` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `order`
--

INSERT INTO `order` (`id`, `customer_email`, `waktu`, `alamat`, `status`, `rating`) VALUES
(1, 'test@email.com', '2021-11-25 14:35:42', 'Pluto', 0, 0),
(3, 'test@gmail.com', '0000-00-00 00:00:00', 'bumi', 0, 5),
(4, 'testt@gmail.com', '0000-00-00 00:00:00', 'bumi', 0, 5),
(5, 'testt@gmail.com', '2021-11-29 23:11:28', 'bumi', 0, 5);

-- --------------------------------------------------------

--
-- Struktur dari tabel `orderdetail`
--

CREATE TABLE `orderdetail` (
  `id` int(50) NOT NULL,
  `pizza_id` int(50) NOT NULL,
  `order_id` int(50) NOT NULL,
  `quantity` int(50) NOT NULL,
  `total_harga` decimal(10,0) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `orderdetail`
--

INSERT INTO `orderdetail` (`id`, `pizza_id`, `order_id`, `quantity`, `total_harga`) VALUES
(1, 1, 0, 2, '100000'),
(35, 1, 0, 1, '0'),
(36, 2, 0, 2, '0'),
(37, 1, 0, 1, '0'),
(38, 2, 0, 2, '0'),
(39, 1, 0, 1, '0'),
(40, 2, 0, 2, '0');

-- --------------------------------------------------------

--
-- Struktur dari tabel `payment`
--

CREATE TABLE `payment` (
  `id` varchar(255) NOT NULL,
  `order_id` int(11) DEFAULT NULL,
  `status_pembayaran` int(10) NOT NULL,
  `total_pembayaran` decimal(10,0) NOT NULL,
  `waktu_pembayaran` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `payment`
--

INSERT INTO `payment` (`id`, `order_id`, `status_pembayaran`, `total_pembayaran`, `waktu_pembayaran`) VALUES
('1', 0, 1, '100000', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Struktur dari tabel `pizza`
--

CREATE TABLE `pizza` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `deskripsi` text NOT NULL,
  `harga` decimal(11,0) NOT NULL,
  `gambar` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `pizza`
--

INSERT INTO `pizza` (`id`, `nama`, `deskripsi`, `harga`, `gambar`) VALUES
(1, 'Beef Delight', 'Lorem Ipsum Dolor Sit Amet', '115000', 'beef-delight.png'),
(2, 'Beef Pepperoni Feast', 'Lorem Ipsum Dolor Sit Amet', '125000', 'beef-pepperoni-feast.png'),
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
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `notelp` varchar(255) NOT NULL,
  `position` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `nama`, `email`, `password`, `notelp`, `position`) VALUES
(0, 'test', 'test@gmail.com', '', '0819238923', 'BOSQU'),
(1, 'Dave', 'dave@gmail.com', 'dave123', '089123123123', 'BOSQU'),
(2, 'Christian', 'christian@gmail.com', 'christian123', '089098098098', 'BOSQU'),
(3, 'Lucky', 'lucky@gmail.com', 'lucky123', '089234234234', 'BOSQU'),
(4, 'Gilbert', 'gilbert@gmail.com', 'gilbert123', '089987987987', 'BOSQU'),
(5, 'MelvinSebastian', 'melvin@gmail.com', 'melvin123', '087785705296', 'BOSQU');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`email`);

--
-- Indeks untuk tabel `order`
--
ALTER TABLE `order`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idCustomer` (`customer_email`);

--
-- Indeks untuk tabel `orderdetail`
--
ALTER TABLE `orderdetail`
  ADD PRIMARY KEY (`id`),
  ADD KEY `menu` (`pizza_id`);

--
-- Indeks untuk tabel `payment`
--
ALTER TABLE `payment`
  ADD PRIMARY KEY (`id`),
  ADD KEY `order` (`order_id`);

--
-- Indeks untuk tabel `pizza`
--
ALTER TABLE `pizza`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `order`
--
ALTER TABLE `order`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=54;

--
-- AUTO_INCREMENT untuk tabel `orderdetail`
--
ALTER TABLE `orderdetail`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;

--
-- AUTO_INCREMENT untuk tabel `pizza`
--
ALTER TABLE `pizza`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
