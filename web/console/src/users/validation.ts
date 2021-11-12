// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** implementation of user auth validation */
export class Validator {
    /** static method for email field validation */
    static isEmail(email: string): boolean {
        const re = new RegExp(/^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/, 'i');

        if (!email) {
            return false;
        };

        return re.test(String(email).toLowerCase());
    };
    /** static method for password field validation */
    static isPassword(password: string): boolean {
        /** same validation from back-end:
         * min 8 letter password, with at least a symbol,
         * upper and lower case letters and a number */
        const re = new RegExp(/^(?=.*\d)(?=.*[!@#$%^&*])(?=.*[a-z])(?=.*[A-Z]).{8,}$/, 'g');

        return re.test(password);
    };
    /** static method for all string form fields validation,
     * except password and email */
    static isName(name: string): boolean {
        return !!name;
    };
};
