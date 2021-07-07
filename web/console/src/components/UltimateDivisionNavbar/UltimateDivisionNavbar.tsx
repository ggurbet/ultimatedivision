/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { NavLink } from 'react-router-dom';

import './UltimateDivisionNavbar.scss';

import ultimate from '../../img/MarketPlacePage/MarketPlaceNavbar/ultimate.png';

import { RouteConfig } from '../../routes';

export const UltimateDivisionNavbar: React.FC = () => {
    return (
        <div className="ultimatedivision-navbar">
            <img className="ultimatedivision-navbar__logo"
                src={ultimate}
                alt={ultimate} />
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
                        to={RouteConfig.MarketPlace.path}
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
        </div >
    );
};
