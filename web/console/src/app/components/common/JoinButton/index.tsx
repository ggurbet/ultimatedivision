// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Link } from 'react-router-dom';

import { AuthRouteConfig } from '@/app/routes';

import './index.scss';

export const JoinButton: React.FC = () =>
    <Link
        className="ultimatedivision-join-btn"
        to={AuthRouteConfig.SignIn.path}
    >
        <button className="ultimatedivision-join-btn">
            <span className="ultimatedivision-join-btn__text">JOIN BETA</span>
        </button>
    </Link>;

