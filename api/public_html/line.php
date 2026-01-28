<?php
require_once($_SERVER['DOCUMENT_ROOT'] . '/../private_html/initialize.php');

$sql = "select id, surah, aya, manzil, juz, hizb, ruku, page from line";
echo select($sql, true);
?>
