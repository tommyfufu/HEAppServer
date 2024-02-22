<?php
ini_set('display_errors', 1);
error_reporting(E_ALL);
require 'auth/authentication.php';

$serverCommunicator = new ServerCommunicator();


$twoWeeksAgo = date('Y-m-d ', strtotime("-7 days")) . "00:00:00";
$currTime = date('Y-m-d ', strtotime("+1 days")) . "00:00:00";
//$currTime = date('Y-m-d H:i:s');

// Fetch data for the last two weeks
$dataArray = $serverCommunicator->dataWithTimePeriod($twoWeeksAgo, $currTime);

$deviceId = null;
$latestHb = null;
$latestBp = null;
$latestSpo2 = null;
$latestStep = null;


foreach ($dataArray['daily_data'] as $dailyData) {
	$deviceId = $dailyData['deviceid'];
    foreach ($dailyData['hb'] as $hb) {
        if (!$latestHb || strtotime($hb['time']) > strtotime($latestHb['time'])) {
            $latestHb = $hb;
        }
    }

    // Process blood pressure data
    foreach ($dailyData['bp'] as $bp) {
        if (!$latestBp || strtotime($bp['time']) > strtotime($latestBp['time'])) {
            $latestBp = $bp;
        }
    }

    // Process spo2 data
    foreach ($dailyData['spo2'] as $spo2) {
        if (!$latestSpo2 || strtotime($spo2['time']) > strtotime($latestSpo2['time'])) {
            $latestSpo2 = $spo2;
        }
    }

    // Process step data
    foreach ($dailyData['step'] as $step) {
        if (!$latestStep || strtotime($step['time']) > strtotime($latestStep['time'])) {
            $latestStep = $step;
        }
    }
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
