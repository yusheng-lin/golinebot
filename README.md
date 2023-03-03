## About The Project

golang linebot exercise 

### Installation
+ HTTP framework: https://github.com/gin-gonic/gin
+ Config: https://github.com/spf13/viper
+ Mongo driver: https://github.com/mongodb/mongo-go-driver
+ Line Bot SDK: https://github.com/line/line-bot-sdk-go
+ Wire: https://github.com/google/wire
+ gin-swagger: https://github.com/swaggo/gin-swagger

### Run 
1. clone
```
git clone git@github.com:yusheng-lin/golinebot.git
cd golinebot
```
2. make
```
make
```
![Alt text](./demo/demo01.png)
3. ngrok
```
ngrok http 8080
```
![Alt text](./demo/demo02.png)

4. set webhook
![Alt text](./demo/demo03.png)

5. join the channel and key some message

![Alt text](./demo/demo04.png)

6. check db table

![Alt text](./demo/demo05.png)

7. swagger test push message

![Alt text](./demo/demo06.png)

8. swagger fetch messages

![Alt text](./demo/demo07.png)