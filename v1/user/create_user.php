<?php
require_once('../db_config.php');
$data = json_decode(file_get_contents("php://input"), true);

$userEmail = $data['email'];
$userName = $data['name'];
$userIdentity = $data['identity'];
$userBirthday = $data['birthday'];
$userGender = $data['gender'];

$sql = "INSERT INTO `user_table` (`email`, `name`, `identity`, `birthday`, `gender`) VALUES (?, ?, ?, ?, ?)";
$stmt = $connDB->prepare($sql);

if ($stmt !== false) {
    $stmt->bind_param("sssss", $userEmail, $userName, $userIdentity, $userBirthday, $userGender);

    if ($stmt->execute()) {
        echo json_encode(["success" => true, "exist" => false, "id" => $stmt->insert_id, "email" => $userEmail, "name" => $userName, "identity" => $userIdentity, "birthday" => $userBirthday, "gender" => $userGender]);
    } else {
        echo json_encode(["success" => false, "exist" => true]);
    }
    $stmt->close();
} else {
    echo json_encode(["success" => false, "message" => "Database error"]);
}
?>
