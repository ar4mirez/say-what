# it_IT

## Pre-requisites
* PHP 7.4+

## Running

### Local
Run `php -S 127.0.0.1:8080`.

### Docker
Generating image executing `docker build ./ -t <name>:<version>`.
Running Docker container executing `docker run -p <local_port>:80 --name=<name_for_container> <name>:<version>`.

## Test
* Install PHPUnit. (https://phpunit.readthedocs.io/en/9.2/installation.html)
* Runnig some like that `php phpunit-9.2.phar Tests/TranslateTest.php`. (According to way you have instaled PHPUnit).

## Author
* Santiago Yepes <zetogk@gmail.com>