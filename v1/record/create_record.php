<?php
require_once('../db_config.php');

$data = json_decode(file_get_contents("php://input"), true);


// php::empty() is true if var == null or 0
// our gameId, gameTime, and score may be 0
// So do not use.
if ($data['userId'] === null || $data['gameId'] === null || $data['gameTime'] === null || $data['score'] === null) {
    echo json_encode(array("success" => false, "message" => "All fields are required"));
    error_log("empty data, $userId, $gameId, $gameTime, $score", 3, "/var/tmp/php_errors.log");
    exit;
}

$fk_user_id = $data['userId'];
$game_id = $data['gameId'];
$game_time = $data['gameTime'];
$score = $data['score'];
$sql = "INSERT INTO `record_table` (`fk_user_id`, `game_id`, `game_time`, `score`) VALUES (?, ?, ?, ?)";

$stmt = $connDB->prepare($sql);

if ($stmt) {
    $stmt->bind_param("iisi", $fk_user_id, $game_id, $game_time, $score);
    if ($stmt->execute()) {
        $insertedId = $stmt->insert_id;

        // Perform a query to get the newly inserted record, including the game_date_time
        $selectSql = "SELECT game_date_time FROM record_table WHERE record_id = ?";
        $selectStmt = $connDB->prepare($selectSql);

        if ($selectStmt) {
            $selectStmt->bind_param("i", $insertedId);
            if ($selectStmt->execute()) {
                $result = $selectStmt->get_result();
                $row = $result->fetch_assoc();

                $game_date_time = $row['game_date_time'];
                $response = array(
                    "success" => true,
                    "record_id" => $insertedId,
                    "fk_user_id" => $fk_user_id,
                    "game_id" => $game_id,
                    "game_date_time" => $game_date_time,
                    "game_time" => $game_time,
                    "score" => $score
                );

                echo json_encode($response);
            } else {
                // Handle failed select
                echo json_encode(array("success" => false, "message" => "Failed to retrieve record after insert"));
            }
            $selectStmt->close();
        } else {
            // Handle failed statement preparation for select
            echo json_encode(array("success" => false, "message" => "Failed to prepare select statement"));
        }
    } else {
        echo json_encode(array("success" => false, "message" => "Failed to create record"));
    }
    $stmt->close();
} else {
    echo json_encode(array("success" => false, "message" => "Statement preparation failed"));
}

$connDB->close();
?>
