<?php
require_once 'db_config.php';

function createUser($email, $name, $identity)
{
    global $connDB;
    $email = mysqli_real_escape_string($connDB, $email);
    $name = mysqli_real_escape_string($connDB, $name);
    $identity = mysqli_real_escape_string($connDB, $identity);

    $sql = "INSERT INTO `users` (`email`, `name`, `identity`) VALUES ('$email', '$name', '$identity')";
    if ($connDB->query($sql) === TRUE) {
        $userId = $connDB->insert_id;
        return $userId;
    } else {
        error_log("Error creating user: " . $connDB->error);
        return null;
    }
}
function getAllUser($email)
{
    $dbdata = array();
    global $connDB;
    $email = mysqli_real_escape_string($connDB, $email);

    $sql = "SELECT * FROM users WHERE email = '$email'";
    $result = $connDB->query($sql);

    if ($result->num_rows > 0) {
        while ($row = $result->fetch_assoc()) {
            $dbdata[] = $row;
        }
        return $dbdata;
    } else {
        return null;
    }
}

function getUser($email)
{
    global $connDB;
    $email = mysqli_real_escape_string($connDB, $email);

    $sql = "SELECT * FROM users WHERE email = '$email' LIMIT 1";
    $result = $connDB->query($sql);

    if ($result->num_rows > 0) {
        $row = $result->fetch_assoc();
        return $row;
    } else {
        return null;
    }
}

function updateUser($userId, $newEmail)
{
    global $connDB;
    $userId = mysqli_real_escape_string($connDB, $userId);
    $newEmail = mysqli_real_escape_string($connDB, $newEmail);

    $sql = "UPDATE users SET email = '$newEmail' WHERE id = $userId";
    if ($connDB->query($sql) === TRUE) {
        return true;
    } else {
        return false;
    }
}

function deleteUser($userId)
{
    global $connDB;
    $userId = mysqli_real_escape_string($connDB, $userId);

    $sql = "DELETE FROM users WHERE id = $userId";
    if ($connDB->query($sql) === TRUE) {
        return true;
    } else {
        return false;
    }
}

function getAllRecords($userId)
{
    global $connDB;
    $userId = mysqli_real_escape_string($connDB, $userId);

    $sql = "SELECT * FROM records WHERE user_id = $userId";
    $result = $connDB->query($sql);

    $records = array();
    while ($row = $result->fetch_assoc()) {
        $records[] = $row;
    }

    return $records;
}

function createRecord($userId, $playTimestamp, $gameTime)
{
    global $connDB;
    $userId = mysqli_real_escape_string($connDB, $userId);
    $playTimestamp = mysqli_real_escape_string($connDB, $playTimestamp);
    $gameTime = mysqli_real_escape_string($connDB, $gameTime);

    $sql = "INSERT INTO records (user_id, play_timestamp, game_time) VALUES ($userId, '$playTimestamp', '$gameTime')";
    if ($connDB->query($sql) === TRUE) {
        $recordId = $connDB->insert_id;
        return $recordId;
    } else {
        return null;
    }
}

function getRecord($recordId)
{
    global $connDB;
    $recordId = mysqli_real_escape_string($connDB, $recordId);

    $sql = "SELECT * FROM records WHERE record_id = $recordId LIMIT 1";
    $result = $connDB->query($sql);

    if ($result->num_rows > 0) {
        $row = $result->fetch_assoc();
        return $row;
    } else {
        return null;
    }
}

function updateRecord($recordId, $newPlayTimestamp, $newGameTime)
{
    global $connDB;
    $recordId = mysqli_real_escape_string($connDB, $recordId);
    $newPlayTimestamp = mysqli_real_escape_string($connDB, $newPlayTimestamp);
    $newGameTime = mysqli_real_escape_string($connDB, $newGameTime);

    $sql = "UPDATE records SET play_timestamp = '$newPlayTimestamp', game_time = '$newGameTime' WHERE record_id = $recordId";
    if ($connDB->query($sql) === TRUE) {
        return true;
    } else {
        return false;
    }
}

function deleteRecord($recordId)
{
    global $connDB;
    $recordId = mysqli_real_escape_string($connDB, $recordId);

    $sql = "DELETE FROM records WHERE record_id = $recordId";
    if ($connDB->query($sql) === TRUE) {
        return true;
    } else {
        return false;
    }
}
?>
