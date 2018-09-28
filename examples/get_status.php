<?php
require_once __DIR__."/../vendor/autoload.php";

$app_id = '8000016';

$app_secret = '********';

$rpc = new \zhongwu\protocol\RpcClient($app_id, $app_secret, 'http://api.zhongwuyun.com');

$Zprinter = new \zhongwu\Printer($rpc);

$device_id = '1111111';

$device_secret = '11111111';

try {

    $rs = $Zprinter->set_args($device_id, $device_secret)->get_status();

    var_dump($rs);

} catch (Exception $e) {

}