CREATE DATABASE  IF NOT EXISTS `puserver` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `puserver`;
-- MySQL dump 10.13  Distrib 5.1.40, for Win32 (ia32)
--
-- Host: localhost    Database: puserver
-- ------------------------------------------------------
-- Server version	5.1.48-community

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
  CONSTRAINT `fk_encounter_slot_encounter1` FOREIGN KEY (`idencounter`) REFERENCES `encounter` (`idencounter`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  `introduced_in_version_group_id` int(11) DEFAULT NULL,
  `is_default` tinyint(1) NOT NULL,
  `is_battle_only` tinyint(1) NOT NULL,
  `order` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `pokemon_id` (`pokemon_id`),
  KEY `introduced_in_version_group_id` (`introduced_in_version_group_id`),
  CONSTRAINT `pokemon_forms_ibfk_1` FOREIGN KEY (`pokemon_id`) REFERENCES `pokemon` (`id`),
  CONSTRAINT `pokemon_forms_ibfk_2` FOREIGN KEY (`introduced_in_version_group_id`) REFERENCES `version_groups` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=728 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  KEY `generation_id` (`generation_id`),
  CONSTRAINT `abilities_ibfk_1` FOREIGN KEY (`generation_id`) REFERENCES `generations` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=165 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `move`
--

DROP TABLE IF EXISTS `move`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `move` (
  `idmove` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`idmove`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `encounter`
--

DROP TABLE IF EXISTS `encounter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `encounter` (
  `idencounter` int(11) NOT NULL AUTO_INCREMENT,
  `idlocation_section` int(11) NOT NULL,
  `idencounter_condition` int(11) NOT NULL,
  `rate` int(11) DEFAULT NULL,
  PRIMARY KEY (`idencounter`),
  KEY `fk_encounter_encounter_condition1` (`idencounter_condition`),
  CONSTRAINT `fk_encounter_encounter_condition1` FOREIGN KEY (`idencounter_condition`) REFERENCES `encounter_condition` (`idencounter_condition`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `move_method`
--

DROP TABLE IF EXISTS `move_method`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `move_method` (
  `idmove_method` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `description` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`idmove_method`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  `color_id` int(11) NOT NULL,
  `shape_id` int(11) NOT NULL,
  `habitat_id` int(11) DEFAULT NULL,
  `gender_rate` int(11) NOT NULL,
  `capture_rate` int(11) NOT NULL,
  `base_happiness` int(11) NOT NULL,
  `is_baby` tinyint(1) NOT NULL,
  `hatch_counter` int(11) NOT NULL,
  `has_gender_differences` tinyint(1) NOT NULL,
  `growth_rate_id` int(11) NOT NULL,
  `forms_switchable` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `generation_id` (`generation_id`),
  KEY `evolves_from_species_id` (`evolves_from_species_id`),
  KEY `evolution_chain_id` (`evolution_chain_id`),
  KEY `color_id` (`color_id`),
  KEY `shape_id` (`shape_id`),
  KEY `habitat_id` (`habitat_id`),
  KEY `growth_rate_id` (`growth_rate_id`),
  CONSTRAINT `pokemon_species_ibfk_1` FOREIGN KEY (`generation_id`) REFERENCES `generations` (`id`),
  CONSTRAINT `pokemon_species_ibfk_2` FOREIGN KEY (`evolves_from_species_id`) REFERENCES `pokemon_species` (`id`),
  CONSTRAINT `pokemon_species_ibfk_3` FOREIGN KEY (`evolution_chain_id`) REFERENCES `evolution_chains` (`id`),
  CONSTRAINT `pokemon_species_ibfk_4` FOREIGN KEY (`color_id`) REFERENCES `pokemon_colors` (`id`),
  CONSTRAINT `pokemon_species_ibfk_5` FOREIGN KEY (`shape_id`) REFERENCES `pokemon_shapes` (`id`),
  CONSTRAINT `pokemon_species_ibfk_6` FOREIGN KEY (`habitat_id`) REFERENCES `pokemon_habitats` (`id`),
  CONSTRAINT `pokemon_species_ibfk_7` FOREIGN KEY (`growth_rate_id`) REFERENCES `growth_rates` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=650 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  KEY `generation_id` (`generation_id`),
  KEY `type_id` (`type_id`),
  KEY `target_id` (`target_id`),
  KEY `damage_class_id` (`damage_class_id`),
  KEY `effect_id` (`effect_id`),
  KEY `contest_type_id` (`contest_type_id`),
  KEY `contest_effect_id` (`contest_effect_id`),
  KEY `super_contest_effect_id` (`super_contest_effect_id`),
  CONSTRAINT `moves_ibfk_1` FOREIGN KEY (`generation_id`) REFERENCES `generations` (`id`),
  CONSTRAINT `moves_ibfk_2` FOREIGN KEY (`type_id`) REFERENCES `types` (`id`),
  CONSTRAINT `moves_ibfk_3` FOREIGN KEY (`target_id`) REFERENCES `move_targets` (`id`),
  CONSTRAINT `moves_ibfk_4` FOREIGN KEY (`damage_class_id`) REFERENCES `move_damage_classes` (`id`),
  CONSTRAINT `moves_ibfk_5` FOREIGN KEY (`effect_id`) REFERENCES `move_effects` (`id`),
  CONSTRAINT `moves_ibfk_6` FOREIGN KEY (`contest_type_id`) REFERENCES `contest_types` (`id`),
  CONSTRAINT `moves_ibfk_7` FOREIGN KEY (`contest_effect_id`) REFERENCES `contest_effects` (`id`),
  CONSTRAINT `moves_ibfk_8` FOREIGN KEY (`super_contest_effect_id`) REFERENCES `super_contest_effects` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10019 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=223 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `group`
--

DROP TABLE IF EXISTS `group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `group` (
  `idgroup` int(11) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `flags` int(11) DEFAULT NULL,
  `priority` int(11) DEFAULT NULL,
  PRIMARY KEY (`idgroup`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `location_encounter`
--

DROP TABLE IF EXISTS `location_encounter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `location_encounter` (
  `idencounter` int(11) NOT NULL,
  `idlocation` int(11) NOT NULL,
  PRIMARY KEY (`idencounter`,`idlocation`),
  KEY `fk_location_encounter` (`idlocation`),
  CONSTRAINT `fk_location_encounter_encounter` FOREIGN KEY (`idencounter`) REFERENCES `encounter` (`idencounter`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  PRIMARY KEY (`idplayer_pokemon_move`),
  KEY `fk_player_pokemon_move_player_pokemon1` (`idplayer_pokemon`),
  KEY `fk_player_pokemon_move_move1` (`idmove`),
  CONSTRAINT `fk_player_pokemon_move_move1` FOREIGN KEY (`idmove`) REFERENCES `move` (`idmove`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_pokemon_move_player_pokemon1` FOREIGN KEY (`idplayer_pokemon`) REFERENCES `player_pokemon` (`idplayer_pokemon`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  PRIMARY KEY (`id`),
  KEY `generation_id` (`generation_id`),
  KEY `damage_class_id` (`damage_class_id`),
  CONSTRAINT `types_ibfk_1` FOREIGN KEY (`generation_id`) REFERENCES `generations` (`id`),
  CONSTRAINT `types_ibfk_2` FOREIGN KEY (`damage_class_id`) REFERENCES `move_damage_classes` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10003 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tile`
--

DROP TABLE IF EXISTS `tile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tile` (
  `idtile` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `x` int(11) NOT NULL,
  `y` int(11) NOT NULL,
  `z` int(11) NOT NULL,
  `idlocation` int(11) NOT NULL,
  `idmap` int(11) NOT NULL,
  `movement` int(11) DEFAULT NULL,
  `script` varchar(128) DEFAULT NULL,
  `idteleport` int(11) DEFAULT NULL,
  PRIMARY KEY (`idtile`),
  KEY `fk_tile_location1` (`idlocation`),
  KEY `position_key` (`x`,`y`,`z`),
  KEY `fk_tile_map1` (`idmap`),
  KEY `fk_tile_teleport1` (`idteleport`),
  CONSTRAINT `fk_tile_location1` FOREIGN KEY (`idlocation`) REFERENCES `location` (`idlocation`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_tile_map1` FOREIGN KEY (`idmap`) REFERENCES `map` (`idmap`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=1129442 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  KEY `fk_tile_layer_tile1` (`idtile`)
) ENGINE=InnoDB AUTO_INCREMENT=1541022 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `map`
--

DROP TABLE IF EXISTS `map`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `map` (
  `idmap` int(11) NOT NULL,
  `name` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`idmap`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=10814 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  PRIMARY KEY (`id`),
  KEY `damage_class_id` (`damage_class_id`),
  CONSTRAINT `stats_ibfk_1` FOREIGN KEY (`damage_class_id`) REFERENCES `move_damage_classes` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  CONSTRAINT `pokemon_evolution_ibfk_4` FOREIGN KEY (`location_id`) REFERENCES `locations` (`id`),
  CONSTRAINT `pokemon_evolution_ibfk_5` FOREIGN KEY (`held_item_id`) REFERENCES `items` (`id`),
  CONSTRAINT `pokemon_evolution_ibfk_6` FOREIGN KEY (`known_move_id`) REFERENCES `moves` (`id`),
  CONSTRAINT `pokemon_evolution_ibfk_7` FOREIGN KEY (`party_species_id`) REFERENCES `pokemon_species` (`id`),
  CONSTRAINT `pokemon_evolution_ibfk_8` FOREIGN KEY (`trade_species_id`) REFERENCES `pokemon_species` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=326 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=18258 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  PRIMARY KEY (`idplayer_pokemon`),
  KEY `fk_player_pokemon_player` (`idplayer`),
  KEY `fk_player_pokemon_pokemon` (`idpokemon`),
  CONSTRAINT `fk_player_pokemon_player` FOREIGN KEY (`idplayer`) REFERENCES `player` (`idplayer`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

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
  `level` int(11) NOT NULL AUTO_INCREMENT,
  `order` int(11) DEFAULT NULL,
  PRIMARY KEY (`pokemon_id`,`version_group_id`,`move_id`,`pokemon_move_method_id`,`level`),
  KEY `idx_autoinc_level` (`level`),
  KEY `ix_pokemon_moves_version_group_id` (`version_group_id`),
  KEY `ix_pokemon_moves_level` (`level`),
  KEY `ix_pokemon_moves_pokemon_id` (`pokemon_id`),
  KEY `ix_pokemon_moves_pokemon_move_method_id` (`pokemon_move_method_id`),
  KEY `ix_pokemon_moves_move_id` (`move_id`),
  CONSTRAINT `pokemon_moves_ibfk_1` FOREIGN KEY (`pokemon_id`) REFERENCES `pokemon` (`id`),
  CONSTRAINT `pokemon_moves_ibfk_2` FOREIGN KEY (`version_group_id`) REFERENCES `version_groups` (`id`),
  CONSTRAINT `pokemon_moves_ibfk_3` FOREIGN KEY (`move_id`) REFERENCES `moves` (`id`),
  CONSTRAINT `pokemon_moves_ibfk_4` FOREIGN KEY (`pokemon_move_method_id`) REFERENCES `pokemon_move_methods` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=150209 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=114 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2011-11-06 14:41:44
