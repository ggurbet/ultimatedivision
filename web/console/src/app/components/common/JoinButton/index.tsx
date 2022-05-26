// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Link } from 'react-router-dom';

import { AuthRouteConfig, RouteConfig } from '@/app/routes';
import { useLocalStorage } from '@/app/hooks/useLocalStorage';

import './index.scss';

export const JoinButton: React.FC = () => {
    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();

    /* Boolean value from localstorge that indicates whether the user is logged in or not. */
    // @ts-ignore .
    const isLoggined = JSON.parse(getLocalStorageItem('IS_LOGGINED'));

    return (
        <Link
            className="ultimatedivision-join-btn"
            to={isLoggined ? RouteConfig.MarketPlace.path : AuthRouteConfig.SignIn.path}
        >
            <button className="ultimatedivision-join-btn">
                <span className="ultimatedivision-join-btn__text">
                    JOIN BETA
                </span>
            </button>
        </Link>
    );
};
