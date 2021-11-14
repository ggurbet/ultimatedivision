// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from 'react';
import { useHistory } from 'react-router';
import { toast } from 'react-toastify';

import { UsersClient } from '@/api/users';
import { useQueryToken } from '@/app/hooks/useQueryToken';
import { AuthRouteConfig } from '@/app/routes';
import { UsersService } from '@/users/service';

const ConfirmEmail: React.FC = () => {
    const token = useQueryToken();
    const history = useHistory();

    const usersClient = new UsersClient();
    const usersService = new UsersService(usersClient);

    /** DELAY is the delay time in milliseconds for redirect to SignIn page. */
    const DELAY: number = 3000;

    /** catches error if token is not valid */
    async function checkEmailToken() {
        try {
            await usersService.checkEmailToken(token);
            toast.success(`Your email has been successfully verified.
            You will be redirected to the sign-in page in 3 seconds.`, {
                position: toast.POSITION.TOP_RIGHT,
            });
            await setTimeout(() => {
                history.push(AuthRouteConfig.SignIn.path);
            }, DELAY);
        } catch (error: any) {
            toast.error('Email verification failed', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        };
    };

    useEffect(() => {
        checkEmailToken();
    }, []);

    return <div className="confirm-email"/>;
};

export default ConfirmEmail;
