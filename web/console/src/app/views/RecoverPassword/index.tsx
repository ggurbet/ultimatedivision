// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';
import { Dispatch } from 'redux';

import { UserClient } from '@/api/user';
import { UserService } from '@/user/service';
import { Validator } from '@/user/validation';
import { RouteConfig } from '@/app/router';

import { useQueryToken } from '@/app/hooks/useQueryToken';

import { recoverPassword } from '@/app/store/actions/users';

import { UserDataArea } from '@components/common/UserDataArea';

import ultimate from '@static/img/registerPage/ultimate.svg';

import './index.scss';

const RecoverPassword: React.FC = () => {
    const dispatch = useDispatch();
    const token = useQueryToken();

    const [errorMessage, setErrorMessage]
        = useState<SetStateAction<null | string>>(null);

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
    useEffect(() => {
        checkRecoverToken();
    }, []);

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
    /** implements recover of user password */
    const recoverUserPassword = (password: string) =>
        async function(dispatch: Dispatch) {
            try {
                await users.recoverPassword(password);
                dispatch(recoverPassword(password));
                location.pathname = RouteConfig.SignIn.path;
            } catch (error: any) {
                /** TODO: it will be reworked with notification system */
            }
        };
    /** sign in user data */
    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };

        try {
            await dispatch(recoverUserPassword(password));
            location.pathname = RouteConfig.SignIn.path;
        } catch (error) {
            /** TODO: it will be reworked with notification system */
        };
    };
    /** user datas for recover password */
    const passwords = [
        {
            value: password,
            placeHolder: 'Enter a new password',
            onChange: setPassword,
            className: 'register__recover__sign-form__password',
            type: 'password',
            error: passwordError,
            clearError: setPasswordError,
            validate: Validator.password,
        },
        {
            value: confirmedPassword,
            placeHolder: 'Enter a new password again',
            onChange: setConfirmedPassword,
            className: 'register__recover__sign-form__password',
            type: 'password',
            error: confirmedPasswordError,
            clearError: setConfirmedPasswordError,
            validate: Validator.password,
        },
    ];

    if (errorMessage) {
        return <h1>{errorMessage}</h1>;
    };

    return (
        <div className="register">
            <div className="register__represent-reset">
                <img
                    src={ultimate}
                    alt="utlimate division logo"
                    className="register__represent-reset__ultimate"
                />
            </div>
            <div className="register__recover">
                <h1 className="register__recover__title">Recover password</h1>
                <form
                    className="register__recover__sign-form"
                    onSubmit={handleSubmit}
                >
                    {passwords.map((password, index) => <UserDataArea
                        key={index}
                        value={password.value}
                        placeHolder={password.placeHolder}
                        onChange={password.onChange}
                        className={password.className}
                        type={password.type}
                        error={password.error}
                        clearError={password.clearError}
                        validate={password.validate}
                    />)}
                    <input
                        className="register__recover__sign-form__confirm"
                        value="RECOVER YOUR PASSWORD"
                        type="submit"
                    />
                </form>
            </div >
        </div>
    );
};

export default RecoverPassword;
