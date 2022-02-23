// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect } from 'react';
import { NavLink, useLocation } from 'react-router-dom';

import HomeNavbar from '@components/home/HomeNavbar';

import { CloseDropdownIcon, DropdownIcon } from '@/app/static/img/Navbar';
import ultimate from '@static/img/Navbar/ultimate.svg';

import { RouteConfig } from '@/app/routes';

import './index.scss';

const Navbar: React.FC = () => {
    const [isDropdownActive, setIsDropdownActive] = useState<boolean>(false);

    const [isHomePath, setIsHomePath] = useState<boolean>(false);

    const location = useLocation();

    /** Сlass visibility for navbar items. */
    const visibleClassName = isDropdownActive ? '-active' : '';

    /** TODO: DIVISIONS will be replaced with id parameter */
    const navbarItems: Array<{ name: string; path: string }> = [
        { name: 'HOME', path: RouteConfig.Home.path },
        { name: 'STORE', path: RouteConfig.Store.path },
        { name: 'CARDS', path: RouteConfig.Cards.path },
        { name: 'FIELD', path: RouteConfig.Field.path },
        { name: 'DIVISIONS', path: RouteConfig.Division.path },
    ];

    useEffect(() => {
        location.pathname === '/home'
            ? setIsHomePath(true)
            : setIsHomePath(false);
    }, [location]);

    return (
        <>
            {isHomePath ?
                <HomeNavbar />
                :
                <div className="ultimatedivision-navbar">
                    <a href={RouteConfig.Home.path}>
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
                        {isDropdownActive ?
                            <CloseDropdownIcon />
                            :
                            <DropdownIcon />
                        }
                    </div>
                    <ul
                        className={`ultimatedivision-navbar__list${visibleClassName}`}
                    >
                        {navbarItems.map((item, index) =>
                            <li
                                key={index}
                                className={`ultimatedivision-navbar__list${visibleClassName}__item`}
                            >
                                <NavLink
                                    key={index}
                                    to={item.path}
                                    className={`ultimatedivision-navbar__list${visibleClassName}__item__active`}
                                    onClick={() => setIsDropdownActive(false)}
                                >
                                    {item.name}
                                </NavLink>
                            </li>
                        )}
                    </ul>
                </div>
            }
        </>
    );
};

export default Navbar;
