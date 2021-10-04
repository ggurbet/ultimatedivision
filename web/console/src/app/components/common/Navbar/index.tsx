// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import { useState } from 'react';
import { NavLink } from 'react-router-dom';
import { RouteConfig } from '@/app/router';
import { DropdownIcon } from '@/app/static/img/Navbar';
import ultimate from '@static/img/Navbar/ultimate.svg';

import './index.scss';

export const Navbar: React.FC = () => {
    const [isDropdownActive, setIsDropdownActive] = useState<boolean>(false);

    const navbarItems: Array<{ name: string; path: string }> = [
        { name: 'HOME', path: RouteConfig.Summary.path },
        { name: 'STORE', path: RouteConfig.Store.path },
        { name: 'MARKETPLACE', path: RouteConfig.MarketPlace.path },
        { name: 'CARDS', path: RouteConfig.Club.path },
        { name: 'FIELD', path: RouteConfig.FootballField.path },
    ];

    return (
        <div className="ultimatedivision-navbar">
            <a href="/">
                <img
                    className="ultimatedivision-navbar__logo"
                    src={ultimate}
                    alt="UltimateDivision logo"
                />
            </a>
            <div
                className="ultimatedivision-navbar__dropdown"
                onClick={() => setIsDropdownActive(!isDropdownActive)}
            >
                <DropdownIcon />
            </div>
            <ul
                className={`ultimatedivision-navbar__list${
                    isDropdownActive ? '-active' : ''
                }`}
            >
                {navbarItems.map((item, index) =>
                    <li className="ultimatedivision-navbar__item">
                        <NavLink
                            key={index}
                            to={item.path}
                            className="ultimatedivision-navbar__item__active"
                            onClick={() => setIsDropdownActive(false)}
                        >
                            {item.name}
                        </NavLink>
                    </li>
                )}
            </ul>
        </div>
    );
};
