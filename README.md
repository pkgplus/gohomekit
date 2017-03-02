# homekit
## 简介
小米智能设备aqara支持Homekit连接工具，实现ios10手机`家庭`控制，且可以使用siri语音控制。

## 实现功能
目前支持功能包括：
* 小米智能网关灯的开关操作；
* 小米智能网关灯的颜色控制；
* 小米温湿度传感器数据显示；

## 使用方法
### 准备工作
1. 一台24小时开机的服务器：可以是Windows，也可以是树莓派等Linux服务器；
2. 一部iOS10手机；
3. 小米智能家庭硬件（绿米设备）；
4. 一部安卓手机（开启智能网关的局域网通信使用）；

### 开启网关局域网通信功能
[开启方法说明](http://bbs.xiaomi.cn/t-13198850)

### 部署gohomekit
1. 根据系统下载gohomekit并解压，[下载页](http://download.bingbaba.com/gohomekit/)
2. 修改配置文件conf/app.json
3. 运行`gohomekit`，可后台运行： `nohup ./gohomekit 2>&1 >./gohomekit.log &`

> 配置文件配置说明：platforms中sid为网关MAC地址，password为网关加密口令，可从安卓手机中查询到

### 手机增加配件
打开`家庭`应用，增加配件，PIN码为conf/app.json配置文件配置的PIN码`12365478`，然后一路下一步就OK了
这时候就可以使用siri控制灯了，比如说“开灯”、“关灯”、“把灯调为红色”等等

## TODO
* 支持人体感应器
* 支持门窗状态
* 增加智能开关
* 增加86火线开关等

## 截图
控制中心效果图：
![手机](http://download.bingbaba.com/images/homekit_iphone.png)
