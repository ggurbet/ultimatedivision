// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useState } from 'react';
import { Link } from 'react-router-dom';
import { useDispatch } from 'react-redux';

import { RouteConfig } from '@/app/router';

import { Validator } from '@/user/validation';

import { loginUser } from '@/app/store/actions/users';

import { UserDataArea } from '@components/common/UserDataArea';

import facebook from '@static/images/registerPage/facebook_logo.svg';
import google from '@static/images/registerPage/google_logo.svg'
import ultimate from '@static/images/registerPage/ultimate.svg';

import './index.scss';

const SignIn: React.FC = () => {
    const dispatch = useDispatch();
    /** controlled values for form inputs */
    const [email, setEmail] = useState('');
    const [emailError, setEmailError] = useState<SetStateAction<null | string>>(null);
    const [password, setPassword] = useState('');
    const [passwordError, setPasswordError] = useState<SetStateAction<null | string>>(null);
    const [isRemember, setIsRemember] = useState(false);
    /** TODO: rework remember me implementation  */
    const handleIsRemember = () => setIsRemember(prev => !prev);
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

        return isValidForm;
    };
    /** user data that will send to server */
    const handleSubmit = (e: any) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };

        dispatch(loginUser(email, password));
    };
    /** user datas for registration */
    const signInDatas = [
        {
            value: email,
            placeHolder: 'E-mail',
            onChange: setEmail,
            className: 'register__sign-in__sign-form__email',
            type: 'email',
            error: emailError,
            clearError: setEmailError,
            validate: Validator.email,
        },
        {
            value: password,
            placeHolder: 'Password',
            onChange: setPassword,
            className: 'register__sign-in__sign-form__password',
            type: 'password',
            error: passwordError,
            clearError: setPasswordError,
            validate: Validator.password,
        },
        {
            value: 'Remember me',
            placeHolder: '',
            onChange: handleIsRemember,
            className: 'register__sign-in__sign-form__remember-me',
            type: 'radio',
            error: null,
            clearError: null,
            validate: () => false,
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
            <div className="register__sign-in">
                <h1 className="register__sign-in__title">SIGN IN</h1>
                <form
                    className="register__sign-in__sign-form"
                    onSubmit={handleSubmit}
                >
                    {signInDatas.map((data, index) => {
                        return data.type === 'radio' ? <div key={index}>
                            <UserDataArea
                                value={data.value}
                                placeHolder={data.placeHolder}
                                onChange={data.onChange}
                                className={data.className}
                                type={data.type}
                                error={data.error}
                                clearError={data.clearError}
                                validate={data.validate}
                            />
                            <Link
                                to={RouteConfig.ResetPassword.path}
                                className="register__sign-in__sign-form__forgot-password"
                            >
                                Forgot Password?
                            </Link>
                        </div> : <UserDataArea
                            key={index}
                            value={data.value}
                            placeHolder={data.placeHolder}
                            onChange={data.onChange}
                            className={data.className}
                            type={data.type}
                            error={data.error}
                            clearError={data.clearError}
                            validate={data.validate}
                        />;
                    })}
                    <div className="register__sign-in__sign-form__auth-internal">
                        <input
                            className="register__sign-in__sign-form__confirm"
                            value="SIGN IN"
                            type="submit"
                        />
                        or
                        <div className="register__sign-in__sign-form__logos">
                            <img
                                src={google}
                                alt="Google logo"
                                className="register__sign-in__sign-form__logos__google"
                            />
                            <img
                                src={facebook}
                                alt="Facebook logo"
                                className="register__sign-in__sign-form__logos__facebook"
                            />
                        </div>
                    </div>
                </form>
                <div className="register__sign-in__description">
                    <p className="register__sign-in__description__information">
                        Don't have an account?
                        <Link
                            className="register__sign-in__description__information__sign"
                            to={RouteConfig.SignUp.path}
                        >
                            Sign up
                        </Link>
                    </p>
                </div>
            </div >
        </div>
    );
};

export default SignIn;