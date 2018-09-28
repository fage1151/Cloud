<?php
/**
 * Created by PhpStorm.
 * User: heqian
 * Date: 18-9-28
 * Time: 下午11:40
 */

namespace zhongwu;

use zhongwu\protocol\RpcClient;


class App
{
    private $client;

    /**
     * printer constructor.
     * @param RpcClient $rpc
     */
    public function __construct($rpc)
    {
        $this->client = $rpc;
    }

    /**
     * @desc
     * @param $dataid
     * @return mixed
     * @throws \Exception
     * @throws exceptions\BusinessException
     * @throws exceptions\InvalidTimestampException
     */
    public function cancel_one($dataid)
    {
        return $this->client->call('cancelone', compact("dataid"));

    }
}