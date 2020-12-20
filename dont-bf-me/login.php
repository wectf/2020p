<?php
include "constant.php";

parse_str($_SERVER["QUERY_STRING"]);

// check args
if (!isset($password) || !isset($_GET['g-recaptcha-response'])) {
    echo "Missing args :(";
    die();
}

// check recaptcha
$recaptcha_resp = json_decode(file_get_contents($RECAPTCHA_URL.$_GET['g-recaptcha-response']), true);
if(!$recaptcha_resp || !$recaptcha_resp["success"]) {
    echo "Bad recaptcha :(";
    die();
}

if ($recaptcha_resp["score"] < 0.8) {
    echo "Stop! Big hacker";
    die();
}

// check password
if($password == $CORRECT_PASSWORD) {
    echo $FLAG;
} else {
    echo "Wrong password :(";
}