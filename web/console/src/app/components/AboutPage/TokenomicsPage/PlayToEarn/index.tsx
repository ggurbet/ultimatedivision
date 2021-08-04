// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';

const PlayToEarn: React.FC = () => {
    return (
        <div className="play-to-earn">
            <h1 className="play-to-earn__title">
                Play to Earn
            </h1>
            <p className="play-to-earn__description">
                Playing Ultimate Division puts users into weekly competitions.
                The allocated tokens are divided every week between players, based on their performance.
                <br /><br />
                By buying in-game items and player packs, users contribute to the initial P2E fund (20% of total UDT).
                <br /><br />
                The initial Play to Earn tokens will be distributed gradually at a decreasing pace among the players for competing in UD, and will be replaced by the proceeds of the game shop.
            </p>
        </div>
    )
}

export default PlayToEarn;
