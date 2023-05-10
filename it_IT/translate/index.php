<?php

    require_once('../utils/index.php');

    header("Content-Type: application/json; charset=UTF-8");

    $word = $_GET['word'];
    $response = translateWord($word);

    $data = array('word' => $response);

    echo json_encode($data);
