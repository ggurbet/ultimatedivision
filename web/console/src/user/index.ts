// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** user domain entity implementation */
export class User {
    /** base user domain entiry constructor */
    public constructor(
        public email: string,
        public password: string,
        public nickName: string,
        public firstName: string,
        public lastName: string,
    ) { };
};
