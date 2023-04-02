package ozon

type Order string

const (
	Ascending  Order = "ASC"
	Descending Order = "DESC"
)

type GetAnalyticsDataFilterOperation string

const (
	Equal        GetAnalyticsDataFilterOperation = "EQ"
	Greater      GetAnalyticsDataFilterOperation = "GT"
	GreaterEqual GetAnalyticsDataFilterOperation = "GTE"
	Lesser       GetAnalyticsDataFilterOperation = "LT"
	LesserEqual  GetAnalyticsDataFilterOperation = "LTE"
)

type GetAnalyticsDataFilterMetric string

const (
	UnknownMetric         GetAnalyticsDataFilterMetric = "unknown_metric"
	HitsViewSearch        GetAnalyticsDataFilterMetric = "hits_view_search"
	HistViewPDP           GetAnalyticsDataFilterMetric = "hits_view_pdp"
	HitsView              GetAnalyticsDataFilterMetric = "hist_view"
	HitsToCartSearch      GetAnalyticsDataFilterMetric = "hits_tocart_search"
	HitsToCartPDP         GetAnalyticsDataFilterMetric = "hits_tocart_pdp"
	SessionViewSearch     GetAnalyticsDataFilterMetric = "session_view_search"
	SessionViewPDP        GetAnalyticsDataFilterMetric = "session_view_pdp"
	SessionView           GetAnalyticsDataFilterMetric = "session_view"
	ConvToCartSearch      GetAnalyticsDataFilterMetric = "conv_tocart_search"
	ConvToCartPDP         GetAnalyticsDataFilterMetric = "conv_tocart_pdp"
	ConvToCart            GetAnalyticsDataFilterMetric = "conv_tocart"
	Revenue               GetAnalyticsDataFilterMetric = "revenue"
	ReturnsMetric         GetAnalyticsDataFilterMetric = "returns"
	CancellationsMetric   GetAnalyticsDataFilterMetric = "cancellations"
	OrderedUnits          GetAnalyticsDataFilterMetric = "ordered_units"
	DeliveredUnits        GetAnalyticsDataFilterMetric = "delivered_units"
	AdvViewPDP            GetAnalyticsDataFilterMetric = "adv_view_pdp"
	AdvViewSearchCategory GetAnalyticsDataFilterMetric = "adv_view_search_category"
	AdvViewAll            GetAnalyticsDataFilterMetric = "adv_view_all"
	AdvSumAll             GetAnalyticsDataFilterMetric = "adv_sum_all"
	PositionCategory      GetAnalyticsDataFilterMetric = "position_category"
	PostingsMetric        GetAnalyticsDataFilterMetric = "postings"
	PostingsPremium       GetAnalyticsDataFilterMetric = "postings_premium"
)

type WarehouseType string

const (
	// Ozon warehouses with Fresh delivery
	ExpressDarkStore WarehouseType = "EXPRESS_DARK_STORE"

	// Ozon warehouses without Fresh delivery
	NotExressDarkStore WarehouseType = "NOT_EXPRESS_DARK_STORE"

	// All Ozon warehouses
	ALLWarehouseType WarehouseType = "ALL"
)

type Language string

const (
	Default Language = "DEFAULT"
	Russian Language = "RU"
	English Language = "EN"
	Turkish Language = "TR"
	Chinese Language = "ZH_HANS"
)

type AttributeType string

const (
	All      AttributeType = "ALL"
	Required AttributeType = "REQUIRED"
	Optional AttributeType = "OPTIONAL"
)

type ListDiscountRequestsStatus string

const (
	New            ListDiscountRequestsStatus = "NEW"
	Seen           ListDiscountRequestsStatus = "SEEN"
	Approved       ListDiscountRequestsStatus = "APPROVED"
	PartlyApproved ListDiscountRequestsStatus = "PARTLY_APPROVED"
	Declined       ListDiscountRequestsStatus = "DECLINED"
	AutoDeclined   ListDiscountRequestsStatus = "AUTO_DECLINED"
	DeclinedByUser ListDiscountRequestsStatus = "DECLINED_BY_USER"
	Coupon         ListDiscountRequestsStatus = "COUPON"
	Purchased      ListDiscountRequestsStatus = "PURCHASED"
)

type WorkingDay string

const (
	Mon WorkingDay = "1"
	Tue WorkingDay = "2"
	Wed WorkingDay = "3"
	Thu WorkingDay = "4"
	Fri WorkingDay = "5"
	Sat WorkingDay = "6"
	Sun WorkingDay = "7"
)

type GetAnalyticsDataDimension string

const (
	UnknownDimension   GetAnalyticsDataDimension = "unknownDimension"
	SKUDimension       GetAnalyticsDataDimension = "sku"
	SPUDimension       GetAnalyticsDataDimension = "spu"
	DayDimension       GetAnalyticsDataDimension = "day"
	WeekDimension      GetAnalyticsDataDimension = "week"
	MonthDimension     GetAnalyticsDataDimension = "month"
	YearDimension      GetAnalyticsDataDimension = "year"
	Category1Dimension GetAnalyticsDataDimension = "category1"
	Category2Dimension GetAnalyticsDataDimension = "category2"
	Category3Dimension GetAnalyticsDataDimension = "category3"
	Category4Dimension GetAnalyticsDataDimension = "category4"
	BrandDimension     GetAnalyticsDataDimension = "brand"
	ModelIDDimension   GetAnalyticsDataDimension = "modelID"
)
