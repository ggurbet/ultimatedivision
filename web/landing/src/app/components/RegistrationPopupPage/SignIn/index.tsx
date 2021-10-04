// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useState } from 'react';
import { useDispatch } from 'react-redux';

import { loginUser } from '@/app/store/actions/users';

import { Validator } from '@/user/validation';

import { UserDataArea } from '@components/common/UserDataArea';

import ultimate from '@static/images/registerPage/ultimate.svg';

export const SignIn: React.FC<{
    handleResetPassword: any,
    handleSignUp: any
}> = ({ handleResetPassword, handleSignUp }) => {
    const dispatch = useDispatch();
    /** controlled values for form inputs */
    const [email, setEmail] = useState('');
    const [emailError, setEmailError] = useState<SetStateAction<null | string>>(null);
    const [password, setPassword] = useState('');
    const [passwordError, setPasswordError]
        = useState<SetStateAction<null | string>>(null);
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
            handleChange: handleIsRemember,
            className: 'register__sign-in__sign-form__remember-me',
            type: 'radio',
            error: null,
            clearError: null,
            validate: () => false,
        },
    ];

    return <div className="register">
        <div className="register__represent">
            <img
                alt="utlimate division logo"
                src={ultimate}
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
                        <span
                            onClick={() => handleResetPassword()}
                            className="register__sign-in__sign-form__forgot-password"
                        >
                            Forgot Password?
                        </span>
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
                <input
                    className="register__sign-in__sign-form__confirm"
                    value="SIGN IN"
                    type="submit"
                />
            </form>
            <div className="register__sign-in__description">
                <p className="register__sign-up__description__information">
                    Don't have an account?
                    <span
                        className="register__sign-in__description__information__sign"
                        onClick={() => handleSignUp()}
                    >
                        sign up
                    </span>
                </p>
            </div>
        </div >
    </div>;
};
