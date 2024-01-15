<?php
	require_once('../db_config.php');
	$data = json_decode(file_get_contents("php://input"), true);
	$userEmail = $data['email'];
	$userName = $data['name'];
	$userIdentity = $data['identity'];

	$sql = "INSERT INTO `user_table` (`email`, `name`, `identity`) VALUES (?, ?, ?)";
	$stmt = $connDB->prepare($sql);
	$stmt->bind_param("sss", $userEmail, $userName, $userIdentity);

	if ($stmt->execute()) {
        	error_log("sql create user success",3,"/var/tmp/php_errors.log");
    		echo json_encode(array("success" => true, "exist" => false, "id" => $stmt->insert_id, "email" => $userEmail, "name" => $userName, "identity" => $userIdentity));
	} else {
        	error_log("sql create user fail",3,"/var/tmp/php_errors.log");
    		echo json_encode(array("success" => false, "exist" => true));
	}

	$stmt->close();
?>

