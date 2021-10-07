// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';
import { RouteConfig } from '@/app/router';
import { DropdownStyle } from '@/app/utils/dropdownStyle';
import ultimate from '@static/img/Navbar/ultimate.svg';
import triangle from '@static/img/FootballFieldPage/triangle.svg';

import './index.scss';

export const AboutMenu = () => {
    const [whitePaperVisibility, changeWhitePaperVisibility] = useState(false);
    const [tokenomicsVisibility, changeTokenomicsVisibility] = useState(false);

    const whitePaperStyle = new DropdownStyle(whitePaperVisibility);
    const tokenomicsStyle = new DropdownStyle(tokenomicsVisibility);

    const path = useLocation().pathname;
    const shouldBeShowed = path.includes('tokenomics') || path.includes('whitepaper');

    const menuFields = {
        whitepaper:
            [
                'Summary',
                'Game Mechanics',
                'Play to Earn and Economy',
                'Technology',
                'Team',
            ],
        tokenomics:
            [
                'UDT Spending',
                'Play to Earn',
                'Staking',
                'UD DAO Fund',
            ],
    };

    return (
        shouldBeShowed ?
            <div className="about-menu">
                <div className="about-menu__logo-wrapper">
                    <img src={ultimate} alt="ultimate logo" />
                </div>
                <div
                    className="about-menu__whitepaper"
                    onClick={() => changeWhitePaperVisibility(prev => !prev)}
                >
                    <h2>Whitepaper</h2>
                    <img
                        className="about-menu__whitepaper-image"
                        src={triangle}
                        style={{ transform: whitePaperStyle.triangleRotate }}
                        alt="triangle img"
                    />
                </div>
                <ul
                    className="about-menu__whitepaper-list"
                    style={{ height: whitePaperStyle.listHeight }}
                >
                    {
                        RouteConfig.Whitepaper.children &&
                        RouteConfig.Whitepaper.children.map((item, index) =>
                            <li
                                key={index}
                                className="about-menu__whitepaper-item"
                            >
                                <Link
                                    to={item.path}
                                    className="about-menu__whitepaper-link"
                                >
                                    {menuFields.whitepaper[index]}
                                </Link>
                            </li>
                        )
                    }
                </ul>
                <div
                    className="about-menu__tokenomics"
                    onClick={() => changeTokenomicsVisibility(prev => !prev)}
                >
                    <h2>Tokenomics</h2>
                    <img
                        className="about-menu__whitepaper-image"
                        src={triangle}
                        style={{ transform: tokenomicsStyle.triangleRotate }}
                        alt="triangle img"
                    />
                </div>
                <ul
                    className="about-menu__tokenomics-list"
                    style={{ height: tokenomicsStyle.listHeight }}
                >
                    {
                        RouteConfig.Tokenomics.children &&
                        RouteConfig.Tokenomics.children.map((item, index) =>
                            <li
                                key={index}
                                className="about-menu__tokenomics-item"
                            >
                                <Link
                                    to={item.path}
                                    className="about-menu__tokenomics-link"
                                >
                                    {menuFields.tokenomics[index]}
                                </Link>
                            </li>
                        )
                    }
                </ul>
            </div>
            :
            <></>
    );
};
