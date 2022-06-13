package order_events

type EventType uint8

const (
	EventTypeDefault EventType = iota + 1
	EventTypeNewOrder
	EventTypeCancelOrder
	EventTypeOrderFailed
	EventTypeBcTxFound
	EventTypeReOrder
	EventTypeOrderSuccessfullyProcessed
)

const (
	EventTypeDefaultName                    = "event_default"
	EventTypeNewOrderName                   = "event_new_order"
	EventTypeCancelOrderName                = "event_cancel_order"
	EventTypeOrderFailedName                = "event_order_failed"
	EventTypeBcTxFoundName                  = "event_bc_tx_found"
	EventTypeReOrderName                    = "event_re_order"
	EventTypeOrderSuccessfullyProcessedName = "event_order_successfully_processed"
)

func (d EventType) String() string {
	return [...]string{"",
		EventTypeDefaultName,
		EventTypeNewOrderName,
		EventTypeCancelOrderName,
		EventTypeOrderFailedName,
		EventTypeBcTxFoundName,
		EventTypeReOrderName,
		EventTypeOrderSuccessfullyProcessedName,
	}[d]
}
