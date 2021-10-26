// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { SetStateAction, useState } from 'react';

import { Validator } from '@/user/validation';
import { UserClient } from '@/api/user';
import { UserService } from '@/user/service';

import { UserDataArea } from '@components/common/UserDataArea';

import './index.scss';

export const Modal: React.FC<{ handleModal: () => void }> = ({
    handleModal
}) => {
    const [email, setEmail] = useState('');
    const [emailError, setEmailError]
        = useState<SetStateAction<null | string>>(null);
    const [successEmailLabel, setSuccessEmailLabel]
        = useState<SetStateAction<null | string>>(null);
    /** checks if value does't valid then set an error message */
    const validateForm: () => boolean = () => {
        let isValidForm = true;

        if (!Validator.email(email)) {
            setEmailError('Email is not valid');
            isValidForm = false;
        };

        if (!email) {
            setEmailError('Please, enter your email');
            isValidForm = false;
        };

        return isValidForm;
    };

    const user = new UserClient();
    const users = new UserService(user);

    /** describes the logic of user subscription to news */
    async function getNotifications() {
        try {
            await users.getNotifications(email);
            setEmailError(null);
            setSuccessEmailLabel('The email was sent successfully');
        } catch (error: any) {
            setSuccessEmailLabel(null);
            setEmailError(error.message);
        };
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        if (!validateForm()) {
            return;
        };

        getNotifications();
    };

    const clearLabel = () => {
        setEmailError(null);
        setSuccessEmailLabel(null);
    };

    const formValue = {
        value: email,
        placeHolder: 'Your Email',
        onChange: setEmail,
        className: 'launch-date-modal__notification__email',
        type: 'text',
        error: emailError,
        clearLabel,
        validate: Validator.email,
        successLabel: successEmailLabel,
    };

    return <section className="launch-date-modal">
        <div className="launch-date-modal__window">
            <div
                onClick={handleModal}
                className="launch-date-modal__close"
            >
                <p className="launch-date-modal__close__text">
                    &#215;
                </p>
            </div>
            <h1 className="launch-date-modal__description">
                Get notified on the launch
            </h1>
            <div>
                <form
                    className="launch-date-modal__notification"
                    onSubmit={handleSubmit}
                >
                    <div className="launch-date-modal__email-wrapper">
                        <label
                            htmlFor={formValue.value}
                            className="launch-date-modal__email-label"
                        >
                            Email
                        </label>
                        <UserDataArea {...formValue} />
                    </div>
                    <input
                        value="SEND"
                        className="launch-date-modal__notification__confirm"
                        type="submit"
                    />
                </form>
            </div>
            <div className="launch-date-modal__wrapper"
            />
        </div>
    </section>;
};
