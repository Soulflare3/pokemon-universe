-- MySQL Administrator dump 1.4
--
-- ------------------------------------------------------
-- Server version	5.1.49-1ubuntu8.1


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


--
-- Create schema pumaster
--

CREATE DATABASE IF NOT EXISTS pumaster;
USE pumaster;

--
-- Definition of table `pumaster`.`characters`
--

DROP TABLE IF EXISTS `pumaster`.`characters`;
CREATE TABLE  `pumaster`.`characters` (
  `idcharacter` int(11) NOT NULL AUTO_INCREMENT,
  `idaccount` int(11) NOT NULL,
  `idserver` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`idcharacter`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


--
-- Definition of table `pumaster`.`servers`
--

DROP TABLE IF EXISTS `pumaster`.`servers`;
CREATE TABLE  `pumaster`.`servers` (
  `idserver` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `address` varchar(20) NOT NULL,
  PRIMARY KEY (`idserver`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


--
-- Definition of table `pumaster`.`users`
--

DROP TABLE IF EXISTS `pumaster`.`users`;
CREATE TABLE  `pumaster`.`users` (
  `iduser` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(250) NOT NULL,
  `password_salt` varchar(250) NOT NULL,
  PRIMARY KEY (`iduser`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
