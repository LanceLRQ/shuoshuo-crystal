const path = require('path');
const { ESLINT_MODES } = require("@craco/craco");
module.exports = {
  eslint: {
    mode: ESLINT_MODES.file
  },
  webpack: {
    alias: {
      "@": path.resolve(__dirname, "src"),
      "@styles": path.resolve(__dirname, "src/styles"),
      "@images": path.resolve(__dirname, "src/images"),
      "@apps": path.resolve(__dirname, "src/apps"),
    }
  },
  jest: {
    configure: {
      moduleNameMapper: {
        "@": "<rootDir>/src/$1",
        "@styles": "<rootDir>/src/styles/$1",
        "@images": "<rootDir>/src/images/$1",
        "@apps": "<rootDir>/src/apps/$1",
      }
    }
  }
}
