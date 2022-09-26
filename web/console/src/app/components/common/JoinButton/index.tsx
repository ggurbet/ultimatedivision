// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { useHistory } from 'react-router-dom';

import { RouteConfig } from '@/app/routes';
import { useLocalStorage } from '@/app/hooks/useLocalStorage';
import { RegistrationPopup } from '../Registration';

import './index.scss';

export const JoinButton: React.FC = () => {
    const history = useHistory();
    const [setLocalStorageItem, getLocalStorageItem] = useLocalStorage();
    const [isRegistrationRequired, setIsRegistrationRequired] = useState(false);

    /* Boolean value from localstorge that indicates whether the user is logged in or not. */
    // @ts-ignore .
    const isLoggined = JSON.parse(getLocalStorageItem('IS_LOGGINED'));

    /** Closes Registration popup componnet. */
    const closeRegistrationPopup = () => {
        setIsRegistrationRequired(false);
    };

    const joinButton = () => {
        if (isLoggined) {
            history.push(RouteConfig.Store.path);
        } else {
            setIsRegistrationRequired(true);
        }
    };

    return (
        <>
            {isRegistrationRequired && <RegistrationPopup closeRegistrationPopup={closeRegistrationPopup} />}
            <button className="ultimatedivision-join-btn" onClick={joinButton}>
                <span className="ultimatedivision-join-btn__text">JOIN BETA</span>
            </button>
        </>
    );
};
