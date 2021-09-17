// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useEffect, useState } from 'react';

import { UserClient } from '@/api/user';
import { UserService } from '@/user/service';

import { useQueryToken } from '@/app/hooks/useQueryToken';
import { RouteConfig } from '@/app/router';

/** TODO: Rework this view after design solution */
const ConfirmEmail: React.FC = () => {
    useEffect(() => {
        checkEmailToken();
    }, []);
    const token = useQueryToken();

    const [errorMessage, setErrorMessage] =
        useState<SetStateAction<null | string>>(null);

    const userClient = new UserClient();
    const users = new UserService(userClient);
    /** catches error if token is not valid */
    async function checkEmailToken() {
        try {
            await users.checkEmailToken(token);
            await setTimeout(() => {
                location.pathname = RouteConfig.SignIn.path;
            }, 3000);
        } catch (error: any) {
            /** TODO: handles error */
            setErrorMessage('Email verification failed');
        };
    };

    if (errorMessage) {
        return <h1>{errorMessage}</h1>
    };

    return <div>
        <h1>
            Your email has been successfully verified.
            You will be redirected to the sign-in page in 3 seconds.
        </h1>
    </div>;
};

export default ConfirmEmail;