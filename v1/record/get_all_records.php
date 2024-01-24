<?php
require_once('../db_config.php');

if (!isset($_GET['userId'])) {
    echo json_encode(array("success" => false, "message" => "User ID is required"));
    exit;
}

$userId = $_GET['userId'];

$sql = "SELECT * FROM record_table WHERE fk_user_id = ?";
$stmt = $connDB->prepare($sql);

if ($stmt) {
    // Bind the user ID parameter and execute
    $stmt->bind_param("i", $userId);
    if ($stmt->execute()) {
        $result = $stmt->get_result();
        $records = array();

        // Fetch each row and add to records array
        while ($row = $result->fetch_assoc()) {
            $records[] = $row;
        }

        echo json_encode(array("success" => true, "records" => $records));
    } else {
        echo json_encode(array("success" => false, "message" => "Failed to execute query"));
    }

    $stmt->close();
} else {
    echo json_encode(array("success" => false, "message" => "Failed to prepare statement"));
}

$connDB->close();
?>
