"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
/* eslint-disable @typescript-eslint/naming-convention */
const coins_1 = require("./coins");
const encoding_1 = require("./encoding");
const testutils_spec_1 = require("./testutils.spec");
describe("encoding", () => {
    describe("sortedJsonStringify", () => {
        it("leaves non-objects unchanged", () => {
            expect(encoding_1.sortedJsonStringify(true)).toEqual(`true`);
            expect(encoding_1.sortedJsonStringify(false)).toEqual(`false`);
            expect(encoding_1.sortedJsonStringify("aabbccdd")).toEqual(`"aabbccdd"`);
            expect(encoding_1.sortedJsonStringify(75)).toEqual(`75`);
            expect(encoding_1.sortedJsonStringify(null)).toEqual(`null`);
            expect(encoding_1.sortedJsonStringify([5, 6, 7, 1])).toEqual(`[5,6,7,1]`);
            expect(encoding_1.sortedJsonStringify([5, ["a", "b"], true, null, 1])).toEqual(`[5,["a","b"],true,null,1]`);
        });
        it("sorts objects by key", () => {
            // already sorted
            expect(encoding_1.sortedJsonStringify({})).toEqual(`{}`);
            expect(encoding_1.sortedJsonStringify({ a: 3 })).toEqual(`{"a":3}`);
            expect(encoding_1.sortedJsonStringify({ a: 3, b: 2, c: 1 })).toEqual(`{"a":3,"b":2,"c":1}`);
            // not yet sorted
            expect(encoding_1.sortedJsonStringify({ b: 2, a: 3, c: 1 })).toEqual(`{"a":3,"b":2,"c":1}`);
            expect(encoding_1.sortedJsonStringify({ aaa: true, aa: true, a: true })).toEqual(`{"a":true,"aa":true,"aaa":true}`);
        });
        it("sorts nested objects", () => {
            // already sorted
            expect(encoding_1.sortedJsonStringify({ x: { y: { z: null } } })).toEqual(`{"x":{"y":{"z":null}}}`);
            // not yet sorted
            expect(encoding_1.sortedJsonStringify({ b: { z: true, x: true, y: true }, a: true, c: true })).toEqual(`{"a":true,"b":{"x":true,"y":true,"z":true},"c":true}`);
        });
        it("sorts objects in arrays", () => {
            // already sorted
            expect(encoding_1.sortedJsonStringify([1, 2, { x: { y: { z: null } } }, 4])).toEqual(`[1,2,{"x":{"y":{"z":null}}},4]`);
            // not yet sorted
            expect(encoding_1.sortedJsonStringify([1, 2, { b: { z: true, x: true, y: true }, a: true, c: true }, 4])).toEqual(`[1,2,{"a":true,"b":{"x":true,"y":true,"z":true},"c":true},4]`);
        });
    });
    describe("makeSignDoc", () => {
        it("works", () => {
            const chainId = "testspace-12";
            const msg1 = {
                type: "cosmos-sdk/MsgDelegate",
                value: {
                    delegator_address: testutils_spec_1.faucet.address0,
                    validator_address: testutils_spec_1.launchpad.validator.address,
                    amount: coins_1.coin(1234, "ustake"),
                },
            };
            const msg2 = {
                type: "cosmos-sdk/MsgSend",
                value: {
                    from_address: testutils_spec_1.faucet.address0,
                    to_address: testutils_spec_1.makeRandomAddress(),
                    amount: coins_1.coins(1234567, "ucosm"),
                },
            };
            const fee = {
                amount: coins_1.coins(2000, "ucosm"),
                gas: "180000",
            };
            const memo = "Use your power wisely";
            const accountNumber = 15;
            const sequence = 16;
            const signDoc = encoding_1.makeSignDoc([msg1, msg2], fee, chainId, memo, accountNumber, sequence);
            expect(signDoc).toEqual({
                msgs: [msg1, msg2],
                fee: fee,
                chain_id: chainId,
                account_number: accountNumber.toString(),
                sequence: sequence.toString(),
                memo: memo,
            });
        });
        it("works with undefined memo", () => {
            const chainId = "testspace-12";
            const msg1 = {
                type: "cosmos-sdk/MsgDelegate",
                value: {
                    delegator_address: testutils_spec_1.faucet.address0,
                    validator_address: testutils_spec_1.launchpad.validator.address,
                    amount: coins_1.coin(1234, "ustake"),
                },
            };
            const msg2 = {
                type: "cosmos-sdk/MsgSend",
                value: {
                    from_address: testutils_spec_1.faucet.address0,
                    to_address: testutils_spec_1.makeRandomAddress(),
                    amount: coins_1.coins(1234567, "ucosm"),
                },
            };
            const fee = {
                amount: coins_1.coins(2000, "ucosm"),
                gas: "180000",
            };
            const accountNumber = 15;
            const sequence = 16;
            const signDoc = encoding_1.makeSignDoc([msg1, msg2], fee, chainId, undefined, accountNumber, sequence);
            expect(signDoc).toEqual({
                msgs: [msg1, msg2],
                fee: fee,
                chain_id: chainId,
                account_number: accountNumber.toString(),
                sequence: sequence.toString(),
                memo: "",
            });
        });
    });
});
//# sourceMappingURL=encoding.spec.js.map