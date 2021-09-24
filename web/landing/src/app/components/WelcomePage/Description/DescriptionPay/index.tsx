// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import radar from '@static/images/Description/radar.svg';

import './index.scss';

export const DescriptionPay = () => {

    return (
        <div className="description-pay">
            <img
                className="description-pay__radar"
                src={radar}
                alt="radar diagram"
            />
            <div className="description-pay__text-area">
                <h2 className="description-pay__title">Play-to-Earn</h2>
                <p className="description-pay__text">
                    Players who own a founder player card will have access
                    to the game 2 weeks earlier, will be marked as UD founders
                    in game and will receive special NFT drops to further boost
                    their clubs. UD Founders will start the competition in the
                    top division and will maximise their Play-to-Earn profits.
                </p>
            </div>
        </div>
    );
};
