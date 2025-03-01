"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getSearchedQuote = exports.getRandomQuote = exports.getAllQuotes = void 0;
var lodash_1 = __importDefault(require("lodash"));
var quotes_with_id_json_1 = __importDefault(require("../quotes-with-id.json"));
//load the quotes JSON
//START OF YOUR CODE...
var getAllQuotes = function (req, res) {
    res.send(quotes_with_id_json_1.default);
};
exports.getAllQuotes = getAllQuotes;
var getRandomQuote = function (req, res) {
    res.send(pickFromArray(quotes_with_id_json_1.default));
};
exports.getRandomQuote = getRandomQuote;
var getSearchedQuote = function (req, res) {
    var _a, _b;
    var searchTerm = (_b = (_a = req.query.term) === null || _a === void 0 ? void 0 : _a.toString()) !== null && _b !== void 0 ? _b : '';
    if (req.query.term) {
        var filtered = quotes_with_id_json_1.default.filter(function (_a) {
            var quote = _a.quote, author = _a.author;
            return (quote.toLowerCase().includes(searchTerm.toString()) ||
                author.toLowerCase().includes(searchTerm.toString()));
        });
        res.send(filtered);
    }
};
exports.getSearchedQuote = getSearchedQuote;
//...END OF YOUR CODE
//You can use this function to pick one element at random from a given array
//example: pickFromArray([1,2,3,4]), or
//example: pickFromArray(myContactsArray)
//
function pickFromArray(arr) {
    return lodash_1.default.sample(arr);
}
