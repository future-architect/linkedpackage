require('./sourcemap-register.js');module.exports =
/******/ (() => { // webpackBootstrap
/******/ 	var __webpack_modules__ = ({

/***/ 955:
/***/ ((module, __unused_webpack_exports, __nccwpck_require__) => {

"use strict";


var defaultDayjs = __nccwpck_require__(112);
var customParseFormatPlugin = __nccwpck_require__(148);
var localizedFormatPlugin = __nccwpck_require__(442);
var isBetweenPlugin = __nccwpck_require__(987);

function _interopDefaultLegacy (e) { return e && typeof e === 'object' && 'default' in e ? e : { 'default': e }; }

var defaultDayjs__default = /*#__PURE__*/_interopDefaultLegacy(defaultDayjs);
var customParseFormatPlugin__default = /*#__PURE__*/_interopDefaultLegacy(customParseFormatPlugin);
var localizedFormatPlugin__default = /*#__PURE__*/_interopDefaultLegacy(localizedFormatPlugin);
var isBetweenPlugin__default = /*#__PURE__*/_interopDefaultLegacy(isBetweenPlugin);

defaultDayjs__default['default'].extend(customParseFormatPlugin__default['default']);
defaultDayjs__default['default'].extend(localizedFormatPlugin__default['default']);
defaultDayjs__default['default'].extend(isBetweenPlugin__default['default']);
var withLocale = function (dayjs, locale) {
    return !locale ? dayjs : function () {
        var args = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            args[_i] = arguments[_i];
        }
        return dayjs.apply(void 0, args).locale(locale);
    };
};
var defaultFormats = {
    normalDateWithWeekday: "ddd, MMM D",
    normalDate: "D MMMM",
    shortDate: "MMM D",
    monthAndDate: "MMMM D",
    dayOfMonth: "D",
    year: "YYYY",
    month: "MMMM",
    monthShort: "MMM",
    monthAndYear: "MMMM YYYY",
    weekday: "dddd",
    weekdayShort: "ddd",
    minutes: "mm",
    hours12h: "hh",
    hours24h: "HH",
    seconds: "ss",
    fullTime: "LT",
    fullTime12h: "hh:mm A",
    fullTime24h: "HH:mm",
    fullDate: "ll",
    fullDateWithWeekday: "dddd, LL",
    fullDateTime: "lll",
    fullDateTime12h: "ll hh:mm A",
    fullDateTime24h: "ll HH:mm",
    keyboardDate: "L",
    keyboardDateTime: "L LT",
    keyboardDateTime12h: "L hh:mm A",
    keyboardDateTime24h: "L HH:mm",
};
var DayjsUtils = /** @class */ (function () {
    function DayjsUtils(_a) {
        var _this = this;
        var _b = _a === void 0 ? {} : _a, locale = _b.locale, formats = _b.formats, instance = _b.instance;
        this.lib = "dayjs";
        this.is12HourCycleInCurrentLocale = function () {
            var _a, _b;
            /* istanbul ignore next */
            return /A|a/.test((_b = (_a = _this.rawDayJsInstance.Ls[_this.locale || "en"]) === null || _a === void 0 ? void 0 : _a.formats) === null || _b === void 0 ? void 0 : _b.LT);
        };
        this.getCurrentLocaleCode = function () {
            return _this.locale || "en";
        };
        this.getFormatHelperText = function (format) {
            // @see https://github.com/iamkun/dayjs/blob/dev/src/plugin/localizedFormat/index.js
            var localFormattingTokens = /(\[[^\[]*\])|(\\)?(LTS|LT|LL?L?L?)|./g;
            return format
                .match(localFormattingTokens)
                .map(function (token) {
                var _a, _b;
                var firstCharacter = token[0];
                if (firstCharacter === "L") {
                    /* istanbul ignore next */
                    return (_b = (_a = _this.rawDayJsInstance.Ls[_this.locale || "en"]) === null || _a === void 0 ? void 0 : _a.formats[token]) !== null && _b !== void 0 ? _b : token;
                }
                return token;
            })
                .join("")
                .replace(/a/gi, "(a|p)m")
                .toLocaleLowerCase();
        };
        this.parse = function (value, format) {
            if (value === "") {
                return null;
            }
            return _this.dayjs(value, format, _this.locale);
        };
        this.date = function (value) {
            if (value === null) {
                return null;
            }
            return _this.dayjs(value);
        };
        this.toJsDate = function (value) {
            return value.toDate();
        };
        this.isValid = function (value) {
            return _this.dayjs(value).isValid();
        };
        this.isNull = function (date) {
            return date === null;
        };
        this.getDiff = function (date, comparing, units) {
            return date.diff(comparing, units);
        };
        this.isAfter = function (date, value) {
            return date.isAfter(value);
        };
        this.isBefore = function (date, value) {
            return date.isBefore(value);
        };
        this.isAfterDay = function (date, value) {
            return date.isAfter(value, "day");
        };
        this.isBeforeDay = function (date, value) {
            return date.isBefore(value, "day");
        };
        this.isBeforeYear = function (date, value) {
            return date.isBefore(value, "year");
        };
        this.isAfterYear = function (date, value) {
            return date.isAfter(value, "year");
        };
        this.startOfDay = function (date) {
            return date.clone().startOf("day");
        };
        this.endOfDay = function (date) {
            return date.clone().endOf("day");
        };
        this.format = function (date, formatKey) {
            return _this.formatByString(date, _this.formats[formatKey]);
        };
        this.formatByString = function (date, formatString) {
            return _this.dayjs(date).format(formatString);
        };
        this.formatNumber = function (numberToFormat) {
            return numberToFormat;
        };
        this.getHours = function (date) {
            return date.hour();
        };
        this.addSeconds = function (date, count) {
            return count < 0
                ? date.subtract(Math.abs(count), "second")
                : date.add(count, "second");
        };
        this.addMinutes = function (date, count) {
            return count < 0
                ? date.subtract(Math.abs(count), "minute")
                : date.add(count, "minute");
        };
        this.addHours = function (date, count) {
            return count < 0 ? date.subtract(Math.abs(count), "hour") : date.add(count, "hour");
        };
        this.addDays = function (date, count) {
            return count < 0 ? date.subtract(Math.abs(count), "day") : date.add(count, "day");
        };
        this.addWeeks = function (date, count) {
            return count < 0 ? date.subtract(Math.abs(count), "week") : date.add(count, "week");
        };
        this.addMonths = function (date, count) {
            return count < 0 ? date.subtract(Math.abs(count), "month") : date.add(count, "month");
        };
        this.setMonth = function (date, count) {
            return date.set("month", count);
        };
        this.setHours = function (date, count) {
            return date.set("hour", count);
        };
        this.getMinutes = function (date) {
            return date.minute();
        };
        this.setMinutes = function (date, count) {
            return date.clone().set("minute", count);
        };
        this.getSeconds = function (date) {
            return date.second();
        };
        this.setSeconds = function (date, count) {
            return date.clone().set("second", count);
        };
        this.getMonth = function (date) {
            return date.month();
        };
        this.getDaysInMonth = function (date) {
            return date.daysInMonth();
        };
        this.isSameDay = function (date, comparing) {
            return date.isSame(comparing, "day");
        };
        this.isSameMonth = function (date, comparing) {
            return date.isSame(comparing, "month");
        };
        this.isSameYear = function (date, comparing) {
            return date.isSame(comparing, "year");
        };
        this.isSameHour = function (date, comparing) {
            return date.isSame(comparing, "hour");
        };
        this.getMeridiemText = function (ampm) {
            return ampm === "am" ? "AM" : "PM";
        };
        this.startOfMonth = function (date) {
            return date.clone().startOf("month");
        };
        this.endOfMonth = function (date) {
            return date.clone().endOf("month");
        };
        this.startOfWeek = function (date) {
            return date.clone().startOf("week");
        };
        this.endOfWeek = function (date) {
            return date.clone().endOf("week");
        };
        this.getNextMonth = function (date) {
            return date.clone().add(1, "month");
        };
        this.getPreviousMonth = function (date) {
            return date.clone().subtract(1, "month");
        };
        this.getMonthArray = function (date) {
            var firstMonth = date.clone().startOf("year");
            var monthArray = [firstMonth];
            while (monthArray.length < 12) {
                var prevMonth = monthArray[monthArray.length - 1];
                monthArray.push(_this.getNextMonth(prevMonth));
            }
            return monthArray;
        };
        this.getYear = function (date) {
            return date.year();
        };
        this.setYear = function (date, year) {
            return date.clone().set("year", year);
        };
        this.mergeDateAndTime = function (date, time) {
            return date.hour(time.hour()).minute(time.minute()).second(time.second());
        };
        this.getWeekdays = function () {
            var start = _this.dayjs().startOf("week");
            return [0, 1, 2, 3, 4, 5, 6].map(function (diff) {
                return _this.formatByString(start.add(diff, "day"), "dd");
            });
        };
        this.isEqual = function (value, comparing) {
            if (value === null && comparing === null) {
                return true;
            }
            return _this.dayjs(value).isSame(comparing);
        };
        this.getWeekArray = function (date) {
            var start = _this.dayjs(date).clone().startOf("month").startOf("week");
            var end = _this.dayjs(date).clone().endOf("month").endOf("week");
            var count = 0;
            var current = start;
            var nestedWeeks = [];
            while (current.isBefore(end)) {
                var weekNumber = Math.floor(count / 7);
                nestedWeeks[weekNumber] = nestedWeeks[weekNumber] || [];
                nestedWeeks[weekNumber].push(current);
                current = current.clone().add(1, "day");
                count += 1;
            }
            return nestedWeeks;
        };
        this.getYearRange = function (start, end) {
            var startDate = _this.dayjs(start).startOf("year");
            var endDate = _this.dayjs(end).endOf("year");
            var years = [];
            var current = startDate;
            while (current.isBefore(endDate)) {
                years.push(current);
                current = current.clone().add(1, "year");
            }
            return years;
        };
        this.isWithinRange = function (date, _a) {
            var start = _a[0], end = _a[1];
            return date.isBetween(start, end, null, "[]");
        };
        this.rawDayJsInstance = instance || defaultDayjs__default['default'];
        this.dayjs = withLocale(this.rawDayJsInstance, locale);
        this.locale = locale;
        this.formats = Object.assign({}, defaultFormats, formats);
    }
    return DayjsUtils;
}());

module.exports = DayjsUtils;


/***/ }),

/***/ 65:
/***/ ((module, exports) => {

exports = module.exports = trim;

function trim(str){
  if (str.trim) return str.trim();
  return exports.right(exports.left(str));
}

exports.left = function(str){
  if (str.trimLeft) return str.trimLeft();

  return str.replace(/^\s\s*/, '');
};

exports.right = function(str){
  if (str.trimRight) return str.trimRight();

  var whitespace_pattern = /\s/,
      i = str.length;
  while (whitespace_pattern.test(str.charAt(--i)));

  return str.slice(0, i + 1);
};


/***/ }),

/***/ 399:
/***/ ((__unused_webpack_module, __webpack_exports__, __nccwpck_require__) => {

"use strict";
__nccwpck_require__.r(__webpack_exports__);
/* harmony import */ var trim__WEBPACK_IMPORTED_MODULE_0__ = __nccwpck_require__(65);
/* harmony import */ var trim__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__nccwpck_require__.n(trim__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var _date_io_dayjs__WEBPACK_IMPORTED_MODULE_1__ = __nccwpck_require__(955);
/* harmony import */ var _date_io_dayjs__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__nccwpck_require__.n(_date_io_dayjs__WEBPACK_IMPORTED_MODULE_1__);


console.log(trim__WEBPACK_IMPORTED_MODULE_0___default()('    hello world    '));
console.log(process.env);
const dateFns = new (_date_io_dayjs__WEBPACK_IMPORTED_MODULE_1___default())();
const date = dateFns.date("2021/04/06");
console.log(dateFns.format(date, "fullDateTime24h"));


/***/ }),

/***/ 112:
/***/ ((module) => {

module.exports = eval("require")("dayjs");


/***/ }),

/***/ 148:
/***/ ((module) => {

module.exports = eval("require")("dayjs/plugin/customParseFormat");


/***/ }),

/***/ 987:
/***/ ((module) => {

module.exports = eval("require")("dayjs/plugin/isBetween");


/***/ }),

/***/ 442:
/***/ ((module) => {

module.exports = eval("require")("dayjs/plugin/localizedFormat");


/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __nccwpck_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		if(__webpack_module_cache__[moduleId]) {
/******/ 			return __webpack_module_cache__[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			// no module.id needed
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		var threw = true;
/******/ 		try {
/******/ 			__webpack_modules__[moduleId](module, module.exports, __nccwpck_require__);
/******/ 			threw = false;
/******/ 		} finally {
/******/ 			if(threw) delete __webpack_module_cache__[moduleId];
/******/ 		}
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/compat get default export */
/******/ 	(() => {
/******/ 		// getDefaultExport function for compatibility with non-harmony modules
/******/ 		__nccwpck_require__.n = (module) => {
/******/ 			var getter = module && module.__esModule ?
/******/ 				() => module['default'] :
/******/ 				() => module;
/******/ 			__nccwpck_require__.d(getter, { a: getter });
/******/ 			return getter;
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/define property getters */
/******/ 	(() => {
/******/ 		// define getter functions for harmony exports
/******/ 		__nccwpck_require__.d = (exports, definition) => {
/******/ 			for(var key in definition) {
/******/ 				if(__nccwpck_require__.o(definition, key) && !__nccwpck_require__.o(exports, key)) {
/******/ 					Object.defineProperty(exports, key, { enumerable: true, get: definition[key] });
/******/ 				}
/******/ 			}
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/hasOwnProperty shorthand */
/******/ 	(() => {
/******/ 		__nccwpck_require__.o = (obj, prop) => Object.prototype.hasOwnProperty.call(obj, prop)
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__nccwpck_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/compat */
/******/ 	
/******/ 	__nccwpck_require__.ab = __dirname + "/";/************************************************************************/
/******/ 	// module exports must be returned from runtime so entry inlining is disabled
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	return __nccwpck_require__(399);
/******/ })()
;
//# sourceMappingURL=index.js.map