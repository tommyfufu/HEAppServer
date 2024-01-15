<?php

    $servername = "database";
    $username = "rtes913";
    $password = "MYSQLrtes913"; //docker-compose env
    $dbname = "HEdb";

    $connDB = new mysqli($servername, $username, $password, $dbname);

    if (mysqli_connect_error()) {
        error_log("mysql connect error ",3,"/var/tmp/php_errors.log");
    }

    // Create 'user_table' Table if not exists
    $sql = "CREATE TABLE IF NOT EXISTS `user_table` (
                `id` int unsigned NOT NULL AUTO_INCREMENT,
                `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                `identity` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                PRIMARY KEY (`id`),
                UNIQUE KEY `email_UNIQUE` (`email`),
                UNIQUE KEY `id_UNIQUE` (`id`)
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_bin";

    if ($connDB->query($sql) === TRUE) {
        error_log("create user table success ",3,"/var/tmp/php_errors.log");
    } else {
        error_log("create user table fail ",3,"/var/tmp/php_errors.log");
    }

    // Create 'record_table' Table if not exists
    $sql = "CREATE TABLE IF NOT EXISTS `record_table` (
            `record_id` int unsigned NOT NULL AUTO_INCREMENT,
            `fk_user_id` int unsigned NOT NULL,
            `game_id` int unsigned NOT NULL,
            `game_date_time` datetime(5) NOT NULL DEFAULT CURRENT_TIMESTAMP(5),
            `game_time` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
            PRIMARY KEY (`record_id`),
            UNIQUE KEY `record_id_UNIQUE` (`record_id`),
            KEY `fk_user_id_idx` (`fk_user_id`),
            CONSTRAINT `fk_user_id` FOREIGN KEY (`fk_user_id`) REFERENCES `user_table` (`id`)
        ) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_bin";

    if ($connDB->query($sql) === TRUE) {
        error_log("create record table success ",3,"/var/tmp/php_errors.log");
    } else {
        error_log("create record table success ",3,"/var/tmp/php_errors.log");
    }
    
?>
