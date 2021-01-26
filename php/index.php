<?php

/**
 * Fetches an array of applications from the GET endpoint
 * @return array
 */
function getApplications() {
    // make the request, read data from GET response
    $response = file_get_contents('https://engineering-task.elancoapps.com/api/applications');
    // decode json
    return json_decode($response);
}

/**
 * Pretty prints the elements of the array as '- element0\n...'
 *
 * @param array $array The array to pretty print
 */
function prettyPrintArray($array) {
    // loop through elements
    foreach ($array as &$value) {
        echo "- {$value}\n";
    }
}

$applications = getApplications();
prettyPrintArray($applications);

?>
