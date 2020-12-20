<?php

/*
Plugin Name: Userpage
Plugin URI: https://wectf.io
Description: Create page of user :)
Version: 1.0
Author: shou
Author URI: https://scf.so/
License: i dont care
*/


// Some must-have stuffs
if (!defined('WPINC')) { die; }
// Activation
function activate_user_page() {}
register_activation_hook( __FILE__, 'activate_user_page' );

// Deactivation
function deactivate_user_page() {}
register_deactivation_hook( __FILE__, 'deactivate_user_page' );

require_once "tmpl.php";
function setup_menu() {
    add_menu_page( 'Edit uPage',
        'Edit uPage',
        'read',
        'edit_upage',
        'edit_tmpl' );
    add_menu_page( 'My uPage',
        'My uPage',
        'read',
        'my_upage',
        'my_tmpl' );
}
add_action('admin_menu', 'setup_menu');
