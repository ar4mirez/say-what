<?php include '../../utils/index.php';?>
<?php

    header('Content-type: application/json');

    $word = $_GET['word'];
    $response = translateWord($word);

    $data = [ 'word' => $response ];

    echo json_encode( $data );

?>