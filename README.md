# 圈子账单导出为钱迹模板

将圈子账本的字符串 json 账本转换为钱迹可接受的模板 excel 文件导入钱迹

## 环境依赖
go 1.12

## 使用方式
### 获取账本 json 字符串数据，存入 data.json 文件中
1. 访问[圈子账本](https://www.jizhangapp.com/account)登录账号
2. F12 检查 network 的请求记录，然后找到类似：
https://quanzi.jizhangapp.com/weixin/h5/book/5e56a76b-7a67-4167-a1c3-77559a8e52d1?list=5e56a76b-7a67-4167-a1c3-77559a8e52d1 的请求，得到后面的一大串字符串。
内容类似于
```json
{
    "data": {
      "membersCount": 1,
      "pageStats": [
        {
          "pageExpenses": [
            {
              "date": "2022-03-22",
              "totalIncome": 0,
              "totalCost": 21,
              "expenses": [
                {
                  "categoryColor": "F7756D",
                  "cost": 6,
                  "cts": "2022-03-22 09:11:29",
                  "accountName": "支付宝",
                  "creatorName": "XXXXXX",
                  "creatorId": "XXXXXX",
                  "accountUuid": "XXXXXX",
                  "remark": "#早餐#",
                  "type": 0,
                  "uuid": "XXXXXX",
                  "categoryName": "餐饮",
                  "lastCtsTimestamp": 1647911489000,
                  "categoryUuid": "XXXXXX",
                  "action": 0,
                  "iconUrl": "https://cdn.jizhangapp.com/app/category/1.png"
                }]
            }]
        }]
    }
}
```
### 运行脚本生成导出数据
1. 运行 export_bill.sh 脚本
脚本会自动下载 go1.12.17 版本，代码依赖于此版本的 go
```bash
bash -x ./export_bill.sh
```
### 将数据导入钱迹
1. 按照[钱迹账本模板导入数据指引](http://docs.qianjiapp.com/other/import_templete.html)将账本数据导入钱迹
