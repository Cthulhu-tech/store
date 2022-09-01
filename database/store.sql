-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Хост: localhost:3306
-- Время создания: Сен 01 2022 г., 17:36
-- Версия сервера: 8.0.30-0ubuntu0.20.04.2
-- Версия PHP: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `store`
--

DELIMITER $$
--
-- Процедуры
--
CREATE DEFINER=`thrackerzod`@`localhost` PROCEDURE `sp_email` (IN `email_value` CHAR(255), IN `url_value` CHAR(255), IN `rand_value` INT(4))  NO SQL
BEGIN

SET @USER_FIND = (SELECT case when COUNT(user.id) = 1 
      THEN 'FIND' 
      ELSE 'OK'
END USER_FIND 
      FROM user 
      WHERE 
      user.email = email_value
     );

IF (@USER_FIND = 'OK') THEN

    INSERT INTO user (email) VALUE (email_value);

    INSERT INTO url_ckeck 
    (user_id, url_ckeck, secret_code)
        VALUE (
            (SELECT LAST_INSERT_ID()),
            url_value, rand_value
        );

ELSE

    SET @URL_INFO = (SELECT case when COUNT(url_ckeck.user_id) = 1 
          THEN 'FIND' 
          ELSE 'NOTFOUND'
    END USER_FIND 
          FROM url_ckeck 
          WHERE 
          url_ckeck.user_id = (
              SELECT user.id 
              FROM user 
              WHERE user.email = email_value)
         );
         
	IF (@URL_INFO = 'NOTFOUND') THEN
    	
        INSERT INTO url_ckeck
        (user_id, url_ckeck, secret_code)
        VALUES
        (
            (SELECT user.id 
              FROM user 
              WHERE user.email = email_value),
            url_value,
            rand_value
        );
        
    ELSE
    
    	UPDATE url_ckeck SET 
        url_ckeck = url_value,
        secret_code = rand_value
        WHERE 
        url_ckeck.user_id = (
              SELECT user.id 
              FROM user 
              WHERE user.email = email_value);
    
    END IF;

END IF;
    
END$$

--
-- Функции
--
CREATE DEFINER=`thrackerzod`@`localhost` FUNCTION `sp_cofnfirm_vk` (`user_login_value` CHAR(55), `user_password_value` CHAR(255), `user_code_value` CHAR(255), `user_secret_value` TINYINT(4)) RETURNS INT NO SQL
BEGIN


SET @USER_FIND = (SELECT case when COUNT(url_ckeck.user_id) = 1 
      THEN 1
      ELSE 0
END USER_FIND 
      FROM url_ckeck 
      WHERE 
      url_ckeck.url_ckeck = user_code_value
      AND
      url_ckeck.secret_code = user_secret_value
     );

IF (@USER_FIND = 1) THEN

    UPDATE 
        user
    SET 
    	user.confirme = 1,
        login = user_login_value,
        password = user_password_value
    WHERE
        user.id = (SELECT url_ckeck.user_id
                  FROM url_ckeck 
                  WHERE 
                   url_ckeck.url_ckeck = user_code_value
                   AND 
                   url_ckeck.secret_code = user_secret_value
                  );
                  
    DELETE FROM url_ckeck WHERE url_ckeck = user_code_value;
    
    RETURN 1;

ELSE

    RETURN 0;

END IF;
        
END$$

CREATE DEFINER=`thrackerzod`@`localhost` FUNCTION `sp_confirm` (`user_code_value` CHAR(255), `user_secret_value` TINYINT(4)) RETURNS INT NO SQL
BEGIN


SET @USER_FIND = (SELECT case when COUNT(url_ckeck.user_id) = 1 
      THEN 1
      ELSE 0
END USER_FIND 
      FROM url_ckeck 
      WHERE 
      url_ckeck.url_ckeck = user_code_value
      AND
      url_ckeck.secret_code = user_secret_value
     );

IF (@USER_FIND = 1) THEN

    UPDATE 
        user
    SET 
        user.confirme = 1
    WHERE
        user.id = (SELECT url_ckeck.user_id
                  FROM url_ckeck 
                  WHERE 
                   url_ckeck.url_ckeck = user_code_value
                   AND 
                   url_ckeck.secret_code = user_secret_value
                  );
                  
    DELETE FROM url_ckeck WHERE url_ckeck = user_code_value;
    
    RETURN 1;

ELSE

    RETURN 0;

END IF;
        
END$$

CREATE DEFINER=`thrackerzod`@`localhost` FUNCTION `sp_registation` (`user_login_value` CHAR(55), `user_password_value` CHAR(255), `user_email_value` CHAR(255)) RETURNS INT NO SQL
BEGIN

	 	SET @USER_FIND = (SELECT case when COUNT(user.id) = 1 
                      THEN 'FIND' 
                      ELSE 'OK'
                END USER_FIND 
                      FROM user 
                      WHERE 
                      user.email = user_email_value 
                      OR 
                      user.login = user_login_value
                     );
        
    IF (@USER_FIND = 'OK') THEN

        INSERT INTO user 
        (login, email, password)
        VALUES
        (user_login_value, user_email_value, user_password_value);
        
        RETURN 1;
        
    ELSE
    	
        RETURN 0;
        
    END IF;
        
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Структура таблицы `token`
--

CREATE TABLE `token` (
  `user_id` int NOT NULL,
  `jwt` char(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `url_ckeck`
--

CREATE TABLE `url_ckeck` (
  `user_id` int NOT NULL,
  `url_ckeck` char(255) NOT NULL,
  `secret_code` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Структура таблицы `user`
--

CREATE TABLE `user` (
  `id` int NOT NULL,
  `login` char(55) DEFAULT NULL,
  `email` char(255) NOT NULL,
  `password` char(255) DEFAULT NULL,
  `confirme` tinyint(1) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `url_ckeck`
--
ALTER TABLE `url_ckeck`
  ADD KEY `url_ckeck_ibfk_1` (`user_id`);

--
-- Индексы таблицы `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `user`
--
ALTER TABLE `user`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=39;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `url_ckeck`
--
ALTER TABLE `url_ckeck`
  ADD CONSTRAINT `url_ckeck_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
