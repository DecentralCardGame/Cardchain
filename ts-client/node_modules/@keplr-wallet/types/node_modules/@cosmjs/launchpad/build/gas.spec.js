"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const math_1 = require("@cosmjs/math");
const gas_1 = require("./gas");
describe("GasPrice", () => {
    it("can be constructed", () => {
        const inputs = ["3.14", "3", "0.14"];
        inputs.forEach((input) => {
            const gasPrice = new gas_1.GasPrice(math_1.Decimal.fromUserInput(input, 18), "utest");
            expect(gasPrice.amount.toString()).toEqual(input);
            expect(gasPrice.denom).toEqual("utest");
        });
    });
    it("can be constructed from a config string", () => {
        const inputs = ["3.14", "3", "0.14"];
        inputs.forEach((input) => {
            const gasPrice = gas_1.GasPrice.fromString(`${input}utest`);
            expect(gasPrice.amount.toString()).toEqual(input);
            expect(gasPrice.denom).toEqual("utest");
        });
    });
});
//# sourceMappingURL=gas.spec.js.map