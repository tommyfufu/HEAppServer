<?php
require_once('../db_config.php');

$data = json_decode(file_get_contents("php://input"), true);
$userEmail = $data['email'];

$sql = "SELECT * FROM `user_table` WHERE `email` = ?";
$stmt = $connDB->prepare($sql);

if ($stmt) {
    $stmt->bind_param("s", $userEmail);
    if ($stmt->execute()) {
        $result = $stmt->get_result();
        
        if ($result->num_rows > 0) {
            // User exists, fetch data
            $row = $result->fetch_assoc();
	    error_log("row result: $row ", 3, "/var/tmp/php_errors.log");
            $userName = $row['name'];
            $userIdentity = $row['identity']; 
            echo json_encode(array(
                "success" => true,
                "exist" => true,
                "id" => $row['id'],
                "email" => $userEmail,
                "name" => $userName,
                "identity" => $userIdentity
            ));
        } else {
            // User does not exist
            echo json_encode(array(
                "success" => true,
                "exist" => false
            ));
        }
    } else {
        // SQL execution failed
        error_log("SQL execution failed: ", 3, "/var/tmp/php_errors.log");
        echo json_encode(array("success" => false, "exist" => false));
    }

    $stmt->close();
} else {
    // Statement preparation failed
    error_log("Statement preparation failed: ", 3, "/var/tmp/php_errors.log");
    echo json_encode(array("success" => false, "exist" => false));
}

$connDB->close();
?>
