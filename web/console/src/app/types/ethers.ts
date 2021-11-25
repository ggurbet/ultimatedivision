// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Class desctibes parameters for transaction */
export class TransactionIdentificators {
    /** includes wallet address, and card id */
    constructor(
        public walletAddress: string,
        public cardId: string
    ) {}
}
