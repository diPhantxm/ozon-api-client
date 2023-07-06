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

type SupplyRequestState string

const (
	// request draft. Only for supplies via vDC
	Draft SupplyRequestState = "DRAFT"

	// selecting supply options. Only for supplies via vDC
	SupplyVariantsArranging SupplyRequestState = "SUPPLY_VARIANTS_ARRANGING"

	// no supply options, the request is archived. Only for supplies via vDC
	HasNoSupplyVariantsArchive SupplyRequestState = "HAS_NO_SUPPLY_VARIANTS_ARCHIVE"

	// no supply options. Only for supplies via vDC
	HasNoSupplyVariantsNew SupplyRequestState = "HAS_NO_SUPPLY_VARIANTS_NEW"

	// supply being approved. Only for supplies via vDC
	SupplyVariantsConfirmation SupplyRequestState = "SUPPLY_VARIANTS_CONFIRMATION"

	// time reservation
	TimeslotBooking SupplyRequestState = "TIMESLOT_BOOKING"

	// filling in the data
	DATA_FILLING SupplyRequestState = "DATA_FILLING"

	// ready for shipment
	ReadyToSupply SupplyRequestState = "READY_TO_SUPPLY"

	// accepted at the shipping point
	AcceptedAtSupplyWarehouse SupplyRequestState = "ACCEPTED_AT_SUPPLY_WAREHOUSE"

	// on the way
	InTransit SupplyRequestState = "IN_TRANSIT"

	// acceptance at the warehouse
	AcceptanceAtStorageWarehouse SupplyRequestState = "ACCEPTANCE_AT_STORAGE_WAREHOUSE"

	// acts being approved
	ReportsConfirmationAwaiting SupplyRequestState = "REPORTS_CONFIRMATION_AWAITING"

	// dispute
	ReportRejected SupplyRequestState = "REPORT_REJECTED"

	// completed
	Completed SupplyRequestState = "COMPLETED"

	// refused acceptance
	RejectedAtSupplyWarehouse SupplyRequestState = "REJECTED_AT_SUPPLY_WAREHOUSE"

	// cancelled
	Cancelled SupplyRequestState = "CANCELLED"

	// overdue
	Overdue SupplyRequestState = "OVERDUE"
)

type ShipmentStatus string

const (
	// acceptance is in progress
	AcceptanceInProgress ShipmentStatus = "acceptance_in_progress"

	// arbitration
	Arbitration ShipmentStatus = "arbitration"

	// awaiting confirmation
	AwaitingApprove ShipmentStatus = "awaiting_approve"

	// awaiting shipping
	AwaitingDeliver ShipmentStatus = "awaiting_deliver"

	// awaiting packaging
	AwaitingPackaging ShipmentStatus = "awaiting_packaging"

	// created
	AwaitingVerification ShipmentStatus = "awaiting_verification"

	// cancelled
	CancelledSubstatus ShipmentStatus = "cancelled"

	// delivered
	Delivered ShipmentStatus = "delivered"

	// delivery is in progress
	Delivering ShipmentStatus = "delivering"

	// picked up by driver
	DriverPickup ShipmentStatus = "driver_pickup"

	// not accepted at the sorting center
	NotAccepted ShipmentStatus = "not_accepted"

	// sent by the seller
	SentBySeller ShipmentStatus = "sent_by_seller"
)

type ShipmentSubstatus string

const (
	// acceptance in progress
	PostingAcceptanceInProgress ShipmentStatus = "posting_acceptance_in_progress"

	// arbitrage
	PostingInArbitration ShipmentStatus = "posting_in_arbitration"

	// created
	PostingCreated ShipmentStatus = "posting_created"

	// in the freight
	PostingInCarriage ShipmentStatus = "posting_in_carriage"

	// not added to the freight
	PostingNotInCarriage ShipmentStatus = "posting_not_in_carriage"

	// registered
	PostingRegistered ShipmentStatus = "posting_registered"

	// is handed over to the delivery service
	PostingTransferringToDelivery ShipmentStatus = "posting_transferring_to_delivery"

	// waiting for passport data
	PostingAwaitingPassportData ShipmentStatus = "posting_awaiting_passport_data"

	// created
	PostingCreatedSubstatus ShipmentStatus = "posting_created"

	// awaiting registration
	PostingAwaitingRegistration ShipmentStatus = "posting_awaiting_registration"

	// registration error
	PostingRegistrationError ShipmentStatus = "posting_registration_error"

	// created
	PostingSplitPending ShipmentStatus = "posting_split_pending"

	// canceled
	PostingCancelled ShipmentStatus = "posting_canceled"

	// customer delivery arbitrage
	PostingInClientArbitration ShipmentStatus = "posting_in_client_arbitration"

	// delivered
	PostingDelivered ShipmentStatus = "posting_delivered"

	// recieved
	PostingReceived ShipmentStatus = "posting_received"

	// presumably delivered
	PostingConditionallyDelivered ShipmentStatus = "posting_conditionally_delivered"

	// courier on the way
	PostingInCourierService ShipmentStatus = "posting_in_courier_service"

	// at the pick-up point
	PostingInPickupPoint ShipmentStatus = "posting_in_pickup_point"

	// on the way to the city
	PostingOnWayToCity ShipmentStatus = "posting_on_way_to_city"

	// on the way to the pick-up point
	PostingOnWayToPickupPoint ShipmentStatus = "posting_on_way_to_pickup_point"

	// returned to the warehouse
	PostingReturnedToWarehouse ShipmentStatus = "posting_returned_to_warehouse"

	// is handed over to the courier
	PostingTransferredToCourierService ShipmentStatus = "posting_transferred_to_courier_service"

	// handed over to the driver
	PostingDriverPickup ShipmentStatus = "posting_driver_pick_up"

	// not accepted at the sorting center
	PostingNotInSortCenter ShipmentStatus = "posting_not_in_sort_center"

	// sent by the seller
	SentBySellerSubstatus ShipmentStatus = "sent_by_seller"
)

type TPLIntegrationType string

const (
	// delivery by the Ozon logistics
	OzonTPLType TPLIntegrationType = "ozon"

	// delivery by a third-party service, Ozon registers the order
	AggregatorTPLType TPLIntegrationType = "aggregator"

	// delivery by a third-party service, the seller registers the order
	TrackingTPLType TPLIntegrationType = "3pl_tracking"

	// delivery by the seller
	NonIntegratedTPLType TPLIntegrationType = "non_integrated"
)

type DetailsDeliveryItemName string

const (
	DirectFlowLogisticSumDetailsDeliveryItemName DetailsDeliveryItemName = "MarketplaceServiceItemDirectFlowLogisticSum"
	DropoffDetailsDeliveryItemName               DetailsDeliveryItemName = "MarketplaceServiceItemDropoff"
	DelivToCustomerDetailsDeliveryItemName       DetailsDeliveryItemName = "MarketplaceServiceItemDelivToCustomer"
)

type DetailsReturnServiceName string

const (
	ReturnAfterDelivToCustomerDetailsReturnServiceName DetailsReturnServiceName = "MarketplaceServiceItemReturnAfterDelivToCustomer"
	ReturnPartGoodsCustomerDetailsReturnServiceName    DetailsReturnServiceName = "MarketplaceServiceItemReturnPartGoodsCustomer"
	ReturnNotDelivToCustomerDetailsReturnServiceName   DetailsReturnServiceName = "MarketplaceServiceItemReturnNotDelivToCustomer"
	ReturnFlowLogisticDetailsReturnServiceName         DetailsReturnServiceName = "MarketplaceServiceItemReturnFlowLogistic"
)

type DetailsServiceItemName string

const (
	OtherMarketAndTech                          DetailsServiceItemName = "MarketplaceServiceItemOtherMarketAndTechService"
	ReturnStorageServiceAtThePickupPointFbsItem DetailsServiceItemName = "MarketplaceReturnStorageServiceAtThePickupPointFbsItem"
	SaleReviewsItem                             DetailsServiceItemName = "MarketplaceSaleReviewsItem"
	ServicePremiumCashbackIndividualPoints      DetailsServiceItemName = "MarketplaceServicePremiumCashbackIndividualPoints"
	ServiceStorageItem                          DetailsServiceItemName = "MarketplaceServiceStorageItem"
	ServiceStockDisposal                        DetailsServiceItemName = "MarketplaceServiceStockDisposal"
	ReturnDisposalServiceFbsItem                DetailsServiceItemName = "MarketplaceReturnDisposalServiceFbsItem"
	ServiceItemFlexiblePaymentSchedule          DetailsServiceItemName = "MarketplaceServiceItemFlexiblePaymentSchedule"
	ServiceProcessingSpoilage                   DetailsServiceItemName = "MarketplaceServiceProcessingSpoilage"
	ServiceProcessingIdentifiedSurplus          DetailsServiceItemName = "MarketplaceServiceProcessingIdentifiedSurplus"
	ServiceProcessingIdentifiedDiscrepancies    DetailsServiceItemName = "MarketplaceServiceProcessingIdentifiedDiscrepancies"
	ServiceItemInternetSiteAdvertising          DetailsServiceItemName = "MarketplaceServiceItemInternetSiteAdvertising"
	ServiceItemPremiumSubscribtion              DetailsServiceItemName = "MarketplaceServiceItemPremiumSubscribtion"
	AgencyFeeAggregator3PLGlobalItem            DetailsServiceItemName = "MarketplaceAgencyFeeAggregator3PLGlobalItem"
)

type DetailsOtherItemName string

const (
	RedistributionOfAcquiringOperation                   DetailsOtherItemName = "MarketplaceRedistributionOfAcquiringOperation"
	CompensationLossOfGoodsOperation                     DetailsOtherItemName = "MarketplaceSellerCompensationLossOfGoodsOperation"
	CorrectionOperation                                  DetailsOtherItemName = "MarketplaceSellerCorrectionOperation"
	OperationCorrectionSeller                            DetailsOtherItemName = "OperationCorrectionSeller"
	OperationMarketplaceWithHoldingForUndeliverableGoods DetailsOtherItemName = "OperationMarketplaceWithHoldingForUndeliverableGoods"
	OperationClaim                                       DetailsOtherItemName = "OperationClaim"
)

type StrategyType string

const (
	MinExtPrice StrategyType = "MIN_EXT_PRICE"
	CompPrice   StrategyType = "COMP_PRICE"
)

type StrategyUpdateType string

const (
	StrategyEnabled          StrategyUpdateType = "strategyEnabled"
	StrategyDisabled         StrategyUpdateType = "strategyDisabled"
	StrategyChanged          StrategyUpdateType = "strategyChanged"
	StrategyCreated          StrategyUpdateType = "strategyCreated"
	StrategyItemsListChanged StrategyUpdateType = "strategyItemsListChanged"
)

type ShipmentCertificateFilterStatus string

const (
	// new
	ShitmentCertificateFilterNew ShipmentCertificateFilterStatus = "new"

	// retry creation
	ShitmentCertificateFilterAwaitingRetry ShipmentCertificateFilterStatus = "awaiting-retry"

	// is being packaged
	ShitmentCertificateFilterInProcess ShipmentCertificateFilterStatus = "in_process"

	// created
	ShitmentCertificateFilterSuccess ShipmentCertificateFilterStatus = "success"

	// creation error
	ShitmentCertificateFilterError ShipmentCertificateFilterStatus = "error"

	// sent
	ShitmentCertificateFilterSend ShipmentCertificateFilterStatus = "sent"

	// received
	ShitmentCertificateFilterReceived ShipmentCertificateFilterStatus = "received"

	// packaged
	ShitmentCertificateFilterFormed ShipmentCertificateFilterStatus = "formed"

	// canceled
	ShitmentCertificateFilterCancelled ShipmentCertificateFilterStatus = "cancelled"

	// in the queue for packaging
	ShitmentCertificateFilterPending ShipmentCertificateFilterStatus = "pending"

	// in the queue for completion
	ShitmentCertificateFilterCompletionEnqueued ShipmentCertificateFilterStatus = "completion_enqueued"

	// in the process of completion
	ShitmentCertificateFilterCompletionProcessing ShipmentCertificateFilterStatus = "completion_processing"

	// completion error
	ShitmentCertificateFilterCompletionFailed ShipmentCertificateFilterStatus = "completion_failed"

	// in the queue for cancellation
	ShitmentCertificateFilterCancelationEnqueued ShipmentCertificateFilterStatus = "cancelation_enqueued"

	// in the process of cancellation
	ShitmentCertificateFilterCancelationProcessing ShipmentCertificateFilterStatus = "cancelation_processing"

	// cancellation error
	ShitmentCertificateFilterCancelationFailed ShipmentCertificateFilterStatus = "cancelation_failed"

	// completed
	ShitmentCertificateFilterCompleted ShipmentCertificateFilterStatus = "completed"

	// closed
	ShitmentCertificateFilterClosed ShipmentCertificateFilterStatus = "closed"
)

type PRROptionStatus string

const (
	// carrying the bulky product using the elevator
	PRROptionLift PRROptionStatus = "lift"

	// carrying the bulky product upstairs
	PRROptionStairs PRROptionStatus = "stairs"

	// the customer canceled the service,
	// you don't need to lift the shipment
	PRROptionNone PRROptionStatus = "none"

	// delivery is included in the price.
	// According to the offer you need to
	// deliver products to the floor
	PRROptionDeliveryDefault PRROptionStatus = "delivery_default"
)

type GetFBSReturnsFilterStatus string

const (
	ReturnedToSeller          GetFBSReturnsFilterStatus = "returned_to_seller"
	WaitingForSeller          GetFBSReturnsFilterStatus = "waiting_for_seller"
	AcceptedFromCustomer      GetFBSReturnsFilterStatus = "accepted_from_customer"
	CancelledWithCompensation GetFBSReturnsFilterStatus = "cancelled_with_compensation"
	ReadyForShipment          GetFBSReturnsFilterStatus = "ready_for_shipment"
)

type GetFBOReturnsFilterStatus string

const (
	GetFBOReturnsFilterStatusReturnedToOzon GetFBOReturnsFilterStatus = "ReturnedToOzon"
	GetFBOReturnsFilterStatusCancelled      GetFBOReturnsFilterStatus = "Cancelled"
)

type GetFBOReturnsReturnStatus string

const (
	GetFBOReturnsReturnStatusCancelled            GetFBOReturnsReturnStatus = "Возврат отменен"
	GetFBOReturnsReturnStatusAcceptedFromCustomer GetFBOReturnsReturnStatus = "Принят от покупателя"
	GetFBOReturnsReturnStatusReceivedAtOzon       GetFBOReturnsReturnStatus = "Получен в Ozon"
)
