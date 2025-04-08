1.用户登录验证
钱包登录：
UserLoginHandler -> UserLogin -> getUserLoginMsgCacheKey ->CacheUserToken
GetLoginMessageHandler -> GetUserLoginMsg -> getUserLoginMsgCacheKey
GetSigStatusHandler -> GetSigStatusMsg ->GetUserSigStatus

2.根据条件获取nft以及集合的最佳出价情况
OrderInfosHandler -> GetOrderInfos -> QueryItemsBestBids -> QueryCollectionTopNBid -> processBids

3.查询用户在指定 NFT 集合中已上架的物品数量

4.NFT交易平台的活动记录查询，可以查询用户的交易历史、NFT的转移记录等活动信息
全部ActivityMultiChainHandler -> GetAllChainActivities -> QueryAllChainActivities -> QueryAllChainActivityExternalInfo
特定ActivityMultiChainHandler -> GetMultiChainActivities -> QueryMultiChainActivities -> QueryMultiChainActivityExternalInfo

5.获取集合出价，包含某个NFT集合当前有多少个不同价格的有效出价以及NFT集合的出价情况(出价价格、该价格下的总交易额、在该价格下出价的不同买家数量)
CollectionBidsHandler -> GetBids -> QueryCollectionBids 

6.查询集合中的NFT列表，包含可以直接购买的，有报价的，同时满足两种情况的
CollectionItemsHandler -> GetItems -> QueryCollectionItemOrder

查询集合中的NFT列表，以及该NFT列表多个维度的补充信息（订单信息(ordersInfo) - 获取挂单详情，外部信息(ItemsExternal) - 获取图片/视频等媒体信息，用户持有数量(userItemCount) - 获取每个用户持有的 NFT 数量，最后成交价(lastSales) - 获取每个 NFT 最后的成交价格，最高出价(bestBids) - 获取每个 NFT 的最高出价信息，集合最高出价(collectionBestBid) - 获取整个集合的最高出价）
CollectionItemsHandler -> GetItems

7.查看某个NFT集合的单个NFT竞价信息
CollectionItemBidsHandler->GetItemBidsInfo->QueryItemBids

8.获取特定 NFT 集合中某个 NFT token 的详细信息(集合信息(collection),物品基本信息(item),物品挂单信息(itemListInfo),物品外部信息(图片/视频等)(ItemExternals),最近成交价格(lastSales),最高出价信息(bestBids),集合最高出价(collectionBestBid))
ItemDetailHandler->GetItem

9.获取 NFT 特征(Trait)相关的价格排名信息
ItemTopTraitPriceHandler-> GetItemTopTraitPrice

10.获取不同时间段内的NFT集合价格变化信息
HistorySalesHandler -> GetHistorySalesPrice

11.获取特定NFT的上架/挂单信息
ItemListingHandler->GetItemListingInfo

12.获取单个NFT特征(Trait)信息
ItemTraitsHandler -> GetItemTraits

13.获取并维护NFT的最新所有权信息
ItemOwnerHandler -> GetItemOwner

14.获取NFT图片信息
GetItemImageHandler -> GetItemImage

15.将 NFT 加入刷新队列
ItemMetadataRefreshHandler-> RefreshItemMetadata

16.展示 NFT 集合的概览页面,帮助用户了解该集合的基本情况和市场表现
CollectionDetailHandler-> GetCollectionDetail

17.查询用户在多个区块链上的NFT收藏品合集信息
UserMultiChainCollectionsHandler -> GetMultiChainUserCollections

18.查询用户在多个区块链上的具体NFT物品信息
UserMultiChainItemsHandler -> GetMultiChainUserItems

19.查询用户在多个区块链上的NFT竞价出价信息
UserMultiChainListingsHandler -> GetMultiChainUserListings

20.查询用户在多个区块链上挂单出售的NFT信息
UserMultiChainListingsHandler -> GetMultiChainUserListings

21.查询NFT集合交易排行榜
TopRankingHandler -> GetTopRanking

EasyswapSync已知功能

1.处理MakeEvent事件
SyncOrderBookEventLoop -> handleMakeEvent

2.处理CancelEvent事件
SyncOrderBookEventLoop -> handleCancelEvent

3.处理MatchEvent事件
SyncOrderBookEventLoop  -> handleMatchEvent

4.维护NFT集合地板价信息
UpKeepingCollectionFloorChangeLoop