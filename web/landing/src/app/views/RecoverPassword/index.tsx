// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';

import { UserClient } from '@/api/user';
import { UserService } from '@/user/service';
import { Validator } from '@/user/validation';

import { useQueryToken } from '@/app/hooks/useQueryToken';

import { recoverUserPassword } from '@/app/store/actions/users';

import { UserDataArea } from '@components/common/UserDataArea';

import ultimate from '@static/images/registerPage/ultimate.svg';

import './index.scss';


const RecoverPassword: React.FC = () => {
    useEffect(() => {
        checkRecoverToken();
    }, []);
    const dispatch = useDispatch();
    const token = useQueryToken();

    const [errorMessage, setErrorMessage] =
        useState<SetStateAction<null | string>>(null);

    const userClient = new UserClient();
    const users = new UserService(userClient);
    /** catches error if token is not valid */
    async function checkRecoverToken() {
        try {
            await users.checkRecoverToken(token);
        } catch (error: any) {
            /** TODO: handles error */
            setErrorMessage('Cannot get access');
        };
    };
    /** controlled values for form inputs */
    const [password, setPassword] = useState('');
    const [passwordError, setPasswordError]
        = useState<SetStateAction<null | string>>(null);
    const [confirmedPassword, setConfirmedPassword] = useState('');
    const [confirmedPasswordError, setConfirmedPasswordError]
        = useState<SetStateAction<null | string>>(null);
    /** checks if values does't valid then set an error messages */
    const validateForm: () => boolean = () => {
        let isValidForm = true;

        if (!Validator.password(password)) {
            setPasswordError('Password is not valid');
            isValidForm = false;
        };

        if (!Validator.password(confirmedPassword)) {
            setConfirmedPasswordError('Confirmed password is not valid');
            isValidForm = false;
        };

        if (password !== confirmedPassword) {
            setConfirmedPasswordError('Passwords does not match, please try again');
            isValidForm = false;
        }

        return isValidForm;
    };

    /** sign in user data */
    const handleSubmit = (e: any) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };

        dispatch(recoverUserPassword(password));

    };
    /** user datas for recover password */
    const passwords = [
        {
            value: password,
            placeHolder: 'Enter a new password',
            handleChange: setPassword,
            className: 'register__recover__sign-form__password',
            type: 'password',
            error: passwordError,
            clearError: setPasswordError,
        },
        {
            value: confirmedPassword,
            placeHolder: 'Enter a new password again',
            handleChange: setConfirmedPassword,
            className: 'register__recover__sign-form__password',
            type: 'password',
            error: confirmedPasswordError,
            clearError: setConfirmedPasswordError,
        },
    ];

    if (errorMessage) {
        return <h1>{errorMessage}</h1>
    };

    return (
        <div className="register">
            <div className="register__represent-reset">
                <img
                    src={ultimate}
                    alt="utlimate division logo"
                    className="register__represent-reset__ultimate-recover"
                />
            </div>
            <div className="register__recover">
                <h1 className="register__recover__title">RECOVER PASSWORD</h1>
                <form
                    className="register__recover__sign-form"
                    onSubmit={handleSubmit}
                >
                    {passwords.map((password, index) => {
                        return <UserDataArea
                            key={index}
                            value={password.value}
                            placeHolder={password.placeHolder}
                            handleChange={password.handleChange}
                            className={password.className}
                            type={password.type}
                            error={password.error}
                            clearError={password.clearError}
                        />;
                    })}
                    <input
                        className="register__recover__sign-form__confirm"
                        value="RECOVER PASSWORD"
                        type="submit"
                    />
                </form>
            </div >
        </div>
    );
};

export default RecoverPassword;