// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useState } from 'react';
import { Link } from 'react-router-dom';
import { useDispatch } from 'react-redux';


import { registerUser } from '@/app/store/actions/users';
import { RouteConfig } from '@/app/router';

import { Validator } from '@/user/validation';

import { UserDataArea } from '@components/common/UserDataArea';

import ultimate from '@static/img/registerPage/ultimate.svg';

import './index.scss';

const SignUp: React.FC = () => {
    const dispatch = useDispatch();
    /** controlled values for form inputs */
    const [firstName, setFirstName] = useState('');
    const [firstNameError, setFirstNameError]
        = useState<SetStateAction<null | string>>(null);
    const [lastName, setLastName] = useState('');
    const [lastNameError, setLastNameError]
        = useState<SetStateAction<null | string>>(null);
    const [email, setEmail] = useState('');
    const [emailError, setEmailError]
        = useState<SetStateAction<null | string>>(null);
    const [password, setPassword] = useState('');
    const [passwordError, setPasswordError]
        = useState<SetStateAction<null | string>>(null);
    const [nickName, setNickName] = useState('');
    const [nickNameError, setNickNameError]
        = useState<SetStateAction<null | string>>(null);
    /** checks if values does't valid then set an error messages */
    const validateForm: () => boolean = () => {
        let isValidForm = true;

        if (!Validator.email(email)) {
            setEmailError('Email is not valid');
            isValidForm = false;
        };

        if (!Validator.password(password)) {
            setPasswordError('Password is not valid');
            isValidForm = false;
        };

        if (!Validator.field(lastName)) {
            setLastNameError('LastName is not valid');
            isValidForm = false;
        };

        if (!Validator.field(firstName)) {
            setFirstNameError('FirstName is not valid');
            isValidForm = false;
        };

        if (!Validator.field(nickName)) {
            setNickNameError('NickName is not valid');
            isValidForm = false;
        };

        return isValidForm;
    };

    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };

        try {
            await dispatch(registerUser({
                email,
                password,
                nickName,
                firstName,
                lastName,
            }));
            location.pathname = RouteConfig.SignIn.path;
        } catch (error) {
            /** TODO: it will be reworked with notification system */
        };
    };

    /** user datas for registration */
    const signUpDatas = [
        {
            value: firstName,
            placeHolder: 'Name',
            onChange: setFirstName,
            className: 'register__sign-up__sign-form__name',
            type: 'text',
            error: firstNameError,
            clearError: setFirstNameError,
            validate: Validator.field,
        },
        {
            value: lastName,
            placeHolder: 'Surname',
            onChange: setLastName,
            className: 'register__sign-up__sign-form__surname',
            type: 'text',
            error: lastNameError,
            clearError: setLastNameError,
            validate: Validator.field,
        },
        {
            value: email,
            placeHolder: 'E-mail',
            onChange: setEmail,
            className: 'register__sign-up__sign-form__email',
            type: 'email',
            error: emailError,
            clearError: setEmailError,
            validate: Validator.email,
        },
        {
            value: password,
            placeHolder: 'Password',
            onChange: setPassword,
            className: 'register__sign-up__sign-form__password',
            type: 'password',
            error: passwordError,
            clearError: setPasswordError,
            validate: Validator.password,
        },
        {
            value: nickName,
            placeHolder: 'Nickname',
            onChange: setNickName,
            className: 'register__sign-up__sign-form__name',
            type: 'text',
            error: nickNameError,
            clearError: setNickNameError,
            validate: Validator.field,
        },
    ];

    return (
        <div className="register">
            <div className="register__represent">
                <img
                    src={ultimate}
                    alt="utlimate division logo"
                    className="register__represent__ultimate"
                />
            </div>
            <div className="register__sign-up">
                <h1 className="register__sign-up__title">SIGN UP</h1>
                <div className="register__sign-up__description">
                    <h2 className="register__sign-up__description__title">
                        Hello!
                    </h2>
                    <p className="register__sign-up__description__information">
                        Sign up to get access to incredible
                        emotions with Ultimate Division
                    </p>
                </div>
                <form
                    className="register__sign-up__sign-form"
                    onSubmit={handleSubmit}
                >
                    {signUpDatas.map((data, index) => <UserDataArea
                        key={index}
                        value={data.value}
                        placeHolder={data.placeHolder}
                        onChange={data.onChange}
                        className={data.className}
                        type={data.type}
                        error={data.error}
                        clearError={data.clearError}
                        validate={data.validate}
                    />)}
                    <input
                        className="register__sign-up__sign-form__confirm"
                        value="SIGN UP"
                        type="submit"
                    />
                </form>

                <div className="register__sign-up__description">
                    <p className="register__sign-up__description__information">
                        Already have an account?
                        <Link
                            className="register__sign-up__description__information__sign"
                            to={RouteConfig.SignIn.path}
                        >
                            Sign in
                        </Link>
                    </p>
                </div>
            </div >
        </div>
    );
};

export default SignUp;
