<?php
function select($query, $setEncoding = false){
	// Create connection
	$conn = new mysqli(DB_SERVER, DB_USER, DB_PASS, DB_NAME);
	// Check connection
	if ($conn->connect_error){
		die("Connection failed");
	}

	// need to set the character set to display unicode data
	if($setEncoding){
		$conn->query('set character set utf8');
	}

	$rows = array();
	// output data of each row
	$result = $conn->query($query);
	while ($row = $result->fetch_assoc()) {
		$rows[] = $row;
	}
	$conn->close();

	// return data as json
	return json_encode($rows);
}
?>
