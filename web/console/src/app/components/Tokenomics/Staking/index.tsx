// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';

const Staking: React.FC = () =>
    <div className="staking">
        <h1 className="staking__title">
            Staking
        </h1>
        <p className="staking__description">
            UDT holders will be able to stake their tokens for periods of 1,3,6,12 months.
            Stakers will be able to claim their rewards at the same weekly intervals as division competitions progress.
            <br /><br />
            For extra rewards on the staked coins, stakers will be required to participate in the weekly UD competitions.
            The minimum play requirement will scale with the stake amount.
            <br /><br />
            For stakers who are not interested in playing the game as much, an option to form a contract
            with other players will be available.
            By hiring players, the stakers will contribute towards the Play to Earn factor of the game.
        </p>
    </div>;


export default Staking;
