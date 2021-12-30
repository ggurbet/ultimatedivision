// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { NavLink } from 'react-router-dom';

import discord from '@static/img/gameLanding/footer/discord.svg';
import twitter from '@static/img/gameLanding/footer/twitter.svg';

import { RouteConfig } from '@/app/routes';

import './index.scss';

export const Footer: React.FC = () => {
    const socialList = [
        {
            id: 1,
            path: 'https://discord.com/invite/ultimatedivision',
            img: discord,
            text: 'Discord',
        },
        {
            id: 2,
            path: 'https://twitter.com/UltimateDivnft',
            img: twitter,
            text: 'Twitter',
        },
    ];

    return (
        <footer className="footer">
            <div className="footer__wrapper">
                <div className="footer__links">
                    <a className="footer__link" href={RouteConfig.Spending.path}>
                    Whitepaper
                    </a>
                    <a className="footer__link" href={RouteConfig.Summary.path}>
                        FAQ
                    </a>
                </div>
                <div className="footer__social">
                    <ul className="footer__list">
                        {socialList.map((social) =>
                            <a
                                key={social.id}
                                className="footer__social-item"
                                id={`link-${social.id}`}
                                href={social.path}
                                target="_blank"
                                rel="noreferrer"
                            >
                                <div className="footer__social-link">
                                    <img
                                        className="footer__social-img"
                                        src={social.img}
                                        alt="social logo"
                                    />
                                    <span className="footer__social-text">
                                        {social.text}
                                    </span>
                                </div>
                            </a>
                        )}
                    </ul>
                </div>
            </div>
        </footer>
    );
};
