SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL';

CREATE SCHEMA IF NOT EXISTS `puserver` DEFAULT CHARACTER SET utf8 ;
USE `puserver` ;

-- -----------------------------------------------------
-- Table `puserver`.`abilities`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`abilities` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`abilities` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `identifier` VARCHAR(24) NOT NULL ,
  `generation_id` INT(11) NOT NULL ,
  PRIMARY KEY (`id`) ,
  INDEX `generation_id` (`generation_id` ASC) )
ENGINE = InnoDB
AUTO_INCREMENT = 165
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`ability_messages`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`ability_messages` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`ability_messages` (
  `idability_messages` INT(11) NOT NULL AUTO_INCREMENT ,
  `ability_id` INT(11) NOT NULL ,
  `message` VARCHAR(255) NULL DEFAULT NULL ,
  PRIMARY KEY (`idability_messages`) )
ENGINE = InnoDB
AUTO_INCREMENT = 60
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`contest_effects`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`contest_effects` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`contest_effects` (
  `id` INT(11) NULL DEFAULT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`contest_types`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`contest_types` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`contest_types` (
  `id` INT(11) NULL DEFAULT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`encounter_condition`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`encounter_condition` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`encounter_condition` (
  `idencounter_condition` INT(11) NOT NULL ,
  `name` VARCHAR(250) NULL DEFAULT NULL ,
  `default` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idencounter_condition`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`encounter`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`encounter` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`encounter` (
  `idencounter` INT(11) NOT NULL AUTO_INCREMENT ,
  `idencounter_condition` INT(11) NOT NULL ,
  `rate` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idencounter`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`pokemon_species`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`pokemon_species` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_species` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `identifier` VARCHAR(20) NOT NULL ,
  `generation_id` INT(11) NULL DEFAULT NULL ,
  `evolves_from_species_id` INT(11) NULL DEFAULT NULL ,
  `evolution_chain_id` INT(11) NULL DEFAULT NULL ,
  `gender_rate` INT(11) NOT NULL ,
  `capture_rate` INT(11) NOT NULL ,
  `base_happiness` INT(11) NOT NULL ,
  `is_baby` TINYINT(1) NOT NULL ,
  `hatch_counter` INT(11) NOT NULL ,
  `has_gender_differences` TINYINT(1) NOT NULL ,
  `growth_rate_id` INT(11) NOT NULL ,
  `forms_switchable` TINYINT(1) NOT NULL ,
  `color_id` INT(11) NULL DEFAULT NULL ,
  `shape_id` INT(11) NULL DEFAULT NULL ,
  `habitat_id` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`id`) ,
  INDEX `evolution_chain_id` (`evolution_chain_id` ASC) ,
  INDEX `growth_rate_id` (`growth_rate_id` ASC) ,
  INDEX `fk_pokemon_species_pokemon_colors1` (`color_id` ASC) ,
  INDEX `fk_pokemon_species_pokemon_shapes1` (`shape_id` ASC) ,
  INDEX `fk_pokemon_species_pokemon_habitats1` (`habitat_id` ASC) )
ENGINE = InnoDB
AUTO_INCREMENT = 650
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`pokemon`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`pokemon` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `species_id` INT(11) NULL DEFAULT NULL ,
  `height` INT(11) NOT NULL ,
  `weight` INT(11) NOT NULL ,
  `base_experience` INT(11) NOT NULL ,
  `order` INT(11) NOT NULL ,
  `is_default` TINYINT(1) NOT NULL ,
  PRIMARY KEY (`id`) ,
  INDEX `ix_pokemon_is_default` (`is_default` ASC) ,
  INDEX `ix_pokemon_order` (`order` ASC) )
ENGINE = InnoDB
AUTO_INCREMENT = 668
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`encounter_slot`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`encounter_slot` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`encounter_slot` (
  `idencounter_slot` INT(11) NOT NULL ,
  `idencounter` INT(11) NOT NULL ,
  `idpokemon` INT(11) NOT NULL ,
  `gender_rate` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idencounter_slot`) ,
  INDEX `fk_encounter_idpokemon` (`idpokemon` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`group`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`group` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`group` (
  `idgroup` INT(11) NOT NULL ,
  `name` VARCHAR(45) NULL DEFAULT NULL ,
  `flags` INT(11) NULL DEFAULT NULL ,
  `priority` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idgroup`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`growth_rates`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`growth_rates` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`growth_rates` (
  `id` INT(11) NULL DEFAULT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_pockets`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_pockets` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_pockets` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `identifier` VARCHAR(16) NOT NULL ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 9
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_categories`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_categories` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_categories` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `pocket_id` INT(11) NOT NULL ,
  `identifier` VARCHAR(16) NOT NULL ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 45
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_category_prose`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_category_prose` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_category_prose` (
  `item_category_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `name` VARCHAR(16) NOT NULL ,
  INDEX `local_language_id` (`local_language_id` ASC) ,
  INDEX `ix_item_category_prose_name` (`name` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_flags`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_flags` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_flags` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `identifier` VARCHAR(24) NOT NULL ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 9
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_flag_map`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_flag_map` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_flag_map` (
  `item_id` INT(11) NOT NULL ,
  `item_flag_id` INT(11) NOT NULL ,
  PRIMARY KEY (`item_id`, `item_flag_id`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_flag_prose`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_flag_prose` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_flag_prose` (
  `item_flag_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `name` VARCHAR(24) NULL DEFAULT NULL ,
  `description` VARCHAR(64) NULL DEFAULT NULL ,
  INDEX `ix_item_flag_prose_name` (`name` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_flavor_summaries`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_flavor_summaries` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_flavor_summaries` (
  `item_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `flavor_summary` VARCHAR(512) NULL DEFAULT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_flavor_text`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_flavor_text` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_flavor_text` (
  `item_id` INT(11) NOT NULL ,
  `version_group_id` INT(11) NOT NULL ,
  `language_id` INT(11) NOT NULL ,
  `flavor_text` VARCHAR(255) NOT NULL ,
  INDEX `version_group_id` (`version_group_id` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_fling_effects`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_fling_effects` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_fling_effects` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 8
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_fling_effect_prose`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_fling_effect_prose` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_fling_effect_prose` (
  `item_fling_effect_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `effect` VARCHAR(255) NOT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_game_indices`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_game_indices` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_game_indices` (
  `item_id` INT(11) NOT NULL ,
  `generation_id` INT(11) NOT NULL ,
  `game_index` INT(11) NOT NULL ,
  INDEX `generation_id` (`generation_id` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_names`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_names` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_names` (
  `item_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `name` VARCHAR(20) NOT NULL ,
  INDEX `ix_item_names_name` (`name` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_pocket_names`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_pocket_names` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_pocket_names` (
  `item_pocket_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `name` VARCHAR(16) NOT NULL ,
  INDEX `ix_item_pocket_names_name` (`name` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`item_prose`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`item_prose` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_prose` (
  `item_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `short_effect` VARCHAR(256) NULL DEFAULT NULL ,
  `effect` VARCHAR(5120) NULL DEFAULT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`music`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`music` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`music` (
  `idmusic` INT(11) NOT NULL AUTO_INCREMENT ,
  `title` VARCHAR(45) NULL DEFAULT NULL ,
  `filename` VARCHAR(45) NULL DEFAULT NULL ,
  PRIMARY KEY (`idmusic`) )
ENGINE = InnoDB
AUTO_INCREMENT = 2
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`pokecenter`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`pokecenter` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokecenter` (
  `idpokecenter` INT(11) NOT NULL ,
  `position` BIGINT(20) NOT NULL ,
  `description` VARCHAR(250) NULL DEFAULT NULL ,
  PRIMARY KEY (`idpokecenter`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`location`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`location` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`location` (
  `idlocation` INT(11) NOT NULL ,
  `name` VARCHAR(45) NULL DEFAULT NULL ,
  `idpokecenter` INT(11) NULL DEFAULT NULL ,
  `idmusic` INT(11) NOT NULL ,
  PRIMARY KEY (`idlocation`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`location_section`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`location_section` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`location_section` (
  `idlocation_section` INT(11) NOT NULL AUTO_INCREMENT ,
  `idlocation` INT(11) NOT NULL ,
  `name` VARCHAR(250) NULL DEFAULT NULL ,
  PRIMARY KEY (`idlocation_section`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`location_encounter`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`location_encounter` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`location_encounter` (
  `idencounter` INT(11) NOT NULL ,
  `idlocation_section` INT(11) NOT NULL ,
  INDEX `fk_location_encounter_section` (`idlocation_section` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`map`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`map` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`map` (
  `idmap` INT(11) NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(128) NULL DEFAULT NULL ,
  PRIMARY KEY (`idmap`) )
ENGINE = InnoDB
AUTO_INCREMENT = 4
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`mapchange`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`mapchange` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`mapchange` (
  `idmapchange` INT(11) NOT NULL AUTO_INCREMENT ,
  `start_x` INT(11) NULL DEFAULT NULL ,
  `start_y` INT(11) NULL DEFAULT NULL ,
  `width` INT(11) NULL DEFAULT NULL ,
  `height` INT(11) NULL DEFAULT NULL ,
  `username` VARCHAR(45) NULL DEFAULT NULL ,
  `description` LONGTEXT NULL DEFAULT NULL ,
  `submit_date` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,
  `status` INT(11) NULL DEFAULT '0' ,
  PRIMARY KEY (`idmapchange`) )
ENGINE = InnoDB
AUTO_INCREMENT = 114
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`mapchange_account`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`mapchange_account` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`mapchange_account` (
  `idmapchange_account` INT(11) NOT NULL AUTO_INCREMENT ,
  `username` VARCHAR(45) NULL DEFAULT NULL ,
  `password` VARCHAR(255) NULL DEFAULT NULL ,
  PRIMARY KEY (`idmapchange_account`) )
ENGINE = InnoDB
AUTO_INCREMENT = 54
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`mapchange_tile`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`mapchange_tile` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`mapchange_tile` (
  `idmapchange_tile` INT(11) NOT NULL AUTO_INCREMENT ,
  `idmapchange` INT(11) NOT NULL ,
  `x` INT(11) NULL DEFAULT NULL ,
  `y` INT(11) NULL DEFAULT NULL ,
  `z` INT(11) NULL DEFAULT NULL ,
  `movement` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idmapchange_tile`) )
ENGINE = InnoDB
AUTO_INCREMENT = 10814
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`mapchange_layer`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`mapchange_layer` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`mapchange_layer` (
  `idmapchange_layer` INT(11) NOT NULL AUTO_INCREMENT ,
  `idmapchange_tile` INT(11) NOT NULL ,
  `index` INT(11) NULL DEFAULT NULL ,
  `sprite` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idmapchange_layer`) )
ENGINE = InnoDB
AUTO_INCREMENT = 18258
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`move_flavor_text`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`move_flavor_text` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`move_flavor_text` (
  `id_move` INT(11) NOT NULL ,
  `version_group_id` INT(11) NOT NULL ,
  `language_id` INT(11) NOT NULL ,
  `flavor_text` VARCHAR(255) NOT NULL ,
  PRIMARY KEY (`id_move`, `version_group_id`, `language_id`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`move_messages`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`move_messages` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`move_messages` (
  `move_effect_id` INT(11) NULL DEFAULT NULL ,
  `message` VARCHAR(255) NULL DEFAULT NULL )
ENGINE = MyISAM
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`types`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`types` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`types` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `identifier` VARCHAR(12) NOT NULL ,
  `generation_id` INT(11) NOT NULL ,
  `damage_class_id` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 10003
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`moves`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`moves` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`moves` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `identifier` VARCHAR(24) NOT NULL ,
  `generation_id` INT(11) NOT NULL ,
  `type_id` INT(11) NOT NULL ,
  `power` SMALLINT(6) NOT NULL ,
  `pp` SMALLINT(6) NULL DEFAULT NULL ,
  `accuracy` SMALLINT(6) NULL DEFAULT NULL ,
  `priority` SMALLINT(6) NOT NULL ,
  `target_id` INT(11) NOT NULL ,
  `damage_class_id` INT(11) NOT NULL ,
  `effect_id` INT(11) NOT NULL ,
  `effect_chance` INT(11) NULL DEFAULT NULL ,
  `contest_type_id` INT(11) NULL DEFAULT NULL ,
  `contest_effect_id` INT(11) NULL DEFAULT NULL ,
  `super_contest_effect_id` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`id`) ,
  INDEX `target_id` (`target_id` ASC) ,
  INDEX `damage_class_id` (`damage_class_id` ASC) ,
  INDEX `contest_type_id` (`contest_type_id` ASC) ,
  INDEX `contest_effect_id` (`contest_effect_id` ASC) ,
  INDEX `super_contest_effect_id` (`super_contest_effect_id` ASC) )
ENGINE = InnoDB
AUTO_INCREMENT = 10019
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`npc_outfit`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`npc_outfit` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`npc_outfit` (
  `idnpc` INT(11) NOT NULL ,
  `head` INT(11) NULL DEFAULT NULL ,
  `nek` INT(11) NULL DEFAULT NULL ,
  `upper` INT(11) NULL DEFAULT NULL ,
  `lower` INT(11) NULL DEFAULT NULL ,
  `feet` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idnpc`) )
ENGINE = MyISAM
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`npc`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`npc` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`npc` (
  `idnpc` INT(11) NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(45) NULL DEFAULT NULL ,
  `script_name` VARCHAR(45) NULL DEFAULT NULL ,
  `position` BIGINT(20) NULL DEFAULT NULL ,
  `idmap` INT(11) NULL DEFAULT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`npc_pokemon`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`npc_pokemon` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`npc_pokemon` (
  `idnpc_pokemon` INT(11) NOT NULL ,
  `idpokemon` INT(11) NULL DEFAULT NULL ,
  `idnpc` INT(11) NULL DEFAULT NULL ,
  `iv_hp` TINYINT(4) NULL DEFAULT NULL ,
  `iv_attack` TINYINT(4) NULL DEFAULT NULL ,
  `iv_attack_spec` TINYINT(4) NULL DEFAULT NULL ,
  `iv_defence` TINYINT(4) NULL DEFAULT NULL ,
  `iv_defence_spec` TINYINT(4) NULL DEFAULT NULL ,
  `iv_speed` TINYINT(4) NULL DEFAULT NULL ,
  `gender` TINYINT(4) NULL DEFAULT NULL ,
  `held_item` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idnpc_pokemon`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`npc_pokemon_move`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`npc_pokemon_move` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`npc_pokemon_move` (
  `idnpc_pokemon_move` INT(11) NOT NULL ,
  `idnpc_pokemon` INT(11) NULL DEFAULT NULL ,
  `idmove` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idnpc_pokemon_move`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`player`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`player` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`player` (
  `idplayer` INT(11) NOT NULL AUTO_INCREMENT ,
  `idaccount` INT(11) NULL DEFAULT NULL ,
  `name` VARCHAR(20) NULL DEFAULT NULL ,
  `password` VARCHAR(45) NULL DEFAULT NULL ,
  `password_salt` VARCHAR(45) NULL DEFAULT NULL ,
  `position` BIGINT(20) NULL DEFAULT NULL COMMENT 'x;y;z' ,
  `movement` SMALLINT(6) NULL DEFAULT NULL ,
  `idpokecenter` INT(11) NULL DEFAULT NULL ,
  `money` INT(11) NULL DEFAULT NULL ,
  `idlocation` INT(11) NOT NULL ,
  PRIMARY KEY (`idplayer`) )
ENGINE = InnoDB
AUTO_INCREMENT = 3
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`player_backpack`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`player_backpack` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_backpack` (
  `idplayer_backpack` INT(11) NOT NULL ,
  `idplayer` INT(11) NOT NULL ,
  `iditem` INT(11) NOT NULL ,
  `count` INT(11) NULL DEFAULT '1' ,
  `slot` INT(11) NOT NULL ,
  PRIMARY KEY (`idplayer_backpack`) )
ENGINE = MyISAM
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`player_group`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`player_group` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_group` (
  `player_idplayer` INT(11) NOT NULL ,
  `group_idgroup` INT(11) NOT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`player_items`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`player_items` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_items` (
  `idplayer_items` INT(11) NOT NULL ,
  `idplayer` INT(11) NULL DEFAULT NULL ,
  `iditem` INT(11) NULL DEFAULT NULL ,
  `count` INT(11) NULL DEFAULT NULL ,
  `slot` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idplayer_items`) ,
  INDEX `fk_player_items_player1` (`idplayer` ASC) ,
  INDEX `fk_player_items_items1` (`iditem` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`player_outfit`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`player_outfit` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_outfit` (
  `idplayer` INT(11) NOT NULL ,
  `head` INT(11) NULL DEFAULT NULL ,
  `nek` INT(11) NULL DEFAULT NULL ,
  `upper` INT(11) NULL DEFAULT NULL ,
  `lower` INT(11) NULL DEFAULT NULL ,
  `feet` INT(11) NULL DEFAULT NULL ,
  UNIQUE INDEX `idplayer_UNIQUE` (`idplayer` ASC) ,
  INDEX `fk_player_outfit_player1` (`idplayer` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`player_pokemon`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`player_pokemon` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_pokemon` (
  `idplayer_pokemon` INT(11) NOT NULL AUTO_INCREMENT ,
  `idpokemon` INT(11) NOT NULL ,
  `idplayer` INT(11) NOT NULL ,
  `nickname` VARCHAR(45) NULL DEFAULT NULL ,
  `bound` TINYINT(4) NULL DEFAULT '0' COMMENT '1 if pokemon is bound to player' ,
  `experience` INT(10) UNSIGNED NULL DEFAULT NULL ,
  `iv_hp` TINYINT(3) UNSIGNED NULL DEFAULT NULL ,
  `iv_attack` TINYINT(3) UNSIGNED NULL DEFAULT NULL ,
  `iv_attack_spec` TINYINT(3) UNSIGNED NULL DEFAULT NULL ,
  `iv_defence` TINYINT(3) UNSIGNED NULL DEFAULT NULL ,
  `iv_defence_spec` TINYINT(3) UNSIGNED NULL DEFAULT NULL ,
  `iv_speed` TINYINT(3) UNSIGNED NULL DEFAULT NULL ,
  `happiness` TINYINT(3) UNSIGNED NULL DEFAULT NULL ,
  `gender` TINYINT(4) NULL DEFAULT NULL COMMENT '-1 None\n0 Male\n1 Female' ,
  `in_party` TINYINT(1) NULL DEFAULT NULL ,
  `party_slot` TINYINT(1) NULL DEFAULT NULL ,
  `held_item` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idplayer_pokemon`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`player_pokemon_move`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`player_pokemon_move` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_pokemon_move` (
  `idplayer_pokemon_move` INT(11) NOT NULL ,
  `idplayer_pokemon` INT(11) NOT NULL ,
  `idmove` INT(11) NOT NULL ,
  `pp_used` SMALLINT(6) NULL DEFAULT NULL ,
  PRIMARY KEY (`idplayer_pokemon_move`) ,
  INDEX `fk_player_pokemon_move_moves1` (`idmove` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`quests`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`quests` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`quests` (
  `idquests` INT(11) NOT NULL ,
  `name` VARCHAR(45) NULL DEFAULT NULL ,
  PRIMARY KEY (`idquests`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`player_quests`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`player_quests` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_quests` (
  `idplayer_quests` INT(11) NOT NULL ,
  `idplayer` INT(11) NULL DEFAULT NULL ,
  `idquest` INT(11) NULL DEFAULT NULL ,
  `status` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idplayer_quests`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`pokemon_abilities`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`pokemon_abilities` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_abilities` (
  `pokemon_id` INT(11) NOT NULL ,
  `ability_id` INT(11) NOT NULL ,
  `is_dream` TINYINT(1) NOT NULL ,
  `slot` INT(11) NOT NULL ,
  INDEX `ix_pokemon_abilities_is_dream` (`is_dream` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`pokemon_evolution`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`pokemon_evolution` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_evolution` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `evolved_species_id` INT(11) NOT NULL ,
  `evolution_trigger_id` INT(11) NOT NULL ,
  `trigger_item_id` INT(11) NULL DEFAULT NULL ,
  `minimum_level` INT(11) NULL DEFAULT NULL ,
  `gender` ENUM('male','female') NULL DEFAULT NULL ,
  `location_id` INT(11) NULL DEFAULT NULL ,
  `held_item_id` INT(11) NULL DEFAULT NULL ,
  `time_of_day` ENUM('day','night') NULL DEFAULT NULL ,
  `known_move_id` INT(11) NULL DEFAULT NULL ,
  `minimum_happiness` INT(11) NULL DEFAULT NULL ,
  `minimum_beauty` INT(11) NULL DEFAULT NULL ,
  `relative_physical_stats` INT(11) NULL DEFAULT NULL ,
  `party_species_id` INT(11) NULL DEFAULT NULL ,
  `trade_species_id` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`id`) ,
  INDEX `evolution_trigger_id` (`evolution_trigger_id` ASC) ,
  INDEX `location_id` (`location_id` ASC) )
ENGINE = InnoDB
AUTO_INCREMENT = 326
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`pokemon_forms`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`pokemon_forms` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_forms` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `form_identifier` VARCHAR(16) NULL DEFAULT NULL ,
  `pokemon_id` INT(11) NOT NULL ,
  `is_default` TINYINT(1) NOT NULL ,
  `is_battle_only` TINYINT(1) NOT NULL ,
  `order` INT(11) NOT NULL ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 728
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`pokemon_moves`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`pokemon_moves` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_moves` (
  `pokemon_id` INT(11) NOT NULL ,
  `version_group_id` INT(11) NOT NULL ,
  `move_id` INT(11) NOT NULL ,
  `pokemon_move_method_id` INT(11) NOT NULL ,
  `level` INT(11) NOT NULL ,
  `order` INT(11) NULL DEFAULT NULL ,
  INDEX `idx_autoinc_level` (`level` ASC) ,
  INDEX `ix_pokemon_moves_version_group_id` (`version_group_id` ASC) ,
  INDEX `ix_pokemon_moves_level` (`level` ASC) ,
  INDEX `ix_pokemon_moves_pokemon_id` (`pokemon_id` ASC) ,
  INDEX `ix_pokemon_moves_pokemon_move_method_id` (`pokemon_move_method_id` ASC) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`stats`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`stats` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`stats` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `damage_class_id` INT(11) NULL DEFAULT NULL ,
  `identifier` VARCHAR(16) NOT NULL ,
  `is_battle_only` TINYINT(1) NOT NULL ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 9
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`pokemon_stats`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`pokemon_stats` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_stats` (
  `pokemon_id` INT(11) NOT NULL ,
  `stat_id` INT(11) NOT NULL ,
  `base_stat` INT(11) NOT NULL ,
  `effort` INT(11) NOT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`pokemon_types`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`pokemon_types` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_types` (
  `pokemon_id` INT(11) NOT NULL ,
  `type_id` INT(11) NOT NULL ,
  `slot` INT(11) NOT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`super_contest_effects`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`super_contest_effects` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`super_contest_effects` (
  `id` INT(11) NULL DEFAULT NULL )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`teleport`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`teleport` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`teleport` (
  `idteleport` INT(11) NOT NULL AUTO_INCREMENT ,
  `x` INT(11) NULL DEFAULT NULL ,
  `y` INT(11) NULL DEFAULT NULL ,
  `z` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idteleport`) )
ENGINE = InnoDB
AUTO_INCREMENT = 224
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`tile`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`tile` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`tile` (
  `idtile` INT(11) NOT NULL AUTO_INCREMENT ,
  `x` INT(11) NOT NULL ,
  `y` INT(11) NOT NULL ,
  `z` INT(11) NOT NULL ,
  `idlocation` INT(11) NOT NULL ,
  `movement` INT(11) NULL DEFAULT NULL ,
  `script` VARCHAR(128) NULL DEFAULT NULL ,
  `idteleport` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idtile`) ,
  INDEX `position_key` (`x` ASC, `y` ASC, `z` ASC) )
ENGINE = InnoDB
AUTO_INCREMENT = 1132711
DEFAULT CHARACTER SET = utf8;


-- -----------------------------------------------------
-- Table `puserver`.`tile_layer`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `puserver`.`tile_layer` ;

CREATE  TABLE IF NOT EXISTS `puserver`.`tile_layer` (
  `idtile_layer` INT(11) NOT NULL AUTO_INCREMENT ,
  `idtile` INT(11) NOT NULL ,
  `sprite` INT(11) NULL DEFAULT NULL ,
  `layer` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idtile_layer`, `idtile`) )
ENGINE = InnoDB
AUTO_INCREMENT = 1549597
DEFAULT CHARACTER SET = utf8;



SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
