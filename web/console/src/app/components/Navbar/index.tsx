/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import { NavLink } from 'react-router-dom';

import { RouteConfig } from '../../routes';

import ultimate from '../../static/img/Navbar/ultimate.png';

import './index.scss';

export const Navbar: React.FC = () =>
    <div className="ultimatedivision-navbar">
        <img
            className="ultimatedivision-navbar__logo"
            src={ultimate}
            alt="UltimateDivision logo"
        />
        <ul className="ultimatedivision-navbar__list">
            <li className="ultimatedivision-navbar__item">
                <NavLink
                    to={RouteConfig.Default.path}
                    className="ultimatedivision-navbar__item__active"
                >
                    HOME
                </NavLink>
            </li>
            <li className="ultimatedivision-navbar__item">
                <NavLink
                    to={RouteConfig.MarketPlace.path}
                    className="ultimatedivision-navbar__item__active"
                >
                    MARKETPLACE
                </NavLink>
            </li>
            <li className="ultimatedivision-navbar__item">
                <NavLink
                    to={RouteConfig.MyCards.path}
                    className="ultimatedivision-navbar__item__active"
                >
                    CLUB
                </NavLink>
            </li>
            <li className="ultimatedivision-navbar__item">
                <NavLink
                    to={RouteConfig.FootballField.path}
                    className="ultimatedivision-navbar__item__active"
                >
                    FIELD
                </NavLink>
            </li>
        </ul >
    </div >;
