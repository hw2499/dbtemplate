# dbTemplate
为用户提供快速搭建访问数据库的web服务端,支持http 及 grpc协议访问.
## 适用场景
- 集成到嵌入式设备,用于设备基本信息配置使用.
- 快速搭建小型服务端,适用于快速交付产品场景.

## 调用方式
- 服务端以命令行方式运行,按配置监听端口,启动一个http或grpc服务
- 客户端(前端)通过http或grpc协议,调用一个统一API接口操作设备端或服务端数据库
- 前后端通过传递json对象实现信息交互

## 配置方式
- 对数据库表的(增|删|改|查)操作只需要配置一个文件,部署简单方便,服务端人员无需编写代码,只需要设计业务表结构.
- 支持配置http协议及grpc协议及端口
- 支持配置防SQL注入正则表达式,服务端人员事先设计好配置文件中sql语句及交互参数.
- 支持配置导入SQL初始化文件创建库表结构
- 支持配置不同数据源连接

## 授权方式
- 永久免费使用,申请授权正式使用

## 如何使用
### 运行服务端
```shell
dbtemplate -fileUrl dbtemplate.xml
```

### 前端http协议调用
```shell
  postman调试说明:
	Method:
            POST
	URL:
	     http://<ip>:<port>/postCommonMethod

	Header:
	   KEY为
	     Content-Type
	   VALUE为
	     application/x-www-form-urlencoded

	Body:
	   KEY为
	     jsondata
	   VALUE为
              {
                  "id":"1",
                  "method":"sel",
                  "cols":"*",
                  "where":" and d_ip like '%192.168.' ",
          
                  "pageSize":10,
                  "currentPage":0
              }


```
### grpc协议调用
[参看样本](https://github.com/hw2499/dbtemplate/tree/main/service/hwgrpcclient)

## 接口说明
- http协议
  <br>调用方式 POST
  <br>接口名称 postCommonMethod
  <br>接口入参 jsondata 
  <br>接口出参 jsondata结果集
- grpc协议
  [参看样本](https://github.com/hw2499/dbtemplate/tree/main/service/hwgrpcclient)

## 配置文件说明

### 全局变量及关键字
配置文件中充当占位符,用于替换前端传递的对应值
- ${cols}  语义列变量,但可在任意位置出现
- ${values} 语义值变量,但可在任意位置出现
- ${where} 语义条件变量,但可在任意位置出现
- ${pageSize}  语义每页记录数变量,但可在任意位置出现
- ${currentPage}  语义当前页变量,但可在任意位置出现
- ${customKeys}  语义自定义变量Key名称,可在任意位置出现,多个key之间用分号;分隔
- ${customValues} 语义自定义变量Value名称,可在任意位置出现,多个value之间用分号;分隔
  ${customKeys} 和 ${customValues}非常灵活,可配置在任意位置.
### 全局函数

- ${timestamp10} 生成10位数字时间戳
- ${timestamp13} 生成13位数字时间戳
- ${timestamp16} 生成16位数字时间戳
- ${timestamp19} 生成19位数字时间戳
- ${uuid32} 生成32位小写UUID字符串
- ${UUID32} 生成32位大写UUID字符串
- ${YYYY-MM-DD hh:mm:ss} 生成对应的时间格式字符串
- ${YYYYMMDDhhmmss} 生成对应的时间格式字符串
- ${YYYY-MM-DD} 生成对应的时间格式字符串
- ${hh:mm:ss} 生成对应的时间格式字符串
- ${YYYY} 生成对应的时间格式字符串
- ${MM} 生成对应的时间格式字符串
- ${DD} 生成对应的时间格式字符串
- ${hh} 生成对应的时间格式字符串
- ${mm} 生成对应的时间格式字符串
- ${ss} 生成对应的时间格式字符串

### DBTemplate 

| 属性           | 说明                               |
|--------------|----------------------------------|
| serverPort          | 服务端监听端口                          |
| serverFlag         | 协议类型, http 或 grpc                |
| isBasicAuth       | 是否开启http协议的BasicAuth认证功能,true是开启 |
| basicUserName    | BasicAuth认证功能用户名称                |
| basicPassword | BasicAuth认证功能用户密码                |

### 样本

```shell
<DBTemplate  
    serverPort="8089" 
    isBasicAuth="true" 
    basicUserName="u1" 
    basicPassword="123456" 
    serverFlag="http" >
</DBTemplate>
```

### SqlInjectFilter

| 属性           | 说明          |
|--------------|-------------|
|           | 防SQL注入正则表达式 |

### 样本

```shell
  <!--
     .*(').*|.*(\b(select|insert|into|update|delete|from|where|and|or|trancate|drop|execute|like|grant|use|union|order|by)\b).*
    -->

    <SqlInjectFilter>
        .*(\b(select|insert|into|update|delete|from|trancate|drop|execute|grant|use|union|order|by)\b).*
    </SqlInjectFilter>

```

### DBSelect


| 属性           | 说明                       |
|--------------|--------------------------|
| id          | 唯一标示,前端通过调用该ID映射对应的SQL语句 |
| connId          | 数据源连接ID,通过该ID映射数据源                      |
| Sql          | 查询数据表语句                  |

### 样本

- 配置文件
```shell
 <DBSelect id="1" connId="conn_01" >
        <Sql>
            <![CDATA[
                    select ${cols}
                    from t_device_info
                    where 1=1
                    ${where} and d_width>30
                    limit  ${currentPage},${pageSize}
                 ]]>
        </Sql>
   </DBSelect>
```
- 前端传递JSON

```shell
      {
            "id":"1",
            "method":"sel",
            "cols":"d_id,d_name,d_ip,d_setup_address",
            "where":" and d_ip like '192.168.%' ",
            "pageSize":10,
            "currentPage":0
        }
```



### DBInsert

| 属性           | 说明                       |
|--------------|--------------------------|
| id          | 唯一标示,前端通过调用该ID映射对应的SQL语句 |
| connId          | 数据源连接ID,通过该ID映射数据源                      |
| Sql          | 插入数据表语句                  |

### 样本
- 配置文件

```shell
  <DBInsert id="3" connId="conn_01">
        <Sql>
            <![CDATA[
                  insert into t_device_info ( ${cols} ) values ( ${values} )
                   ]]>
        </Sql>
 </DBInsert>
```
- 前端传递JSON
```shell
 
      {
            "id":"3",
            "method":"ins",
            "cols":"d_id,d_name,d_setup_address,d_ip,d_desc,d_width,d_height,d_writetime",
            "values":"'${timestamp13}','路由器01','北京市大安小区15栋2单101室','192.168.0.110','Hx科技',20,40,'${YYYY-MM-DD hh:mm:ss}' ",

        }

```

### DBDelete

| 属性           | 说明                       |
|--------------|--------------------------|
| id          | 唯一标示,前端通过调用该ID映射对应的SQL语句 |
| connId          | 数据源连接ID,通过该ID映射数据源                     |
| Sql          | 删除数据表语句                  |

### 样本
- 配置文件

```shell
    <DBDelete id="6" connId="conn_01">
        <Sql>
            <![CDATA[
                delete from t_device_info where 1=1 ${custom_k1} ${custom_k2}
                   ]]>
        </Sql>
    </DBDelete>
```
- 前端传递JSON
```shell
{
            "id":"6",
            "method":"del",
            "customKeys":"custom_k1;custom_k2",
            "customValues":"and  d_desc='100M' ; and  d_name='路由器03' ",
        } 

```

### DBUpdate

| 属性           | 说明                       |
|--------------|--------------------------|
| id          | 唯一标示,前端通过调用该ID映射对应的SQL语句 |
| connId          | 数据源连接ID,通过该ID映射数据源       |
| Sql          | 更新数据表语句                  |


### 样本
- 配置文件

```shell

   <DBUpdate id="8" connId="conn_01">
        <Sql>
            <![CDATA[
                update t_device_info set d_width=${custom_k1},d_desc='${custom_k2}' where 1=1 ${where}

               ]]>
        </Sql>
  </DBUpdate>
```
- 前端传递JSON
```shell
      {
            "id":"8",
            "method":"upd",
            "customKeys":"custom_k1;custom_k2",
            "customValues":"60;100M",
            "where":"and d_name='路由器03' and d_ip='192.168.0.111'",
        } 

```


### DBConnection


| 属性           | 说明                  |
|--------------|---------------------|
| id          | 唯一标示,数据源连接ID        |
| dbType          | 数据库类型 mysql sqlite3 |
| dbURL          | 数据库地址               |
| database          | 数据库名称               |
| username          | 数据库用户名称             |
| password          | 数据库用户密码             |
| maxIdleConns          | 最大空闲连接数             |
| maxOpenConns          | 最大激活连接数             |
| connMaxLifetime          | 连接最大存活周期            |

### 样本
- 配置文件
```shell
    <DBConnection id="conn_011" dbType="mysql" dbURL="127.0.0.1:3306" database="db2" username="root" password="******" maxIdleConns="5" maxOpenConns="10" connMaxLifetime="300"   />
    <DBConnection id="conn_01" dbType="sqlite3" database="./sqlite3.db"  maxIdleConns="5" maxOpenConns="10" connMaxLifetime="300" />
 

```


### DBScript

| 属性           | 说明                   |
|--------------|----------------------|
| connId          | 数据源连接ID,通过该ID映射数据源   |
| fileUrl          | 初始化表结构文件位置           |
| isRun          | 服务启动时是否执行该脚本,true是执行 |

- sql 文件样本
```shell
create table if not exists  t_device_info (
  d_id varchar(32) not null,
  d_name varchar(50),
  d_setup_address varchar(100),
  d_ip varchar(20),
  d_desc varchar(100),
  d_width float, 
  d_height float,
  d_writetime varchar(19),
  PRIMARY KEY (d_id)
); 

```

