// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** User describes user domain entity. */
export class User {
    /** User domain casperWallet, email, id, lastLogin, nickName, registerData, wallet, walletType. */
    public constructor(
        public casperWalletId: string = '',
        public casperWallet: string = '',
        public email: string = '',
        public id: string = '',
        public lastLogin: string = '',
        public nickname: string = '',
        public registerDate: string = '',
        public wallet: string = '',
        public walletType: string = ''
    ) { };
};
