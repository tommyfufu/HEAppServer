<?php
require 'vendor/autoload.php'; // If using Composer to manage dependencies

use GuzzleHttp\Client;
use GuzzleHttp\Exception\GuzzleException;

class ServerCommunicator {
    private $client;
    private $bearerToken;
    private $apiUrl = 'https://aocc-api-mgm.azure-api.net';
    private $config;
    private $subkey;
    private $sourcedomain;
    private $apiid;
    private $authtoken;
    public function __construct() {
        $this->client = new Client(); // Initialize the Guzzle client
	$this->config = require '/var/www/config/asus/asus_vivowatch_config.php';
    	$this->subkey = $this->config['subkey'];
    	$this->sourcedomain = $this->config['sourcedomain'];
	$this->apiid = $this->config['apiid'];
	$this->authtoken = $this->config['authtoken'];
    }

    public function getAuth() {
        try {
            $response = $this->client->request('POST', $this->apiUrl . "/vivowatch/bpapi/GetAuth", [
		    'headers' => [
	            'Content-Type' => 'application/json',
                    'Authorization' => 'Bearer ' . $this->authtoken,
                    'Ocp-Apim-Subscription-Key' => $this->subkey
		],
		'body' => json_encode([
                    'source_domain' => $this->sourcedomain,
                    'apiid' => $this->apiid
                ])
	    ]);
	    $statusCode = $response->getStatusCode();

            if ($statusCode == 200) {
                $data = json_decode($response->getBody(), true);
                $this->bearerToken = $data['token'];
	    } else if ($statusCode == 400){
                echo "Wrong Parameters Error: HTTP status code " . $statusCode;
            } else if ($statusCode == 401){
                echo "Token Wrong Error: HTTP status code " . $statusCode;
	    }  else if ($statusCode == 500){
                echo "Internal Serious Error: HTTP status code " . $statusCode;
	    } else {
                echo "Error: HTTP status code " . $statusCode;
	    }
        } catch (GuzzleException $e) {
            echo $e->getMessage();
        }
    }

    public function dataWithTimePeriod($startTime, $endTime) {
        if (!$this->bearerToken) {
            $this->getAuth();
        }

        try {
            $response = $this->client->request('POST', $this->apiUrl . "/vivowatch_2b_data/v1/bpapi/DataWithTimePeriod", [
                'headers' => [
	            'Content-Type' => 'application/json', // Add this line
                    'Authorization' => 'Bearer ' . $this->bearerToken,
                    'Ocp-Apim-Subscription-Key' => $this->subkey
                ],
		'body' => json_encode([
                    'start_time' => $startTime,
                    'end_time' => $endTime
                ])
            ]);

	    $statusCode = $response->getStatusCode();

            if ($statusCode == 200) {
		$data = json_decode($response->getBody(), true);
		return $data;
	    } else if ($statusCode == 204){
                echo "Time Period no data Error: HTTP status code " . $statusCode;
            } else if ($statusCode == 400){
                echo "Wrong Parameters Error: HTTP status code " . $statusCode;
            } else if ($statusCode == 401){
                echo "Token Wrong Error: HTTP status code " . $statusCode;
	    }  else if ($statusCode == 500){
                echo "Internal Serious Error: HTTP status code " . $statusCode;
	    } else {
                echo "Error: HTTP status code " . $statusCode;
	    }
        } catch (GuzzleException $e) {
            echo $e->getMessage();
	}

    }
}
?>
