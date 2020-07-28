<?php include './utils/index.php';?>
<?php

use PHPUnit\Framework\TestCase;

class TranslateTest extends TestCase
{
    public function testSuccessWord()
    {
        $this->assertSame('ciao', translateWord('hi'));
    }

    public function testWrongWord()
    {
        $this->assertSame('', translateWord('no-valid-word'));
    }
}

?>