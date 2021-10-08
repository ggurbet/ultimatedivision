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
                    Club Owners who hold a Founder Collection NFT will be
                    awarded the in-game title of UD Founder. The UD Founders
                    will receive exclusive airdrops and will start
                    the game in UD’s top division.
                </p>
            </div>
        </div>
    );
};
