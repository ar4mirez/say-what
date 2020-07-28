<?php
function translateWord($word) {
  
    $response = '';

    switch ($word) {
        case "hello":
            $response = "buongiorno1";
            break;
        case "hi":
            $response = "ciao";
            break;
        case "bye":
            $response = "ciao";
            break;
        case "goodbye":
            $response = "arrivederci";
            break;
    }

    return $response;

}

?>