<?php
require_once('../db_config.php');
header('Content-Type: application/json');


if (!isset($_GET['email'])) {
    echo json_encode(array("success" => false, "exist" => false, "error" => "Email parameter not provided"));
    exit;
}

$userEmail = $_GET['email'];
$sql = "SELECT * FROM `user_table` WHERE `email` = ?";
$stmt = $connDB->prepare($sql);

if ($stmt) {
    $stmt->bind_param("s", $userEmail);
    if ($stmt->execute()) {
        $result = $stmt->get_result();
        
        if ($result->num_rows > 0) {
            // User exists, fetch data
            $row = $result->fetch_assoc();
            $userName = $row['name'];
            $userIdentity = $row['identity']; 
	    $userId = $row['id'];
	    $userBirthday = $row['birthday'];
	    $userGender = $row['gender'];
	    $arr = array("success" => true, "exist" => true, "id" => $userId, "email" => $userEmail, "name" => $userName, "identity" => $userIdentity, "birthday" => $userBirthday, "gender" => $userGender);
	    echo json_encode($arr);
	    error_log("success => true, exist => true, id => $userId, email => $userEmail, name => $userName, identity => $userIdentity, birthday => $userBirthday, gender => $userGender", 3, "/var/tmp/php_errors.log");
        } else {
            // User does not exist
            echo json_encode(array(
                "success" => true,
                "exist" => false
            ));
	    error_log("success => true, exist => false, id => $userId, email => $userEmail, name => $userName, identity => $userIdentity", 3, "/var/tmp/php_errors.log");
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
