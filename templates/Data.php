

<?php
$servername = "127.0.0.1";
$username = "root";
$password = "heizhenzhu";
$semester = "2018 Spring" // $_POST['Semester']
$database_name = $semester . "_Course_Evl";

$con = new mysqli($servername,$username,$password);
if($con -> connect_error){
	die("Connection failed: " . $conn->connect_error);
}

//$sql = "Create DATABASE springf($database_name)";
//echo $sql;


?>
<!DOCTYPE html>
<html>
<head>
	<title></title>
</head>
<body>

</body>
</html>