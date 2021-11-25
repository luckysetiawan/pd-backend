-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 25 Nov 2021 pada 12.12
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
-- Database: `pd_backend_2`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `menus`
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
-- Dumping data untuk tabel `menus`
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
-- Struktur dari tabel `order`
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
-- Dumping data untuk tabel `order`
--

INSERT INTO `order` (`id`, `idCustomer`, `idOrderDetail`, `invoice`, `waktu`, `alamat`, `status`) VALUES
(1, 1, 1, 'P-2021-11-25 14:35:42', '2021-11-25 14:35:42', 'Pluto', 0),
(40, 2, 2, 'P-2021-11-24 19:03:15', '2021-11-24 19:03:15', 'MARSSSSSSS', 2),
(41, 2, 2, 'P-2021-11-24 20:39:19', '2021-11-24 20:39:19', 'Moon', 1),
(45, 2, 2, 'P-2021-11-25 16:47:36', '2021-11-25 16:47:36', 'Pluto', 1),
(46, 2, 2, 'P-2021-11-25 16:51:56', '2021-11-25 16:51:56', 'MARSSSSSSS', 0),
(47, 2, 2, 'P-2021-11-25 16:54:07', '2021-11-25 16:54:07', 'Angkasa', 2),
(49, 2, 1, 'P-2021-11-25 17:02', '2021-11-25 17:02:25', 'MARSSSSSSS', 0),
(50, 2, 1, 'P-2021-11-25 17:22', '2021-11-25 17:22:03', 'MARSSSSSSS', 0);

-- --------------------------------------------------------

--
-- Struktur dari tabel `orderdetail`
--

CREATE TABLE `orderdetail` (
  `id` int(50) NOT NULL,
  `menu` int(50) NOT NULL,
  `rating` int(50) NOT NULL,
  `quantity` int(50) NOT NULL,
  `totalPesanan` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `orderdetail`
--

INSERT INTO `orderdetail` (`id`, `menu`, `rating`, `quantity`, `totalPesanan`) VALUES
(1, 1, 5, 2, 100000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `payment`
--

CREATE TABLE `payment` (
  `invoice` varchar(255) NOT NULL,
  `statusPembayaran` int(10) NOT NULL,
  `totalHarga` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `payment`
--

INSERT INTO `payment` (`invoice`, `statusPembayaran`, `totalHarga`) VALUES
('P-2021-11-25 14:35:42', 1, 100000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `notelp` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `nama`, `email`, `password`, `notelp`) VALUES
(0, 'test', 'test@gmail.com', '', '0819238923'),
(1, 'Dave', 'dave@gmail.com', 'dave123', '089123123123'),
(2, 'Christian', 'christian@gmail.com', 'christian123', '089098098098'),
(3, 'Lucky', 'lucky@gmail.com', 'lucky123', '089234234234'),
(4, 'Gilbert', 'gilbert@gmail.com', 'gilbert123', '089987987987'),
(5, 'MelvinSebastian', 'melvin@gmail.com', 'melvin123', '087785705296');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `menus`
--
ALTER TABLE `menus`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `order`
--
ALTER TABLE `order`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idOrderDetail` (`idOrderDetail`),
  ADD KEY `invoice` (`invoice`),
  ADD KEY `idCustomer` (`idCustomer`);

--
-- Indeks untuk tabel `orderdetail`
--
ALTER TABLE `orderdetail`
  ADD PRIMARY KEY (`id`),
  ADD KEY `menu` (`menu`);

--
-- Indeks untuk tabel `payment`
--
ALTER TABLE `payment`
  ADD PRIMARY KEY (`invoice`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `menus`
--
ALTER TABLE `menus`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT untuk tabel `order`
--
ALTER TABLE `order`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=52;

--
-- AUTO_INCREMENT untuk tabel `orderdetail`
--
ALTER TABLE `orderdetail`
  MODIFY `id` int(50) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `orderdetail`
--
ALTER TABLE `orderdetail`
  ADD CONSTRAINT `menu` FOREIGN KEY (`menu`) REFERENCES `menus` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
