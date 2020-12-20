<?php
$flag = getenv("FLAG");
waf_echo($_SERVER, $flag);
show_source(__FILE__);