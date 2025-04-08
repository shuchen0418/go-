EasySwap NFT交易平台 - 后端开发工程师
2023.10 - 2024.03

项目描述：
EasySwap是一个去中心化NFT交易平台，支持多链(ETH、Optimism、Arbitrum等)的NFT挂单、出价、交易等功能。负责开发和维护平台的区块链事件同步服务和后端API服务。

核心职责：
1. 设计并实现区块链事件同步系统(EasySwapSync)
- **开发智能合约事件监听和处理模块，实现NFT挂单(List)、出价(Bid)、交易(Match)等链上事件的实时同步**

  （挂单/出价事件(Make)

  处理用户挂单出售NFT或对NFT发起出价

  支持针对单个NFT和整个Collection的操作

  将订单信息存入数据库并加入订单管理队列

  

  交易撮合事件(Match)

  处理NFT交易完成时的撮合信息

  更新买卖双方订单状态

  记录交易活动并更新NFT所有权

  

  取消订单事件(Cancel)

  处理用户取消挂单或出价的操作

  更新订单状态并记录相关活动）

  

- **设计并实现Collection地板价更新机制，通过定时任务维护NFT集合的最新地板价格**
  1.设置定时器每天查询一次地板价
  2.将查到的地板价存到数据库中
  3.清理过期的地板价记录

- **支持多链并行处理**

  每个链的处理逻辑封装在Service结构体中。

  使用threading.GoSafe并行启动多个服务实例。

  每个服务实例独立运行其事件同步和地板价格更新的循环。

  通过传入不同的链ID和配置，确保每个实例可以处理不同链的数据

2. 开发交易平台核心功能(EasySwapBackend)
- **实现NFT交易相关的API接口，包括订单管理、用户资产、交易历史等功能**

  

- **设计并实现订单撮合引擎，支持自动匹配买卖订单**

- **开发用户钱包集成模块，支持多链资产管理**

### . 用户相关功能

用户登录验证

钱包登录：UserLoginHandler -> UserLogin -> getUserLoginMsgCacheKey -> CacheUserToken

获取登录信息：GetLoginMessageHandler -> GetUserLoginMsg -> getUserLoginMsgCacheKey

获取签名状态：GetSigStatusHandler -> GetSigStatusMsg -> GetUserSigStatus

用户NFT信息查询

查询用户在指定NFT集合中已上架的物品数量

查询用户在多个区块链上的NFT收藏品合集信息：UserMultiChainCollectionsHandler -> GetMultiChainUserCollections

查询用户在多个区块链上的具体NFT物品信息：UserMultiChainItemsHandler -> GetMultiChainUserItems

查询用户在多个区块链上的NFT竞价出价信息：UserMultiChainListingsHandler -> GetMultiChainUserListings

查询用户在多个区块链上挂单出售的NFT信息：UserMultiChainListingsHandler -> GetMultiChainUserListings

### 2. NFT及集合相关功能

NFT集合信息查询

获取集合出价：CollectionBidsHandler -> GetBids -> QueryCollectionBids

查询集合中的NFT列表：CollectionItemsHandler -> GetItems -> QueryCollectionItemOrder

查询集合中的NFT列表及补充信息：CollectionItemsHandler -> GetItems

获取特定NFT集合中某个NFT token的详细信息：ItemDetailHandler -> GetItem

获取NFT特征(Trait)相关的价格排名信息：ItemTopTraitPriceHandler -> GetItemTopTraitPrice

获取不同时间段内的NFT集合价格变化信息：HistorySalesHandler -> GetHistorySalesPrice

查询NFT集合交易排行榜：TopRankingHandler -> GetTopRanking

### 3. NFT交易及活动记录

NFT交易记录查询

NFT交易平台的活动记录查询：ActivityMultiChainHandler -> GetAllChainActivities -> QueryAllChainActivities

特定活动查询：ActivityMultiChainHandler -> GetMultiChainActivities -> QueryMultiChainActivities

NFT竞价信息

查看某个NFT集合的单个NFT竞价信息：CollectionItemBidsHandler -> GetItemBidsInfo -> QueryItemBids

### 4. NFT管理与维护

NFT所有权与状态维护

获取并维护NFT的最新所有权信息：ItemOwnerHandler -> GetItemOwner

将NFT加入刷新队列：ItemMetadataRefreshHandler -> RefreshItemMetadata

获取NFT图片信息：GetItemImageHandler -> GetItemImage

### 5. 事件处理功能（EasyswapSync）

事件处理

处理MakeEvent事件：SyncOrderBookEventLoop -> handleMakeEvent

处理CancelEvent事件：SyncOrderBookEventLoop -> handleCancelEvent

处理MatchEvent事件：SyncOrderBookEventLoop -> handleMatchEvent

地板价维护

维护NFT集合地板价信息：UpKeepingCollectionFloorChangeLoop

技术亮点：
1. 采用Go语言开发高性能区块链事件同步系统，通过批处理和并发优化，单链TPS达到1000+
2. 设计实现分布式定时任务系统，确保多节点部署时地板价更新的一致性
3. 使用GORM进行数据库操作，优化查询性能，单表数据量超过1000万条
4. 基于Redis实现分布式缓存，提升热点数据访问性能
5. 应用zap日志库进行系统监控，实现错误追踪和性能分析

技术栈：
- 后端：Go、Go-Zero、GORM、Redis
- 区块链：Web3、以太坊、智能合约、多链开发
- 数据库：MySQL
- 工具：Git、Docker

项目成果：
1. 系统支持ETH、Optimism、Arbitrum等6条主流公链
2. 平台日交易额突破100ETH
3. 服务可用性达到99.9%
4. 链上事件同步延迟<3秒