<?php
require_once __DIR__."/../vendor/autoload.php";

$app_id = '8000013';

$app_secret = '8e2dff6734643fcc50a12b6aefff3451';

$device_id = '123457';

$device_secret = 'jnxiaer7';

$rpc = new \zhongwu\protocol\RpcClient($app_id, $app_secret, 'http://api.zhongwuyun.com:9999');

$Zprinter = new \zhongwu\Printer($rpc);



$printdata = '中午云拥有自主研发的云打印机，提供稳定高效，高可用的云打印方案';

try {

    var_dump($Zprinter->set_args($device_id, $device_secret)->cloud_print($printdata));

} catch (Exception $e) {
    echo $e;
}