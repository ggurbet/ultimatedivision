// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { NavLink } from 'react-router-dom';

import { RouteConfig } from '@/app/routes';

import { setScrollAble } from '@/app/internal/setScrollAble';

import { CloseDropdownIcon, DropdownIcon } from '@/app/static/img/Navbar';

import './index.scss';

const HomeNavbar: React.FC = () => {
    const [isDropdownActive, setIsDropdownActive] = useState<boolean>(false);

    /** Сlass visibility for navbar items. */
    const navbarClassName = isDropdownActive ? '-active' : '';
    const navbarWrapperClassName = isDropdownActive ? 'wrapper--active' : '';

    const changeNavbarDropdownActivity = () => {
        setIsDropdownActive(!isDropdownActive);
        setScrollAble();
    };

    /** TODO: DIVISIONS will be replaced with id parameter */
    const navbarItems: Array<{ pageName: string; path: string }> = [
        { pageName: 'Home', path: RouteConfig.Home.path },
        { pageName: 'My Club', path: RouteConfig.Field.path },
        { pageName: 'Store', path: RouteConfig.Store.path },
        { pageName: 'FAQ', path: RouteConfig.Summary.path },
    ];

    return (
        <nav className="ultimatedivision-home-navbar">
            <div className={`wrapper ${navbarWrapperClassName}`}>
                <div className="ultimatedivision-home-navbar__wrapper">
                    {isDropdownActive && <p className="ultimatedivision-home-navbar__text">Menu</p>}
                    <div
                        className="ultimatedivision-home-navbar__dropdown"
                        onClick={() => changeNavbarDropdownActivity()}
                    >
                        {isDropdownActive ? <CloseDropdownIcon /> : <DropdownIcon />}
                    </div>
                </div>

                <ul
                    className={`ultimatedivision-home-navbar__list${navbarClassName}`}
                >
                    {navbarItems.map((item, index) =>
                        <li
                            key={index}
                            className={`ultimatedivision-home-navbar__list${navbarClassName}__item`}
                        >
                            <NavLink
                                key={index}
                                to={item.path}
                                className={`ultimatedivision-home-navbar__list${navbarClassName}__item__active`}
                                onClick={() => changeNavbarDropdownActivity()}
                            >
                                {item.pageName}
                            </NavLink>
                        </li>
                    )}
                </ul>
            </div>
        </nav>
    );
};

export default HomeNavbar;
