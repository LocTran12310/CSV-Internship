-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Máy chủ: 127.0.0.1
-- Thời gian đã tạo: Th8 13, 2021 lúc 09:08 AM
-- Phiên bản máy phục vụ: 10.4.20-MariaDB
-- Phiên bản PHP: 8.0.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Cơ sở dữ liệu: `golang`
--

-- --------------------------------------------------------

--
-- Cấu trúc bảng cho bảng `profiles`
--

CREATE TABLE `profiles` (
  `id` int(11) NOT NULL,
  `employee_id` varchar(15) COLLATE utf8_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `birthday` datetime DEFAULT NULL,
  `position_id` int(11) NOT NULL,
  `department_id` int(11) NOT NULL,
  `status` tinyint(1) DEFAULT NULL COMMENT '1: Available, 2: Unavailable',
  `address` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `telephone` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `mobile` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL,
  `official_date` date DEFAULT NULL,
  `probation_date` date DEFAULT NULL,
  `gender` tinyint(1) DEFAULT NULL COMMENT '1:nam, 2: nữ',
  `image` blob DEFAULT NULL,
  `del_flag` tinyint(1) DEFAULT 0 COMMENT 'mặc định bằng 0\r\n1: bị xóa',
  `created_time` datetime DEFAULT NULL,
  `created_user` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'users login id',
  `updated_time` datetime DEFAULT NULL,
  `updated_user` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'users login id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Đang đổ dữ liệu cho bảng `profiles`
--

INSERT INTO `profiles` (`id`, `employee_id`, `name`, `email`, `birthday`, `position_id`, `department_id`, `status`, `address`, `telephone`, `mobile`, `official_date`, `probation_date`, `gender`, `image`, `del_flag`, `created_time`, `created_user`, `updated_time`, `updated_user`) VALUES
(1, '090909', 'Tran Phuoc Loc', 'email@mail.com', '2021-08-13 12:13:14', 1, 1, 1, 'abca acb, dhc, Vietnam', '0809123456', '0123456789', '2021-08-09', '2021-08-09', 1, '', 0, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', ''),
(2, '234567', 'Employee So Mot', 'email@mail.com', '0000-00-00 00:00:00', 1, 1, 0, '', '', '', '0000-00-00', '0000-00-00', 0, '', 0, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', ''),
(3, '101422', 'Tran Phuoc Loc', 'email@mail.com', '2021-08-13 12:13:14', 1, 1, 1, 'abca acb, dhc, Vietnam', '0809123456', '0123456789', '2021-08-09', '2021-08-09', 1, '', 0, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', ''),
(4, '555555', 'Tran Phuoc Loc', 'email@mail.com', '2021-08-13 12:13:14', 1, 1, 1, 'abca acb, dhc, Vietnam', '0809123456', '0123456789', '2021-08-09', '2021-08-09', 1, '', 0, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', ''),
(5, '444444', 'Tran Phuoc Loc', 'email@mail.com', '2021-08-13 12:13:14', 1, 1, 1, 'abca acb, dhc, Vietnam', '0809123456', '0123456789', '2021-08-09', '2021-08-09', 1, '', 0, '0000-00-00 00:00:00', '', '0000-00-00 00:00:00', ''),
(6, '666666', 'Tran Phuoc Loc', 'email@mail.com', '2021-08-13 12:13:14', 1, 1, 1, 'abca acb, dhc, Vietnam', '0809123456', '0123456789', '2021-08-09', '2021-08-09', 1, '', 0, '2021-08-13 13:54:51', '', '0000-00-00 00:00:00', ''),
(7, '77777', 'Tran Phuoc Loc', 'email@mail.com', '2021-08-13 12:13:14', 1, 1, 1, 'abca acb, dhc, Vietnam', '0809123456', '0123456789', '2021-08-09', '2021-08-09', 1, '', 0, '2021-08-13 13:56:05', '', '0000-00-00 00:00:00', ''),
(8, '888888', 'Tran Phuoc Loc', 'email@mail.com', '2021-08-13 12:13:14', 1, 1, 1, 'abca acb, dhc, Vietnam', '0809123456', '0123456789', '2021-08-09', '2021-08-09', 1, '', 0, '2021-08-13 13:58:51', '', '0000-00-00 00:00:00', ''),
(9, '999999', 'Edit 2 Tran Phuoc Loc', 'email@mail.com', '2021-08-13 12:13:14', 1, 2, 1, 'abca acb, dhc, Vietnam', '0809123456', '0123456789', '2021-08-09', '2021-08-09', 1, '', 0, '2021-08-13 14:00:40', '', '2021-08-13 14:02:02', '');

--
-- Chỉ mục cho các bảng đã đổ
--

--
-- Chỉ mục cho bảng `profiles`
--
ALTER TABLE `profiles`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT cho các bảng đã đổ
--

--
-- AUTO_INCREMENT cho bảng `profiles`
--
ALTER TABLE `profiles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
