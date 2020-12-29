package payment

type Order struct {
	payment *Payment
}

func (o *Order) Unify(params map[string]string, contract ...bool) {
	if _, ok := params["spbill_create_ip"]; !ok {
		//logger
		o.payment.Logger.Errorf("lost payment.order.Unify params [spbill_create_ip]")
	}
	params["appid"] = o.payment.Config.AppID
	if _, ok := params["notify_url"]; !ok {
		params["notify_url"] = o.payment.Config.NotifyUrl
	}

	var isContract bool
	if len(contract) > 0 {
		isContract = contract[0]

	}
	if isContract {
		//@todo contract
		// params["contract_appid"] = o.payment.Config.AppID
		// params["contract_mchid"] = o.payment.Config.MchID

	}

}
