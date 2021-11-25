// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { Capability } from '@components/WelcomePage/Capabilities/Capability';

import club from '@static/images/Capabilities/club-icon.png';
import marketplace
    from '@static/images/Capabilities/marketplace-icon.png';
import weekly
    from '@static/images/Capabilities/weekly-icon.png';

import './index.scss';

export const Capabilities: React.FC = () => {
    const capabilities = [
        {
            title: 'build your club',
            description: `You are in charge of bringing
                your team to success on the field.
                Hire a manager or run it yourself
                and earn yield on your crypto while playing.
                You will need a squad and gameplay tactics
                that will work for your team. Treat and train
                every player to increase performance.
                Bring in ad contracts, develop player academy,
                work with your fanbase and stadiums`,
            icon: club,
            id: 1,
        },
        {
            title: 'participate in weekly competition',
            description: `All clubs are placed in 1 of 10
                divisions and ranks are updated weekly.
                Win games to get promoted to the ULTIMATE
                division. Stake your UDT (ultimate division
                token) to get a percentage on your yield.
                Playing in higher divisions brings more coins,
                unique rewards and more opportunities.`,
            icon: weekly,
            id: 2,
        },
        {
            title: 'marketplace & economics',
            description: `All clubs and players are unique
                NFT items on Flow - the world's most powerful
                blockchain protocol for NFTs and on-chain games,
                founded by Dapper Labs (CryptoKitties & NBA
                Top Shots creators). You can also accept
                smart-contract jobs as a manager for established
                clubs to earn coins, all within our game.
                Tokenize your gameplay with Ultimate Division,
                the most fair e-sport DAO powered game.
                Leverage DAO principles to vote for the future of UD`,
            icon: marketplace,
            id: 3,
        },
    ];

    return (
        <section className="group-capabilities">
            {
                capabilities.map((capability) => (
                    <Capability
                        {...capability}
                        key={capability.id}
                    />
                ))
            }
        </section>
    );
};
