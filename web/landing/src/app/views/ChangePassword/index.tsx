// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useState } from 'react';
import { Link } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { useHistory } from 'react-router';

import { changeUserPassword } from '@/app/store/actions/users';
import { RouteConfig } from '@/app/router';
import { Validator } from '@/user/validation';

import { UserDataArea } from '@components/common/UserDataArea';

import ultimate from '@static/images/registerPage/ultimate.svg';
import goBack from '@static/images/registerPage/goback.svg';

import './index.scss';

const ChangePassword: React.FC = () => {
    const dispatch = useDispatch();
    const browserHistory = useHistory();
    /** controlled values for form inputs */
    const [password, setPassword] = useState('');
    const [passwordError, setPasswordError]
        = useState<SetStateAction<null | string>>(null);
    const [newPassword, setNewPassword] = useState('');
    const [newPasswordError, setNewPasswordError]
        = useState<SetStateAction<null | string>>(null);
    /** checks if values does't valid then set an error messages */
    const validateForm: () => boolean = () => {
        let isValidForm = true;

        if (!Validator.password(password)) {
            setPasswordError('Old password is not valid');
            isValidForm = false;
        };

        if (!Validator.password(newPassword)) {
            setNewPasswordError('New password is not valid');
            isValidForm = false;
        };

        return isValidForm;
    };

    /** sign in user data */
    const handleSubmit = (e: any) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };

        dispatch(changeUserPassword(password, newPassword));
    };
    /** user datas for registration */
    const resetPasswordDatas = [
        {
            value: password,
            placeHolder: 'Enter the old password',
            handleChange: setPassword,
            className: 'register__reset__sign-form__password',
            type: 'password',
            error: passwordError,
            clearError: setPasswordError,
        },
        {
            value: newPassword,
            placeHolder: 'Enter a new password',
            handleChange: setNewPassword,
            className: 'register__reset__sign-form__password',
            type: 'password',
            error: newPasswordError,
            clearError: setNewPasswordError,
        },
    ];

    return (
        <div className="register">
            <div className="register__represent-reset">
                <img
                    src={ultimate}
                    alt="utlimate division logo"
                    className="register__represent-reset__ultimate"
                />
            </div>
            <div className="register__reset">
                <Link
                    className="register__reset__go-back"
                    to={RouteConfig.SignIn.path}>
                    <img
                        alt="go back"
                        src={goBack}
                    />
                    <span className="register__reset__go-back__title">
                        GO BACK
                    </span>
                </Link>
                <h1 className="register__reset__title">Change your password</h1>
                <form
                    className="register__reset__sign-form"
                    onSubmit={handleSubmit}
                >
                    {resetPasswordDatas.map((data, index) => {
                        return <UserDataArea
                            key={index}
                            value={data.value}
                            placeHolder={data.placeHolder}
                            handleChange={data.handleChange}
                            className={data.className}
                            type={data.type}
                            error={data.error}
                            clearError={data.clearError}
                        />;
                    })}
                    <input
                        className="register__reset__sign-form__confirm"
                        value="CHANGE PASSWORD"
                        type="submit"
                    />
                </form>
            </div >
        </div>
    );
};

export default ChangePassword;
