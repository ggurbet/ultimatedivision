/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';
import { NavLink } from 'react-router-dom';

import './MarketPlaceNavbar.scss';

import ultimate
    from '../../../img/MarketPlacePage/MarketPlaceNavbar/ultimate.png';

export const MarketPlaceNavbar = () => {
    return (
        <div className="marketplace-navbar">
            <img className="marketplace-navbar__logo"
                src={ultimate}
                alt={ultimate} />
            <ul className="marketplace-navbar__list">
                <li className="marketplace-navbar__item">
                    <NavLink to="/ud"
                        className="marketplace-navbar__item__active">
                        HOME
                    </NavLink>
                </li>
                <li className="marketplace-navbar__item">
                    <NavLink to="/ud/marketplace"
                        className="marketplace-navbar__item__active">
                        MARKETPLACE
                    </NavLink>
                </li>
                <li className="marketplace-navbar__item">
                    <NavLink className="marketplace-navbar__item__active"
                        to="/ud/club">
                        CLUB
                    </NavLink>
                </li>
            </ul >
        </div>
    );
};
