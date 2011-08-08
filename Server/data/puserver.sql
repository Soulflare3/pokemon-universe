SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL';

CREATE SCHEMA IF NOT EXISTS `puserver` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci ;
USE `puserver` ;

-- -----------------------------------------------------
-- Table `puserver`.`pokecenter`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`pokecenter` (
  `idpokecenter` INT NOT NULL AUTO_INCREMENT ,
  `position` BIGINT NOT NULL ,
  `description` VARCHAR(250) NULL ,
  PRIMARY KEY (`idpokecenter`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`music`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`music` (
  `idmusic` INT NOT NULL AUTO_INCREMENT ,
  `title` VARCHAR(45) NULL ,
  `filename` VARCHAR(45) NULL ,
  PRIMARY KEY (`idmusic`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`location`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`location` (
  `idlocation` INT NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(45) NULL ,
  `idpokecenter` INT NOT NULL ,
  `idmusic` INT NOT NULL ,
  PRIMARY KEY (`idlocation`) ,
  INDEX `fk_location_pokecenter1` (`idpokecenter` ASC) ,
  INDEX `fk_location_music1` (`idmusic` ASC) ,
  CONSTRAINT `fk_location_pokecenter1`
    FOREIGN KEY (`idpokecenter` )
    REFERENCES `puserver`.`pokecenter` (`idpokecenter` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_location_music1`
    FOREIGN KEY (`idmusic` )
    REFERENCES `puserver`.`music` (`idmusic` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`player`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`player` (
  `idplayer` INT NOT NULL AUTO_INCREMENT ,
  `idaccount` INT NULL ,
  `name` VARCHAR(20) NULL ,
  `password` VARCHAR(45) NULL ,
  `password_salt` VARCHAR(45) NULL ,
  `position` BIGINT NULL COMMENT 'x;y;z' ,
  `movement` SMALLINT NULL ,
  `idpokecenter` INT NULL ,
  `money` INT NULL ,
  `idlocation` INT NOT NULL ,
  PRIMARY KEY (`idplayer`) ,
  INDEX `fk_player_location1` (`idlocation` ASC) ,
  CONSTRAINT `fk_player_location1`
    FOREIGN KEY (`idlocation` )
    REFERENCES `puserver`.`location` (`idlocation` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`pokemon`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon` (
  `idpokemon` INT NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(45) NULL ,
  `idevolution_chain` INT NULL ,
  PRIMARY KEY (`idpokemon`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`player_pokemon`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`player_pokemon` (
  `idplayer_pokemon` INT NOT NULL AUTO_INCREMENT ,
  `idpokemon` INT NOT NULL ,
  `idplayer` INT NOT NULL ,
  `nickname` VARCHAR(45) NULL ,
  `bound` TINYINT NULL DEFAULT 0 COMMENT '1 if pokemon is bound to player' ,
  `experience` INT NULL ,
  `iv_hp` TINYINT NULL ,
  `iv_attack` TINYINT NULL ,
  `iv_attack_spec` TINYINT NULL ,
  `iv_defence` TINYINT NULL ,
  `iv_defence_spec` TINYINT NULL ,
  `iv_speed` TINYINT NULL ,
  `happiness` TINYINT NULL ,
  `gender` TINYINT NULL COMMENT '-1 None\n0 Male\n1 Female' ,
  PRIMARY KEY (`idplayer_pokemon`) ,
  INDEX `fk_player_pokemon_player` (`idplayer` ASC) ,
  INDEX `fk_player_pokemon_pokemon` (`idpokemon` ASC) ,
  CONSTRAINT `fk_player_pokemon_player`
    FOREIGN KEY (`idplayer` )
    REFERENCES `puserver`.`player` (`idplayer` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_pokemon_pokemon`
    FOREIGN KEY (`idpokemon` )
    REFERENCES `puserver`.`pokemon` (`idpokemon` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`location_section`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`location_section` (
  `idlocation_section` INT NOT NULL AUTO_INCREMENT ,
  `idlocation` INT NOT NULL ,
  `name` VARCHAR(250) NULL ,
  PRIMARY KEY (`idlocation_section`) ,
  INDEX `fk_location_sections_location1` (`idlocation` ASC) ,
  CONSTRAINT `fk_location_sections_location1`
    FOREIGN KEY (`idlocation` )
    REFERENCES `puserver`.`location` (`idlocation` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`move`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`move` (
  `idmove` INT NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(45) NULL ,
  PRIMARY KEY (`idmove`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`move_method`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`move_method` (
  `idmove_method` INT NOT NULL AUTO_INCREMENT ,
  `name` VARCHAR(45) NULL ,
  `description` VARCHAR(250) NULL ,
  PRIMARY KEY (`idmove_method`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`pokemon_move`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`pokemon_move` (
  `idpokemon_move` INT NOT NULL AUTO_INCREMENT ,
  `idpokemon` INT NOT NULL ,
  `idmove` INT NOT NULL ,
  `idmove_method` INT NULL ,
  `level` INT NULL ,
  `order` INT NULL ,
  PRIMARY KEY (`idpokemon_move`) ,
  INDEX `fk_pokemon_move_move_method1` (`idmove_method` ASC) ,
  INDEX `fk_pokemon_move_move1` (`idmove` ASC) ,
  INDEX `fk_pokemon_move_pokemon1` (`idpokemon` ASC) ,
  CONSTRAINT `fk_pokemon_move_move_method1`
    FOREIGN KEY (`idmove_method` )
    REFERENCES `puserver`.`move_method` (`idmove_method` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_pokemon_move_move1`
    FOREIGN KEY (`idmove` )
    REFERENCES `puserver`.`move` (`idmove` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_pokemon_move_pokemon1`
    FOREIGN KEY (`idpokemon` )
    REFERENCES `puserver`.`pokemon` (`idpokemon` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`player_pokemon_move`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`player_pokemon_move` (
  `idplayer_pokemon_move` INT NOT NULL ,
  `idplayer_pokemon` INT NOT NULL ,
  `idmove` INT NOT NULL ,
  PRIMARY KEY (`idplayer_pokemon_move`) ,
  INDEX `fk_player_pokemon_move_player_pokemon1` (`idplayer_pokemon` ASC) ,
  INDEX `fk_player_pokemon_move_move1` (`idmove` ASC) ,
  CONSTRAINT `fk_player_pokemon_move_player_pokemon1`
    FOREIGN KEY (`idplayer_pokemon` )
    REFERENCES `puserver`.`player_pokemon` (`idplayer_pokemon` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_pokemon_move_move1`
    FOREIGN KEY (`idmove` )
    REFERENCES `puserver`.`move` (`idmove` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`encounter_condition`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`encounter_condition` (
  `idencounter_condition` INT NOT NULL ,
  `name` VARCHAR(250) NULL ,
  `default` INT NULL ,
  PRIMARY KEY (`idencounter_condition`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`encounter`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`encounter` (
  `idencounter` INT NOT NULL AUTO_INCREMENT ,
  `idlocation_section` INT NOT NULL ,
  `idencounter_condition` INT NOT NULL ,
  `rate` INT NULL ,
  PRIMARY KEY (`idencounter`) ,
  INDEX `fk_encounter_encounter_condition1` (`idencounter_condition` ASC) ,
  CONSTRAINT `fk_encounter_encounter_condition1`
    FOREIGN KEY (`idencounter_condition` )
    REFERENCES `puserver`.`encounter_condition` (`idencounter_condition` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`encounter_slot`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`encounter_slot` (
  `idencounter_slot` INT NOT NULL ,
  `idencounter` INT NOT NULL ,
  `idpokemon` INT NOT NULL ,
  `gender_rate` INT NULL ,
  PRIMARY KEY (`idencounter_slot`) ,
  INDEX `fk_encounter_slot_encounter1` (`idencounter` ASC) ,
  INDEX `fk_encounter_slot_pokemon1` (`idpokemon` ASC) ,
  CONSTRAINT `fk_encounter_slot_encounter1`
    FOREIGN KEY (`idencounter` )
    REFERENCES `puserver`.`encounter` (`idencounter` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_encounter_slot_pokemon1`
    FOREIGN KEY (`idpokemon` )
    REFERENCES `puserver`.`pokemon` (`idpokemon` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`location_encounter`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`location_encounter` (
  `idencounter` INT NOT NULL ,
  `idlocation_section` INT NOT NULL ,
  PRIMARY KEY (`idencounter`, `idlocation_section`) ,
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
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`player_outfit`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`player_outfit` (
  `idplayer` INT NOT NULL ,
  `head` INT NULL ,
  `nek` INT NULL ,
  `upper` INT NULL ,
  `lower` INT NULL ,
  `feet` INT NULL ,
  INDEX `fk_player_outfit_player1` (`idplayer` ASC) ,
  UNIQUE INDEX `idplayer_UNIQUE` (`idplayer` ASC) ,
  CONSTRAINT `fk_player_outfit_player1`
    FOREIGN KEY (`idplayer` )
    REFERENCES `puserver`.`player` (`idplayer` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`group`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`group` (
  `idgroup` INT NOT NULL ,
  `name` VARCHAR(45) NULL ,
  `flags` INT NULL ,
  `priority` INT NULL ,
  PRIMARY KEY (`idgroup`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`player_group`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`player_group` (
  `player_idplayer` INT NOT NULL ,
  `group_idgroup` INT NOT NULL ,
  PRIMARY KEY (`player_idplayer`, `group_idgroup`) ,
  INDEX `fk_player_has_group_group1` (`group_idgroup` ASC) ,
  CONSTRAINT `fk_player_has_group_player1`
    FOREIGN KEY (`player_idplayer` )
    REFERENCES `puserver`.`player` (`idplayer` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_player_has_group_group1`
    FOREIGN KEY (`group_idgroup` )
    REFERENCES `puserver`.`group` (`idgroup` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`map`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`map` (
  `idmap` INT NOT NULL ,
  `name` VARCHAR(128) NULL ,
  PRIMARY KEY (`idmap`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`teleport`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`teleport` (
  `idteleport` INT NOT NULL ,
  `x` INT NULL ,
  `y` INT NULL ,
  `z` INT NULL ,
  PRIMARY KEY (`idteleport`) )
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`tile`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`tile` (
  `idtile` INT UNSIGNED NOT NULL AUTO_INCREMENT ,
  `x` INT NOT NULL ,
  `y` INT NOT NULL ,
  `z` INT NOT NULL ,
  `idlocation` INT NOT NULL ,
  `idmap` INT NOT NULL ,
  `movement` INT NULL ,
  `script` VARCHAR(128) NULL ,
  `idteleport` INT NOT NULL ,
  INDEX `fk_tile_location1` (`idlocation` ASC) ,
  PRIMARY KEY (`idtile`) ,
  INDEX `position_key` (`x` ASC, `y` ASC, `z` ASC) ,
  INDEX `fk_tile_map1` (`idmap` ASC) ,
  INDEX `fk_tile_teleport1` (`idteleport` ASC) ,
  CONSTRAINT `fk_tile_location1`
    FOREIGN KEY (`idlocation` )
    REFERENCES `puserver`.`location` (`idlocation` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tile_map1`
    FOREIGN KEY (`idmap` )
    REFERENCES `puserver`.`map` (`idmap` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_tile_teleport1`
    FOREIGN KEY (`idteleport` )
    REFERENCES `puserver`.`teleport` (`idteleport` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `puserver`.`tile_layer`
-- -----------------------------------------------------
CREATE  TABLE IF NOT EXISTS `puserver`.`tile_layer` (
  `idtile_layer` INT NOT NULL ,
  `idtile` INT NOT NULL ,
  `sprite` INT NULL ,
  PRIMARY KEY (`idtile_layer`, `idtile`) ,
  INDEX `fk_tile_layer_tile1` (`idtile` ASC) ,
  CONSTRAINT `fk_tile_layer_tile1`
    FOREIGN KEY (`idtile` )
    REFERENCES `puserver`.`tile` (`idtile` )
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;



SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
