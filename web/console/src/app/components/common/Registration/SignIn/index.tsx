// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import MetaMaskOnboarding from '@metamask/onboarding';
import { useMemo, SetStateAction, useState } from 'react';
import { useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { toast } from 'react-toastify';

import { UserDataArea } from '@components/common/UserDataArea';

import facebook from '@static/img/registerPage/facebook_logo.svg';
import google from '@static/img/registerPage/google_logo.svg';
import metamask from '@static/img/registerPage/metamask.svg';
import ultimate from '@static/img/registerPage/ultimate.svg';

import { useLocalStorage } from '@/app/hooks/useLocalStorage';
import { RouteConfig } from '@/app/routes';
import { ServicePlugin } from '@/app/plugins/service';
import { loginUser } from '@/app/store/actions/users';
import { Validator } from '@/users/validation';

// TODO: it will be reworked on wrapper with children props.
export const SignIn: React.FC<{
    showResetPasswordComponent: () => void;
    showSignUpComponent: () => void;
}> = ({ showResetPasswordComponent, showSignUpComponent }) => {
    const onboarding = useMemo(() => new MetaMaskOnboarding(), []);

    /** Creates ethers provider. */
    const service = ServicePlugin.create();

    const dispatch = useDispatch();
    const history = useHistory();

    /** Controlled values for form inputs */
    const [email, setEmail] = useState('');
    const [emailError, setEmailError] =
        useState<SetStateAction<null | string>>(null);
    const [password, setPassword] = useState('');
    const [passwordError, setPasswordError] =
        useState<SetStateAction<null | string>>(null);

    const [isRemember, setIsRemember] = useState(false);

    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();

    /** TODO: rework remember me implementation  */
    const handleIsRemember = () => setIsRemember((prev) => !prev);

    /** Checks if values does't valid. */
    const validateForm: () => boolean = () => {
        let isFormValid = true;

        if (!Validator.isEmail(email)) {
            setEmailError('Email is not valid');
            isFormValid = false;
        }

        if (!Validator.isPassword(password)) {
            setPasswordError('Password is not valid');
            isFormValid = false;
        }

        return isFormValid;
    };

    /** Submits form values. */
    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        }

        try {
            await dispatch(loginUser(email, password));

            setLocalStorageItem('IS_LOGGINED', true);

            history.push(RouteConfig.MarketPlace.path);
        } catch (error: any) {
            toast.error('Incorrect email or password', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        }
    };

    /** Exposes form values. */
    const formValues = [
        {
            value: email,
            placeHolder: 'E-mail',
            onChange: setEmail,
            className: 'register__sign-in__sign-form__email',
            type: 'email',
            error: emailError,
            clearError: setEmailError,
            validate: Validator.isEmail,
        },
        {
            value: password,
            placeHolder: 'Password',
            onChange: setPassword,
            className: 'register__sign-in__sign-form__password',
            type: 'password',
            error: passwordError,
            clearError: setPasswordError,
            validate: Validator.isPassword,
        },
    ];

    /** Logins with matamask. */
    const metamaskLogin = async() => {
        /** Error code which indicates that 'eth_requestAccounts' already processing. */
        const METAMASK_RPC_ERROR_CODE = -32002;
        if (MetaMaskOnboarding.isMetaMaskInstalled()) {
            try {
                // @ts-ignore
                await window.ethereum.request({
                    method: 'eth_requestAccounts',
                });

                await service.login();

                setLocalStorageItem('IS_LOGGINED', true);

                history.push(RouteConfig.MarketPlace.path);
            } catch (error: any) {
                error.code === METAMASK_RPC_ERROR_CODE
                    ? toast.error('Please open metamask manually!', {
                        position: toast.POSITION.TOP_RIGHT,
                        theme: 'colored',
                    })
                    : toast.error('Something went wrong', {
                        position: toast.POSITION.TOP_RIGHT,
                        theme: 'colored',
                    });
            }
        } else {
            onboarding.startOnboarding();
        }
    };

    return (
        <div className="register">
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
                    {formValues.map((formValue, index) =>
                        <UserDataArea
                            key={index}
                            value={formValue.value}
                            placeHolder={formValue.placeHolder}
                            onChange={formValue.onChange}
                            className={formValue.className}
                            type={formValue.type}
                            error={formValue.error}
                            clearError={formValue.clearError}
                            validate={formValue.validate}
                        />
                    )}
                    <div className="register__sign-in__sign-form__checkbox-wrapper">
                        <input
                            id="register-sign-in-checkbox"
                            className="register__sign-in__sign-form__remember-me"
                            type="checkbox"
                        />
                        <label
                            className="register__sign-in__sign-form__remember-me__text"
                            htmlFor="register-sign-in-checkbox"
                        >
                            Remember me
                        </label>
                        <span
                            onClick={showResetPasswordComponent}
                            className="register__sign-in__sign-form__forgot-password"
                        >
                            Forgot Password?
                        </span>
                    </div>
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
                            <img
                                src={metamask}
                                alt="Metamask logo"
                                className="register__sign-in__sign-form__logos__metamask"
                                onClick={metamaskLogin}
                            />
                        </div>
                    </div>
                </form>
                <div className="register__sign-in__description">
                    <p className="register__sign-up__description__information">
                        Don't have an account?
                        <span
                            className="register__sign-in__description__information__sign"
                            onClick={showSignUpComponent}
                        >
                            Sign up
                        </span>
                    </p>
                </div>
            </div>
        </div>
    );
};
