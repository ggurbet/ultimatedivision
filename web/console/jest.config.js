// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

module.exports = {
    injectGlobals: true,
    /** To test HTML DOM add in spec.[ts | tsx] next line:
     * @jest-environment jsdom. */
    testEnvironment: 'jsdom',
    /**
     * Test enviroment options needed to make manipulation with HTML DOM.
     * Could be extended with such fields as html, url, and userAgent.
     * To make this options visible add in spec.[ts | tsx] next line:
     * @jest-environment-options.
     */
    testEnvironmentOptions: {
        html: `
            <html lang="en">
                <head>
                    <meta name="description" content="test environment options"/>
                    <meta name="gateway-address" content="http://localhost:8089">
                <head/>
            </html>
        `,
    },
    testMatch: ['**/__tests__/**/*.+(ts|tsx|js)', '**/?(*.)+(spec|test).+(ts|tsx|js)'],
    moduleNameMapper: {
        'app/(.*)': '<rootDir>/src/app/$1',
        '^@/(.*)$': '<rootDir>/src/$1',
        '^@static/(.*)$': '<rootDir>/src/app/static/$1',
        '^@components/(.*)$': '<rootDir>/src/app/components/$1',
        '^@views/(.*)$': '<rootDir>/src/app/views/$1',
        '^@store/(.*)$': '<rootDir>/src/app/store/$1',
        '^@utils/(.*)$': '<rootDir>/src/app/utils/$1',
        '\\.(css|less|sass|scss)$': 'identity-obj-proxy'
    },
    'roots': [
        '<rootDir>'
    ],
    transform: {
        '^.+\\.(ts|tsx)?$': 'ts-jest',
        '\\.svg': 'jest-transform-stub'
    },
    moduleFileExtensions: ['ts', 'js', 'tsx', 'jsx', 'json'],
    moduleDirectories: ["node_modules", "bower_components", "src"],
    collectCoverage: true,
    clearMocks: true,
    coverageDirectory: "coverage",
};
