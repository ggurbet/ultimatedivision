import club from '../../img/CapabilitiesPage/capabilities/club-icon.png';
import marketplace
    from '../../img/CapabilitiesPage/capabilities/marketplace-icon.png';
import weekly from '../../img/CapabilitiesPage/capabilities/weekly-icon.png';

const groupOfCapabilities = [
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
        id: '1',
    },
    {
        title: 'participate in weekly competition',
        description: `All clubs are placed in 1 of 10
        division and ranks are updated weekly.
        Win games to get promoted to the ULTIMATE
        division. Stake your UDT (ultimate division
        token) to get a percentage on your yield.
        Playing in higher divisions brings more coins,
        unique rewards and more opportunities.`,
        icon: weekly,
        id: '2',
    },
    {
        title: 'marketplace & economics',
        description: `All clubs and players are unique
        NFT items on OnFlow - the world most powerful
        blockchain protocol for NFTs and on-chain game,
        founded by Dapper Labs (CryptoKitties & NBA
        Top Shots creators). You can also accept
        smart-contract jobs as a manager for established
        clubs to earn coins, all within our game.
        Tokenize your gameplay with Ultimate Division,
        the most fair e-sport DAO powered game,
        leverage DAO principles to vote for a
        future of the project.`,
        icon: marketplace,
        id: '3',
    },
];

/* eslint-disable */

export const capabilitiesReducer = (capabilitiesState = groupOfCapabilities, action) => {
    return capabilitiesState;
};
