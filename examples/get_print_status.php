<?php

$app_id = '******';

$app_secret = '********';

$rpc = new \zhongwu\protocol\RpcClient($app_id, $app_secret, 'http://api.zhongwuyun.com');

$Zprinter = new \zhongwu\Printer($rpc);

$device_id = '1111111';

$device_secret = '11111111';

$id = 123;
try {

    $Zprinter->set_args($device_id, $device_secret)->get_print_status($id);

} catch (Exception $e) {

}