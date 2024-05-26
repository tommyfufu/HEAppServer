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

$twoWeeksAgo = date('Y-m-d ', strtotime("-7 days")) . "00:00:00";
$currTime = date('Y-m-d ', strtotime("+1 days")) . "00:00:00";

// Fetch data for the last two weeks
$dataArray = $serverCommunicator->dataWithTimePeriod($twoWeeksAgo, $currTime);

$latestHb = null;
$latestBp = null;
$latestSpo2 = null;
$latestStep = null;

foreach ($dataArray['daily_data'] as $dailyData) {
    if ($dailyData['deviceid'] !== $deviceId) {
        continue; // Skip data not related to the specified device ID
    }

    foreach ($dailyData['hb'] as $hb) {
        if (!$latestHb || strtotime($hb['time']) > strtotime($latestHb['time'])) {
            $latestHb = $hb;
        }
    }

    foreach ($dailyData['bp'] as $bp) {
        if (!$latestBp || strtotime($bp['time']) > strtotime($latestBp['time'])) {
            $latestBp = $bp;
        }
    }

    foreach ($dailyData['spo2'] as $spo2) {
        if (!$latestSpo2 || strtotime($spo2['time']) > strtotime($latestSpo2['time'])) {
            $latestSpo2 = $spo2;
        }
    }

    foreach ($dailyData['step'] as $step) {
        if (!$latestStep || strtotime($step['time']) > strtotime($latestStep['time'])) {
            $latestStep = $step;
        }
    }
}

if ($latestHb === null && $latestBp === null && $latestSpo2 === null && $latestStep === null) {
    echo json_encode(['error' => 'No data found for the specified device ID']);
    exit;
}

$latestData = [
    'deviceId' => $deviceId,
    'latestHb' => $latestHb,
    'latestBp' => $latestBp,
    'latestSpO2' => $latestSpo2,
    'latestStep' => $latestStep,
];

$jsonLatestData = json_encode($latestData, JSON_PRETTY_PRINT);
echo $jsonLatestData;
?>
