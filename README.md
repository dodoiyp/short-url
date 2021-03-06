# URL Shortener
## Requirement
This project is for dacard pre-interview homework 
  
Design and implement (with unit tests) an URL shortener using Go programming language.  
Criteria:  
 1. URL shortener has 2 APIs, please follow API example to implement:  
   a. A RESTful API to upload a URL with its expired date and response with a shorten URL.  
   b. b.An API to serve shorten URLs responded by upload API, and redirect to original URL. If URL is expired,please response with status 404.
 1. Please feel free to use any external libs if needed.  
 1. It is also free to use following external storage including:  
   a. Relational database (MySQL, PostgreSQL, SQLite).  
   b. Cache storage (Redis, Memcached). 
 1. Please implement reasonable constrains and error handling of these 2 APIs.
 1. You do not need to consider auth.
 1. Many clients might access shorten URL simultaneously or try to access with non-existent shorten URL, please take. 
performance into account.  

API example :
```
# Upload URL API
curl -X POST -H "Content-Type:application/json" http://localhost/api/v1/urls -d '{
   "url": "<original_url>",
"expireAt": "2021-02-08T09:20:41Z"
}'
# Response
{
   "id": "<url_id>",
   "shortUrl": "http://localhost/<url_id>"
}
# ------------------
# Redirect URL API
curl -L -X GET http://localhost/<url_id> => REDIRECT to original URL
```
Constraints:  
1. QPS:  
  a. Write:500.  
  b. Read:100k. 
1. url max length:none. 
1. Is shorturl need to be analysis :no.
1. shorturl expire time limit : none.  
1. Availability>Consistency

## QuickStart
1. 架設環境 mysql
> docker-compose up -d

2. 運行服務
> go run main.go 

3.測試api是否打通


## SystemDesign
思路 首先先完成一個可以執行的服務需要提供正常的上傳及轉址功能，目前想到兩種做法  
1. 將使用者輸入的資料做加密存在資料庫裡，短網址查詢時直接匹配
   缺點：
   1. 有可能會碰撞 
   2. 資料量大的時候演算對伺服器產生的壓力較大 
   解決辦法： 
   1. 需要在創建時多增加邏輯檢查是否存在已經存在的話就將加密解果做二次加密
   2. 花錢升級機器 
2. 由於短網址的產生的過程不需要與原網址有直接關聯所以可以用其他方式產生短網址後再將關聯存起來
![image](https://github.com/dodoiyp/short-url/blob/main/doc/v1/short_url_system%20design.jpg)  

上傳網址流程  先建立兩張表 1張為儲存url的表 1張用來產生key的表sequence. 
   每次要產生新的短網址時就先插入sequence 表將插入的id 做 base62 後對應給原網址  
![image](https://github.com/dodoiyp/short-url/blob/main/doc/v1/short_url-set-url.jpg)  
短網址轉址流程  
![image](https://github.com/dodoiyp/short-url/blob/main/doc/v1/short_url_get_url.jpg)  
##資料庫及三方套件選擇：
  資料庫：mysql reason: heavyRead system 通常還是使用關聯資料庫居多 由於沒有時間去測試 nosql(mongoDB) 所以先選用mysql 實作
  三方套件/框架/演算法選擇：  
    1.base62 :不論是之後需要真的對短網址壓縮或是用發號器製作key 計算相對簡單比較不消耗效能
    2.zap/log: 性能高
    3.caache/lru: 由於短網址在轉址的需求上量會較大 故將熱點更新至local 緩存 選用lru 讓越多人點擊的短網址能持續存在緩存內（雖然現在還沒有實現其他緩存清理）
    4.gorm:  為了快速實現(利用auto create table 功能 以及減少撰寫sql細節 專心在業務實現上)
    5.gin ：https://hackernoon.com/the-myth-about-golang-frameworks-and-external-libraries-93cb4b7da50f。
    6.viper: 先使用較熟悉的套件做 本題目與此套件較無直接關聯


接下來優化的方向為下圖
![image](https://github.com/dodoiyp/short-url/blob/main/doc/v2/short_url_system_design_%20optimization.jpg)
1. 加入外部cache (ex:redis,memcache). 
2. 將產生key的db/服務 與 存url的db/服務分開(但是會提升系統複雜度). 
3. 增加預先產生key的功能，可以加速上傳url 的速度. 
4. 增加schedule service 移除過期的數據
5. 批次新增key 並將部份key 存入緩存或是程式變數讓請求key的次數減少ex: 一次發1000把key 
6. 如果資料量龐大可以分表分庫 或是讀寫分離 

此設計之後要面臨的問題及目前的想法做討論  
1. 由於上傳的網址有可能會相同但是過期時間不一定相同可能會導致很多短網址都是轉導到同個原網址  
2. 短時間內一直查詢不存在的短網址 消耗系統效能(緩存穿透)  
3. 緩存的熱門資料過期導致請求全部導入主資料庫(緩存擊穿)  
3. 大量緩存的資料過期時間相同導致在某一時間出現大量請求送向主資料庫(緩存雪崩)  
3. 緩存的更新機制需要注意(ex 資料庫資料已經過期 緩存尚未被更新 會導致使用者還是可以轉址)。  
  
解法：   
1. 這需要從需求面下去討論若是希望短網址與原網址是1:1的話 則不能讓使用者輸入過期時間由服務設計方設計 若有人重複上傳相同網址更新過期時間延長  
2. 限制流量或是將查詢不存在的暫存起來 只要查詢到就在緩存層直接回覆 等該短網址真的有被用到或是過了過期時間再將其從黑名單移除  
3. 查詢緩存沒有資料後加入lock 在主資料庫查完後解除lock一次只讓一個人查主資料庫查完後趕快寫入緩存 或是緩存熱點不移除
4. 在建立緩存時緩存的過期時間加減隨機時間降低"同時"失效的機率
5. 只要接近過期時間就要檢查還存資料或是將主資料庫的過期時間一併紀錄在緩存上。
