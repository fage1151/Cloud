<?php

require_once __DIR__."/../vendor/autoload.php";

$app_id = '8000000';

$app_secret = '********';

$device_id = '123457';

$device_secret = 'jnxiaer7';

$rpc = new \zhongwu\protocol\RpcClient($app_id, $app_secret, 'http://api.zhongwuyun.com');

$Zprinter = new \zhongwu\Printer($rpc);


try {

    $rs = $Zprinter->set_args($device_id, $device_secret)->set_sound(1);

    var_dump($rs);

} catch (Exception $e) {

}