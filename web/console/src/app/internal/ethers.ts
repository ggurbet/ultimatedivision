// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

const FIRST_ELEMENT_INDEX = 0;
/** Creates needed hash by replacing last part of hashBase by value. */
export const buildHash = (hashSignature: number | string) => {
    const HASH_BASE = '0000000000000000000000000000000000000000000000000000000000000000';

    return HASH_BASE.slice(FIRST_ELEMENT_INDEX, HASH_BASE.length - hashSignature.toString().length) + hashSignature;
};
