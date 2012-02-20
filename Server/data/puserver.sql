SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL';

CREATE  TABLE IF NOT EXISTS `puserver`.`abilities` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `identifier` VARCHAR(24) NOT NULL ,
  `generation_id` INT(11) NOT NULL ,
  PRIMARY KEY (`id`) ,
  INDEX `generation_id` (`generation_id` ASC) )
ENGINE = InnoDB
AUTO_INCREMENT = 165
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`encounter` (
  `idencounter` INT(11) NOT NULL AUTO_INCREMENT ,
  `idencounter_condition` INT(11) NOT NULL ,
  `rate` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idencounter`) ,
  INDEX `fk_encounter_encounter_condition1` (`idencounter_condition` ASC) ,
  CONSTRAINT `fk_encounter_encounter_condition1`
    FOREIGN KEY (`idencounter_condition` )
    REFERENCES `puserver`.`encounter_condition` (`idencounter_condition` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`encounter_condition` (
  `idencounter_condition` INT(11) NOT NULL ,
  `name` VARCHAR(250) NULL DEFAULT NULL ,
  `default` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idencounter_condition`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`encounter_slot` (
  `idencounter_slot` INT(11) NOT NULL ,
  `idencounter` INT(11) NOT NULL ,
  `idpokemon` INT(11) NOT NULL ,
  `gender_rate` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idencounter_slot`) ,
  INDEX `fk_encounter_slot_encounter1` (`idencounter` ASC) ,
  INDEX `fk_encounter_slot_pokemon1` (`idpokemon` ASC) ,
  INDEX `fk_encounter_idpokemon` (`idpokemon` ASC) ,
  CONSTRAINT `fk_encounter_slot_encounter1`
    FOREIGN KEY (`idencounter` )
    REFERENCES `puserver`.`encounter` (`idencounter` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_encounter_idpokemon`
    FOREIGN KEY (`idpokemon` )
    REFERENCES `puserver`.`pokemon` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`group` (
  `idgroup` INT(11) NOT NULL ,
  `name` VARCHAR(45) NULL DEFAULT NULL ,
  `flags` INT(11) NULL DEFAULT NULL ,
  `priority` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idgroup`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`location` (
  `idlocation` INT(11) NOT NULL ,
  `name` VARCHAR(45) NULL DEFAULT NULL ,
  `idpokecenter` INT(11) NULL DEFAULT NULL ,
  `idmusic` INT(11) NOT NULL ,
  PRIMARY KEY (`idlocation`) ,
  INDEX `fk_location_pokecenter1` (`idpokecenter` ASC) ,
  INDEX `fk_location_music1` (`idmusic` ASC) ,
  CONSTRAINT `fk_location_music1`
    FOREIGN KEY (`idmusic` )
    REFERENCES `puserver`.`music` (`idmusic` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_location_pokecenter1`
    FOREIGN KEY (`idpokecenter` )
    REFERENCES `puserver`.`pokecenter` (`idpokecenter` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`location_encounter` (
  `idencounter` INT(11) NOT NULL ,
  `idlocation_section` INT(11) NOT NULL ,
  PRIMARY KEY (`idencounter`, `idlocation_section`) ,
  INDEX `fk_location_encounter` (`idlocation_section` ASC) ,
  INDEX `fk_location_encounter_section` (`idlocation_section` ASC) ,
  CONSTRAINT `fk_location_encounter_encounter`
    FOREIGN KEY (`idencounter` )
    REFERENCES `puserver`.`encounter` (`idencounter` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_location_encounter_section`
    FOREIGN KEY (`idlocation_section` )
    REFERENCES `puserver`.`location_section` (`idlocation_section` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`location_section` (
  `idlocation_section` INT(11) NOT NULL AUTO_INCREMENT ,
  `idlocation` INT(11) NOT NULL ,
  `name` VARCHAR(250) NULL DEFAULT NULL ,
  PRIMARY KEY (`idlocation_section`) ,
  INDEX `fk_location_sections_location1` (`idlocation` ASC) ,
  CONSTRAINT `fk_location_sections_location1`
    FOREIGN KEY (`idlocation` )
    REFERENCES `puserver`.`location` (`idlocation` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`map` (
  `idmap` INT(11) NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(128) NULL DEFAULT NULL ,
  PRIMARY KEY (`idmap`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

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
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`mapchange_account` (
  `idmapchange_account` INT(11) NOT NULL AUTO_INCREMENT ,
  `username` VARCHAR(45) NULL DEFAULT NULL ,
  `password` VARCHAR(255) NULL DEFAULT NULL ,
  PRIMARY KEY (`idmapchange_account`) )
ENGINE = InnoDB
AUTO_INCREMENT = 54
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`mapchange_layer` (
  `idmapchange_layer` INT(11) NOT NULL AUTO_INCREMENT ,
  `idmapchange_tile` INT(11) NOT NULL ,
  `index` INT(11) NULL DEFAULT NULL ,
  `sprite` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idmapchange_layer`) ,
  INDEX `fk_mapchange_layer_mapchange_tile1` (`idmapchange_tile` ASC) ,
  CONSTRAINT `fk_mapchange_layer_mapchange_tile1`
    FOREIGN KEY (`idmapchange_tile` )
    REFERENCES `puserver`.`mapchange_tile` (`idmapchange_tile` )
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 18258
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`mapchange_tile` (
  `idmapchange_tile` INT(11) NOT NULL AUTO_INCREMENT ,
  `idmapchange` INT(11) NOT NULL ,
  `x` INT(11) NULL DEFAULT NULL ,
  `y` INT(11) NULL DEFAULT NULL ,
  `z` INT(11) NULL DEFAULT NULL ,
  `movement` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idmapchange_tile`) ,
  INDEX `fk_mapchange_tile_mapchange1` (`idmapchange` ASC) ,
  CONSTRAINT `fk_mapchange_tile_mapchange1`
    FOREIGN KEY (`idmapchange` )
    REFERENCES `puserver`.`mapchange` (`idmapchange` )
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 10814
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

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
  INDEX `type_id` (`type_id` ASC) ,
  INDEX `target_id` (`target_id` ASC) ,
  INDEX `damage_class_id` (`damage_class_id` ASC) ,
  INDEX `effect_id` (`effect_id` ASC) ,
  INDEX `contest_type_id` (`contest_type_id` ASC) ,
  INDEX `contest_effect_id` (`contest_effect_id` ASC) ,
  INDEX `super_contest_effect_id` (`super_contest_effect_id` ASC) ,
  CONSTRAINT `moves_ibfk_2`
    FOREIGN KEY (`type_id` )
    REFERENCES `puserver`.`types` (`id` ))
ENGINE = InnoDB
AUTO_INCREMENT = 10019
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`music` (
  `idmusic` INT(11) NOT NULL AUTO_INCREMENT ,
  `title` VARCHAR(45) NULL DEFAULT NULL ,
  `filename` VARCHAR(45) NULL DEFAULT NULL ,
  PRIMARY KEY (`idmusic`) )
ENGINE = InnoDB
AUTO_INCREMENT = 2
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

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
  PRIMARY KEY (`idplayer`) ,
  INDEX `fk_player_location1` (`idlocation` ASC) ,
  CONSTRAINT `fk_player_location1`
    FOREIGN KEY (`idlocation` )
    REFERENCES `puserver`.`location` (`idlocation` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 3
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_group` (
  `player_idplayer` INT(11) NOT NULL ,
  `group_idgroup` INT(11) NOT NULL ,
  PRIMARY KEY (`player_idplayer`, `group_idgroup`) ,
  INDEX `fk_player_has_group_group1` (`group_idgroup` ASC) ,
  CONSTRAINT `fk_player_has_group_group1`
    FOREIGN KEY (`group_idgroup` )
    REFERENCES `puserver`.`group` (`idgroup` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_has_group_player1`
    FOREIGN KEY (`player_idplayer` )
    REFERENCES `puserver`.`player` (`idplayer` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_outfit` (
  `idplayer` INT(11) NOT NULL ,
  `head` INT(11) NULL DEFAULT NULL ,
  `nek` INT(11) NULL DEFAULT NULL ,
  `upper` INT(11) NULL DEFAULT NULL ,
  `lower` INT(11) NULL DEFAULT NULL ,
  `feet` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idplayer`) ,
  UNIQUE INDEX `idplayer_UNIQUE` (`idplayer` ASC) ,
  INDEX `fk_player_outfit_player1` (`idplayer` ASC) ,
  CONSTRAINT `fk_player_outfit_player1`
    FOREIGN KEY (`idplayer` )
    REFERENCES `puserver`.`player` (`idplayer` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

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
  PRIMARY KEY (`idplayer_pokemon`) ,
  INDEX `fk_player_pokemon_player` (`idplayer` ASC) ,
  INDEX `fk_player_pokemon_pokemon` (`idpokemon` ASC) ,
  INDEX `fk_player_pokemon_items1` (`held_item` ASC) ,
  CONSTRAINT `fk_player_pokemon_player`
    FOREIGN KEY (`idplayer` )
    REFERENCES `puserver`.`player` (`idplayer` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_pokemon_pokemon1`
    FOREIGN KEY (`idpokemon` )
    REFERENCES `puserver`.`pokemon` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_pokemon_items1`
    FOREIGN KEY (`held_item` )
    REFERENCES `puserver`.`items` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_pokemon_move` (
  `idplayer_pokemon_move` INT(11) NOT NULL ,
  `idplayer_pokemon` INT(11) NOT NULL ,
  `idmove` INT(11) NOT NULL ,
  `pp_used` SMALLINT(6) NULL DEFAULT NULL ,
  PRIMARY KEY (`idplayer_pokemon_move`) ,
  INDEX `fk_player_pokemon_move_player_pokemon1` (`idplayer_pokemon` ASC) ,
  INDEX `fk_player_pokemon_move_move1` (`idmove` ASC) ,
  INDEX `fk_player_pokemon_move_moves1` (`idmove` ASC) ,
  CONSTRAINT `fk_player_pokemon_move_player_pokemon1`
    FOREIGN KEY (`idplayer_pokemon` )
    REFERENCES `puserver`.`player_pokemon` (`idplayer_pokemon` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_pokemon_move_moves1`
    FOREIGN KEY (`idmove` )
    REFERENCES `puserver`.`moves` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokecenter` (
  `idpokecenter` INT(11) NOT NULL ,
  `position` BIGINT(20) NOT NULL ,
  `description` VARCHAR(250) NULL DEFAULT NULL ,
  PRIMARY KEY (`idpokecenter`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `species_id` INT(11) NULL DEFAULT NULL ,
  `height` INT(11) NOT NULL ,
  `weight` INT(11) NOT NULL ,
  `base_experience` INT(11) NOT NULL ,
  `order` INT(11) NOT NULL ,
  `is_default` TINYINT(1) NOT NULL ,
  PRIMARY KEY (`id`) ,
  INDEX `species_id` (`species_id` ASC) ,
  INDEX `ix_pokemon_is_default` (`is_default` ASC) ,
  INDEX `ix_pokemon_order` (`order` ASC) ,
  CONSTRAINT `pokemon_ibfk_1`
    FOREIGN KEY (`species_id` )
    REFERENCES `puserver`.`pokemon_species` (`id` ))
ENGINE = InnoDB
AUTO_INCREMENT = 668
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_abilities` (
  `pokemon_id` INT(11) NOT NULL ,
  `ability_id` INT(11) NOT NULL ,
  `is_dream` TINYINT(1) NOT NULL ,
  `slot` INT(11) NOT NULL ,
  PRIMARY KEY (`pokemon_id`, `slot`) ,
  INDEX `ability_id` (`ability_id` ASC) ,
  INDEX `ix_pokemon_abilities_is_dream` (`is_dream` ASC) ,
  CONSTRAINT `pokemon_abilities_ibfk_1`
    FOREIGN KEY (`pokemon_id` )
    REFERENCES `puserver`.`pokemon` (`id` ),
  CONSTRAINT `pokemon_abilities_ibfk_2`
    FOREIGN KEY (`ability_id` )
    REFERENCES `puserver`.`abilities` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

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
  INDEX `evolved_species_id` (`evolved_species_id` ASC) ,
  INDEX `evolution_trigger_id` (`evolution_trigger_id` ASC) ,
  INDEX `trigger_item_id` (`trigger_item_id` ASC) ,
  INDEX `location_id` (`location_id` ASC) ,
  INDEX `held_item_id` (`held_item_id` ASC) ,
  INDEX `known_move_id` (`known_move_id` ASC) ,
  INDEX `party_species_id` (`party_species_id` ASC) ,
  INDEX `trade_species_id` (`trade_species_id` ASC) ,
  CONSTRAINT `pokemon_evolution_ibfk_1`
    FOREIGN KEY (`evolved_species_id` )
    REFERENCES `puserver`.`pokemon_species` (`id` ),
  CONSTRAINT `pokemon_evolution_ibfk_2`
    FOREIGN KEY (`evolution_trigger_id` )
    REFERENCES `puserver`.`evolution_triggers` (`id` ),
  CONSTRAINT `pokemon_evolution_ibfk_3`
    FOREIGN KEY (`trigger_item_id` )
    REFERENCES `puserver`.`items` (`id` ),
  CONSTRAINT `pokemon_evolution_ibfk_5`
    FOREIGN KEY (`held_item_id` )
    REFERENCES `puserver`.`items` (`id` ),
  CONSTRAINT `pokemon_evolution_ibfk_6`
    FOREIGN KEY (`known_move_id` )
    REFERENCES `puserver`.`moves` (`id` ),
  CONSTRAINT `pokemon_evolution_ibfk_7`
    FOREIGN KEY (`party_species_id` )
    REFERENCES `puserver`.`pokemon_species` (`id` ),
  CONSTRAINT `pokemon_evolution_ibfk_8`
    FOREIGN KEY (`trade_species_id` )
    REFERENCES `puserver`.`pokemon_species` (`id` ))
ENGINE = InnoDB
AUTO_INCREMENT = 326
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_forms` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `form_identifier` VARCHAR(16) NULL DEFAULT NULL ,
  `pokemon_id` INT(11) NOT NULL ,
  `is_default` TINYINT(1) NOT NULL ,
  `is_battle_only` TINYINT(1) NOT NULL ,
  `order` INT(11) NOT NULL ,
  PRIMARY KEY (`id`) ,
  INDEX `pokemon_id` (`pokemon_id` ASC) ,
  CONSTRAINT `pokemon_forms_ibfk_1`
    FOREIGN KEY (`pokemon_id` )
    REFERENCES `puserver`.`pokemon` (`id` ))
ENGINE = InnoDB
AUTO_INCREMENT = 728
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_moves` (
  `pokemon_id` INT(11) NOT NULL ,
  `version_group_id` INT(11) NOT NULL ,
  `move_id` INT(11) NOT NULL ,
  `pokemon_move_method_id` INT(11) NOT NULL ,
  `level` INT(11) NOT NULL ,
  `order` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`pokemon_id`, `version_group_id`, `move_id`, `pokemon_move_method_id`, `level`) ,
  INDEX `idx_autoinc_level` (`level` ASC) ,
  INDEX `ix_pokemon_moves_version_group_id` (`version_group_id` ASC) ,
  INDEX `ix_pokemon_moves_level` (`level` ASC) ,
  INDEX `ix_pokemon_moves_pokemon_id` (`pokemon_id` ASC) ,
  INDEX `ix_pokemon_moves_pokemon_move_method_id` (`pokemon_move_method_id` ASC) ,
  INDEX `ix_pokemon_moves_move_id` (`move_id` ASC) ,
  CONSTRAINT `pokemon_moves_ibfk_1`
    FOREIGN KEY (`pokemon_id` )
    REFERENCES `puserver`.`pokemon` (`id` ),
  CONSTRAINT `pokemon_moves_ibfk_3`
    FOREIGN KEY (`move_id` )
    REFERENCES `puserver`.`moves` (`id` ),
  CONSTRAINT `pokemon_moves_ibfk_4`
    FOREIGN KEY (`pokemon_move_method_id` )
    REFERENCES `puserver`.`pokemon_move_methods` (`id` ))
ENGINE = InnoDB
AUTO_INCREMENT = 150209
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

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
  PRIMARY KEY (`id`) ,
  INDEX `evolves_from_species_id` (`evolves_from_species_id` ASC) ,
  INDEX `evolution_chain_id` (`evolution_chain_id` ASC) ,
  INDEX `growth_rate_id` (`growth_rate_id` ASC) ,
  CONSTRAINT `pokemon_species_ibfk_2`
    FOREIGN KEY (`evolves_from_species_id` )
    REFERENCES `puserver`.`pokemon_species` (`id` ),
  CONSTRAINT `pokemon_species_ibfk_3`
    FOREIGN KEY (`evolution_chain_id` )
    REFERENCES `puserver`.`evolution_chains` (`id` ))
ENGINE = InnoDB
AUTO_INCREMENT = 650
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_stats` (
  `pokemon_id` INT(11) NOT NULL ,
  `stat_id` INT(11) NOT NULL ,
  `base_stat` INT(11) NOT NULL ,
  `effort` INT(11) NOT NULL ,
  PRIMARY KEY (`pokemon_id`, `stat_id`) ,
  INDEX `stat_id` (`stat_id` ASC) ,
  CONSTRAINT `pokemon_stats_ibfk_1`
    FOREIGN KEY (`pokemon_id` )
    REFERENCES `puserver`.`pokemon` (`id` ),
  CONSTRAINT `pokemon_stats_ibfk_2`
    FOREIGN KEY (`stat_id` )
    REFERENCES `puserver`.`stats` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_types` (
  `pokemon_id` INT(11) NOT NULL ,
  `type_id` INT(11) NOT NULL ,
  `slot` INT(11) NOT NULL ,
  PRIMARY KEY (`pokemon_id`, `slot`) ,
  INDEX `type_id` (`type_id` ASC) ,
  CONSTRAINT `pokemon_types_ibfk_1`
    FOREIGN KEY (`pokemon_id` )
    REFERENCES `puserver`.`pokemon` (`id` ),
  CONSTRAINT `pokemon_types_ibfk_2`
    FOREIGN KEY (`type_id` )
    REFERENCES `puserver`.`types` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`stats` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `damage_class_id` INT(11) NULL DEFAULT NULL ,
  `identifier` VARCHAR(16) NOT NULL ,
  `is_battle_only` TINYINT(1) NOT NULL ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 9
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`teleport` (
  `idteleport` INT(11) NOT NULL AUTO_INCREMENT ,
  `x` INT(11) NULL DEFAULT NULL ,
  `y` INT(11) NULL DEFAULT NULL ,
  `z` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idteleport`) )
ENGINE = InnoDB
AUTO_INCREMENT = 223
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

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
  INDEX `fk_tile_location` (`idlocation` ASC) ,
  INDEX `position_key` (`x` ASC, `y` ASC, `z` ASC) ,
  INDEX `fk_tile_map` (`z` ASC) ,
  INDEX `fk_tile_teleport` (`idteleport` ASC) ,
  CONSTRAINT `fk_tile_location1`
    FOREIGN KEY (`idlocation` )
    REFERENCES `puserver`.`location` (`idlocation` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tile_map1`
    FOREIGN KEY (`z` )
    REFERENCES `puserver`.`map` (`idmap` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tile_teleport1`
    FOREIGN KEY (`idteleport` )
    REFERENCES `puserver`.`teleport` (`idteleport` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 1129442
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`tile_layer` (
  `idtile_layer` INT(11) NOT NULL AUTO_INCREMENT ,
  `idtile` INT(11) NOT NULL ,
  `sprite` INT(11) NULL DEFAULT NULL ,
  `layer` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idtile_layer`, `idtile`) ,
  INDEX `fk_tile_layer_tileid` (`idtile` ASC) ,
  CONSTRAINT `fk_tile_layer_tileid`
    FOREIGN KEY (`idtile` )
    REFERENCES `puserver`.`tile` (`idtile` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 1541022
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`types` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `identifier` VARCHAR(12) NOT NULL ,
  `generation_id` INT(11) NOT NULL ,
  `damage_class_id` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 10003
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`npc` (
  `idnpc` INT(11) NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(45) NULL DEFAULT NULL ,
  `script_name` VARCHAR(45) NULL DEFAULT NULL ,
  `position` BIGINT(20) NULL DEFAULT NULL ,
  `idmap` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idnpc`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_prose` (
  `item_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `short_effect` VARCHAR(256) NULL DEFAULT NULL ,
  `effect` VARCHAR(5120) NULL DEFAULT NULL ,
  PRIMARY KEY (`item_id`, `local_language_id`) ,
  INDEX `local_language_id` (`local_language_id` ASC) ,
  CONSTRAINT `item_prose_ibfk_1`
    FOREIGN KEY (`item_id` )
    REFERENCES `puserver`.`items` (`id` ),
  CONSTRAINT `item_prose_ibfk_2`
    FOREIGN KEY (`local_language_id` )
    REFERENCES `puserver`.`languages` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_pockets` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `identifier` VARCHAR(16) NOT NULL ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 9
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_pocket_names` (
  `item_pocket_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `name` VARCHAR(16) NOT NULL ,
  PRIMARY KEY (`item_pocket_id`, `local_language_id`) ,
  INDEX `local_language_id` (`local_language_id` ASC) ,
  INDEX `ix_item_pocket_names_name` (`name` ASC) ,
  CONSTRAINT `item_pocket_names_ibfk_1`
    FOREIGN KEY (`item_pocket_id` )
    REFERENCES `puserver`.`item_pockets` (`id` ),
  CONSTRAINT `item_pocket_names_ibfk_2`
    FOREIGN KEY (`local_language_id` )
    REFERENCES `puserver`.`languages` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_names` (
  `item_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `name` VARCHAR(20) NOT NULL ,
  PRIMARY KEY (`item_id`, `local_language_id`) ,
  INDEX `local_language_id` (`local_language_id` ASC) ,
  INDEX `ix_item_names_name` (`name` ASC) ,
  CONSTRAINT `item_names_ibfk_1`
    FOREIGN KEY (`item_id` )
    REFERENCES `puserver`.`items` (`id` ),
  CONSTRAINT `item_names_ibfk_2`
    FOREIGN KEY (`local_language_id` )
    REFERENCES `puserver`.`languages` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_game_indices` (
  `item_id` INT(11) NOT NULL ,
  `generation_id` INT(11) NOT NULL ,
  `game_index` INT(11) NOT NULL ,
  PRIMARY KEY (`item_id`, `generation_id`) ,
  INDEX `generation_id` (`generation_id` ASC) ,
  CONSTRAINT `item_game_indices_ibfk_1`
    FOREIGN KEY (`item_id` )
    REFERENCES `puserver`.`items` (`id` ),
  CONSTRAINT `item_game_indices_ibfk_2`
    FOREIGN KEY (`generation_id` )
    REFERENCES `puserver`.`generations` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_fling_effects` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 8
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_fling_effect_prose` (
  `item_fling_effect_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `effect` VARCHAR(255) NOT NULL ,
  PRIMARY KEY (`item_fling_effect_id`, `local_language_id`) ,
  INDEX `local_language_id` (`local_language_id` ASC) ,
  CONSTRAINT `item_fling_effect_prose_ibfk_1`
    FOREIGN KEY (`item_fling_effect_id` )
    REFERENCES `puserver`.`item_fling_effects` (`id` ),
  CONSTRAINT `item_fling_effect_prose_ibfk_2`
    FOREIGN KEY (`local_language_id` )
    REFERENCES `puserver`.`languages` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_flavor_text` (
  `item_id` INT(11) NOT NULL ,
  `version_group_id` INT(11) NOT NULL ,
  `language_id` INT(11) NOT NULL ,
  `flavor_text` VARCHAR(255) NOT NULL ,
  PRIMARY KEY (`item_id`, `version_group_id`, `language_id`) ,
  INDEX `version_group_id` (`version_group_id` ASC) ,
  INDEX `language_id` (`language_id` ASC) ,
  CONSTRAINT `item_flavor_text_ibfk_1`
    FOREIGN KEY (`item_id` )
    REFERENCES `puserver`.`items` (`id` ),
  CONSTRAINT `item_flavor_text_ibfk_2`
    FOREIGN KEY (`version_group_id` )
    REFERENCES `puserver`.`version_groups` (`id` ),
  CONSTRAINT `item_flavor_text_ibfk_3`
    FOREIGN KEY (`language_id` )
    REFERENCES `puserver`.`languages` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_flavor_summaries` (
  `item_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `flavor_summary` VARCHAR(512) NULL DEFAULT NULL ,
  PRIMARY KEY (`item_id`, `local_language_id`) ,
  INDEX `local_language_id` (`local_language_id` ASC) ,
  CONSTRAINT `item_flavor_summaries_ibfk_1`
    FOREIGN KEY (`item_id` )
    REFERENCES `puserver`.`items` (`id` ),
  CONSTRAINT `item_flavor_summaries_ibfk_2`
    FOREIGN KEY (`local_language_id` )
    REFERENCES `puserver`.`languages` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_flags` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `identifier` VARCHAR(24) NOT NULL ,
  PRIMARY KEY (`id`) )
ENGINE = InnoDB
AUTO_INCREMENT = 9
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_flag_prose` (
  `item_flag_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `name` VARCHAR(24) NULL DEFAULT NULL ,
  `description` VARCHAR(64) NULL DEFAULT NULL ,
  PRIMARY KEY (`item_flag_id`, `local_language_id`) ,
  INDEX `local_language_id` (`local_language_id` ASC) ,
  INDEX `ix_item_flag_prose_name` (`name` ASC) ,
  CONSTRAINT `item_flag_prose_ibfk_1`
    FOREIGN KEY (`item_flag_id` )
    REFERENCES `puserver`.`item_flags` (`id` ),
  CONSTRAINT `item_flag_prose_ibfk_2`
    FOREIGN KEY (`local_language_id` )
    REFERENCES `puserver`.`languages` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_flag_map` (
  `item_id` INT(11) NOT NULL ,
  `item_flag_id` INT(11) NOT NULL ,
  PRIMARY KEY (`item_id`, `item_flag_id`) ,
  INDEX `item_flag_id` (`item_flag_id` ASC) ,
  CONSTRAINT `item_flag_map_ibfk_1`
    FOREIGN KEY (`item_id` )
    REFERENCES `puserver`.`items` (`id` ),
  CONSTRAINT `item_flag_map_ibfk_2`
    FOREIGN KEY (`item_flag_id` )
    REFERENCES `puserver`.`item_flags` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_category_prose` (
  `item_category_id` INT(11) NOT NULL ,
  `local_language_id` INT(11) NOT NULL ,
  `name` VARCHAR(16) NOT NULL ,
  PRIMARY KEY (`item_category_id`, `local_language_id`) ,
  INDEX `local_language_id` (`local_language_id` ASC) ,
  INDEX `ix_item_category_prose_name` (`name` ASC) ,
  CONSTRAINT `item_category_prose_ibfk_1`
    FOREIGN KEY (`item_category_id` )
    REFERENCES `puserver`.`item_categories` (`id` ),
  CONSTRAINT `item_category_prose_ibfk_2`
    FOREIGN KEY (`local_language_id` )
    REFERENCES `puserver`.`languages` (`id` ))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`item_categories` (
  `id` INT(11) NOT NULL AUTO_INCREMENT ,
  `pocket_id` INT(11) NOT NULL ,
  `identifier` VARCHAR(16) NOT NULL ,
  PRIMARY KEY (`id`) ,
  INDEX `pocket_id` (`pocket_id` ASC) ,
  CONSTRAINT `item_categories_ibfk_1`
    FOREIGN KEY (`pocket_id` )
    REFERENCES `puserver`.`item_pockets` (`id` ))
ENGINE = InnoDB
AUTO_INCREMENT = 45
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`ability_messages` (
  `idability_messages` INT(11) NOT NULL AUTO_INCREMENT ,
  `ability_id` INT(11) NOT NULL ,
  `message` VARCHAR(255) NULL DEFAULT NULL ,
  PRIMARY KEY (`idability_messages`) ,
  INDEX `fk_ability_messages_abilities1` (`ability_id` ASC) ,
  CONSTRAINT `fk_ability_messages_abilities1`
    FOREIGN KEY (`ability_id` )
    REFERENCES `puserver`.`abilities` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
AUTO_INCREMENT = 60
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_items` (
  `idplayer_items` INT(11) NOT NULL ,
  `idplayer` INT(11) NULL DEFAULT NULL ,
  `iditem` INT(11) NULL DEFAULT NULL ,
  `amount` INT(11) NULL DEFAULT NULL ,
  `is_bound` BIT(1) NULL DEFAULT 0 ,
  PRIMARY KEY (`idplayer_items`) ,
  INDEX `fk_player_items_player1` (`idplayer` ASC) ,
  INDEX `fk_player_items_items1` (`iditem` ASC) ,
  CONSTRAINT `fk_player_items_player1`
    FOREIGN KEY (`idplayer` )
    REFERENCES `puserver`.`player` (`idplayer` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_items_items1`
    FOREIGN KEY (`iditem` )
    REFERENCES `puserver`.`items` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`quests` (
  `idquests` INT(11) NOT NULL ,
  `name` VARCHAR(45) NULL DEFAULT NULL ,
  PRIMARY KEY (`idquests`) )
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`player_quests` (
  `idplayer_quests` INT(11) NOT NULL ,
  `idplayer` INT(11) NULL DEFAULT NULL ,
  `idquest` INT(11) NULL DEFAULT NULL ,
  `status` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idplayer_quests`) ,
  INDEX `fk_player_quests_player1` (`idplayer` ASC) ,
  INDEX `fk_player_quests_quests1` (`idquest` ASC) ,
  CONSTRAINT `fk_player_quests_player1`
    FOREIGN KEY (`idplayer` )
    REFERENCES `puserver`.`player` (`idplayer` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_quests_quests1`
    FOREIGN KEY (`idquest` )
    REFERENCES `puserver`.`quests` (`idquests` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

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
  PRIMARY KEY (`idnpc_pokemon`) ,
  INDEX `fk_npc_pokemon_npc1` (`idnpc` ASC) ,
  INDEX `fk_npc_pokemon_pokemon1` (`idpokemon` ASC) ,
  INDEX `fk_npc_pokemon_items1` (`held_item` ASC) ,
  CONSTRAINT `fk_npc_pokemon_npc1`
    FOREIGN KEY (`idnpc` )
    REFERENCES `puserver`.`npc` (`idnpc` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_npc_pokemon_pokemon1`
    FOREIGN KEY (`idpokemon` )
    REFERENCES `puserver`.`pokemon` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_npc_pokemon_items1`
    FOREIGN KEY (`held_item` )
    REFERENCES `puserver`.`items` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;

CREATE  TABLE IF NOT EXISTS `puserver`.`npc_pokemon_move` (
  `idnpc_pokemon_move` INT(11) NOT NULL ,
  `idnpc_pokemon` INT(11) NULL DEFAULT NULL ,
  `idmove` INT(11) NULL DEFAULT NULL ,
  PRIMARY KEY (`idnpc_pokemon_move`) ,
  INDEX `fk_npc_pokemon_move_npc_pokemon1` (`idnpc_pokemon` ASC) ,
  INDEX `fk_npc_pokemon_move_moves1` (`idmove` ASC) ,
  CONSTRAINT `fk_npc_pokemon_move_npc_pokemon1`
    FOREIGN KEY (`idnpc_pokemon` )
    REFERENCES `puserver`.`npc_pokemon` (`idnpc_pokemon` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_npc_pokemon_move_moves1`
    FOREIGN KEY (`idmove` )
    REFERENCES `puserver`.`moves` (`id` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
