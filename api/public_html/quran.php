<?php
require_once($_SERVER['DOCUMENT_ROOT'] . '/../private_html/initialize.php');

$sql = "select id, language, name, name_english, quran_type from quran";
echo select($sql, true);
?>
