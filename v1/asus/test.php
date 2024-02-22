<?php
ini_set('display_errors', 1);
error_reporting(E_ALL);
require 'auth/authentication.php';

$serverCommunicator = new ServerCommunicator();

$twoWeeksAgo = '2024-02-10 00:00:00';
$currTime = date('Y-m-d H:i:s');

// Fetch data for the last two weeks
//$dataArray = $serverCommunicator->dataWithTimePeriod($twoWeeksAgo, $currTime);

echo $currTime;
?>
