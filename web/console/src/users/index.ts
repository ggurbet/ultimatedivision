// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** User describes user domain entity. */
export class User {
    /** User domain entity contains email, password, nickName, firstName, lastName. */
    public constructor(
        public email: string,
        public password: string,
        public nickName: string,
        public firstName: string,
        public lastName: string,
    ) { };
};
