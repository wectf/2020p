let express = require('express');
let router = express.Router();
const {v4} = require("uuid");
const validate = require('bitcoin-address-validation');

let WALLETS = {}; // temp database, it is mapping from sha256(user token) => wallet object
///////////////////////////// Utils /////////////////////////////
function sha256(message) {
  return require("crypto").createHash("sha256").update(message).digest("hex"); // convert message to sha256(message)
}
// Wrapper for web framework
class WebObject {
  // get all fields from children, selfishly believe all private fields starts with "__"
  get keys(){ return Reflect.ownKeys(this).filter(key => !key.match(/__/)); }
  get dict(){
    let result = {};
    this.keys.map(key => { result[key] = this[key]; })
    return result;
  }
  sendRaw(res){
    res.set("Content-Type", "text/plain");
    res.send(this.keys.map(key => { return `${key}=${this[key]}`; }).join(","));
  } // send raw form
  renderColor(res){ res.render(`${this.constructor.name}_color`, this.dict); } // render data with color template
  renderDark(res){ res.render(`${this.constructor.name}_dark`, this.dict); } // render data with dark template
}
///////////////////////////// Model /////////////////////////////
class Wallet extends WebObject {
  constructor(addresses) {
    super(); // js nonsense
    this.style = "Color Theme"; // theme
    this.addresses = addresses; // bitcoin addresses
  }
}
///////////////////////////// Middleware /////////////////////////////
router.use(async function (req, res, next) {
  if (JSON.stringify(WALLETS).length > 10000) WALLETS = {}; // prevent WALLETS obj from getting too big, not really part of the challenge
  let {user_token} = req.cookies; // get token from cookie
  let wallet = WALLETS[sha256(user_token || "")];
  if (!user_token || !wallet) WALLETS[sha256(user_token = v4())] = wallet = new Wallet([]); // initialize wallet
  req.wallet = wallet; // put user's wallet into req to pass down to other handlers
  res.cookie("user_token", user_token); // set cookie with user's token
  next(); // go to next handler
})
///////////////////////////// Route / /////////////////////////////
router.get('/', function(req, res, next) {
  if (/Dark Theme/.test(req.wallet.style)) return req.wallet.renderDark(res); // user wants to have dark theme
  if (/Raw/.test(req.wallet.style)) return req.wallet.sendRaw(res); // user is robot and wants to use raw theme
  return req.wallet.renderColor(res); // if user demands any other themes, we give them colorful theme
});
///////////////////////////// Route POST /address /////////////////////////////
router.post('/address', function(req, res, next) {
  req.wallet.addresses.push(validate(req.body.address).address || `Unknown Address: ${req.body.address}`); // add the address to temp db
  return res.redirect("/"); // redirect to /
});
///////////////////////////// Route POST /style /////////////////////////////
router.post('/style', function(req, res, next) {
  req.wallet.style = req.body.style; // update the style to temp db
  return res.redirect("/"); // redirect to /
});
module.exports = router;