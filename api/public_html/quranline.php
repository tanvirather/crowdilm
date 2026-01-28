<?php
require_once($_SERVER['DOCUMENT_ROOT'] . '/../private_html/initialize.php');

$quran_id = array_change_key_case($_GET, CASE_LOWER)['quran_id'];
$sql = <<<EOT
select
    line.surah,
    line.aya,
    line.manzil,
    line.juz,
    line.hizb,
    line.ruku,
    line.page,
    quran_line.line_id,
    quran_line.line_id as line,
    quran_line.text
from quran_line
inner join line on quran_line.line_id = line.id
where quran_line.quran_id = '$quran_id'
EOT;

echo select($sql, true);
?>
