<?php
ini_set('display_errors', 1);
error_reporting(E_ALL);
require 'auth/authentication.php';

// Create an instance of the server communicator
$serverCommunicator = new ServerCommunicator();

if (!isset($_GET['deviceid']) || empty($_GET['deviceid'])) {
    echo json_encode(['error' => 'Device ID is required']);
    exit;
}

$deviceId = $_GET['deviceid'];

// Set the time period to one year ago until today
$oneYearAgo = date('Y-m-d ', strtotime("-1 year")) . "00:00:00";
$currTime = date('Y-m-d ', strtotime("+1 days")) . "00:00:00";

// Fetch data for the last year
$dataArray = $serverCommunicator->dataWithTimePeriod($oneYearAgo, $currTime);

if (empty($dataArray['daily_data'])) {
    echo json_encode(['error' => 'No data found for the specified device ID']);
    exit;
}

// Collect all data for the specified device ID
$allData = [
    'hb' => [],
    'bp' => [],
    'spo2' => [],
    'step' => []
];

foreach ($dataArray['daily_data'] as $dailyData) {
    if ($dailyData['deviceid'] !== $deviceId) {
        continue; // Skip data not related to the specified device ID
    }

    // Merge all hb, bp, spo2, and step data
    $allData['hb'] = array_merge($allData['hb'], $dailyData['hb']);
    $allData['bp'] = array_merge($allData['bp'], $dailyData['bp']);
    $allData['spo2'] = array_merge($allData['spo2'], $dailyData['spo2']);
    $allData['step'] = array_merge($allData['step'], $dailyData['step']);
}

if (empty($allData['hb']) && empty($allData['bp']) && empty($allData['spo2']) && empty($allData['step'])) {
    echo json_encode(['error' => 'No data found for the specified device ID']);
    exit;
}

$jsonData = json_encode(['deviceId' => $deviceId, 'data' => $allData], JSON_PRETTY_PRINT);
echo $jsonData;
?>
