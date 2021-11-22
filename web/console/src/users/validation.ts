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
        const re = new RegExp(/^(?=.*\d)(?=.*[!@#%^<>&_\/\*\;\:\-\^\=\$\.\,\+\(\)\[\]\{\}\?])(?=.*[a-z])(?=.*[A-Z]).{8,}$/, 'g');

        return re.test(password);
    };

    /** static method for firstname and lastname validation */
    static isName(name: string): boolean {
        /** min 2 letter name, consist of uppercase and lowercase letters */
        const re = new RegExp(/^[a-zA-Z]{2,}$/, 'i');

        return re.test(name);
    };

    /** static method for nickname validation */
    static isNickName(nickName: string): boolean {
        /** min 2 letter name, consist of uppercase and lowercase letters
         * and a number */
        const re = new RegExp(/^[a-zA-Z0-9]{2,}$/, 'i');

        return re.test(nickName);
    };
};
