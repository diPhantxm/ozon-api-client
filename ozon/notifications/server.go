package notifications

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Handler func(req interface{}) error

type NotificationServer struct {
	port     int
	handlers map[MessageType]Handler
}

func NewNotificationServer(port int) *NotificationServer {
	return &NotificationServer{
		port:     port,
		handlers: map[MessageType]Handler{},
	}
}

func (ns *NotificationServer) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ns.handler)
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", ns.port),
		Handler: mux,
	}
	return server.ListenAndServe()
}

func (ns *NotificationServer) handler(rw http.ResponseWriter, httpReq *http.Request) {
	mt := &Common{}
	content, err := ioutil.ReadAll(httpReq.Body)
	if err != nil {
		ns.error(rw, http.StatusBadRequest, err)
		return
	}
	if err := json.Unmarshal(content, mt); err != nil {
		ns.error(rw, http.StatusBadRequest, err)
		return
	}

	if mt.MessageType == PingType {
		resp := pingResponse{
			Version: "1.0",
			Name:    "Ozon Seller API",
			Time:    time.Now(),
		}
		respJson, err := json.Marshal(resp)
		if err != nil {
			ns.error(rw, http.StatusInternalServerError, err)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(respJson)
		return
	}

	req, err := ns.unmarshal(mt.MessageType, content)
	if err != nil {
		ns.result(rw, false)
		//ns.error(rw, http.StatusInternalServerError, err)
		return
	}
	h := ns.handlers[mt.MessageType]
	if err := h(req); err != nil {
		ns.result(rw, false)
		//ns.error(rw, http.StatusInternalServerError, err)
		return
	}
}

func (ns *NotificationServer) Register(mt MessageType, handler func(req interface{}) error) {
	ns.handlers[mt] = handler
}

func (ns *NotificationServer) unmarshal(messageType MessageType, content []byte) (interface{}, error) {
	switch messageType {
	case NewPostingType:
		v := &NewPosting{}
		err := json.Unmarshal(content, v)
		return v, err
	case PostingCancelledType:
		v := &PostingCancelled{}
		err := json.Unmarshal(content, v)
		return v, err
	case StateChangedType:
		v := &StateChanged{}
		err := json.Unmarshal(content, v)
		return v, err
	case CutoffDateChangedType:
		v := &CutoffDateChanged{}
		err := json.Unmarshal(content, v)
		return v, err
	case DeliveryDateChangedType:
		v := &DeliveryDateChanged{}
		err := json.Unmarshal(content, v)
		return v, err
	case CreateOrUpdateType:
		v := &CreateOrUpdateItem{}
		err := json.Unmarshal(content, v)
		return v, err
	case PriceIndexChangedType:
		v := &PriceIndexChanged{}
		err := json.Unmarshal(content, v)
		return v, err
	case StocksChangedType:
		v := &StocksChanged{}
		err := json.Unmarshal(content, v)
		return v, err
	case NewMessageType:
		v := &NewMessage{}
		err := json.Unmarshal(content, v)
		return v, err
	case UpdateMessageType:
		v := &UpdateMessage{}
		err := json.Unmarshal(content, v)
		return v, err
	case MessageReadType:
		v := &MessageRead{}
		err := json.Unmarshal(content, v)
		return v, err
	case ChatClosedType:
		v := &ChatClosed{}
		err := json.Unmarshal(content, v)
		return v, err
	default:
		return nil, fmt.Errorf("unsupported type: %s", messageType)
	}
}

func (ns *NotificationServer) error(rw http.ResponseWriter, statusCode int, err error) {
	errResp := errorResponse{
		Data: errorData{
			Code:    fmt.Sprintf("%d", statusCode),
			Message: err.Error(),
		},
	}
	errJson, _ := json.Marshal(errResp)
	rw.WriteHeader(statusCode)
	rw.Write(errJson)
}

func (ns *NotificationServer) result(rw http.ResponseWriter, res bool) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(fmt.Sprintf(`{"result": %t}`, res)))
}
