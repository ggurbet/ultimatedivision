// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useState } from 'react';
import { useDispatch } from 'react-redux';
import { toast } from 'react-toastify';

import { UserDataArea } from '@components/common/UserDataArea';

import ultimate from '@static/img/registerPage/ultimate.svg';

import { BadRequestError } from '@/api';
import { registerUser } from '@/app/store/actions/users';
import { Validator } from '@/users/validation';

// TODO: it will be reworked on wrapper with children props.
export const SignUp: React.FC<{ showSignUpComponent: () => void }> = ({
    showSignUpComponent,
}) => {
    const dispatch = useDispatch();

    /** Controlled form values. */
    const [firstName, setFirstName] = useState('');
    const [firstNameError, setFirstNameError] = useState<SetStateAction<null | string>>(null);
    const [lastName, setLastName] = useState('');
    const [lastNameError, setLastNameError] = useState<SetStateAction<null | string>>(null);
    const [email, setEmail] = useState('');
    const [emailError, setEmailError] = useState<SetStateAction<null | string>>(null);
    const [password, setPassword] = useState('');
    const [passwordError, setPasswordError] = useState<SetStateAction<null | string>>(null);
    const [nickName, setNickName] = useState('');
    const [nickNameError, setNickNameError] = useState<SetStateAction<null | string>>(null);

    /** Checks if values does't valid. */
    const validateForm: () => boolean = () => {
        let isFormValid = true;

        if (!Validator.isEmail(email)) {
            setEmailError('Email is not valid');
            isFormValid = false;
        };

        if (!Validator.isPassword(password)) {
            setPasswordError('Password is not valid');
            isFormValid = false;
        };

        if (!Validator.isName(lastName)) {
            setLastNameError('LastName is not valid');
            isFormValid = false;
        };

        if (!Validator.isName(firstName)) {
            setFirstNameError('FirstName is not valid');
            isFormValid = false;
        };

        if (!Validator.isNickName(nickName)) {
            setNickNameError('NickName is not valid');
            isFormValid = false;
        };

        return isFormValid;
    };

    /** Submits form values. */
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
            toast.success('Successfully! Please, check your mail box and confirm email.', {
                position: toast.POSITION.TOP_RIGHT,
            });
        } catch (error: any) {
            let errorMessage = 'Someting wrong, please, try again.';
            if (error instanceof BadRequestError) {
                errorMessage = 'Email is already in use.';
            };
            toast.error(errorMessage, {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        };
    };

    /** Exposes form values. */
    const formValues = [
        {
            value: firstName,
            placeHolder: 'Name',
            onChange: setFirstName,
            className: 'register__sign-up__sign-form__name',
            type: 'text',
            error: firstNameError,
            clearError: setFirstNameError,
            validate: Validator.isName,
        },
        {
            value: lastName,
            placeHolder: 'Surname',
            onChange: setLastName,
            className: 'register__sign-up__sign-form__surname',
            type: 'text',
            error: lastNameError,
            clearError: setLastNameError,
            validate: Validator.isName,
        },
        {
            value: email,
            placeHolder: 'E-mail',
            onChange: setEmail,
            className: 'register__sign-up__sign-form__email',
            type: 'text',
            error: emailError,
            clearError: setEmailError,
            validate: Validator.isEmail,
        },
        {
            value: password,
            placeHolder: 'Password',
            onChange: setPassword,
            className: 'register__sign-up__sign-form__password',
            type: 'password',
            error: passwordError,
            clearError: setPasswordError,
            validate: Validator.isPassword,
        },
        {
            value: nickName,
            placeHolder: 'Nickname',
            onChange: setNickName,
            className: 'register__sign-up__sign-form__name',
            type: 'text',
            error: nickNameError,
            clearError: setNickNameError,
            validate: Validator.isNickName,
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
                        Sign up to get access to incredible emotions
                        with Ultimate Division
                    </p>
                </div>
                <form
                    className="register__sign-in__sign-form"
                    onSubmit={handleSubmit}
                >
                    {formValues.map((formValue, index) => <UserDataArea
                        key={index}
                        value={formValue.value}
                        placeHolder={formValue.placeHolder}
                        onChange={formValue.onChange}
                        className={formValue.className}
                        type={formValue.type}
                        error={formValue.error}
                        clearError={formValue.clearError}
                        validate={formValue.validate}
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
                        <span
                            className="register__sign-up__description__information__sign"
                            onClick={showSignUpComponent}
                        >
                            Sign in
                        </span>
                    </p>
                </div>
            </div >
        </div>
    );
};
