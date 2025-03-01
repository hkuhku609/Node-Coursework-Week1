"use strict";
// server.js
// This is where your node app starts
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
//load the 'express' module which makes writing webservers easy
require('dotenv').config();
var cors_1 = __importDefault(require("cors"));
var express_1 = __importDefault(require("express"));
var quotes_1 = __importDefault(require("./routes/quotes"));
var app = (0, express_1.default)();
app.use((0, cors_1.default)());
app.use(requestLogger);
// Now register handlers for some routes:
//   /                  - Return some helpful welcome info (text)
//   /quotes            - Should return all quotes (json)
//   /quotes/random     - Should return ONE quote (json)
app.get('/', function (request, response) {
    response.send("Neill's Quote Server!  Ask me for /quotes/random, or /quotes");
});
app.use('/quotes', quotes_1.default);
//Start our server so that it listens for HTTP requests!
var listener = app.listen(process.env.PORT || 4000, function () {
    var port = listener.address().port;
    console.log('Your app is listening on port ' + port);
});
function requestLogger(req, res, next) {
    var startTime = process.hrtime();
    res.on('finish', function () {
        var elapsedTime = process.hrtime(startTime);
        var elapsedTimeInMs = elapsedTime[0] * 1000 + elapsedTime[1] / 1e6;
        console.log("[".concat(req.method, "] [").concat(req.url, "] [").concat(elapsedTimeInMs.toFixed(3), "ms]"));
    });
    next();
}
