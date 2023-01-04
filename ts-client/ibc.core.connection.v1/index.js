"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.registry = exports.queryClient = exports.txClient = exports.msgTypes = exports.Module = void 0;
const module_1 = require("./module");
exports.Module = module_1.default;
const module_2 = require("./module");
Object.defineProperty(exports, "txClient", { enumerable: true, get: function () { return module_2.txClient; } });
Object.defineProperty(exports, "queryClient", { enumerable: true, get: function () { return module_2.queryClient; } });
Object.defineProperty(exports, "registry", { enumerable: true, get: function () { return module_2.registry; } });
const registry_1 = require("./registry");
Object.defineProperty(exports, "msgTypes", { enumerable: true, get: function () { return registry_1.msgTypes; } });
__exportStar(require("./types"), exports);
