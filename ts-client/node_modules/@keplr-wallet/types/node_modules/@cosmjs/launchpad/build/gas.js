"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.buildFeeTable = exports.GasPrice = void 0;
const math_1 = require("@cosmjs/math");
const coins_1 = require("./coins");
class GasPrice {
    constructor(amount, denom) {
        this.amount = amount;
        this.denom = denom;
    }
    static fromString(gasPrice) {
        const matchResult = gasPrice.match(/^(?<amount>.+?)(?<denom>[a-z]+)$/);
        if (!matchResult) {
            throw new Error("Invalid gas price string");
        }
        const { amount, denom } = matchResult.groups;
        if (denom.length < 3 || denom.length > 127) {
            throw new Error("Gas price denomination must be between 3 and 127 characters");
        }
        const fractionalDigits = 18;
        const decimalAmount = math_1.Decimal.fromUserInput(amount, fractionalDigits);
        return new GasPrice(decimalAmount, denom);
    }
}
exports.GasPrice = GasPrice;
function calculateFee(gasLimit, { denom, amount: gasPriceAmount }) {
    const amount = Math.ceil(gasPriceAmount.multiply(new math_1.Uint53(gasLimit)).toFloatApproximation());
    return {
        amount: coins_1.coins(amount, denom),
        gas: gasLimit.toString(),
    };
}
function buildFeeTable(gasPrice, defaultGasLimits, gasLimits) {
    return Object.entries(defaultGasLimits).reduce((feeTable, [type, defaultGasLimit]) => (Object.assign(Object.assign({}, feeTable), { [type]: calculateFee(gasLimits[type] || defaultGasLimit, gasPrice) })), {});
}
exports.buildFeeTable = buildFeeTable;
//# sourceMappingURL=gas.js.map