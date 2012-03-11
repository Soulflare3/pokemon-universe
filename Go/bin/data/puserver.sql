CREATE DATABASE  IF NOT EXISTS `puserver` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `puserver`;
-- MySQL dump 10.13  Distrib 5.1.40, for Win32 (ia32)
--
-- Host: DARK-MYSQL    Database: puserver
-- ------------------------------------------------------
-- Server version	5.1.58-1ubuntu1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `encounter_slot`
--

DROP TABLE IF EXISTS `encounter_slot`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `encounter_slot` (
  `idencounter_slot` int(11) NOT NULL,
  `idencounter` int(11) NOT NULL,
  `idpokemon` int(11) NOT NULL,
  `gender_rate` int(11) DEFAULT NULL,
  PRIMARY KEY (`idencounter_slot`),
  KEY `fk_encounter_slot_encounter1` (`idencounter`),
  KEY `fk_encounter_slot_pokemon1` (`idpokemon`),
  KEY `fk_encounter_idpokemon` (`idpokemon`),
  CONSTRAINT `fk_encounter_slot_encounter1` FOREIGN KEY (`idencounter`) REFERENCES `encounter` (`idencounter`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_encounter_idpokemon` FOREIGN KEY (`idpokemon`) REFERENCES `pokemon` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `encounter_slot`
--

LOCK TABLES `encounter_slot` WRITE;
/*!40000 ALTER TABLE `encounter_slot` DISABLE KEYS */;
/*!40000 ALTER TABLE `encounter_slot` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `abilities`
--

DROP TABLE IF EXISTS `abilities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `abilities` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `identifier` varchar(24) NOT NULL,
  `generation_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `generation_id` (`generation_id`)
) ENGINE=InnoDB AUTO_INCREMENT=165 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `abilities`
--

LOCK TABLES `abilities` WRITE;
/*!40000 ALTER TABLE `abilities` DISABLE KEYS */;
/*!40000 ALTER TABLE `abilities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_flavor_text`
--

DROP TABLE IF EXISTS `item_flavor_text`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_flavor_text` (
  `item_id` int(11) NOT NULL,
  `version_group_id` int(11) NOT NULL,
  `language_id` int(11) NOT NULL,
  `flavor_text` varchar(255) NOT NULL,
  PRIMARY KEY (`item_id`,`version_group_id`,`language_id`),
  KEY `version_group_id` (`version_group_id`),
  KEY `language_id` (`language_id`),
  CONSTRAINT `item_flavor_text_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `item_flavor_text_ibfk_2` FOREIGN KEY (`version_group_id`) REFERENCES `version_groups` (`id`),
  CONSTRAINT `item_flavor_text_ibfk_3` FOREIGN KEY (`language_id`) REFERENCES `languages` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_flavor_text`
--

LOCK TABLES `item_flavor_text` WRITE;
/*!40000 ALTER TABLE `item_flavor_text` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_flavor_text` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `npc_outfit`
--

DROP TABLE IF EXISTS `npc_outfit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `npc_outfit` (
  `idnpc` int(11) NOT NULL,
  `head` int(11) DEFAULT NULL,
  `nek` int(11) DEFAULT NULL,
  `upper` int(11) DEFAULT NULL,
  `lower` int(11) DEFAULT NULL,
  `feet` int(11) DEFAULT NULL,
  PRIMARY KEY (`idnpc`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `npc_outfit`
--

LOCK TABLES `npc_outfit` WRITE;
/*!40000 ALTER TABLE `npc_outfit` DISABLE KEYS */;
/*!40000 ALTER TABLE `npc_outfit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_game_indices`
--

DROP TABLE IF EXISTS `item_game_indices`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_game_indices` (
  `item_id` int(11) NOT NULL,
  `generation_id` int(11) NOT NULL,
  `game_index` int(11) NOT NULL,
  PRIMARY KEY (`item_id`,`generation_id`),
  KEY `generation_id` (`generation_id`),
  CONSTRAINT `item_game_indices_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `item_game_indices_ibfk_2` FOREIGN KEY (`generation_id`) REFERENCES `generations` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_game_indices`
--

LOCK TABLES `item_game_indices` WRITE;
/*!40000 ALTER TABLE `item_game_indices` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_game_indices` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_flag_map`
--

DROP TABLE IF EXISTS `item_flag_map`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_flag_map` (
  `item_id` int(11) NOT NULL,
  `item_flag_id` int(11) NOT NULL,
  PRIMARY KEY (`item_id`,`item_flag_id`),
  KEY `item_flag_id` (`item_flag_id`),
  CONSTRAINT `item_flag_map_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `item_flag_map_ibfk_2` FOREIGN KEY (`item_flag_id`) REFERENCES `item_flags` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_flag_map`
--

LOCK TABLES `item_flag_map` WRITE;
/*!40000 ALTER TABLE `item_flag_map` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_flag_map` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player_quests`
--

DROP TABLE IF EXISTS `player_quests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_quests` (
  `idplayer_quests` int(11) NOT NULL,
  `idplayer` int(11) DEFAULT NULL,
  `idquest` int(11) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  PRIMARY KEY (`idplayer_quests`),
  KEY `fk_player_quests_player1` (`idplayer`),
  KEY `fk_player_quests_quests1` (`idquest`),
  CONSTRAINT `fk_player_quests_player1` FOREIGN KEY (`idplayer`) REFERENCES `player` (`idplayer`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_quests_quests1` FOREIGN KEY (`idquest`) REFERENCES `quests` (`idquests`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player_quests`
--

LOCK TABLES `player_quests` WRITE;
/*!40000 ALTER TABLE `player_quests` DISABLE KEYS */;
/*!40000 ALTER TABLE `player_quests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pokemon`
--

DROP TABLE IF EXISTS `pokemon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pokemon` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `species_id` int(11) DEFAULT NULL,
  `height` int(11) NOT NULL,
  `weight` int(11) NOT NULL,
  `base_experience` int(11) NOT NULL,
  `order` int(11) NOT NULL,
  `is_default` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `species_id` (`species_id`),
  KEY `ix_pokemon_is_default` (`is_default`),
  KEY `ix_pokemon_order` (`order`),
  CONSTRAINT `pokemon_ibfk_1` FOREIGN KEY (`species_id`) REFERENCES `pokemon_species` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=668 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pokemon`
--

LOCK TABLES `pokemon` WRITE;
/*!40000 ALTER TABLE `pokemon` DISABLE KEYS */;
/*!40000 ALTER TABLE `pokemon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pokemon_abilities`
--

DROP TABLE IF EXISTS `pokemon_abilities`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pokemon_abilities` (
  `pokemon_id` int(11) NOT NULL,
  `ability_id` int(11) NOT NULL,
  `is_dream` tinyint(1) NOT NULL,
  `slot` int(11) NOT NULL,
  PRIMARY KEY (`pokemon_id`,`slot`),
  KEY `ability_id` (`ability_id`),
  KEY `ix_pokemon_abilities_is_dream` (`is_dream`),
  CONSTRAINT `pokemon_abilities_ibfk_1` FOREIGN KEY (`pokemon_id`) REFERENCES `pokemon` (`id`),
  CONSTRAINT `pokemon_abilities_ibfk_2` FOREIGN KEY (`ability_id`) REFERENCES `abilities` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pokemon_abilities`
--

LOCK TABLES `pokemon_abilities` WRITE;
/*!40000 ALTER TABLE `pokemon_abilities` DISABLE KEYS */;
/*!40000 ALTER TABLE `pokemon_abilities` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group` (
  `idgroup` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `flags` int(11) DEFAULT NULL,
  `priority` int(11) DEFAULT NULL,
  PRIMARY KEY (`idgroup`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group`
--

LOCK TABLES `group` WRITE;
/*!40000 ALTER TABLE `group` DISABLE KEYS */;

/*!40000 ALTER TABLE `group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player_group`
--

DROP TABLE IF EXISTS `player_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_group` (
  `player_idplayer` int(11) NOT NULL,
  `group_idgroup` int(11) NOT NULL,
  PRIMARY KEY (`player_idplayer`,`group_idgroup`),
  KEY `fk_player_has_group_group1` (`group_idgroup`),
  CONSTRAINT `fk_player_has_group_group1` FOREIGN KEY (`group_idgroup`) REFERENCES `group` (`idgroup`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_has_group_player1` FOREIGN KEY (`player_idplayer`) REFERENCES `player` (`idplayer`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player_group`
--

LOCK TABLES `player_group` WRITE;
/*!40000 ALTER TABLE `player_group` DISABLE KEYS */;

/*!40000 ALTER TABLE `player_group` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `location_encounter`
--

DROP TABLE IF EXISTS `location_encounter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `location_encounter` (
  `idencounter` int(11) NOT NULL,
  `idlocation_section` int(11) NOT NULL,
  PRIMARY KEY (`idencounter`,`idlocation_section`),
  KEY `fk_location_encounter` (`idlocation_section`),
  KEY `fk_location_encounter_section` (`idlocation_section`),
  CONSTRAINT `fk_location_encounter_encounter` FOREIGN KEY (`idencounter`) REFERENCES `encounter` (`idencounter`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_location_encounter_section` FOREIGN KEY (`idlocation_section`) REFERENCES `location_section` (`idlocation_section`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `location_encounter`
--

LOCK TABLES `location_encounter` WRITE;
/*!40000 ALTER TABLE `location_encounter` DISABLE KEYS */;
/*!40000 ALTER TABLE `location_encounter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player_pokemon_move`
--

DROP TABLE IF EXISTS `player_pokemon_move`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_pokemon_move` (
  `idplayer_pokemon_move` int(11) NOT NULL,
  `idplayer_pokemon` int(11) NOT NULL,
  `idmove` int(11) NOT NULL,
  `pp_used` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`idplayer_pokemon_move`),
  KEY `fk_player_pokemon_move_player_pokemon1` (`idplayer_pokemon`),
  KEY `fk_player_pokemon_move_move1` (`idmove`),
  KEY `fk_player_pokemon_move_moves1` (`idmove`),
  CONSTRAINT `fk_player_pokemon_move_player_pokemon1` FOREIGN KEY (`idplayer_pokemon`) REFERENCES `player_pokemon` (`idplayer_pokemon`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_pokemon_move_moves1` FOREIGN KEY (`idmove`) REFERENCES `moves` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player_pokemon_move`
--

LOCK TABLES `player_pokemon_move` WRITE;
/*!40000 ALTER TABLE `player_pokemon_move` DISABLE KEYS */;
/*!40000 ALTER TABLE `player_pokemon_move` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `music`
--

DROP TABLE IF EXISTS `music`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `music` (
  `idmusic` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(45) DEFAULT NULL,
  `filename` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`idmusic`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `music`
--

LOCK TABLES `music` WRITE;
/*!40000 ALTER TABLE `music` DISABLE KEYS */;

/*!40000 ALTER TABLE `music` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `types`
--

DROP TABLE IF EXISTS `types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `types` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `identifier` varchar(12) NOT NULL,
  `generation_id` int(11) NOT NULL,
  `damage_class_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10003 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `types`
--

LOCK TABLES `types` WRITE;
/*!40000 ALTER TABLE `types` DISABLE KEYS */;
/*!40000 ALTER TABLE `types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mapchange_account`
--

DROP TABLE IF EXISTS `mapchange_account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mapchange_account` (
  `idmapchange_account` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(45) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`idmapchange_account`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mapchange_account`
--

LOCK TABLES `mapchange_account` WRITE;
/*!40000 ALTER TABLE `mapchange_account` DISABLE KEYS */;
/*!40000 ALTER TABLE `mapchange_account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tile_layer`
--

DROP TABLE IF EXISTS `tile_layer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tile_layer` (
  `idtile_layer` int(11) NOT NULL AUTO_INCREMENT,
  `idtile` int(11) NOT NULL,
  `sprite` int(11) DEFAULT NULL,
  `layer` int(11) DEFAULT NULL,
  PRIMARY KEY (`idtile_layer`,`idtile`),
  KEY `fk_tile_layer_tileid` (`idtile`),
  CONSTRAINT `fk_tile_layer_tileid` FOREIGN KEY (`idtile`) REFERENCES `tile` (`idtile`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=1549597 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tile_layer`
--

LOCK TABLES `tile_layer` WRITE;
/*!40000 ALTER TABLE `tile_layer` DISABLE KEYS */;
/*!40000 ALTER TABLE `tile_layer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_flag_prose`
--

DROP TABLE IF EXISTS `item_flag_prose`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_flag_prose` (
  `item_flag_id` int(11) NOT NULL,
  `local_language_id` int(11) NOT NULL,
  `name` varchar(24) DEFAULT NULL,
  `description` varchar(64) DEFAULT NULL,
  PRIMARY KEY (`item_flag_id`,`local_language_id`),
  KEY `local_language_id` (`local_language_id`),
  KEY `ix_item_flag_prose_name` (`name`),
  CONSTRAINT `item_flag_prose_ibfk_1` FOREIGN KEY (`item_flag_id`) REFERENCES `item_flags` (`id`),
  CONSTRAINT `item_flag_prose_ibfk_2` FOREIGN KEY (`local_language_id`) REFERENCES `languages` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_flag_prose`
--

LOCK TABLES `item_flag_prose` WRITE;
/*!40000 ALTER TABLE `item_flag_prose` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_flag_prose` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `location`
--

DROP TABLE IF EXISTS `location`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `location` (
  `idlocation` int(11) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `idpokecenter` int(11) DEFAULT NULL,
  `idmusic` int(11) NOT NULL,
  PRIMARY KEY (`idlocation`),
  KEY `fk_location_pokecenter1` (`idpokecenter`),
  KEY `fk_location_music1` (`idmusic`),
  CONSTRAINT `fk_location_music1` FOREIGN KEY (`idmusic`) REFERENCES `music` (`idmusic`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_location_pokecenter1` FOREIGN KEY (`idpokecenter`) REFERENCES `pokecenter` (`idpokecenter`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `location`
--

LOCK TABLES `location` WRITE;
/*!40000 ALTER TABLE `location` DISABLE KEYS */;

/*!40000 ALTER TABLE `location` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `location_section`
--

DROP TABLE IF EXISTS `location_section`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `location_section` (
  `idlocation_section` int(11) NOT NULL AUTO_INCREMENT,
  `idlocation` int(11) NOT NULL,
  `name` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`idlocation_section`),
  KEY `fk_location_sections_location1` (`idlocation`),
  CONSTRAINT `fk_location_sections_location1` FOREIGN KEY (`idlocation`) REFERENCES `location` (`idlocation`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `location_section`
--

LOCK TABLES `location_section` WRITE;
/*!40000 ALTER TABLE `location_section` DISABLE KEYS */;
/*!40000 ALTER TABLE `location_section` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_pocket_names`
--

DROP TABLE IF EXISTS `item_pocket_names`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_pocket_names` (
  `item_pocket_id` int(11) NOT NULL,
  `local_language_id` int(11) NOT NULL,
  `name` varchar(16) NOT NULL,
  PRIMARY KEY (`item_pocket_id`,`local_language_id`),
  KEY `local_language_id` (`local_language_id`),
  KEY `ix_item_pocket_names_name` (`name`),
  CONSTRAINT `item_pocket_names_ibfk_1` FOREIGN KEY (`item_pocket_id`) REFERENCES `item_pockets` (`id`),
  CONSTRAINT `item_pocket_names_ibfk_2` FOREIGN KEY (`local_language_id`) REFERENCES `languages` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_pocket_names`
--

LOCK TABLES `item_pocket_names` WRITE;
/*!40000 ALTER TABLE `item_pocket_names` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_pocket_names` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_prose`
--

DROP TABLE IF EXISTS `item_prose`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_prose` (
  `item_id` int(11) NOT NULL,
  `local_language_id` int(11) NOT NULL,
  `short_effect` varchar(256) DEFAULT NULL,
  `effect` varchar(5120) DEFAULT NULL,
  PRIMARY KEY (`item_id`,`local_language_id`),
  KEY `local_language_id` (`local_language_id`),
  CONSTRAINT `item_prose_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `item_prose_ibfk_2` FOREIGN KEY (`local_language_id`) REFERENCES `languages` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_prose`
--

LOCK TABLES `item_prose` WRITE;
/*!40000 ALTER TABLE `item_prose` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_prose` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `growth_rates`
--

DROP TABLE IF EXISTS `growth_rates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `growth_rates` (
  `id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `growth_rates`
--

LOCK TABLES `growth_rates` WRITE;
/*!40000 ALTER TABLE `growth_rates` DISABLE KEYS */;
/*!40000 ALTER TABLE `growth_rates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `encounter`
--

DROP TABLE IF EXISTS `encounter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `encounter` (
  `idencounter` int(11) NOT NULL AUTO_INCREMENT,
  `idencounter_condition` int(11) NOT NULL,
  `rate` int(11) DEFAULT NULL,
  PRIMARY KEY (`idencounter`),
  KEY `fk_encounter_encounter_condition1` (`idencounter_condition`),
  CONSTRAINT `fk_encounter_encounter_condition1` FOREIGN KEY (`idencounter_condition`) REFERENCES `encounter_condition` (`idencounter_condition`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `encounter`
--

LOCK TABLES `encounter` WRITE;
/*!40000 ALTER TABLE `encounter` DISABLE KEYS */;
/*!40000 ALTER TABLE `encounter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `contest_effects`
--

DROP TABLE IF EXISTS `contest_effects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `contest_effects` (
  `id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contest_effects`
--

LOCK TABLES `contest_effects` WRITE;
/*!40000 ALTER TABLE `contest_effects` DISABLE KEYS */;
/*!40000 ALTER TABLE `contest_effects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player_items`
--

DROP TABLE IF EXISTS `player_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_items` (
  `idplayer_items` int(11) NOT NULL,
  `idplayer` int(11) DEFAULT NULL,
  `iditem` int(11) DEFAULT NULL,
  `count` int(11) DEFAULT NULL,
  `slot` int(11) DEFAULT NULL,
  PRIMARY KEY (`idplayer_items`),
  KEY `fk_player_items_player1` (`idplayer`),
  KEY `fk_player_items_items1` (`iditem`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player_items`
--

LOCK TABLES `player_items` WRITE;
/*!40000 ALTER TABLE `player_items` DISABLE KEYS */;
/*!40000 ALTER TABLE `player_items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pokecenter`
--

DROP TABLE IF EXISTS `pokecenter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pokecenter` (
  `idpokecenter` int(11) NOT NULL,
  `position` bigint(20) NOT NULL,
  `description` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`idpokecenter`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pokecenter`
--

LOCK TABLES `pokecenter` WRITE;
/*!40000 ALTER TABLE `pokecenter` DISABLE KEYS */;

/*!40000 ALTER TABLE `pokecenter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mapchange_layer`
--

DROP TABLE IF EXISTS `mapchange_layer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mapchange_layer` (
  `idmapchange_layer` int(11) NOT NULL AUTO_INCREMENT,
  `idmapchange_tile` int(11) NOT NULL,
  `index` int(11) DEFAULT NULL,
  `sprite` int(11) DEFAULT NULL,
  PRIMARY KEY (`idmapchange_layer`),
  KEY `fk_mapchange_layer_mapchange_tile1` (`idmapchange_tile`),
  CONSTRAINT `fk_mapchange_layer_mapchange_tile1` FOREIGN KEY (`idmapchange_tile`) REFERENCES `mapchange_tile` (`idmapchange_tile`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=18258 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mapchange_layer`
--

LOCK TABLES `mapchange_layer` WRITE;
/*!40000 ALTER TABLE `mapchange_layer` DISABLE KEYS */;
/*!40000 ALTER TABLE `mapchange_layer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player_pokemon`
--

DROP TABLE IF EXISTS `player_pokemon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_pokemon` (
  `idplayer_pokemon` int(11) NOT NULL AUTO_INCREMENT,
  `idpokemon` int(11) NOT NULL,
  `idplayer` int(11) NOT NULL,
  `nickname` varchar(45) DEFAULT NULL,
  `bound` tinyint(4) DEFAULT '0' COMMENT '1 if pokemon is bound to player',
  `experience` int(10) unsigned DEFAULT NULL,
  `iv_hp` tinyint(3) unsigned DEFAULT NULL,
  `iv_attack` tinyint(3) unsigned DEFAULT NULL,
  `iv_attack_spec` tinyint(3) unsigned DEFAULT NULL,
  `iv_defence` tinyint(3) unsigned DEFAULT NULL,
  `iv_defence_spec` tinyint(3) unsigned DEFAULT NULL,
  `iv_speed` tinyint(3) unsigned DEFAULT NULL,
  `happiness` tinyint(3) unsigned DEFAULT NULL,
  `gender` tinyint(4) DEFAULT NULL COMMENT '-1 None\n0 Male\n1 Female',
  `in_party` tinyint(1) DEFAULT NULL,
  `party_slot` tinyint(1) DEFAULT NULL,
  `held_item` int(11) DEFAULT NULL,
  `shiny` tinyint(1) DEFAULT '0',
  `idability` int(11) DEFAULT NULL,
  `damaged_hp` int(11) DEFAULT NULL,
  PRIMARY KEY (`idplayer_pokemon`),
  KEY `fk_player_pokemon_player` (`idplayer`),
  KEY `fk_player_pokemon_pokemon` (`idpokemon`),
  KEY `fk_player_pokemon_items1` (`held_item`),
  CONSTRAINT `fk_player_pokemon_items1` FOREIGN KEY (`held_item`) REFERENCES `items` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_pokemon_player` FOREIGN KEY (`idplayer`) REFERENCES `player` (`idplayer`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_pokemon_pokemon1` FOREIGN KEY (`idpokemon`) REFERENCES `pokemon` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player_pokemon`
--

LOCK TABLES `player_pokemon` WRITE;
/*!40000 ALTER TABLE `player_pokemon` DISABLE KEYS */;
/*!40000 ALTER TABLE `player_pokemon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pokemon_moves`
--

DROP TABLE IF EXISTS `pokemon_moves`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pokemon_moves` (
  `pokemon_id` int(11) NOT NULL,
  `version_group_id` int(11) NOT NULL,
  `move_id` int(11) NOT NULL,
  `pokemon_move_method_id` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  `order` int(11) DEFAULT NULL,
  PRIMARY KEY (`pokemon_id`,`version_group_id`,`move_id`,`pokemon_move_method_id`,`level`),
  KEY `idx_autoinc_level` (`level`),
  KEY `ix_pokemon_moves_version_group_id` (`version_group_id`),
  KEY `ix_pokemon_moves_level` (`level`),
  KEY `ix_pokemon_moves_pokemon_id` (`pokemon_id`),
  KEY `ix_pokemon_moves_pokemon_move_method_id` (`pokemon_move_method_id`),
  KEY `ix_pokemon_moves_move_id` (`move_id`),
  CONSTRAINT `pokemon_moves_ibfk_1` FOREIGN KEY (`pokemon_id`) REFERENCES `pokemon` (`id`),
  CONSTRAINT `pokemon_moves_ibfk_3` FOREIGN KEY (`move_id`) REFERENCES `moves` (`id`),
  CONSTRAINT `pokemon_moves_ibfk_4` FOREIGN KEY (`pokemon_move_method_id`) REFERENCES `pokemon_move_methods` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pokemon_moves`
--

LOCK TABLES `pokemon_moves` WRITE;
/*!40000 ALTER TABLE `pokemon_moves` DISABLE KEYS */;
/*!40000 ALTER TABLE `pokemon_moves` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mapchange`
--

DROP TABLE IF EXISTS `mapchange`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mapchange` (
  `idmapchange` int(11) NOT NULL AUTO_INCREMENT,
  `start_x` int(11) DEFAULT NULL,
  `start_y` int(11) DEFAULT NULL,
  `width` int(11) DEFAULT NULL,
  `height` int(11) DEFAULT NULL,
  `username` varchar(45) DEFAULT NULL,
  `description` longtext,
  `submit_date` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `status` int(11) DEFAULT '0',
  PRIMARY KEY (`idmapchange`)
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mapchange`
--

LOCK TABLES `mapchange` WRITE;
/*!40000 ALTER TABLE `mapchange` DISABLE KEYS */;
/*!40000 ALTER TABLE `mapchange` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pokemon_forms`
--

DROP TABLE IF EXISTS `pokemon_forms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pokemon_forms` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `form_identifier` varchar(16) DEFAULT NULL,
  `pokemon_id` int(11) NOT NULL,
  `is_default` tinyint(1) NOT NULL,
  `is_battle_only` tinyint(1) NOT NULL,
  `order` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `pokemon_id` (`pokemon_id`),
  CONSTRAINT `pokemon_forms_ibfk_1` FOREIGN KEY (`pokemon_id`) REFERENCES `pokemon` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=728 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pokemon_forms`
--

LOCK TABLES `pokemon_forms` WRITE;
/*!40000 ALTER TABLE `pokemon_forms` DISABLE KEYS */;
/*!40000 ALTER TABLE `pokemon_forms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_category_prose`
--

DROP TABLE IF EXISTS `item_category_prose`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_category_prose` (
  `item_category_id` int(11) NOT NULL,
  `local_language_id` int(11) NOT NULL,
  `name` varchar(16) NOT NULL,
  PRIMARY KEY (`item_category_id`,`local_language_id`),
  KEY `local_language_id` (`local_language_id`),
  KEY `ix_item_category_prose_name` (`name`),
  CONSTRAINT `item_category_prose_ibfk_1` FOREIGN KEY (`item_category_id`) REFERENCES `item_categories` (`id`),
  CONSTRAINT `item_category_prose_ibfk_2` FOREIGN KEY (`local_language_id`) REFERENCES `languages` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_category_prose`
--

LOCK TABLES `item_category_prose` WRITE;
/*!40000 ALTER TABLE `item_category_prose` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_category_prose` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player`
--

DROP TABLE IF EXISTS `player`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player` (
  `idplayer` int(11) NOT NULL AUTO_INCREMENT,
  `idaccount` int(11) DEFAULT NULL,
  `name` varchar(20) DEFAULT NULL,
  `password` varchar(45) DEFAULT NULL,
  `password_salt` varchar(45) DEFAULT NULL,
  `position` bigint(20) DEFAULT NULL COMMENT 'x;y;z',
  `movement` smallint(6) DEFAULT NULL,
  `idpokecenter` int(11) DEFAULT NULL,
  `money` int(11) DEFAULT NULL,
  `idlocation` int(11) NOT NULL,
  PRIMARY KEY (`idplayer`),
  KEY `fk_player_location1` (`idlocation`),
  CONSTRAINT `fk_player_location1` FOREIGN KEY (`idlocation`) REFERENCES `location` (`idlocation`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player`
--

LOCK TABLES `player` WRITE;
/*!40000 ALTER TABLE `player` DISABLE KEYS */;

/*!40000 ALTER TABLE `player` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pokemon_species`
--

DROP TABLE IF EXISTS `pokemon_species`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pokemon_species` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `identifier` varchar(20) NOT NULL,
  `generation_id` int(11) DEFAULT NULL,
  `evolves_from_species_id` int(11) DEFAULT NULL,
  `evolution_chain_id` int(11) DEFAULT NULL,
  `gender_rate` int(11) NOT NULL,
  `capture_rate` int(11) NOT NULL,
  `base_happiness` int(11) NOT NULL,
  `is_baby` tinyint(1) NOT NULL,
  `hatch_counter` int(11) NOT NULL,
  `has_gender_differences` tinyint(1) NOT NULL,
  `growth_rate_id` int(11) NOT NULL,
  `forms_switchable` tinyint(1) NOT NULL,
  `color_id` int(11) DEFAULT NULL,
  `shape_id` int(11) DEFAULT NULL,
  `habitat_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `evolves_from_species_id` (`evolves_from_species_id`),
  KEY `evolution_chain_id` (`evolution_chain_id`),
  KEY `growth_rate_id` (`growth_rate_id`),
  KEY `fk_pokemon_species_pokemon_colors1` (`color_id`),
  KEY `fk_pokemon_species_pokemon_shapes1` (`shape_id`),
  KEY `fk_pokemon_species_pokemon_habitats1` (`habitat_id`),
  CONSTRAINT `pokemon_species_ibfk_2` FOREIGN KEY (`evolves_from_species_id`) REFERENCES `pokemon_species` (`id`),
  CONSTRAINT `pokemon_species_ibfk_3` FOREIGN KEY (`evolution_chain_id`) REFERENCES `evolution_chains` (`id`),
  CONSTRAINT `fk_pokemon_species_pokemon_colors1` FOREIGN KEY (`color_id`) REFERENCES `pokemon_colors` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_pokemon_species_pokemon_shapes1` FOREIGN KEY (`shape_id`) REFERENCES `pokemon_shapes` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_pokemon_species_pokemon_habitats1` FOREIGN KEY (`habitat_id`) REFERENCES `pokemon_habitats` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=650 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pokemon_species`
--

LOCK TABLES `pokemon_species` WRITE;
/*!40000 ALTER TABLE `pokemon_species` DISABLE KEYS */;
/*!40000 ALTER TABLE `pokemon_species` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_fling_effect_prose`
--

DROP TABLE IF EXISTS `item_fling_effect_prose`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_fling_effect_prose` (
  `item_fling_effect_id` int(11) NOT NULL,
  `local_language_id` int(11) NOT NULL,
  `effect` varchar(255) NOT NULL,
  PRIMARY KEY (`item_fling_effect_id`,`local_language_id`),
  KEY `local_language_id` (`local_language_id`),
  CONSTRAINT `item_fling_effect_prose_ibfk_1` FOREIGN KEY (`item_fling_effect_id`) REFERENCES `item_fling_effects` (`id`),
  CONSTRAINT `item_fling_effect_prose_ibfk_2` FOREIGN KEY (`local_language_id`) REFERENCES `languages` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_fling_effect_prose`
--

LOCK TABLES `item_fling_effect_prose` WRITE;
/*!40000 ALTER TABLE `item_fling_effect_prose` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_fling_effect_prose` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_categories`
--

DROP TABLE IF EXISTS `item_categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_categories` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pocket_id` int(11) NOT NULL,
  `identifier` varchar(16) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `pocket_id` (`pocket_id`),
  CONSTRAINT `item_categories_ibfk_1` FOREIGN KEY (`pocket_id`) REFERENCES `item_pockets` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_categories`
--

LOCK TABLES `item_categories` WRITE;
/*!40000 ALTER TABLE `item_categories` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ability_messages`
--

DROP TABLE IF EXISTS `ability_messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ability_messages` (
  `idability_messages` int(11) NOT NULL AUTO_INCREMENT,
  `ability_id` int(11) NOT NULL,
  `message` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`idability_messages`),
  KEY `fk_ability_messages_abilities1` (`ability_id`),
  CONSTRAINT `fk_ability_messages_abilities1` FOREIGN KEY (`ability_id`) REFERENCES `abilities` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ability_messages`
--

LOCK TABLES `ability_messages` WRITE;
/*!40000 ALTER TABLE `ability_messages` DISABLE KEYS */;
/*!40000 ALTER TABLE `ability_messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `moves`
--

DROP TABLE IF EXISTS `moves`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `moves` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `identifier` varchar(24) NOT NULL,
  `generation_id` int(11) NOT NULL,
  `type_id` int(11) NOT NULL,
  `power` smallint(6) NOT NULL,
  `pp` smallint(6) DEFAULT NULL,
  `accuracy` smallint(6) DEFAULT NULL,
  `priority` smallint(6) NOT NULL,
  `target_id` int(11) NOT NULL,
  `damage_class_id` int(11) NOT NULL,
  `effect_id` int(11) NOT NULL,
  `effect_chance` int(11) DEFAULT NULL,
  `contest_type_id` int(11) DEFAULT NULL,
  `contest_effect_id` int(11) DEFAULT NULL,
  `super_contest_effect_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `type_id` (`type_id`),
  KEY `target_id` (`target_id`),
  KEY `damage_class_id` (`damage_class_id`),
  KEY `effect_id` (`effect_id`),
  KEY `contest_type_id` (`contest_type_id`),
  KEY `contest_effect_id` (`contest_effect_id`),
  KEY `super_contest_effect_id` (`super_contest_effect_id`),
  CONSTRAINT `moves_ibfk_2` FOREIGN KEY (`type_id`) REFERENCES `types` (`id`),
  CONSTRAINT `fk_moves_move_messages1` FOREIGN KEY (`effect_id`) REFERENCES `move_messages` (`move_effect_id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=10019 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `moves`
--

LOCK TABLES `moves` WRITE;
/*!40000 ALTER TABLE `moves` DISABLE KEYS */;
/*!40000 ALTER TABLE `moves` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `stats`
--

DROP TABLE IF EXISTS `stats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `stats` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `damage_class_id` int(11) DEFAULT NULL,
  `identifier` varchar(16) NOT NULL,
  `is_battle_only` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stats`
--

LOCK TABLES `stats` WRITE;
/*!40000 ALTER TABLE `stats` DISABLE KEYS */;
/*!40000 ALTER TABLE `stats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player_backpack`
--

DROP TABLE IF EXISTS `player_backpack`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_backpack` (
  `idplayer_backpack` int(11) NOT NULL,
  `idplayer` int(11) NOT NULL,
  `iditem` int(11) NOT NULL,
  `count` int(11) DEFAULT '1',
  `slot` int(11) NOT NULL,
  PRIMARY KEY (`idplayer_backpack`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player_backpack`
--

LOCK TABLES `player_backpack` WRITE;
/*!40000 ALTER TABLE `player_backpack` DISABLE KEYS */;
/*!40000 ALTER TABLE `player_backpack` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `npc_pokemon_move`
--

DROP TABLE IF EXISTS `npc_pokemon_move`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `npc_pokemon_move` (
  `idnpc_pokemon_move` int(11) NOT NULL,
  `idnpc_pokemon` int(11) DEFAULT NULL,
  `idmove` int(11) DEFAULT NULL,
  PRIMARY KEY (`idnpc_pokemon_move`),
  KEY `fk_npc_pokemon_move_npc_pokemon1` (`idnpc_pokemon`),
  KEY `fk_npc_pokemon_move_moves1` (`idmove`),
  CONSTRAINT `fk_npc_pokemon_move_npc_pokemon1` FOREIGN KEY (`idnpc_pokemon`) REFERENCES `npc_pokemon` (`idnpc_pokemon`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_npc_pokemon_move_moves1` FOREIGN KEY (`idmove`) REFERENCES `moves` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `npc_pokemon_move`
--

LOCK TABLES `npc_pokemon_move` WRITE;
/*!40000 ALTER TABLE `npc_pokemon_move` DISABLE KEYS */;
/*!40000 ALTER TABLE `npc_pokemon_move` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `super_contest_effects`
--

DROP TABLE IF EXISTS `super_contest_effects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `super_contest_effects` (
  `id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `super_contest_effects`
--

LOCK TABLES `super_contest_effects` WRITE;
/*!40000 ALTER TABLE `super_contest_effects` DISABLE KEYS */;
/*!40000 ALTER TABLE `super_contest_effects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tile`
--

DROP TABLE IF EXISTS `tile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tile` (
  `idtile` int(11) NOT NULL AUTO_INCREMENT,
  `x` int(11) NOT NULL,
  `y` int(11) NOT NULL,
  `z` int(11) NOT NULL,
  `idlocation` int(11) NOT NULL,
  `movement` int(11) DEFAULT NULL,
  `script` varchar(128) DEFAULT NULL,
  `idteleport` int(11) DEFAULT NULL,
  PRIMARY KEY (`idtile`),
  KEY `fk_tile_location` (`idlocation`),
  KEY `position_key` (`x`,`y`,`z`),
  KEY `fk_tile_map` (`z`),
  KEY `fk_tile_teleport` (`idteleport`),
  CONSTRAINT `fk_tile_teleport1` FOREIGN KEY (`idteleport`) REFERENCES `teleport` (`idteleport`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_tile_location1` FOREIGN KEY (`idlocation`) REFERENCES `location` (`idlocation`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_tile_map1` FOREIGN KEY (`z`) REFERENCES `map` (`idmap`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=1132713 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tile`
--

LOCK TABLES `tile` WRITE;
/*!40000 ALTER TABLE `tile` DISABLE KEYS */;

/*!40000 ALTER TABLE `tile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_flags`
--

DROP TABLE IF EXISTS `item_flags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_flags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `identifier` varchar(24) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_flags`
--

LOCK TABLES `item_flags` WRITE;
/*!40000 ALTER TABLE `item_flags` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_flags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_pockets`
--

DROP TABLE IF EXISTS `item_pockets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_pockets` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `identifier` varchar(16) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_pockets`
--

LOCK TABLES `item_pockets` WRITE;
/*!40000 ALTER TABLE `item_pockets` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_pockets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `map`
--

DROP TABLE IF EXISTS `map`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `map` (
  `idmap` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`idmap`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `map`
--

LOCK TABLES `map` WRITE;
/*!40000 ALTER TABLE `map` DISABLE KEYS */;

/*!40000 ALTER TABLE `map` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `contest_types`
--

DROP TABLE IF EXISTS `contest_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `contest_types` (
  `id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `contest_types`
--

LOCK TABLES `contest_types` WRITE;
/*!40000 ALTER TABLE `contest_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `contest_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `teleport`
--

DROP TABLE IF EXISTS `teleport`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `teleport` (
  `idteleport` int(11) NOT NULL AUTO_INCREMENT,
  `x` int(11) DEFAULT NULL,
  `y` int(11) DEFAULT NULL,
  `z` int(11) DEFAULT NULL,
  PRIMARY KEY (`idteleport`)
) ENGINE=InnoDB AUTO_INCREMENT=224 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `teleport`
--

LOCK TABLES `teleport` WRITE;
/*!40000 ALTER TABLE `teleport` DISABLE KEYS */;

/*!40000 ALTER TABLE `teleport` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `encounter_condition`
--

DROP TABLE IF EXISTS `encounter_condition`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `encounter_condition` (
  `idencounter_condition` int(11) NOT NULL,
  `name` varchar(250) DEFAULT NULL,
  `default` int(11) DEFAULT NULL,
  PRIMARY KEY (`idencounter_condition`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `encounter_condition`
--

LOCK TABLES `encounter_condition` WRITE;
/*!40000 ALTER TABLE `encounter_condition` DISABLE KEYS */;
/*!40000 ALTER TABLE `encounter_condition` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `move_messages`
--

DROP TABLE IF EXISTS `move_messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `move_messages` (
  `move_effect_id` int(11) DEFAULT NULL,
  `message` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `move_messages`
--

LOCK TABLES `move_messages` WRITE;
/*!40000 ALTER TABLE `move_messages` DISABLE KEYS */;
/*!40000 ALTER TABLE `move_messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `npc`
--

DROP TABLE IF EXISTS `npc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `npc` (
  `idnpc` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `script_name` varchar(45) DEFAULT NULL,
  `position` bigint(20) DEFAULT NULL,
  `idmap` int(11) DEFAULT NULL,
  PRIMARY KEY (`idnpc`),
  CONSTRAINT `fk_npc_npc_outfit1` FOREIGN KEY (`idnpc`) REFERENCES `npc_outfit` (`idnpc`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `npc`
--

LOCK TABLES `npc` WRITE;
/*!40000 ALTER TABLE `npc` DISABLE KEYS */;
/*!40000 ALTER TABLE `npc` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mapchange_tile`
--

DROP TABLE IF EXISTS `mapchange_tile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mapchange_tile` (
  `idmapchange_tile` int(11) NOT NULL AUTO_INCREMENT,
  `idmapchange` int(11) NOT NULL,
  `x` int(11) DEFAULT NULL,
  `y` int(11) DEFAULT NULL,
  `z` int(11) DEFAULT NULL,
  `movement` int(11) DEFAULT NULL,
  PRIMARY KEY (`idmapchange_tile`),
  KEY `fk_mapchange_tile_mapchange1` (`idmapchange`),
  CONSTRAINT `fk_mapchange_tile_mapchange1` FOREIGN KEY (`idmapchange`) REFERENCES `mapchange` (`idmapchange`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=10814 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mapchange_tile`
--

LOCK TABLES `mapchange_tile` WRITE;
/*!40000 ALTER TABLE `mapchange_tile` DISABLE KEYS */;
/*!40000 ALTER TABLE `mapchange_tile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `move_flavor_text`
--

DROP TABLE IF EXISTS `move_flavor_text`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `move_flavor_text` (
  `id_move` int(11) NOT NULL,
  `version_group_id` int(11) NOT NULL,
  `language_id` int(11) NOT NULL,
  `flavor_text` varchar(255) NOT NULL,
  PRIMARY KEY (`id_move`,`version_group_id`,`language_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `move_flavor_text`
--

LOCK TABLES `move_flavor_text` WRITE;
/*!40000 ALTER TABLE `move_flavor_text` DISABLE KEYS */;
/*!40000 ALTER TABLE `move_flavor_text` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `player_outfit`
--

DROP TABLE IF EXISTS `player_outfit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `player_outfit` (
  `idplayer` int(11) NOT NULL,
  `head` int(11) DEFAULT NULL,
  `nek` int(11) DEFAULT NULL,
  `upper` int(11) DEFAULT NULL,
  `lower` int(11) DEFAULT NULL,
  `feet` int(11) DEFAULT NULL,
  PRIMARY KEY (`idplayer`),
  UNIQUE KEY `idplayer_UNIQUE` (`idplayer`),
  KEY `fk_player_outfit_player1` (`idplayer`),
  CONSTRAINT `fk_player_outfit_player1` FOREIGN KEY (`idplayer`) REFERENCES `player` (`idplayer`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `player_outfit`
--

LOCK TABLES `player_outfit` WRITE;
/*!40000 ALTER TABLE `player_outfit` DISABLE KEYS */;

/*!40000 ALTER TABLE `player_outfit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_names`
--

DROP TABLE IF EXISTS `item_names`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_names` (
  `item_id` int(11) NOT NULL,
  `local_language_id` int(11) NOT NULL,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`item_id`,`local_language_id`),
  KEY `local_language_id` (`local_language_id`),
  KEY `ix_item_names_name` (`name`),
  CONSTRAINT `item_names_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `item_names_ibfk_2` FOREIGN KEY (`local_language_id`) REFERENCES `languages` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_names`
--

LOCK TABLES `item_names` WRITE;
/*!40000 ALTER TABLE `item_names` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_names` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pokemon_evolution`
--

DROP TABLE IF EXISTS `pokemon_evolution`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pokemon_evolution` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `evolved_species_id` int(11) NOT NULL,
  `evolution_trigger_id` int(11) NOT NULL,
  `trigger_item_id` int(11) DEFAULT NULL,
  `minimum_level` int(11) DEFAULT NULL,
  `gender` enum('male','female') DEFAULT NULL,
  `location_id` int(11) DEFAULT NULL,
  `held_item_id` int(11) DEFAULT NULL,
  `time_of_day` enum('day','night') DEFAULT NULL,
  `known_move_id` int(11) DEFAULT NULL,
  `minimum_happiness` int(11) DEFAULT NULL,
  `minimum_beauty` int(11) DEFAULT NULL,
  `relative_physical_stats` int(11) DEFAULT NULL,
  `party_species_id` int(11) DEFAULT NULL,
  `trade_species_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `evolved_species_id` (`evolved_species_id`),
  KEY `evolution_trigger_id` (`evolution_trigger_id`),
  KEY `trigger_item_id` (`trigger_item_id`),
  KEY `location_id` (`location_id`),
  KEY `held_item_id` (`held_item_id`),
  KEY `known_move_id` (`known_move_id`),
  KEY `party_species_id` (`party_species_id`),
  KEY `trade_species_id` (`trade_species_id`),
  CONSTRAINT `pokemon_evolution_ibfk_1` FOREIGN KEY (`evolved_species_id`) REFERENCES `pokemon_species` (`id`),
  CONSTRAINT `pokemon_evolution_ibfk_2` FOREIGN KEY (`evolution_trigger_id`) REFERENCES `evolution_triggers` (`id`),
  CONSTRAINT `pokemon_evolution_ibfk_3` FOREIGN KEY (`trigger_item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `pokemon_evolution_ibfk_5` FOREIGN KEY (`held_item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `pokemon_evolution_ibfk_6` FOREIGN KEY (`known_move_id`) REFERENCES `moves` (`id`),
  CONSTRAINT `pokemon_evolution_ibfk_7` FOREIGN KEY (`party_species_id`) REFERENCES `pokemon_species` (`id`),
  CONSTRAINT `pokemon_evolution_ibfk_8` FOREIGN KEY (`trade_species_id`) REFERENCES `pokemon_species` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=326 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pokemon_evolution`
--

LOCK TABLES `pokemon_evolution` WRITE;
/*!40000 ALTER TABLE `pokemon_evolution` DISABLE KEYS */;
/*!40000 ALTER TABLE `pokemon_evolution` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pokemon_stats`
--

DROP TABLE IF EXISTS `pokemon_stats`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pokemon_stats` (
  `pokemon_id` int(11) NOT NULL,
  `stat_id` int(11) NOT NULL,
  `base_stat` int(11) NOT NULL,
  `effort` int(11) NOT NULL,
  PRIMARY KEY (`pokemon_id`,`stat_id`),
  KEY `stat_id` (`stat_id`),
  CONSTRAINT `pokemon_stats_ibfk_1` FOREIGN KEY (`pokemon_id`) REFERENCES `pokemon` (`id`),
  CONSTRAINT `pokemon_stats_ibfk_2` FOREIGN KEY (`stat_id`) REFERENCES `stats` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pokemon_stats`
--

LOCK TABLES `pokemon_stats` WRITE;
/*!40000 ALTER TABLE `pokemon_stats` DISABLE KEYS */;
/*!40000 ALTER TABLE `pokemon_stats` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_flavor_summaries`
--

DROP TABLE IF EXISTS `item_flavor_summaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_flavor_summaries` (
  `item_id` int(11) NOT NULL,
  `local_language_id` int(11) NOT NULL,
  `flavor_summary` varchar(512) DEFAULT NULL,
  PRIMARY KEY (`item_id`,`local_language_id`),
  KEY `local_language_id` (`local_language_id`),
  CONSTRAINT `item_flavor_summaries_ibfk_1` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `item_flavor_summaries_ibfk_2` FOREIGN KEY (`local_language_id`) REFERENCES `languages` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_flavor_summaries`
--

LOCK TABLES `item_flavor_summaries` WRITE;
/*!40000 ALTER TABLE `item_flavor_summaries` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_flavor_summaries` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `identifier` varchar(45) DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL,
  `cost` int(11) DEFAULT NULL,
  `fling_power` int(11) DEFAULT NULL,
  `fling_effect_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `items`
--

LOCK TABLES `items` WRITE;
/*!40000 ALTER TABLE `items` DISABLE KEYS */;
/*!40000 ALTER TABLE `items` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pokemon_types`
--

DROP TABLE IF EXISTS `pokemon_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pokemon_types` (
  `pokemon_id` int(11) NOT NULL,
  `type_id` int(11) NOT NULL,
  `slot` int(11) NOT NULL,
  PRIMARY KEY (`pokemon_id`,`slot`),
  KEY `type_id` (`type_id`),
  CONSTRAINT `pokemon_types_ibfk_1` FOREIGN KEY (`pokemon_id`) REFERENCES `pokemon` (`id`),
  CONSTRAINT `pokemon_types_ibfk_2` FOREIGN KEY (`type_id`) REFERENCES `types` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pokemon_types`
--

LOCK TABLES `pokemon_types` WRITE;
/*!40000 ALTER TABLE `pokemon_types` DISABLE KEYS */;
/*!40000 ALTER TABLE `pokemon_types` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `quests`
--

DROP TABLE IF EXISTS `quests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `quests` (
  `idquests` int(11) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`idquests`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `quests`
--

LOCK TABLES `quests` WRITE;
/*!40000 ALTER TABLE `quests` DISABLE KEYS */;
/*!40000 ALTER TABLE `quests` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `npc_pokemon`
--

DROP TABLE IF EXISTS `npc_pokemon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `npc_pokemon` (
  `idnpc_pokemon` int(11) NOT NULL,
  `idpokemon` int(11) DEFAULT NULL,
  `idnpc` int(11) DEFAULT NULL,
  `iv_hp` tinyint(4) DEFAULT NULL,
  `iv_attack` tinyint(4) DEFAULT NULL,
  `iv_attack_spec` tinyint(4) DEFAULT NULL,
  `iv_defence` tinyint(4) DEFAULT NULL,
  `iv_defence_spec` tinyint(4) DEFAULT NULL,
  `iv_speed` tinyint(4) DEFAULT NULL,
  `gender` tinyint(4) DEFAULT NULL,
  `held_item` int(11) DEFAULT NULL,
  PRIMARY KEY (`idnpc_pokemon`),
  KEY `fk_npc_pokemon_npc1` (`idnpc`),
  KEY `fk_npc_pokemon_pokemon1` (`idpokemon`),
  KEY `fk_npc_pokemon_items1` (`held_item`),
  CONSTRAINT `fk_npc_pokemon_npc1` FOREIGN KEY (`idnpc`) REFERENCES `npc` (`idnpc`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_npc_pokemon_pokemon1` FOREIGN KEY (`idpokemon`) REFERENCES `pokemon` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_npc_pokemon_items1` FOREIGN KEY (`held_item`) REFERENCES `items` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `npc_pokemon`
--

LOCK TABLES `npc_pokemon` WRITE;
/*!40000 ALTER TABLE `npc_pokemon` DISABLE KEYS */;
/*!40000 ALTER TABLE `npc_pokemon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_fling_effects`
--

DROP TABLE IF EXISTS `item_fling_effects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `item_fling_effects` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_fling_effects`
--

LOCK TABLES `item_fling_effects` WRITE;
/*!40000 ALTER TABLE `item_fling_effects` DISABLE KEYS */;
/*!40000 ALTER TABLE `item_fling_effects` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2012-03-11 19:01:03
