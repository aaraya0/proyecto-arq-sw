-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 16-07-2022 a las 22:33:34
-- Versión del servidor: 10.4.24-MariaDB
-- Versión de PHP: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Base de datos: `sistema`
--
CREATE DATABASE IF NOT EXISTS `sistema` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `sistema`;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `addresses`
--

CREATE TABLE `addresses` (
  `id` int(11) UNSIGNED NOT NULL,
  `user_id` int(11) NOT NULL,
  `cp` int(11) NOT NULL,
  `number` int(11) NOT NULL,
  `street` varchar(550) NOT NULL,
  `city` varchar(550) NOT NULL,
  `country` varchar(550) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `addresses`:
--

--
-- Volcado de datos para la tabla `addresses`
--

INSERT INTO `addresses` VALUES
(1, 2, 5009, 6017, 'branly', 'cordoba', 'argentina'),
(2, 2, 5000, 3232, 'ricardo rojas', 'cordoba', 'argentina');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `categories`
--

CREATE TABLE `categories` (
  `id` smallint(11) UNSIGNED NOT NULL,
  `name` varchar(250) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `categories`:
--

--
-- Volcado de datos para la tabla `categories`
--

INSERT INTO `categories` VALUES
(1, 'Juvenil'),
(2, 'Policial'),
(3, 'Terror'),
(4, 'Fantasia'),
(5, 'Romance'),
(6, 'Infantil'),
(7, 'Thriller'),
(8, 'Historia'),
(9, 'Distopia'),
(10, 'Novela');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `details`
--

CREATE TABLE `details` (
  `Id` smallint(5) UNSIGNED NOT NULL,
  `unit_price` float NOT NULL,
  `quantity` int(11) NOT NULL,
  `total` int(11) NOT NULL,
  `name` varchar(550) NOT NULL,
  `product_id` smallint(5) UNSIGNED NOT NULL,
  `order_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `details`:
--   `Id`
--       `details` -> `Id`
--   `product_id`
--       `products` -> `id`
--   `product_id`
--       `products` -> `id`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `orders`
--

CREATE TABLE `orders` (
  `id` int(11) NOT NULL,
  `date` date NOT NULL,
  `user_id` smallint(5) UNSIGNED NOT NULL,
  `total_amount` float UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `orders`:
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `products`
--

CREATE TABLE `products` (
  `id` smallint(11) UNSIGNED NOT NULL,
  `title` varchar(2500) NOT NULL,
  `author` varchar(250) NOT NULL,
  `category_id` smallint(11) NOT NULL,
  `stock` int(11) NOT NULL,
  `price` float NOT NULL,
  `image` varchar(250) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `products`:
--

--
-- Volcado de datos para la tabla `products`
--

INSERT INTO `products` VALUES
(1, 'Orgullo y Prejucio', 'Jane Austin', 5, 5, 2000, 'orgullo.jpg'),
(2, 'Los juegos del hambre', 'Suzanne Collins', 9, 11, 1000, 'libro1.jpg'),
(3, 'Hamlet', 'William Shakespeare', 10, 4, 1600, 'hamlet.jpg'),
(4, 'Harry Potter', 'J.K. Rowlling', 4, 20, 3000, 'harrypotter.jpg'),
(5, 'Mujercitas', 'Louisa May Alcott', 10, 10, 2500, 'mujercitas.jpg'),
(6, 'Un cuento perfecto', 'Elisabet Benavent', 5, 4, 1800, 'cuento.jpg'),
(7, 'El mentalista', 'Camila Läckberg', 2, 12, 2100, 'mentalista.jpg'),
(8, 'El principito', 'Antoine Saint-Exupéry', 6, 4, 1010, 'principito.jpg'),
(9, 'Cazadores de Sombras', 'Cassandra Clare', 1, 19, 2300, 'cazadores.jpg'),
(10, 'It', 'Stephen King', 3, 10, 3000, 'it.jpg'),
(11, 'El Club de los Psicopatas', 'Jhon Katzenbach', 7, 8, 3100, 'psicopatas.jpg'),
(12, 'Las Señoritas', 'Laura Ramos', 8, 9, 2100, 'señoritas.jpg'),
(13, 'Momentos estelares de la humanidad', 'Stefan Zweig', 8, 3, 3200, 'momentos.jpg'),
(14, 'Los Griegos', 'Isaac Asimov', 8, 5, 3000, 'griegos.jpg'),
(15, 'Breve historia del mundo', 'Ernst H. Gombrich', 8, 1, 4200, 'mundo.jpg'),
(16, 'Cuentos Policiales', 'Edgar Allan Poe', 2, 3, 3200, 'cuentos.jpg'),
(17, 'Un cadaver en la biblioteca', 'Agatha Christie', 2, 5, 3000, 'biblioteca.jpg'),
(18, 'Diez Negritos', 'Agatha Christie', 2, 1, 4200, 'diez.jpg'),
(19, 'El Psicoanalista', 'John Katzenbach', 2, 1, 3200, 'psicoanalista.jpg'),
(20, 'Lo que el viento se llevo', 'Margaret Mitchell', 5, 1, 4200, 'viento.jpg'),
(21, 'Cada Suspiro', 'Nicholas Sparks', 5, 4, 4000, 'suspiro.jpg'),
(22, 'Solo un amor de verano', 'Alexandra Roma', 5, 4, 2200, 'verano.jpg'),
(23, 'Los dos amores de mi vida', 'Taylor Jenkins Reid', 5, 2, 3600, 'amores.jpg'),
(24, 'El campamento', 'Blue Jeans', 1, 1, 4200, 'campamento.jpg'),
(25, 'Rojo Blanco y Sangre Azul', 'Cassy McQuiston', 1, 4, 4000, 'rojo.jpg'),
(26, 'Alguien esta mintiendo', 'Karen M McManus', 1, 5, 2200, 'alguien.jpg'),
(27, 'Buscando a Alaska', 'John Green', 1, 2, 3600, 'alaska.jpg'),
(28, 'En el nombre del viento', 'Patrick Rothfuss', 4, 1, 4200, 'nomviento.jpg'),
(29, 'El señor de los anillos', 'J.R.R. Tolkien', 4, 4, 4000, 'anillos.jpg'),
(30, 'Seis de Cuervos', 'Leigh Bardugo', 4, 5, 2200, 'seis.jpg'),
(31, 'El Imperio Final', 'Brandon Sanderson', 4, 2, 3600, 'imperio.jpg'),
(32, 'Divergente', 'Veronica Roth', 9, 1, 4200, 'divergente.jpg'),
(33, 'Maze Runner', 'James Dashner', 9, 4, 4000, 'maze.jpg'),
(34, 'El año de Gracia', 'Kim Liggett', 9, 5, 2200, 'gracia.jpg'),
(35, '1984', 'George Orwell', 9, 2, 3600, '1984.jpg'),
(36, 'Dracula', 'Bram Stoker', 3, 1, 4200, 'dracula.jpg'),
(37, 'Coraline', 'Neil Gaiman', 3, 4, 4000, 'coraline.jpg'),
(38, 'El Exorcista', 'William Peter Blatty', 3, 5, 2200, 'excor.jpg'),
(39, 'El resplandor', 'Stephen King', 3, 2, 3600, 'resplandor.jpg'),
(40, 'Ricitos de Oro y los tres osos', 'Robert Southey', 6, 0, 4200, 'ricitos.jpg'),
(41, 'Caperucita Roja', 'Charles Perrault', 6, 4, 4000, 'caperucita.jpg'),
(42, 'El Patito Feo', 'Hans Christian Andersen', 6, 5, 2200, 'patito.jpg'),
(43, 'Hansel y Gretel', 'Hermanos Grimm', 6, 2, 3600, 'hansel.jpg'),
(44, 'La chica del tren', 'Paula Hawkins', 7, 2, 3160, 'tren.jpg'),
(45, 'La bestia', 'Carmen Mola', 7, 2, 2800, 'bestia.jpg'),
(46, 'El pecador de Oxford', 'Mar Petryk', 7, 2, 3500, 'pecador.jpg'),
(47, 'A fuego lento', 'Paula Hawkins', 7, 2, 3740, 'fuego.jpg'),
(48, 'Cien años de soledad', 'Gabriel Garcia Marquez', 10, 2, 3160, 'cien.jpg'),
(49, 'Matar a un Ruiseñor', 'Harper Lee', 10, 2, 2800, 'matar.jpg'),
(50, 'El tiempo entre costuras', 'Maria Dueñas', 10, 2, 3500, 'costuras.jpg'),
(51, 'Cumbres Borrascosas', 'Emily Brontë', 10, 2, 3740, 'cumbres.jpg');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `users`
--

CREATE TABLE `users` (
  `id` smallint(11) UNSIGNED NOT NULL,
  `Name` varchar(350) NOT NULL,
  `LastName` varchar(250) NOT NULL,
  `UserName` varchar(150) NOT NULL,
  `Password` varchar(150) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- RELACIONES PARA LA TABLA `users`:
--

--
-- Volcado de datos para la tabla `users`
--

INSERT INTO `users` VALUES
(1, 'Agustina', 'Araya', 'agusao', 'agus123'),
(2, 'Agostina', 'Morellato', 'agosm', 'agos123'),
(3, 'Agostina', 'Henderson', 'agosh', 'agosh123');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `addresses`
--
ALTER TABLE `addresses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `FK_id_user` (`id`);

--
-- Indices de la tabla `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`);

--
-- Indices de la tabla `details`
--
ALTER TABLE `details`
  ADD PRIMARY KEY (`Id`),
  ADD KEY `FK_id_product` (`product_id`),
  ADD KEY `FK_id_orden` (`order_id`);

--
-- Indices de la tabla `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `FK_id_user` (`user_id`);

--
-- Indices de la tabla `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `FK_id_category` (`category_id`);

--
-- Indices de la tabla `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `UserName` (`UserName`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `addresses`
--
ALTER TABLE `addresses`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT de la tabla `categories`
--
ALTER TABLE `categories`
  MODIFY `id` smallint(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT de la tabla `details`
--
ALTER TABLE `details`
  MODIFY `Id` smallint(5) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT de la tabla `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT de la tabla `products`
--
ALTER TABLE `products`
  MODIFY `id` smallint(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=52;

--
-- AUTO_INCREMENT de la tabla `users`
--
ALTER TABLE `users`
  MODIFY `id` smallint(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Restricciones para tablas volcadas
--

--
-- Filtros para la tabla `details`
--
ALTER TABLE `details`
  ADD CONSTRAINT `details_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;