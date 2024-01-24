<?php
require_once('../db_config.php');
header('Content-Type: application/json');

$data = json_decode(file_get_contents("php://input"), true);

if (empty($data['id']) || empty($data['name'])) {
    echo json_encode(array("success" => false, "message" => "ID and name are required"));
    exit;
}

$userId = $data['id'];
$userName = $data['name'];

$sql = "UPDATE `user_table` SET `name` = ? WHERE `id` = ?";
$stmt = $connDB->prepare($sql);

if ($stmt) {
    $stmt->bind_param("si", $userName, $userId);
    if ($stmt->execute()) {
        echo json_encode(array("success" => true, "message" => "User updated successfully"));
    } else {
        echo json_encode(array("success" => false, "message" => "Failed to update user"));
    }

    $stmt->close();
} else {
    echo json_encode(array("success" => false, "message" => "Statement preparation failed"));
}

$connDB->close();
?>
