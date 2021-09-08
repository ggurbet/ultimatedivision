// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useEffect } from 'react';
import Aos from 'aos';

import './index.scss';

export const HeadingInformation: React.FC = () => {
    useEffect(() => {
        Aos.init({
            duration: 1500,
        });
    });

    return (
        <div className="heading-information" data-aos="fade-left">
            <h1 className="heading-information__title">OWN THE GAME</h1>
            <p className="heading-information__description">
                Join the world of community-driven football.
                Build your own club and own it on the blockchain.
                Train and develop your players to win trophies or
                trade them for profit on the NFT marketplace.
                You can build a grand club, become a pro manager
                or earn money as a professional agent.
                Choose your own path and engage in PvP matches
                with other players powered by smart-contracts.
            </p>
        </div>
    );
};
