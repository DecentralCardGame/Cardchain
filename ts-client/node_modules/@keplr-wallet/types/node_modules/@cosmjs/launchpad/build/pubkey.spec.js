"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const encoding_1 = require("@cosmjs/encoding");
const pubkey_1 = require("./pubkey");
describe("pubkey", () => {
    describe("encodeSecp256k1Pubkey", () => {
        it("encodes a compresed pubkey", () => {
            const pubkey = encoding_1.fromBase64("AtQaCqFnshaZQp6rIkvAPyzThvCvXSDO+9AzbxVErqJP");
            expect(pubkey_1.encodeSecp256k1Pubkey(pubkey)).toEqual({
                type: "tendermint/PubKeySecp256k1",
                value: "AtQaCqFnshaZQp6rIkvAPyzThvCvXSDO+9AzbxVErqJP",
            });
        });
        it("throws for uncompressed public keys", () => {
            const pubkey = encoding_1.fromBase64("BE8EGB7ro1ORuFhjOnZcSgwYlpe0DSFjVNUIkNNQxwKQE7WHpoHoNswYeoFkuYpYSKK4mzFzMV/dB0DVAy4lnNU=");
            expect(() => pubkey_1.encodeSecp256k1Pubkey(pubkey)).toThrowError(/public key must be compressed secp256k1/i);
        });
    });
    describe("decodeAminoPubkey", () => {
        it("works for secp256k1", () => {
            const amino = encoding_1.Bech32.decode("cosmospub1addwnpepqd8sgxq7aw348ydctp3n5ajufgxp395hksxjzc6565yfp56scupfqhlgyg5").data;
            expect(pubkey_1.decodeAminoPubkey(amino)).toEqual({
                type: "tendermint/PubKeySecp256k1",
                value: "A08EGB7ro1ORuFhjOnZcSgwYlpe0DSFjVNUIkNNQxwKQ",
            });
        });
        it("works for ed25519", () => {
            // Encoded from `corald tendermint show-validator`
            // Decoded from http://localhost:26657/validators
            const amino = encoding_1.Bech32.decode("coralvalconspub1zcjduepqvxg72ccnl9r65fv0wn3amlk4sfzqfe2k36l073kjx2qyaf6sk23qw7j8wq").data;
            expect(pubkey_1.decodeAminoPubkey(amino)).toEqual({
                type: "tendermint/PubKeyEd25519",
                value: "YZHlYxP5R6olj3Tj3f7VgkQE5VaOvv9G0jKATqdQsqI=",
            });
        });
    });
    describe("decodeBech32Pubkey", () => {
        it("works", () => {
            expect(pubkey_1.decodeBech32Pubkey("cosmospub1addwnpepqd8sgxq7aw348ydctp3n5ajufgxp395hksxjzc6565yfp56scupfqhlgyg5")).toEqual({
                type: "tendermint/PubKeySecp256k1",
                value: "A08EGB7ro1ORuFhjOnZcSgwYlpe0DSFjVNUIkNNQxwKQ",
            });
        });
        it("works for enigma pubkey", () => {
            expect(pubkey_1.decodeBech32Pubkey("enigmapub1addwnpepqw5k9p439nw0zpg2aundx4umwx4nw233z5prpjqjv5anl5grmnchzp2xwvv")).toEqual({
                type: "tendermint/PubKeySecp256k1",
                value: "A6lihrEs3PEFCu8m01ebcas3KjEVAjDIEmU7P9ED3PFx",
            });
        });
        it("works for ed25519", () => {
            // Encoded from `corald tendermint show-validator`
            // Decoded from http://localhost:26657/validators
            const decoded = pubkey_1.decodeBech32Pubkey("coralvalconspub1zcjduepqvxg72ccnl9r65fv0wn3amlk4sfzqfe2k36l073kjx2qyaf6sk23qw7j8wq");
            expect(decoded).toEqual({
                type: "tendermint/PubKeyEd25519",
                value: "YZHlYxP5R6olj3Tj3f7VgkQE5VaOvv9G0jKATqdQsqI=",
            });
        });
    });
    describe("encodeAminoPubkey", () => {
        it("works for secp256k1", () => {
            const pubkey = {
                type: "tendermint/PubKeySecp256k1",
                value: "A08EGB7ro1ORuFhjOnZcSgwYlpe0DSFjVNUIkNNQxwKQ",
            };
            const expected = encoding_1.Bech32.decode("cosmospub1addwnpepqd8sgxq7aw348ydctp3n5ajufgxp395hksxjzc6565yfp56scupfqhlgyg5").data;
            expect(pubkey_1.encodeAminoPubkey(pubkey)).toEqual(expected);
        });
        it("works for ed25519", () => {
            // Decoded from http://localhost:26657/validators
            // Encoded from `corald tendermint show-validator`
            const pubkey = {
                type: "tendermint/PubKeyEd25519",
                value: "YZHlYxP5R6olj3Tj3f7VgkQE5VaOvv9G0jKATqdQsqI=",
            };
            const expected = encoding_1.Bech32.decode("coralvalconspub1zcjduepqvxg72ccnl9r65fv0wn3amlk4sfzqfe2k36l073kjx2qyaf6sk23qw7j8wq").data;
            expect(pubkey_1.encodeAminoPubkey(pubkey)).toEqual(expected);
        });
    });
    describe("encodeBech32Pubkey", () => {
        it("works for secp256k1", () => {
            const pubkey = {
                type: "tendermint/PubKeySecp256k1",
                value: "A08EGB7ro1ORuFhjOnZcSgwYlpe0DSFjVNUIkNNQxwKQ",
            };
            expect(pubkey_1.encodeBech32Pubkey(pubkey, "cosmospub")).toEqual("cosmospub1addwnpepqd8sgxq7aw348ydctp3n5ajufgxp395hksxjzc6565yfp56scupfqhlgyg5");
        });
        it("works for ed25519", () => {
            // Decoded from http://localhost:26657/validators
            // Encoded from `corald tendermint show-validator`
            const pubkey = {
                type: "tendermint/PubKeyEd25519",
                value: "YZHlYxP5R6olj3Tj3f7VgkQE5VaOvv9G0jKATqdQsqI=",
            };
            expect(pubkey_1.encodeBech32Pubkey(pubkey, "coralvalconspub")).toEqual("coralvalconspub1zcjduepqvxg72ccnl9r65fv0wn3amlk4sfzqfe2k36l073kjx2qyaf6sk23qw7j8wq");
        });
    });
});
//# sourceMappingURL=pubkey.spec.js.map