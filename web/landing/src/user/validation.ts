// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** implementation of user auth validation */
export class Validator {
    /** static method for email field validation */
    static email(email: string): boolean {
        const re = new RegExp(/^(([^<>()[\],;:\s@"]+([^<>()[\],;:\s@"]+)*)|(".+"))@(([^<>()[\],;:\s@"]+)+[^<>()[\],;:\s@"]{2,})$/, 'i');

        return re.test(String(email).toLowerCase());
    };
    /** static method for password field validation */
    static password(password: string): boolean {
        /** same validation from back-end:
         * min 8 letter password, with at least a symbol,
         * upper and lower case letters and a number */
        const re = new RegExp(/^(?=.*\d)(?=.*[!@#$%^&*])(?=.*[a-z])(?=.*[A-Z]).{8,}$/, 'i');

        return re.test(String(password).toLowerCase());
    };
    /** static method for all string form fields validation,
     * except password and email */
    static field(field: string): boolean {
        return !!field;
    };
};
