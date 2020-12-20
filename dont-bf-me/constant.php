<?php
// recaptcha
$PUB_KEY = getenv("PUB_KEY");
$PRIV_KEY = getenv("PRIV_KEY");
$RECAPTCHA_URL = "https://www.google.com/recaptcha/api/siteverify?secret=$PRIV_KEY&response=";

// password
$CORRECT_PASSWORD = getenv("PASSWORD");

// flag
$FLAG = getenv("FLAG");

// bug does not exist if we can't see it
error_reporting(0);