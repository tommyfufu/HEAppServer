<?php

    $servername = "database";
    $username = "rtes913";
    $password = "MYSQLrtes913"; //docker-compose env
    $dbname = "HEdb";

    $connDB = new mysqli($servername, $username, $password, $dbname);

    if ($connDB->connect_error) {
    	error_log("Connection error: " . $connDB->connect_error, 3, "/var/tmp/php_errors.log");
    	echo json_encode(["success" => false, "error" => "Database connection error"]);
    	exit;
    }
    $connDB->set_charset("utf8mb4");

    $sql = "CREATE TABLE IF NOT EXISTS `user_table` (
                `id` int NOT NULL AUTO_INCREMENT,
                `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                `identity` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'user',
		`birthday` DATE NOT NULL DEFAULT '1970-01-01',
                `gender` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
		PRIMARY KEY (`id`),
                UNIQUE KEY `email_UNIQUE` (`email`),
                UNIQUE KEY `id_UNIQUE` (`id`)
            ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci";

    if ($connDB->query($sql) === FALSE) {
    	echo json_encode(array("success" => false, "error" => "mysql create user_table failed"));
    }

    $sql = "CREATE TABLE IF NOT EXISTS `record_table` (
    `record_id` int NOT NULL AUTO_INCREMENT,
    `fk_user_id` int NOT NULL,
    `game_id` int NOT NULL,
    `game_date_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `game_time` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `score` int NOT NULL,
    PRIMARY KEY (`record_id`),
    UNIQUE KEY `record_id_UNIQUE` (`record_id`),
    KEY `fk_user_id_idx` (`fk_user_id`),
    CONSTRAINT `fk_user_id` FOREIGN KEY (`fk_user_id`) REFERENCES `user_table` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci";

    if ($connDB->query($sql) === FALSE) {
    	echo json_encode(array("success" => false, "error" => "mysql create record_table failed"));
    }

?>
