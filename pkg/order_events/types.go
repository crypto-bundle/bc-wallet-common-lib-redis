package order_events

type EventType uint8

const (
	EventTypeDefault EventType = iota + 1
	EventTypeOrderAccepted
	EventTypeOrderCreated
	EventTypeCancelOrder
	EventTypeOrderFailed
	EventTypeBcTxFound
	EventTypeReOrder
	EventTypeOrderSuccessfullyProcessed
)

const (
	EventTypeDefaultName                    = "event_default"
	EventTypeAcceptedOrderName              = "event_order_accepted"
	EventTypeCreatedOrderName               = "event_order_created"
	EventTypeCancelOrderName                = "event_cancel_order"
	EventTypeOrderFailedName                = "event_order_failed"
	EventTypeBcTxFoundName                  = "event_bc_tx_found"
	EventTypeReOrderName                    = "event_re_order"
	EventTypeOrderSuccessfullyProcessedName = "event_order_successfully_processed"
)

func (d EventType) String() string {
	return [...]string{"",
		EventTypeDefaultName,
		EventTypeAcceptedOrderName,
		EventTypeCreatedOrderName,
		EventTypeCancelOrderName,
		EventTypeOrderFailedName,
		EventTypeBcTxFoundName,
		EventTypeReOrderName,
		EventTypeOrderSuccessfullyProcessedName,
	}[d]
}
