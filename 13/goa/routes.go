/**
* @api {get} /me/company/txs
* @apiGroup shop me/business
* @apiDescription list companies transactions
* @apiPermission Business
 */
r.Path("/me/company/txs").Methods("POST").Handler(dc.With(txs.CompanyList)) // HL
/**
 * @api {post} /me/orders/encrypted encrypted checkout
 * @apiGroup shop me/business
 * @apiDescription creates order from cart in request body, encrypting personal order info
 * @apiPermission Bank
 */
r.Path("/me/orders/encrypted").Methods("POST").Handler(dcBank.With(carts.EncryptedCheckout)) // HL
/**
 * @api {post} /me/orders/onestep onestep checkout
 * @apiGroup shop me/business
 * @apiDescription creates order from cart in request body
 * @apiPermission Business
 */
r.Path("/me/orders/onestep").Methods("POST").Handler(dc.With(carts.BusinessOnestepCheckout)) // HL
