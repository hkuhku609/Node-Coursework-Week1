"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var express_1 = __importDefault(require("express"));
var quotes_1 = require("../controllers/quotes");
var router = express_1.default.Router();
router.route('/').get(quotes_1.getAllQuotes);
router.route('/random').get(quotes_1.getRandomQuote);
router.route('/search').get(quotes_1.getSearchedQuote);
exports.default = router;
