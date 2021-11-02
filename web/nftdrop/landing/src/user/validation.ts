// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Implementation of user auth validation. */
export class Validator {
    /** Static method for email field validation. */
    static email(email: string): boolean {
        const re = new RegExp(/^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/, 'i');

        return re.test(String(email).toLowerCase());
    };
};
