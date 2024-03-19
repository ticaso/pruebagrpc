-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 17-03-2024 a las 02:57:47
-- Versión del servidor: 10.4.32-MariaDB
-- Versión de PHP: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `aspirante`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `aspirante`
--

CREATE TABLE `aspirante` (
  `id` int(11) NOT NULL,
  `nombre` varchar(255) NOT NULL,
  `apellido` varchar(255) NOT NULL,
  `imagen_principal` varchar(255) DEFAULT NULL,
  `imagen_secundaria` varchar(255) DEFAULT NULL,
  `carrera` varchar(255) DEFAULT NULL,
  `lenguajes_programacion` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `aspirante`
--

INSERT INTO `aspirante` (`id`, `nombre`, `apellido`, `imagen_principal`, `imagen_secundaria`, `carrera`, `lenguajes_programacion`) VALUES
(1, 'Juan', 'Pérez', 'https://img.freepik.com/foto-gratis/hombre-negocios-feliz-vestido-traje-gris-que-encuentran-aisladas-pared-blanca_231208-9211.jpg?size=626&ext=jpg&ga=GA1.1.1319243779.1710374400&semt=ais', 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Python-logo-notext.svg/1869px-Python-logo-notext.svg.png', 'Ingeniería en Sistemas Computacionales', 'C++, Java, Python'),
(2, 'Laura', 'González', 'https://img.freepik.com/foto-gratis/increible-mujer-negocios-alegre-pie-brazos-cruzados_171337-8487.jpg?size=626&ext=jpg&ga=GA1.1.1700460183.1709510400&semt=ais', 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6a/JavaScript-logo.png/768px-JavaScript-logo.png', 'Licenciatura en Informática', 'JavaScript, TypeScript, Rust');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `aspirante`
--
ALTER TABLE `aspirante`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `aspirante`
--
ALTER TABLE `aspirante`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
