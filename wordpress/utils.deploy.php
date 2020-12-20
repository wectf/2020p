<?php

// Get the db cursor
global $wpdb;
require_once(ABSPATH . 'wp-settings.php');

// Table name of where we save data
$usermeta_table_name = $wpdb->prefix . "usermeta";
class Upage {
    var $user_id;
    var $user_info = array();
    var $conf = ABSPATH.'/wp-content/plugins/userpage/upage.conf';
    var $disallowed_words = [];
    var $is_debug;
    // Load configuration of bad words
    private function _load_conf() {
        $conf_content = file_get_contents($this->conf);
        if (strlen($conf_content) < 15 || substr($conf_content, 0, 15) != 'upage_conf_file')
            echo "Config file in incorrect format: " . $conf_content;
        $args = explode("\n", $conf_content);

        $this->disallowed_words = explode(";", $args[1]);
        $this->is_debug = $args[2] == "true";
    }
    // Check whether the content contains bad words
    private function _check_content($content) {
        foreach ($this->disallowed_words as $word)
            if (strpos($content, $word) !== false) {
                echo "Bad content!!!";
                die();
            }
    }
    // Spin up
    public function __construct() {
        $this->_load_conf();
        $this->user_id = get_current_user_id();
        global $wpdb, $usermeta_table_name;
        // Get the user information
        $result = $wpdb->get_results("SELECT meta_value, meta_key FROM $usermeta_table_name
                WHERE user_id = $this->user_id");
        foreach ($result as $r)
            $this->user_info[$r->meta_key] = $r->meta_value;
    }
    public function __wakeup() { $this->_load_conf(); }
    // Set a specific entry of user information
    public function set($key, $value) {
        if ($key == "wp_capabilities"){
            echo "Writing to a read-only slot in table wp_usermeta";
            echo "Please don't try to be admin.";
            die();
        }
        global $wpdb, $usermeta_table_name;
        // check bad words
        $this->_check_content($value);
        if (isset($this->user_info[$key]) > 0) // already in DB
            $wpdb->query("UPDATE $usermeta_table_name SET meta_value='$value'
                    WHERE user_id=$this->user_id AND meta_key='$key'");
        else // not yet in DB
            $wpdb->query("INSERT INTO $usermeta_table_name (meta_value, meta_key, user_id)  
                    VALUES ('$value', '$key', $this->user_id)");
        $this->user_info[$key] = $value;
    }
    // Get a specific entry of user information
    public function __get($key) { return $this->user_info[$key]; }
}