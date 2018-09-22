<?php

namespace zhongwu\protocol;

use InvalidArgumentException;
use Exception;
use zhongwu\exceptions\BusinessException;
use zhongwu\exceptions\InvalidTimestampException;
use zhongwu\exceptions\UnauthorizedException;
use zhongwu\exceptions\ValidationFailedException;

class RpcClient
{
    private $app_id;
    private $app_secret;
    private $remote_url;

    public function __construct($app_id, $app_secret, $remote_url = 'http://api.zhongwuyun.com')
    {
        if (is_null($app_id)
            || is_null($app_secret)
            || is_null($remote_url)
        ) {

            throw new InvalidArgumentException("invalid construct parameters.");
        }

        $this->app_id = $app_id;
        $this->app_secret = $app_secret;
        $this->remote_url = $remote_url;
    }

    /** call server api with nop
     * @param       $action
     * @param array $parameters
     * @return mixed
     * @throws BusinessException
     * @throws Exception
     */
    public function call($action, array $parameters = array())
    {
        $protocol = array_merge(array(
            "appid" => $this->app_id,
            "timestamp" => time(),
        ), $parameters);

        $protocol['sign'] = $this->generate_signature($protocol);

        $result = $this->post($this->remote_url . "/" . $action, $protocol);
        $response = json_decode($result);
        if (is_null($response)) {
            throw new Exception("invalid response.");
        }

        if (isset($response->errNum) && $response->errNum != 0) {
            switch ($response->errNum) {
                case 1:
                    throw new UnauthorizedException($response->errMsg);
                    break;
                case 2:
                    throw new InvalidTimestampException($response->errMsg);
                    break;
                case 3:
                    throw new ValidationFailedException($response->errMsg);
                    break;
                default:
                    throw new BusinessException($response->errMsg);
            }
        }

        return $response;
    }

    private function generate_signature($protocol)
    {
        $stringtoSigned = '';

        ksort($protocol);

        foreach ($protocol as $k => $v) {

            $stringtoSigned .= $k . $v;

        }

        $stringtoSigned .= $this->app_secret;

        return md5($stringtoSigned);
    }

    private function post($url, $data)
    {
        $ch = curl_init($url);
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_HTTPHEADER, array('Expect:'));
        curl_setopt($ch, CURLOPT_POSTFIELDS, http_build_query($data));
        curl_setopt($ch, CURLOPT_TIMEOUT, 10);
        $response = curl_exec($ch);
        if (curl_errno($ch)) {
            throw new Exception(curl_error($ch));
        }

        return $response;
    }
}
