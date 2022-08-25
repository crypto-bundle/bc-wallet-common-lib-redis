package order_events

type EventType uint8

const (
	EventTypeDefault EventType = iota + 1
	EventTypeOrderAccepted
	EventTypeWithdrawOrderCreated
	EventTypeCancelOrder
	EventTypeOrderFailed
	EventTypeBcTxFound
	EventTypeReOrder
	EventTypeOrderSuccessfullyProcessed
	EventTypeDepositOrderCreated
)

const (
	EventTypeDefaultName                    = "event_default"
	EventTypeAcceptedOrderName              = "event_order_accepted"
	EventTypeCreatedWithdrawOrderName       = "event_withdraw_order_created"
	EventTypeCancelOrderName                = "event_cancel_order"
	EventTypeOrderFailedName                = "event_order_failed"
	EventTypeBcTxFoundName                  = "event_bc_tx_found"
	EventTypeReOrderName                    = "event_re_order"
	EventTypeOrderSuccessfullyProcessedName = "event_order_successfully_processed"
	EventTypeCreatedDepositOrderName        = "event_deposit_order_created"
)

func (d EventType) String() string {
	return [...]string{"",
		EventTypeDefaultName,
		EventTypeAcceptedOrderName,
		EventTypeCreatedWithdrawOrderName,
		EventTypeCancelOrderName,
		EventTypeOrderFailedName,
		EventTypeBcTxFoundName,
		EventTypeReOrderName,
		EventTypeOrderSuccessfullyProcessedName,
		EventTypeCreatedDepositOrderName,
	}[d]
}
