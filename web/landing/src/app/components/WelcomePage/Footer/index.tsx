// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import facebook from '@static/images/Footer/facebook.png';
import telegram from '@static/images/Footer/telegram.png';
import twitter from '@static/images/Footer/twitter.png';
import Subtract from '@static/images/Footer/subtract.png';

import './index.scss';

export const Footer: React.FC = () => {
    const socialList = [
        {
            id: 1,
            path: 'https://t.me/ultimatedivision',
            img: telegram,
        },
        {
            id: 2,
            path: 'http://web.facebook.com/groups/ultimatedivision/',
            img: facebook,
        },
        {
            id: 3,
            path: 'https://twitter.com/UltimateDiv',
            img: twitter,
        },
        {
            id: 4,
            path: 'https://www.reddit.com/r/UltimateDivision/',
            img: Subtract,
        },
    ];

    useEffect(() => {
        Aos.init({
            duration: 1500,
        });
    }, []);

    return (
        <footer className="footer">
            <div className="footer__wrapper">
                <div className="footer__links">
                    <a
                        className="footer__link"
                        href="https://ultimatedivision.com/ud/whitepaper/summary"
                    >
                        Whitepaper
                    </a>
                    <a
                        className="footer__link"
                        href="https://ultimatedivision.com/ud/tokenomics/udt-spending"
                    >
                        Tokenomics
                    </a>
                    <a
                        className="footer__link"
                        href="https://ultimatedivision.com/ud/whitepaper/summary"
                    >
                        FAQ
                    </a>
                </div>
                <div className="footer__social">
                    <p>follow us on social media:</p>
                    <ul className="footer__list">
                        {socialList.map((social) => (
                            <li
                                key={social.id}
                                className="footer__social-item"
                            >
                                <a
                                    className="footer__social-link"
                                    href={social.path}
                                    target="_blank"
                                    rel="noreferrer"
                                >
                                    <img
                                        className="footer__social-img"
                                        src={social.img}
                                        alt="social logo"
                                    />
                                </a>
                            </li>
                        ))}
                    </ul>
                </div>
            </div>
        </footer >
    );
};
