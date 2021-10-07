// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import twitter from '@static/images/Footer/twitter.svg';
import discord from '@static/images/Footer/discord.svg';

import './index.scss';

export const Footer: React.FC = () => {
    const socialList = [
        {
            id: 1,
            path: 'https://twitter.com/UltimateDivnft',
            img: twitter,
        },
        {
            id: 2,
            path: 'https://discord.com/invite/ultimatedivision',
            img: discord,
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
                        href="https://ultimatedivision.com/ud/whitepaper/summary"
                    >
                        FAQ
                    </a>
                </div>
                <div className="footer__social">
                    <p className="footer__text">follow us on social media:</p>
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
