<?php
ini_set('display_errors', 1);
error_reporting(E_ALL);
require 'auth/authentication.php'; // Update with the actual path to ServerCommunicator.php

$serverCommunicator = new ServerCommunicator();

// Example: Fetch data for a specific time period
$startTime = '2024-01-20 00:00:00';
$endTime = '2024-01-25 16:00:00';
$data = $serverCommunicator->dataWithTimePeriod($startTime, $endTime);

if ($data) {
    echo '<pre>';
    print_r($data);
    echo '</pre>';
} else {
    echo "No data or error occurred.";
}
?>
