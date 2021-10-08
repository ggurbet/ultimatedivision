// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** implementation of user auth validation */

export class Validator {
    /** static method for email field validation */
    static email(email: string): boolean {
        const re = new RegExp(/^(([^<>()[\],;:\s@"]+([^<>()[\],;:\s@"]+)*)|(".+"))@(([^<>()[\],;:\s@"]+)+[^<>()[\],;:\s@"]{2,})$/, 'i');

        if (!email) {
            return false;
        };

        return re.test(String(email).toLowerCase());
    };
};
