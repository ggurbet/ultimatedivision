//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.

import { useState } from 'react';

import { DropdownStyle } from '@/app/utils/dropdownStyle';

import ultimate from '@static/img/Navbar/ultimate.png';
import triangle from '@static/img/FootballFieldPage/triangle.svg';

import './index.scss';
import { Link } from 'react-router-dom';
import { RouteConfig } from '@/app/routes';

export const AboutMenu = () => {
    const [whitePaperVisibility, changeWhitePaperVisibility] = useState(false);
    const [tokenomicsVisibility, changeTokenomicsVisibility] = useState(false);

    const LIST_HEIGHT = 130;
    const whitePaperStyle = new DropdownStyle(whitePaperVisibility, LIST_HEIGHT);
    const tokenomicsStyle = new DropdownStyle(tokenomicsVisibility, LIST_HEIGHT);

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
                    RouteConfig.WhitePaper.subRoutes &&
                    RouteConfig.WhitePaper.subRoutes.map((item, index) =>
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
                    RouteConfig.Tokenomics.subRoutes &&
                    RouteConfig.Tokenomics.subRoutes.map((item, index) =>
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
    );
};
