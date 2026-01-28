<?php
require_once($_SERVER['DOCUMENT_ROOT'] . '/../private_html/initialize.php');

$sql = "SELECT id,ayas,name_arabic,name_english,revelation_city,revelation_order FROM sura;";
echo select($sql, true);
?>
