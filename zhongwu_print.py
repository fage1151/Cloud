#!/usr/bin/python
#encode: utf8
import requests
import  demjson
import  time
import  base64
import  os.path
from   hashlib import md5


class ZW_print:
    url='http://api.zhongwuyun.com'
    def __init__(self, app_id, device_id,device_secret):
        if  device_id ==''  or device_secret=='' or  app_id=='':
            return '传入参数不能为空'
        self.__appid=app_id
        self.__deviceid=device_id
        self.__devicesecret=device_secret
    # 签名加密算法
    def  sign(self,data):
        return md5(self.sortitemvalues(data).encode('utf-8')).hexdigest()
    # 按键排序
    def sortitemvalues(self,data):
        items=data.keys()
        items= sorted(items)
        tmp=''
        for key  in  items:
              tmp +=key+str( data[key])
        return  tmp


    # 打印发送数据
    def  cloud_print(self,printdata):
        data={}
        data['appid']=self.__appid
        data['deviceid']= self.__deviceid
        data['devicesecret']=self.__devicesecret
        data['printdata']=printdata
        data['timestamp']=int(time.time())
        data['sign']=self.sign(data)
        re=requests.post(self.url,data)
        redata=demjson.decode(re.text)
        if 'errNum' in redata.keys():
            return  redata['retData']['id']
        else:
            return  redata['error']

    # 设置声音
    def  set_sound(self,sound):
        if   int(sound) not in [1,2,3]:
            return 'sound 必须为 1,2,3'
        data={}
        data['appid']=self.__appid
        data['deviceid']= self.__deviceid
        data['devicesecret']=self.__devicesecret
        data['sound']=sound
        data['timestamp']=int(time.time())
        data['sign']=self.sign(data)
        re=requests.post(self.url+'/sound',data)
        redata=demjson.decode(re.text)
        if 'errNum' in redata.keys():
            if int( redata['retData']['status'] )!=1:
                return  'error,设备异常'
            else:
                return  redata['errMsg']
    def devicestatus(self):
        data={}
        data['appid']=self.__appid
        data['deviceid']= self.__deviceid
        data['devicesecret']=self.__devicesecret
        data['timestamp']=int(time.time())
        data['sign']=self.sign(data)
        re=requests.post(self.url+'/status',data)
        redata=demjson.decode(re.text)
        if 'errNum' in redata.keys():
            if int( redata['retData']['status'] ) ==0 :
                return  '离线'
            elif int( redata['retData']['status'] ) ==1 :

                return  '在线'
            elif int( redata['retData']['status'] ) ==2 :
                return  '缺纸'
    # 订单打印状态查询
    def printstatus(self,dataid):
        data={}
        data['appid']=self.__appid
        data['deviceid']= self.__deviceid
        data['devicesecret']=self.__devicesecret
        data['timestamp']=int(time.time())
        data['dataid']= dataid
        data['sign']=self.sign(data)
        re=requests.get(self.url+'/printstatus',data)
        redata=demjson.decode(re.text.encode('utf8'))
        if 'errNum' in redata.keys():
            if redata['retData']['status']==0:
                return '未打印'
            elif  redata['retData']['status']==1:
                return  '已打印'
        else:
           return redata['retMsg']
    # 清空队列
    def emptyqueue(self):
        data={}
        data['appid']=self.__appid
        data['deviceid']= self.__deviceid
        data['devicesecret']=self.__devicesecret
        data['timestamp']=int(time.time())
        data['sign']=self.sign(data)
        re=requests.post(self.url+'/emptyprintqueue',data)
        redata=demjson.decode(re.text)
        if 'errNum' in redata.keys():
            if int( redata['retData']['status']) ==1 :
                return  '清空队列'+redata['retData']['row']
            else:
                return  '打印机缺纸或者不在线'
    # 设置logo
    def  set_logo(self,img):
        if  not  os.path.exists(img):
            return  '文件路径不存在'

        fsize=os.path.getsize(img)
        if fsize >40* 1024:
            return  '文件不能超过40kb'
        f=open(img,'rb')
        logodata=  base64.b64encode(f.read())
        f.close()
        data={}
        data['appid']=self.__appid
        data['deviceid']= self.__deviceid
        data['devicesecret']=self.__devicesecret
        data['timestamp']=int(time.time())
        data['logodata']=logodata
        data['sign']=self.sign(data)
        re=requests.post(self.url+'/logo',data)
        redata=demjson.decode(re.text)
        if 'errNum' in redata.keys():
            if redata['errNum']==0:
                if int( redata['retData']['status']) ==1 :
                    return  '设置logo成功'
                else:
                    return  '打印机缺纸或者不在线'
            else:
                return  redata['retMsg']
        




#  应用id
app_id=''
#  打印机号
device_id=''
# 打印机密钥
device_secret=''
zw= ZW_print(app_id,device_id,device_secret)
# zw.cloud_print(printdata='s')
# zw.set_sound('2')
# zw.devicestatus()
# zw.printstatus(34232881)
# zw.emptyqueue()
# zw.set_logo('/home/shanxs/test.png')
